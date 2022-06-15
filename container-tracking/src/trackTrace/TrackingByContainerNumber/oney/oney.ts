import {
    BaseTrackerByContainerNumber,
    ITrackingArgs,
    TrackingContainerResponse,
    OneTrackingEvent, BaseContainerConstructor
} from "../../base";
import {
    OneyGetContainerSizeSchema,
    OneyGetBillNumberResponse,
    OneyInfoAboutMovingSchema
} from "./oneyApiResponseSchemas";
import {fetchArgs, IRequest} from "../../helpers/requestSender";
import {NotThisShippingLineException} from "../../../exceptions";
import {IUserAgentGenerator} from "../../helpers/userAgentGenerator";
import {IDatetime} from "../../helpers/datetime";
import {CopNo, BkgNo} from "../../../types";
import {config} from "../../../../tests/classesConfigurator";
import RequestsUtils from "../../helpers/utils/requestsUtils";

export class OneyRequest {
    public request: IRequest<fetchArgs>
    protected userAgentGenerator: IUserAgentGenerator;

    public constructor(request: IRequest<fetchArgs>, userAgentGenerator: IUserAgentGenerator) {
        this.request = request
        this.userAgentGenerator = userAgentGenerator
    }

    protected getHeaders(): object {
        return {
            'Host': 'ecomm.one-line.com',
            'Accept': 'application/json, text/javascript, */*; q=0.01',
            'Accept-Language': 'en-US,en;q=0.5',
            'Accept-Encoding': 'gzip, deflate, br',
            'Connection': 'keep-alive',
            'content-type': 'application/x-www-form-urlencoded',
            'origin': 'https://ecomm.one-line.com',
            'referer': 'https://ecomm.one-line.com/ecom/CUP_HOM_3301.do',
            'sec-ch-ua': '" Not;A Brand";v = "99", "Google Chrome";v = "97", "Chromium";v = "97"',
            'sec-ch-ua-mobile': '?0',
            'sec-ch-ua-platform': "macOS",
            'sec-fetch-dest': 'empty',
            'sec-fetch-mode': 'cors',
            'sec-fetch-site': 'same - origin',
            'user-agent': this.userAgentGenerator.generateUserAgent(),
            'x-requested-with': 'XMLHttpRequest'
        }
    }

    protected async sendReqToOneyApiWithJsonBody(jsonBody: object): Promise<any> {
        return await this.request.sendRequestAndGetJson({
            url: "https://ecomm.one-line.com/ecom/CUP_HOM_3301GS.do",
            method: "POST",
            headers: this.getHeaders(),
            body: RequestsUtils.jsonToQueryString(jsonBody)
        })
    }

    public async sendReqToOneyApiAndGetFirstDataResp(container: string): Promise<OneyGetBillNumberResponse> {
        return await this.sendReqToOneyApiWithJsonBody({
            'f_cmd': '122',
            'cust_cd': '',
            'cntr_no': container,
            'search_type': 'C'
        });
    }

    public async sendReqToOneyApiAndGetInfoAboutMovingResponse(container: string, bookingNumber: string, copNo: string): Promise<OneyInfoAboutMovingSchema> {
        return await this.sendReqToOneyApiWithJsonBody({
            'f_cmd': '125',
            'cntr_no': container,
            'bkg_no': bookingNumber,
            'cop_no': copNo
        });
    }

    public async sendReqToOneyApiAndGetContainerSizeResponse(container: string, bkgNo: string, copNo: string): Promise<OneyGetContainerSizeSchema> {
        return await this.sendReqToOneyApiWithJsonBody({
            'f_cmd': '123',
            'cntr_no': container,
            'bkg_no': bkgNo,
            'cop_no': copNo
        });
    }
}

export class OneyCopNoAndBkgNoParser {
    public getCopNoAndBkgNo(billNumberOneyApiResponse: OneyGetBillNumberResponse): [CopNo, BkgNo] {
        return [billNumberOneyApiResponse.list[0].copNo, billNumberOneyApiResponse.list[0].bkgNo]
    }
}

export class OneyInfoAboutMovingParser {
    protected datetime: IDatetime;

    public constructor(datetime: IDatetime) {
        this.datetime = datetime
    }

    public parseInfoAboutMoving(infoAboutMovingApiResp: OneyInfoAboutMovingSchema): OneTrackingEvent[] {
        let infoAboutMoving: OneTrackingEvent[] = []
        for (let item of infoAboutMovingApiResp.list) {
            let oneInfoObject: OneTrackingEvent = {
                time: this.datetime.strptime(item.eventDt, "YYYY-MM-DD HH:mm").getTime(),
                operationName: item.statusNm,
                location: item.placeNm,
                vessel: item.vslEngNm
            }
            infoAboutMoving.push(oneInfoObject)
        }
        return infoAboutMoving
    }
}

export class OneyContainerSizeParser {
    public getContainerSize(oneyContainerSizeApiResp: OneyGetContainerSizeSchema): string {
        return oneyContainerSizeApiResp.list[0].cntrTpszNm
    }
}

export class OneyContainer extends BaseTrackerByContainerNumber<fetchArgs> {
    protected request: OneyRequest;
    protected infoAboutMovingParser: OneyInfoAboutMovingParser;
    protected containerSizeParser: OneyContainerSizeParser;
    protected CopNoAndBkgNoParser: OneyCopNoAndBkgNoParser;

    public constructor(args: BaseContainerConstructor<fetchArgs>) {
        super(args);
        this.request = new OneyRequest(args.requestSender, args.UserAgentGenerator);
        this.infoAboutMovingParser = new OneyInfoAboutMovingParser(args.datetime);
        this.containerSizeParser = new OneyContainerSizeParser();
        this.CopNoAndBkgNoParser = new OneyCopNoAndBkgNoParser();
    }

    public async trackContainer(args: ITrackingArgs): Promise<TrackingContainerResponse> {
        try {
            let [copNo, bkgNo] = this.CopNoAndBkgNoParser.getCopNoAndBkgNo(await this.request.sendReqToOneyApiAndGetFirstDataResp(args.container));
            let containerSize: string = this.containerSizeParser.getContainerSize(await this.request.sendReqToOneyApiAndGetContainerSizeResponse(args.container, bkgNo, copNo));
            let infoAboutMoving: OneTrackingEvent[] = this.infoAboutMovingParser.parseInfoAboutMoving(await this.request.sendReqToOneyApiAndGetInfoAboutMovingResponse(args.container, bkgNo, copNo))
            return {
                container: args.container,
                containerSize: containerSize,
                scac: "ONEY",
                infoAboutMoving: infoAboutMoving
            }
        } catch (e) {
            throw new NotThisShippingLineException()
        }

    }
}

// (async () => {
//     let oney = new OneyContainer({
//         datetime: config.DATETIME,
//         requestSender: config.REQUEST_SENDER,
//         UserAgentGenerator: config.USER_AGENT_GENERATOR
//     })
//     let actualResponse = await oney.trackContainer({container: "GAOU6642924"})
//     console.log(actualResponse)
// })();