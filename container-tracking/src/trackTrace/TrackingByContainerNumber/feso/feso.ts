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
        const FESCO_GRAPHQL_API_URL = "https://tracking.fesco.com/graphql";
        let body = `{\"query\":\"query trackingData($codes: [String!], $fromFile: Boolean, $request: String, $filename: String, $email: String, $forDate: String) {\\n  tracking {\\n    data(codes: $codes, fromFile: $fromFile, request: $request, filename: $filename, email: $email, forDate: $forDate) {\\n      requestKey\\n      containers\\n      missing\\n      __typename\\n    }\\n    __typename\\n  }\\n}\\n\",\"variables\":{\"codes\":[\"${container}\"],\"fromFile\":false,\"request\":null,\"filename\":null,\"email\":null,\"forDate\":null},\"operationName\":\"trackingData\"}`
        let response: FesoApiFullResponseSchema = await this.request.sendRequestAndGetJson({
            url: FESCO_GRAPHQL_API_URL,
            method: "POST",
            headers: {
                "accept": "*/*",
                "accept-language": "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6,zh-CN;q=0.5,zh;q=0.4",
                "content-type": "application/json",
                "sec-ch-ua": "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"101\", \"Google Chrome\";v=\"101\"",
                "sec-ch-ua-mobile": "?0",
                "sec-ch-ua-platform": "\"macOS\"",
                "sec-fetch-dest": "empty",
                "sec-fetch-mode": "cors",
                "sec-fetch-site": "cross-site"
            },
            body: body
        });
        console.log(response)
        return JSON.parse(response.data.tracking.data.containers[0])
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
                operationName: item.operationNameLatin,
                location: item.locNameLatin,
                vessel: item.vessel ? item.vessel : ""
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

