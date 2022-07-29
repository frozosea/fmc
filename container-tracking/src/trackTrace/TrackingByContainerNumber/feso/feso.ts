import {
    BaseContainerConstructor,
    BaseTrackerByContainerNumber,
    ITrackingArgs,
    OneTrackingEvent,
    TrackingContainerResponse
} from "../../base";
import {FesoApiFullResponseSchema, FesoApiResponse} from "./fescoApiResponseSchemas";
import {NotThisShippingLineException} from "../../../exceptions";
import {fetchArgs, IRequest} from "../../helpers/requestSender";
import {IDatetime} from "../../helpers/datetime";


export class FesoRequest {
    protected request: IRequest<fetchArgs>

    public constructor(request: IRequest<fetchArgs>) {
        this.request = request
    }

    public async sendRequestToFescoGraphQlApiAndGetJsonResponse(container: string): Promise<FesoApiResponse> {
        const FESO_API_URL = "https://tracking.fesco.com/api/v1/tracking/get";
        let body = {
            "codes": [
                container
            ],
            "email": null,
            "forDate": null,
            "fromFile": false
        }
        let response: FesoApiFullResponseSchema = await this.request.sendRequestAndGetJson({
            url: FESO_API_URL,
            method: "POST",
            headers: {
                "accept": "application/json, text/plain, */*",
                "accept-language": "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6,zh-CN;q=0.5,zh;q=0.4",
                "content-type": "application/json",
                "sec-ch-ua": "\".Not/A)Brand\";v=\"99\", \"Google Chrome\";v=\"103\", \"Chromium\";v=\"103\"",
                "sec-ch-ua-mobile": "?0",
                "sec-ch-ua-platform": "\"macOS\"",
                "sec-fetch-dest": "empty",
                "sec-fetch-mode": "cors",
                "sec-fetch-site": "cross-site",
                "Referer": "https://www.fesco.ru/",
                "Referrer-Policy": "strict-origin-when-cross-origin"
            },
            body: JSON.stringify(body)
        });
        return JSON.parse(response.containers[0])
    }
}

export class FesoInfoAboutMovingParser {
    protected datetime: IDatetime;

    public constructor(datetime: IDatetime) {
        this.datetime = datetime;
    }

    public getInfoAboutMoving(fescoApiResponse: FesoApiResponse): OneTrackingEvent[] {
        let infoAboutMoving: OneTrackingEvent[] = [];
        let lastEvents = fescoApiResponse.lastEvents
        for (let item of lastEvents) {
            let time = item.time.split("Z")[0]
            let oneOperationObject: OneTrackingEvent = {
                time: this.datetime.strptime(time, "YYYY-MM-DDTHH:mm:ss.SSS").getTime(),
                operationName: item.operationNameLatin.trim(),
                location: item.locNameLatin.trim(),
                vessel: item.vessel ? item.vessel : " "
            }
            infoAboutMoving.push(oneOperationObject)
        }
        return infoAboutMoving
    }
}

export class FesoContainerSizeParser {
    public getContainerSize(fescoJson: FesoApiResponse): string {
        return fescoJson.containerCTCode
    }
}

export class FesoApiParser {
    public infoAboutMovingParser: FesoInfoAboutMovingParser;
    public containerSizeParser: FesoContainerSizeParser;

    public constructor(datetime: IDatetime) {
        this.infoAboutMovingParser = new FesoInfoAboutMovingParser(datetime);
        this.containerSizeParser = new FesoContainerSizeParser();
    }

    public getOutputObjectAndGetEta(fescoApiResponse: FesoApiResponse): TrackingContainerResponse {
        let containerSize: string = this.containerSizeParser.getContainerSize(fescoApiResponse);
        let infoAboutMoving = this.infoAboutMovingParser.getInfoAboutMoving(fescoApiResponse);
        return {
            container: fescoApiResponse.container,
            scac: "FESO",
            containerSize: containerSize,
            infoAboutMoving: infoAboutMoving
        }
    }
}

export class FesoContainer extends BaseTrackerByContainerNumber<fetchArgs> {
    protected request: FesoRequest;
    protected apiParser: FesoApiParser;

    public constructor(args: BaseContainerConstructor<fetchArgs>) {
        super(args);
        this.request = new FesoRequest(args.requestSender)
        this.apiParser = new FesoApiParser(args.datetime);
    }

    public async trackContainer(args: ITrackingArgs): Promise<TrackingContainerResponse> {
        try {
            let response = await this.request.sendRequestToFescoGraphQlApiAndGetJsonResponse(args.number);
            return this.apiParser.getOutputObjectAndGetEta(response)
        } catch (e) {
            throw new NotThisShippingLineException()

        }

    }
}

