import {fetchArgs, IRequest} from "../../helpers/requestSender";
import ZhguApiResponseSchema from "./zhguApiResponseSchema";
import {ITrackingArgs, ITrackingByBillNumberResponse, OneTrackingEvent} from "../../../types";
import {IDatetime} from "../../helpers/datetime";
import {IBillNumberTracker} from "../base";
import {BaseContainerConstructor} from "../../base";
import {IUserAgentGenerator} from "../../helpers/userAgentGenerator";
import {NotThisShippingLineException} from "../../../exceptions";


export class ZhguRequest {
    protected request: IRequest<fetchArgs>;
    protected userAgentGenerator: IUserAgentGenerator

    public constructor(request: IRequest<fetchArgs>, userAgentGenerator: IUserAgentGenerator) {
        this.request = request;
        this.userAgentGenerator = userAgentGenerator;
    }

    public async getApiResponse(args: ITrackingArgs): Promise<ZhguApiResponseSchema> {
        return await this.request.sendRequestAndGetJson({
            url: `http://elines.zhonggu56.com/api/booking/getVoyageInfo`,
            method: "POST",
            headers: {
                "accept": "application/json, text/plain, */*",
                "accept-language": "en",
                "content-type": "application/json;charset=UTF-8",
                "Referer": "http://elines.zhonggu56.com/track",
                "Referrer-Policy": "strict-origin-when-cross-origin",
                "User-Agent": this.userAgentGenerator.generateUserAgent()
            },
            body: JSON.stringify({blNo: args.number})
        })
    }
}

export class ZhguEtdParser {
    public datetime: IDatetime;

    public constructor(datetime: IDatetime) {
        this.datetime = datetime;
    }

    public getEtd(apiResp: ZhguApiResponseSchema): OneTrackingEvent {
        let etd = apiResp.object[0].etd
        if (!etd) {
            throw new Error("cannot find etd")
        }
        return {
            time: this.datetime.strptime(etd, "YYYY-MM-DD").getTime(),
            operationName: "ETD",
            location: apiResp.object[0].portFromName,
            vessel: apiResp.object[0].vesselName !== "" || apiResp.object[0].vesselName !== null ? apiResp.object[0].vesselName : " "
        }
    }
}

export class ZhguAtdParser extends ZhguEtdParser {
    public getAtd(apiResp: ZhguApiResponseSchema): OneTrackingEvent {
        let atd = apiResp.object[0].atd
        if (!atd.length) {
            throw new Error("cannot find atd")
        }
        return {
            time: this.datetime.strptime(atd, "YYYY-MM-DD").getTime(),
            operationName: "ATD",
            location: apiResp.object[0].portFromName,
            vessel: apiResp.object[0].vesselName !== "" || apiResp.object[0].vesselName !== null ? apiResp.object[0].vesselName : " "
        }
    }
}

export class ZhguAtaParser extends ZhguAtdParser {
    public getAta(apiResp: ZhguApiResponseSchema) {
        let ata = apiResp.object[0].ata
        if (!ata.length) {
            throw new Error("cannot find ata")
        }
        return {
            time: this.datetime.strptime(ata, "YYYY-MM-DD").getTime(),
            operationName: "ATA",
            location: apiResp.object[0].portFromName,
            vessel: apiResp.object[0].vesselName !== "" || apiResp.object[0].vesselName !== null ? apiResp.object[0].vesselName : " "
        }
    }
}

export class ZhguEtaParser extends ZhguAtaParser {
    public getEta(apiResp: ZhguApiResponseSchema): number {
        let etaStr = apiResp.object[0].eta
        if (!etaStr.length) {
            throw  new Error("cannot find eta")
        }
        return this.datetime.strptime(etaStr, "YYYY-MM-DD").getTime()
    }
}

export class ZhguInfoAboutMovingParser extends ZhguAtaParser {
    public getInfoAboutMoving(apiResp: ZhguApiResponseSchema): OneTrackingEvent[] {
        let outputArr: OneTrackingEvent[] = []
        try {
            let etd = this.getEtd(apiResp)
            outputArr.push(etd)
        } catch (e) {
        }
        try {
            let atd = this.getAtd(apiResp)
            outputArr.push(atd)
        } catch (e) {
        }
        try {
            let ata = this.getAta(apiResp)
            outputArr.push(ata)
        } catch (e) {
        }
        return outputArr
    }
}


export class ZhguBillNumber implements IBillNumberTracker {
    protected request: ZhguRequest;
    protected infoAboutMovingParser: ZhguInfoAboutMovingParser;
    protected etaParser: ZhguEtaParser;

    public constructor(args: BaseContainerConstructor<fetchArgs>) {
        this.request = new ZhguRequest(args.requestSender, args.UserAgentGenerator)
        this.infoAboutMovingParser = new ZhguInfoAboutMovingParser(args.datetime)
        this.etaParser = new ZhguEtaParser(args.datetime)
    }

    public async trackByBillNumber(args: ITrackingArgs): Promise<ITrackingByBillNumberResponse> {
        try {
            let resp = await this.request.getApiResponse(args)
            let infoAboutMoving = this.infoAboutMovingParser.getInfoAboutMoving(resp)
            try {
                let eta = this.etaParser.getEta(resp)
                return {billNo: args.number, scac: "ZHGU", infoAboutMoving: infoAboutMoving, etaFinalDelivery: eta}
            } catch (e) {
                return {billNo: args.number, scac: "ZHGU", infoAboutMoving: infoAboutMoving, etaFinalDelivery: 1}
            }
        } catch (e) {
            throw new NotThisShippingLineException()
        }

    }
}