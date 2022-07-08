import {
    ISitcBillNumberRequest,
    SitcBillNumber,
    SitcEtaParser
} from "../../src/trackTrace/trackingBybillNumber/sitc/sitc";
import {config} from "../classesConfigurator";
import {SitcBillNumberResponse, sitcExpectedResult} from "./sitcExpectedResult";
import SitcBillNumberApiResponseSchema from "../../src/trackTrace/trackingBybillNumber/sitc/sitcApiResponseSchema";
import {ICaptcha} from "../../src/trackTrace/trackingBybillNumber/sitc/captchaResolver";
import {fetchArgs, IRequest} from "../../src/trackTrace/helpers/requestSender";

const assert = require("assert");

export const requestMoch: IRequest<fetchArgs> = {
    async sendRequestAndGetJson(_: fetchArgs): Promise<any> {
        return sitcExpectedResult.SITU9130070
    },
    async sendRequestAndGetHtml(_: fetchArgs): Promise<string> {
        return ""
    },
    async sendRequestAndGetImage(_: fetchArgs): Promise<any> {
        return ""
    }
}

export const billRequestMoch: ISitcBillNumberRequest = {
    async getApiResponse(_: { billNo: string, solvedCaptcha: string, randomString: string }): Promise<SitcBillNumberApiResponseSchema> {
        return SitcBillNumberResponse
    }
}
export const captchaSolverMoch: ICaptcha = {
    async getSolvedCaptchaAndRandomString(): Promise<[string, string]> {
        return ["1488", "2948919824798142"]
    }
}
describe("SITC tracking by bill number test", () => {
    it("eta parser test", () => {
        let etaParser = new SitcEtaParser(config.DATETIME)
        const expectedEta = 1657684800000
        assert.strictEqual(etaParser.getEta(SitcBillNumberResponse), expectedEta)
    })
    it("main class with moch test", async () => {
        const sitcBill = new SitcBillNumber({
            datetime: config.DATETIME,
            requestSender: requestMoch,
            UserAgentGenerator: config.USER_AGENT_GENERATOR
        }, captchaSolverMoch, billRequestMoch)
        const billNo = "SITDLVK222G951"
        let result = await sitcBill.trackByBillNumber({number: billNo})
        assert.strictEqual(result.scac, "SITC")
        assert.strictEqual(result.billNo, billNo)
        for (let item = 0; item < result.infoAboutMoving.length; item++) {
            assert.strictEqual(result.infoAboutMoving[item].operationName, sitcExpectedResult.SITU9130070.data.list[item].movementNameEn);
            assert.strictEqual(result.infoAboutMoving[item].time, config.DATETIME.strptime(sitcExpectedResult.SITU9130070.data.list[item].eventDate, "YYYY-MM-DD HH:mm:ss").getTime());
            assert.strictEqual(result.infoAboutMoving[item].vessel, sitcExpectedResult.SITU9130070.data.list[item].vesselCode);
            assert.strictEqual(result.infoAboutMoving[item].location, sitcExpectedResult.SITU9130070.data.list[item].eventPort);
        }
    })
})