import {
    BaseContainerConstructor,
    BaseTrackerByContainerNumber,
    ITrackingArgs,
    OneTrackingEvent,
    TrackingContainerResponse
} from "../../base";
import {fetchArgs, IRequest} from "../../helpers/requestSender";
import {MscuApiResponseSchema} from "./mscuApiResponseSchema";
import {GetEtaException, NotThisShippingLineException} from "../../../exceptions";
import {IDatetime} from "../../helpers/datetime";
import {IUserAgentGenerator} from "../../helpers/userAgentGenerator";


export class MscuRequest {
    protected request: IRequest<fetchArgs>;
    protected userAgentGenerator: IUserAgentGenerator;

    public constructor(request: IRequest<fetchArgs>, userAgentGenerator: IUserAgentGenerator) {
        this.request = request;
        this.userAgentGenerator = userAgentGenerator
    }

    public async getApiJson(args: ITrackingArgs): Promise<MscuApiResponseSchema> {
        return await this.request.sendRequestAndGetJson({
            url: "https://www.msc.com/api/feature/tools/TrackingInfo",
            method: "POST",
            headers: {
                "accept": "application/json, text/plain, */*",
                "accept-language": "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6,zh-CN;q=0.5,zh;q=0.4",
                "content-type": "application/json",
                "sec-ch-ua": "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"101\", \"Google Chrome\";v=\"101\"",
                "sec-ch-ua-mobile": "?0",
                "sec-ch-ua-platform": "\"macOS\"",
                "sec-fetch-dest": "empty",
                "sec-fetch-mode": "cors",
                "sec-fetch-site": "same-origin",
                "x-requested-with": "XMLHttpRequest",
                "user-agent": this.userAgentGenerator.generateUserAgent()
            },
            body: JSON.stringify({trackingNumber: args.number, trackingMode: 0})
        })
    }
}


export class MscuContainerSizeParser {
    public getContainerSize(apiResp: MscuApiResponseSchema): string {
        return apiResp.Data.BillOfLadings[0].ContainersInfo[0].ContainerType
    }
}

class BaseMscuParser {
    protected datetime: IDatetime;

    public constructor(datetime: IDatetime) {
        this.datetime = datetime
    }
}

export class MscuInfoAboutMovingParser extends BaseMscuParser {
    public getInfoAboutMoving(apiResp: MscuApiResponseSchema): OneTrackingEvent[] {
        let infoAboutMovingArray: OneTrackingEvent[] = []
        for (let item of apiResp.Data.BillOfLadings[0].ContainersInfo[0].Events) {
            let event: OneTrackingEvent = {
                time: this.datetime.strptime(item.Date, "DD/MM/YYYY").getTime(),
                operationName: item.Description,
                location: item.Location,
                vessel: ""
            }
            infoAboutMovingArray.push(event)
        }
        return infoAboutMovingArray
    }
}

export class MscuEtaParser extends BaseMscuParser {
    public getEta(apiResp: MscuApiResponseSchema): OneTrackingEvent {
        let etaDate: string = apiResp.Data.BillOfLadings[0].GeneralTrackingInfo.FinalPodEtaDate
        if (etaDate === "") {
            throw new GetEtaException()
        }
        return {
            time: this.datetime.strptime(etaDate, "DD/MM/YYYY").getTime(),
            operationName: "ETA",
            location: "",
            vessel: ""
        }
    }
}


export class MscuContainer extends BaseTrackerByContainerNumber<fetchArgs> {
    protected request: MscuRequest;
    protected infoAboutMovingParser: MscuInfoAboutMovingParser;
    protected etaParser: MscuEtaParser;
    protected containerSizeParser: MscuContainerSizeParser;

    public constructor(args: BaseContainerConstructor<fetchArgs>) {
        super(args);
        this.request = new MscuRequest(args.requestSender, args.UserAgentGenerator);
        this.infoAboutMovingParser = new MscuInfoAboutMovingParser(args.datetime)
        this.etaParser = new MscuEtaParser(args.datetime)
        this.containerSizeParser = new MscuContainerSizeParser()
    }

    public async trackContainer(args: ITrackingArgs): Promise<TrackingContainerResponse> {
        let apiResp: MscuApiResponseSchema = await this.request.getApiJson(args)
        try {
            let infoAboutMoving: OneTrackingEvent[] = this.infoAboutMovingParser.getInfoAboutMoving(apiResp)
            try {
                let eta: OneTrackingEvent = this.etaParser.getEta(apiResp)
                infoAboutMoving.push(eta)
            } catch (e) {
            }
            return {
                container: args.number,
                containerSize: this.containerSizeParser.getContainerSize(apiResp),
                scac: "MSCU",
                infoAboutMoving: infoAboutMoving
            };
        } catch (e) {
            throw new NotThisShippingLineException()
        }

    }
}

