import {
    SkluApiParser, SkluContainer,
    SkluEtaParser,
    SkluInfoAboutMovingParser
} from "../../src/trackTrace/TrackingByContainerNumber/sklu/sklu";
import {config} from "../classesConfigurator";
import {
    expectedInfoAboutMoving,
    expectedResult,
    skluApiResponseExample,
    UnlocodesRepoMoch
} from "./skluApiResponseExample";
import {fetchArgs, IRequest} from "../../src/trackTrace/helpers/requestSender";

const assert = require("assert");
const fs = require("fs");
const path = require("path")


const requestMoch: IRequest<fetchArgs> = {
    async sendRequestAndGetJson(_: fetchArgs): Promise<any> {
        return skluApiResponseExample
    },
    async sendRequestAndGetHtml(_: fetchArgs): Promise<string> {
        return fs.readFileSync(path.resolve(__dirname, './skluInfoAboutMovingExampleHtml.html')).toString("utf-8")
    }
}

describe("SKLU container tracking test", () => {
    let skluParser = new SkluApiParser()
    let unlocodesRepoMoch = new UnlocodesRepoMoch()
    const container = "TEMU2094051"
    it("SKLU info about moving parser test", () => {
        let infoAboutMovingParser = new SkluInfoAboutMovingParser(config.DATETIME);
        let data = fs.readFileSync(path.resolve(__dirname, './skluInfoAboutMovingExampleHtml.html'))
        let infoAboutMoving = infoAboutMovingParser.parseInfoAboutMovingPage(data.toString(), container)
        assert.deepEqual(infoAboutMoving, expectedInfoAboutMoving)

    })
    it("SKLU eta parser test", async () => {
        let etaParser = new SkluEtaParser(unlocodesRepoMoch);
        let etaObj = await etaParser.getEtaObject(skluParser.parseSinokorApiJson(skluApiResponseExample))
        assert.strictEqual(etaObj.time, new Date(new Date(skluApiResponseExample[0].ETA).toUTCString()).getTime())
        // assert.strictEqual(etaObj.location, await unlocodesRepoMoch.getUnlocode(""))
        assert.strictEqual(etaObj.operationName, "ETA")
    })
    it("SKLU api parser test", () => {
        let parsedApiResp = skluParser.parseSinokorApiJson(skluApiResponseExample)
        assert.strictEqual(parsedApiResp.eta, new Date(new Date(skluApiResponseExample[0].ETA).toUTCString()).getTime())
        assert.strictEqual(parsedApiResp.billNo, skluApiResponseExample[0].BKNO)
        assert.strictEqual(parsedApiResp.containerSize, skluApiResponseExample[0].CNTR)
    })
    it("SKLU main class with moch test", async () => {
        let sklu = new SkluContainer({
            requestSender: requestMoch,
            datetime: config.DATETIME,
            UserAgentGenerator: config.USER_AGENT_GENERATOR
        }, unlocodesRepoMoch)
        let actualResult = await sklu.trackContainer({container: container})
        if (!actualResult.infoAboutMoving.length) {
            throw new assert.AssertionError()
        }
        assert.strictEqual(actualResult.scac, "SKLU")
        assert.strictEqual(actualResult.container, container)
        assert.strictEqual(actualResult.containerSize, expectedResult.containerSize)
        // assert.deepEqual(actualResult, expectedResult)
    }).timeout(100000000)

})