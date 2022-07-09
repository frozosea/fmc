import {SitcContainer, SitcInfoAboutMovingParser, SitcRequest} from "../../TrackingByContainerNumber/sitc/sitc";
import {IBillNumberTracker} from "../base";
import {ITrackingArgs, ITrackingByBillNumberResponse, OneTrackingEvent} from "../../../types";
import {ICaptcha} from "./captchaResolver";
import {BaseContainerConstructor} from "../../base";
import {fetchArgs} from "../../helpers/requestSender";
import SitcBillNumberApiResponseSchema, {SitcContainerMovementInfoSchema} from "./sitcApiResponseSchema";
import {IDatetime} from "../../helpers/datetime";

export interface ISitcBillNumberRequest {
    getBillNoResponse(args: { billNo: string, solvedCaptcha: string, randomString: string }): Promise<SitcBillNumberApiResponseSchema>;

    getContainerInfo(args: { billNo: string, containerNo: string }): Promise<SitcContainerMovementInfoSchema>;
}

export class SitcBillNumberRequest extends SitcRequest implements ISitcBillNumberRequest {
    public async getBillNoResponse(args: { billNo: string, solvedCaptcha: string, randomString: string }): Promise<SitcBillNumberApiResponseSchema> {
        return await this.request.sendRequestAndGetJson({
            url: `http://api.sitcline.com/doc/cargoTrack/searchApp?blNo=${args.billNo}&code=${args.solvedCaptcha}&randomStr=${args.randomString}`,
            method: "POST",
            headers: {
                "accept": "application/json, text/plain, */*",
                "accept-language": "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6,zh-CN;q=0.5,zh;q=0.4",
                // "cookie": "rememberMe=true; username=flyasea; Hm_lvt_6c6c344c19846b08289aece9f968bd6f=1655363989,1657166543; password=qCbbbVcqDxGb6GMzs6bWLk11MXsF5K0xgTumeXo1GX/4BbpCdWZ8y0ZMRBXntKfj3pr/vwghgsRtcYqLKiLFUA==; Hm_lpvt_6c6c344c19846b08289aece9f968bd6f=1657266306",
                "Referer": "http://api.sitcline.com/app/cargoTrackSearch",
                "Referrer-Policy": "strict-origin-when-cross-origin"
            },
            body: null
        })
    }

    public async getContainerInfo(args: { billNo: string, containerNo: string }): Promise<SitcContainerMovementInfoSchema> {
        return await this.request.sendRequestAndGetJson({
            url: `http://api.sitcline.com/doc/cargoTrack/movementDetailApp?blNo=${args.billNo}&containerNo=${args.containerNo}`,
            method: "POST",
            headers: {
                "accept": "application/json, text/plain, */*",
                "accept-language": "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6,zh-CN;q=0.5,zh;q=0.4",
                // "cookie": "rememberMe=true; username=flyasea; Hm_lvt_6c6c344c19846b08289aece9f968bd6f=1655363989,1657166543; password=qCbbbVcqDxGb6GMzs6bWLk11MXsF5K0xgTumeXo1GX/4BbpCdWZ8y0ZMRBXntKfj3pr/vwghgsRtcYqLKiLFUA==; Hm_lpvt_6c6c344c19846b08289aece9f968bd6f=1657266306",
                "Referer": "http://api.sitcline.com/app/cargoTrack",
                "Referrer-Policy": "strict-origin-when-cross-origin"
            },
            body: null
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

export class SitcBillInfoAboutMovingParser extends SitcInfoAboutMovingParser {
    public parseInfoAboutMoving(response: SitcContainerMovementInfoSchema): OneTrackingEvent[] {
        let events: OneTrackingEvent[] = []
        for (let item of response.data.list) {
            let event: OneTrackingEvent = {
                time: this.datetime.strptime(item.eventdate, "YYYY-MM-DD").getTime(),
                operationName: item.movementnameen, location: item.portname, vessel: ""
            }
            events.push(event)
        }
        return events
    }
}

export class SitcBillNumber extends SitcContainer implements IBillNumberTracker {
    protected captchaSolver: ICaptcha;
    protected billRequest: ISitcBillNumberRequest
    protected containerNumberParser: SitcContainerNumberParser;
    protected etaParser: SitcEtaParser;
    protected infoAboutMovingParser: SitcBillInfoAboutMovingParser

    public constructor(args: BaseContainerConstructor<fetchArgs>, captchaSolver: ICaptcha, sitcBillNumberRequest: ISitcBillNumberRequest) {
        super(args);
        this.billRequest = sitcBillNumberRequest;
        this.captchaSolver = captchaSolver;
        this.containerNumberParser = new SitcContainerNumberParser();
        this.etaParser = new SitcEtaParser(this.datetime);
        this.infoAboutMovingParser = new SitcBillInfoAboutMovingParser(this.datetime);
    }

    public async trackByBillNumber(args: ITrackingArgs): Promise<ITrackingByBillNumberResponse> {
        let [solvedCaptcha, randomString] = await this.captchaSolver.getSolvedCaptchaAndRandomString();
        let response: SitcBillNumberApiResponseSchema = await this.billRequest.getBillNoResponse({
            billNo: args.number,
            solvedCaptcha: solvedCaptcha,
            randomString: randomString
        })
        let eta: number = this.etaParser.getEta(response)
        let containerNumber: string = this.containerNumberParser.getNumber(response)
        let containerNumberTrackingResponse = await this.billRequest.getContainerInfo({
            billNo: args.number,
            containerNo: containerNumber
        })
        let infoAboutMoving: OneTrackingEvent[] = this.infoAboutMovingParser.parseInfoAboutMoving(containerNumberTrackingResponse)
        return {billNo: args.number, scac: "SITC", infoAboutMoving: infoAboutMoving, etaFinalDelivery: eta}
    }
}