import {
    OneTrackingEvent, BaseContainerConstructor,
    BaseTrackerByContainerNumber,
    TrackingContainerResponse,
    ITrackingArgs
} from "../../base";
import {CosuApiResponseSchema, CosuInfoAboutMoving, EtaResponseSchema} from "./cosuApiResponseSchema";
import {NotThisShippingLineException} from "../../../exceptions";
import {fetchArgs, IRequest} from "../../helpers/requestSender";
import {IUserAgentGenerator} from "../../helpers/userAgentGenerator";
import {IDatetime} from "../../helpers/datetime";

export class CosuRequest {
    protected UserAgentGenerator: IUserAgentGenerator;
    protected requestSender: IRequest<fetchArgs>

    public constructor(UserAgentGenerator: IUserAgentGenerator, requestSender: IRequest<fetchArgs>) {
        this.UserAgentGenerator = UserAgentGenerator
        this.requestSender = requestSender
    }

    protected getHeaders(): object {
        return {
            'Accept': '*/*',
            'Accept-Language': 'ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6,zh;q=0.5',
            'Connection': 'keep-alive',
            'Host': 'elines.coscoshipping.com',
            'language': 'en_US',
            'Referer': 'https://elines.coscoshipping.com/ebusiness/cargoTracking',
            'sec-ch-ua': '" Not A;Brand";v="99", "Chromium";v="98", "Google Chrome";v="98"',
            'sec-ch-ua-mobile': '?0',
            'Sec-Fetch-Dest': 'empty',
            'Sec-Fetch-Mode': 'cors',
            'Sec-Fetch-Site': 'same-origin',
            'sys': 'eb',
            'User-Agent': this.UserAgentGenerator.generateUserAgent()

        }
    }

    protected getTimeStamp(): number {
        return new Date().getTime();
    }

    public async getCosuInfoAboutMovingJson(container: string): Promise<CosuApiResponseSchema> {
        return await this.requestSender.sendRequestAndGetJson({
            url: `https://elines.coscoshipping.com/ebtracking/public/containers/${container}?timestamp=${this.getTimeStamp()}`,
            method: "GET",
            headers: this.getHeaders()
        })
    }

    public async getEtaJson(container: string): Promise<EtaResponseSchema> {
        return await this.requestSender.sendRequestAndGetJson({
            url: `https://elines.coscoshipping.com/ebtracking/public/container/eta/${container}?timestamp=${this.getTimeStamp()}`,
            method: "GET",
            headers: this.getHeaders()
        })
    }

    public getPod(cosuEtaResp: CosuApiResponseSchema): string {
        let pod = cosuEtaResp.data.content.containers[0].container.pod
        return pod !== null ? pod : " "
    }

}

export class CosuEtaParser {
    protected datetime: IDatetime;

    public constructor(datetime: IDatetime) {
        this.datetime = datetime
    }

    public getEtaObject(cosuEtaResp: EtaResponseSchema, pod: string): OneTrackingEvent {
        let rawEta: string = cosuEtaResp.data.content
        let etaTimeStamp: number = this.datetime.strptime(rawEta, "YYYY-MM-DD HH:mm").getTime()
        return {time: etaTimeStamp, operationName: "ETA", vessel: " ", location: pod}
    }

}

export class CosuPodParser {
    public getPod(cosuApiResp: CosuApiResponseSchema): string {
        let pod = cosuApiResp.data.content.containers[0].container.pod
        return pod !== null ? pod : " "
    }
}

export class CosuContainerSizeParser {
    public getContainerSize(apiResp: CosuApiResponseSchema): string {
        return apiResp.data.content.containers[0].container.containerType
    }
}

export class CosuInfoAboutMovingParser {
    protected datetime: IDatetime;

    public constructor(datetime: IDatetime) {
        this.datetime = datetime
    }

    public getInfoAboutMoving(apiResp: CosuApiResponseSchema): OneTrackingEvent[] {
        let infoAboutMoving: OneTrackingEvent[] = []
        let rawContainerHistoryInfo: CosuInfoAboutMoving[] = apiResp.data.content.containers[0].containerHistorys
        for (let item of rawContainerHistoryInfo) {
            let oneInfoAboutMovingObject: OneTrackingEvent = {
                time: item.timeOfIssue ? this.datetime.strptime(item.timeOfIssue, "YYYY-MM-DD HH:mm").getTime() : 0,
                operationName: item.containerNumberStatus ? item.containerNumberStatus : " ",
                location: item.location ? item.location : " ",
                vessel: item.transportation ? item.transportation : " "
            }
            infoAboutMoving.push(oneInfoAboutMovingObject)
        }
        return infoAboutMoving.sort()
    }
}

export class CosuContainer extends BaseTrackerByContainerNumber<fetchArgs> {
    protected request: CosuRequest;
    protected infoAboutMovingParser: CosuInfoAboutMovingParser;
    protected etaParser: CosuEtaParser
    protected podParser: CosuPodParser
    protected containerSizeParser: CosuContainerSizeParser;

    public constructor(args: BaseContainerConstructor<fetchArgs>) {
        super(args);
        this.request = new CosuRequest(args.UserAgentGenerator, this.requestSender);
        this.infoAboutMovingParser = new CosuInfoAboutMovingParser(args.datetime);
        this.etaParser = new CosuEtaParser(args.datetime);
        this.podParser = new CosuPodParser();
        this.containerSizeParser = new CosuContainerSizeParser();
    }

    protected checkContainerFound(apiResp: CosuApiResponseSchema, container: string): boolean {
        //{"code":"200","message":"","data":{"content":{"containers":[],"notFound":"SKLU1637340"}}}
        return container !== apiResp.data.content.notFound;
    }

    public async trackContainer(args: ITrackingArgs): Promise<TrackingContainerResponse> {
        let rawInfoAboutMoving: CosuApiResponseSchema = await this.request.getCosuInfoAboutMovingJson(args.number);
        if (this.checkContainerFound(rawInfoAboutMoving, args.number)) {
            let rawEtaResp: EtaResponseSchema = await this.request.getEtaJson(args.number);
            let pod = this.podParser.getPod(rawInfoAboutMoving);
            let etaObject: OneTrackingEvent = this.etaParser.getEtaObject(rawEtaResp, pod);
            let infoAboutMoving: OneTrackingEvent[] = this.infoAboutMovingParser.getInfoAboutMoving(rawInfoAboutMoving);
            infoAboutMoving.push(etaObject)
            let containerSize: string = this.containerSizeParser.getContainerSize(rawInfoAboutMoving);
            return {
                container: args.number,
                scac: "COSU",
                containerSize: containerSize,
                infoAboutMoving: infoAboutMoving
            }
        } else {
            throw new NotThisShippingLineException()
        }
    }
}