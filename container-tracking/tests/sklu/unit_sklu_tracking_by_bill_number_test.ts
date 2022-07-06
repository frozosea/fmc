import {fetchArgs, IRequest} from "../../src/trackTrace/helpers/requestSender"
import {skluApiResponseExample, UnlocodesRepoMoch} from "./skluApiResponseExample";
import {SkluBillNumber, SkluContainerNumberParser} from "../../src/trackTrace/trackingBybillNumber/sklu/sklu";
import {config} from "../classesConfigurator";

const path = require("path")
const assert = require("assert")
const fs = require("fs")
let isContainerInfoAboutMoving = false

const requestMoch: IRequest<fetchArgs> = {
    async sendRequestAndGetHtml(_: fetchArgs): Promise<string> {
        if (!isContainerInfoAboutMoving) return fs.readFileSync(path.resolve(__dirname, './skluBillNumberInfoAboutMovingWithManyContainersExample.html')).toString("utf-8")
        else return fs.readFileSync(path.resolve(__dirname, './skluInfoAboutMovingExampleHtml.html')).toString("utf-8")
    },
    async sendRequestAndGetJson(_: fetchArgs): Promise<any> {
        return skluApiResponseExample
    }
}

describe("SKLU track by bill number test", () => {
    let unlocodesRepoMoch = new UnlocodesRepoMoch()

    it("SKLU container number parser test", () => {
        let data = fs.readFileSync(path.resolve(__dirname, './skluBillNumberInfoAboutMovingWithManyContainersExample.html')).toString("utf-8")
        const containerNumberParser = new SkluContainerNumberParser()
        let expectedContainer = "SKLU1327134"
        assert.strictEqual(containerNumberParser.getContainerNumberByStringHtml(data), expectedContainer)
    })
    it("SKLU main class with moch test", async () => {
        const skluBillNumber = new SkluBillNumber({
            requestSender: requestMoch,
            datetime: config.DATETIME,
            UserAgentGenerator: config.USER_AGENT_GENERATOR
        }, unlocodesRepoMoch)
        isContainerInfoAboutMoving = true
        let result = await skluBillNumber.trackByBillNumber({number: "TEMU2094051"})
        assert.strictEqual(result.etaFinalDelivery,1655424000000)
    })
})