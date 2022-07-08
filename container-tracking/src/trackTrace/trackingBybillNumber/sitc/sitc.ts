import {SitcContainer, SitcRequest} from "../../TrackingByContainerNumber/sitc/sitc";
import {IBillNumberTracker} from "../base";
import {ITrackingArgs, ITrackingByBillNumberResponse, OneTrackingEvent} from "../../../types";
import {ICaptcha} from "./captchaResolver";
import {BaseContainerConstructor} from "../../base";
import {fetchArgs} from "../../helpers/requestSender";
import SitcBillNumberApiResponseSchema from "./sitcApiResponseSchema";
import {IDatetime} from "../../helpers/datetime";
import {NotThisShippingLineException} from "../../../exceptions";

export interface ISitcBillNumberRequest {
    getApiResponse(args: { billNo: string, solvedCaptcha: string, randomString: string }): Promise<SitcBillNumberApiResponseSchema>
}

export class SitcBillNumberRequest extends SitcRequest {
    public async getApiResponse(args: { billNo: string, solvedCaptcha: string, randomString: string }): Promise<SitcBillNumberApiResponseSchema> {
        return await this.request.sendRequestAndGetJson({
            url: `http://api.sitcline.com/doc/cargoTrack/searchApp?blNo=${args.billNo}&code=${args.solvedCaptcha}&randomStr=${args.randomString}`,
            method: "POST",
            headers: {
                "accept": "application/json, text/plain, */*",
                "accept-language": "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6,zh-CN;q=0.5,zh;q=0.4",
                // "cookie": "rememberMe=true; username=flyasea; Hm_lvt_6c6c344c19846b08289aece9f968bd6f=1655363989,1657166543; password=qCbbbVcqDxGb6GMzs6bWLk11MXsF5K0xgTumeXo1GX/4BbpCdWZ8y0ZMRBXntKfj3pr/vwghgsRtcYqLKiLFUA==; Hm_lpvt_6c6c344c19846b08289aece9f968bd6f=1657266306",
                "Referer": "http://api.sitcline.com/app/cargoTrackSearch",
                "Referrer-Policy": "strict-origin-when-cross-origin"
            }
        })
    }
}

export class SitcContainerNumberParser {
    public getNumber(response: SitcBillNumberApiResponseSchema): string {
        return response.data.list3[0].containerNo
    }
}

export class SitcEtaParser {
    protected datetime: IDatetime;

    public constructor(datetime: IDatetime) {
        this.datetime = datetime;
    }

    public getEta(response: SitcBillNumberApiResponseSchema): number {
        let cctb = response.data.list2[response.data.list2.length - 1].cctb
        if (cctb === null) {
            return this.datetime.strptime(response.data.list2[response.data.list2.length - 1].eta, "YYYY-MM-DD HH:mm:SS").getTime()
        }
        return this.datetime.strptime(cctb, "YYYY-MM-DD HH:mm").getTime()
    }
}

export class SitcBillNumber extends SitcContainer implements IBillNumberTracker {
    protected captchaSolver: ICaptcha;
    protected billRequest: ISitcBillNumberRequest
    protected containerNumberParser: SitcContainerNumberParser;
    protected etaParser: SitcEtaParser;

    public constructor(args: BaseContainerConstructor<fetchArgs>, captchaSolver: ICaptcha, sitcBillNumberRequest: ISitcBillNumberRequest) {
        super(args);
        this.billRequest = sitcBillNumberRequest;
        this.captchaSolver = captchaSolver;
        this.containerNumberParser = new SitcContainerNumberParser();
        this.etaParser = new SitcEtaParser(this.datetime);
    }

    public async trackByBillNumber(args: ITrackingArgs): Promise<ITrackingByBillNumberResponse> {
        let solvedCaptcha: string
        let randomString: string
        try {
            [solvedCaptcha, randomString] = await this.captchaSolver.getSolvedCaptchaAndRandomString();
        } catch (e) {
        }
        try {
            [solvedCaptcha, randomString] = await this.captchaSolver.getSolvedCaptchaAndRandomString();
        } catch (e) {
            throw new NotThisShippingLineException();
        }
        let response: SitcBillNumberApiResponseSchema = await this.billRequest.getApiResponse({
            billNo: args.number,
            solvedCaptcha: solvedCaptcha,
            randomString: randomString
        })
        let eta: number = this.etaParser.getEta(response)
        let containerNumber: string = this.containerNumberParser.getNumber(response)
        let containerNumberTrackingResponse = await this.request.getApiResponseJson({number: containerNumber})
        let infoAboutMoving: OneTrackingEvent[] = this.infoAboutMovingParser.getInfoAboutMoving(containerNumberTrackingResponse)
        return {billNo: args.number, scac: "SITC", infoAboutMoving: infoAboutMoving, etaFinalDelivery: eta}

    }
}