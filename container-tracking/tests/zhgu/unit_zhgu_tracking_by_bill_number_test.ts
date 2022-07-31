import {describe} from "mocha";
import {
    ZhguAtdParser,
    ZhguBillNumber,
    ZhguEtaParser,
    ZhguEtdParser,
    ZhguInfoAboutMovingParser
} from "../../src/trackTrace/trackingBybillNumber/zhgu/zhgu";
import {zhguExampleResponseWithoutMistakes} from "./exampleData";
import {baseArgs} from "../tracking/unit_tracking_for_other_countries_test";
import {fetchArgs, IRequest} from "../../src/trackTrace/helpers/requestSender";

const assert = require("assert");

const requestMoch: IRequest<fetchArgs> = {
    async sendRequestAndGetHtml(_: fetchArgs): Promise<string> {
        return ""
    },
    async sendRequestAndGetJson(_: fetchArgs): Promise<any> {
        return zhguExampleResponseWithoutMistakes
    },
    async sendRequestAndGetImage(_: fetchArgs): Promise<any> {
        return ""
    }
}


describe("ZHGU tracking by bill number test", () => {
    it("etd parser test", () => {
        let etdParser = new ZhguEtdParser(baseArgs.datetime)
        let resp = etdParser.getEtd(zhguExampleResponseWithoutMistakes)
        assert.strictEqual(resp.operationName, "ETD")
        assert.strictEqual(resp.vessel, "ZHONG GU BO HAI")
        assert.strictEqual(resp.location, "SHANGHAI")
        assert.strictEqual(resp.time, baseArgs.datetime.strptime("2022-07-22", "YYYY-MM-DD").getTime())
    })
    it("atd parser test", () => {
        let etdParser = new ZhguAtdParser(baseArgs.datetime)
        let resp = etdParser.getAtd(zhguExampleResponseWithoutMistakes)
        assert.strictEqual(resp.operationName, "ATD")
        assert.strictEqual(resp.vessel, "ZHONG GU BO HAI")
        assert.strictEqual(resp.location, "SHANGHAI")
        assert.strictEqual(resp.time, baseArgs.datetime.strptime("2022-07-25", "YYYY-MM-DD").getTime())
    })
    it("eta parser test", () => {
        let etdParser = new ZhguEtaParser(baseArgs.datetime)
        let resp = etdParser.getEta(zhguExampleResponseWithoutMistakes)
        assert.strictEqual(resp, baseArgs.datetime.strptime("2022-07-31", "YYYY-MM-DD").getTime())
    })
    it("infoAboutMovingParser parser test", () => {
        let infoAboutMovingParser = new ZhguInfoAboutMovingParser(baseArgs.datetime)
        let resp = infoAboutMovingParser.getInfoAboutMoving(zhguExampleResponseWithoutMistakes)
        assert.strictEqual(resp.length, 2)
    })
    it("main class with moch test", async () => {
        let zhgu = new ZhguBillNumber({
            requestSender: requestMoch,
            datetime: baseArgs.datetime,
            UserAgentGenerator: baseArgs.UserAgentGenerator
        })
        let resp = await zhgu.trackByBillNumber({number: "ZGSHA0100001921"})
        const expectedInfoAboutMoving = [
            {
                "location": "SHANGHAI",
                "operationName": "ETD",
                "time": 1658448000000,
                "vessel": "ZHONG GU BO HAI"
            },
            {
                "location": "SHANGHAI",
                "operationName": "ATD",
                "time": 1658707200000,
                "vessel": "ZHONG GU BO HAI"
            }
        ]
        assert.strictEqual(resp.scac, "ZHGU")
        assert.strictEqual(resp.billNo, "ZGSHA0100001921")
        assert.deepEqual(resp.infoAboutMoving, expectedInfoAboutMoving)
        assert.strictEqual(resp.etaFinalDelivery, 1659225600000)
    })
})