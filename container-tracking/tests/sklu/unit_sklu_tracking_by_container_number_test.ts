import {
    SkluApiParser,
    SkluContainer,
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
        return fs.readFileSync(path.resolve(__dirname, './skluInfoAboutMovingExampleHtml.txt')).toString("utf-8")
    },
    async sendRequestAndGetImage(_: fetchArgs): Promise<any> {
        return ""
    }
}

export const SkluinfoAboutMovingTest = (pathToHtml: string, container: string) => {
    let infoAboutMovingParser = new SkluInfoAboutMovingParser(config.DATETIME);
    let data = fs.readFileSync(path.resolve(__dirname, pathToHtml))
    let infoAboutMoving = infoAboutMovingParser.parseInfoAboutMovingPage(data.toString(), container)
    assert.deepEqual(infoAboutMoving, expectedInfoAboutMoving)
}
export const SkluEtaParserTest = async (example, unlocodesRepoMoch, skluParser) => {
    let etaParser = new SkluEtaParser(unlocodesRepoMoch);
    let etaObj = await etaParser.getEtaObject(skluParser.parseSinokorApiJson(example))
    assert.strictEqual(etaObj.time, new Date(new Date(example[0].ETA).toUTCString()).getTime())
    assert.strictEqual(etaObj.operationName, "ETA")
}
export const SkluApiParserTest = (skluParser, example) => {
    let parsedApiResp = skluParser.parseSinokorApiJson(example)
    assert.strictEqual(parsedApiResp.eta, new Date(new Date(example[0].ETA).toUTCString()).getTime())
    assert.strictEqual(parsedApiResp.billNo, example[0].BKNO)
    assert.strictEqual(parsedApiResp.containerSize, example[0].CNTR)
}
describe("SKLU container tracking test", () => {
    let skluParser = new SkluApiParser(config.DATETIME)
    let unlocodesRepoMoch = new UnlocodesRepoMoch()
    const container = "TEMU2094051"
    it("SKLU info about moving parser test", () => {
        SkluinfoAboutMovingTest('./skluInfoAboutMovingExampleHtml.txt', container)

    })
    it("SKLU eta parser test", async () => {
        await SkluEtaParserTest(skluApiResponseExample, unlocodesRepoMoch, skluParser)
    })
    it("SKLU api parser test", () => {
        SkluApiParserTest(skluParser,skluApiResponseExample)
    })
    it("SKLU main class with moch test", async () => {
        let sklu = new SkluContainer({
            requestSender: requestMoch,
            datetime: config.DATETIME,
            UserAgentGenerator: config.USER_AGENT_GENERATOR
        }, unlocodesRepoMoch)
        let actualResult = await sklu.trackContainer({number: container})
        if (!actualResult.infoAboutMoving.length) {
            throw new assert.AssertionError()
        }
        assert.strictEqual(actualResult.scac, "SKLU")
        assert.strictEqual(actualResult.container, container)
        assert.strictEqual(actualResult.containerSize, expectedResult.containerSize)
        assert.deepEqual(actualResult, expectedResult)
    }).timeout(100000000)

})