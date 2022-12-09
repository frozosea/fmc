import {
    BaseContainerConstructor,
    BaseTrackerByContainerNumber,
    ITrackingArgs,
    OneTrackingEvent,
    TrackingContainerResponse
} from "../../base";
import {SitcContainerTrackingApiResponseSchema} from "./sitcApiResponseSchema";
import {fetchArgs, IRequest} from "../../helpers/requestSender";
import {IDatetime} from "../../helpers/datetime";
import {NotThisShippingLineException} from "../../../exceptions";

export class SitcRequest {
    public request: IRequest<fetchArgs>;

    public constructor(request: IRequest<fetchArgs>) {
        this.request = request;
    }

    public async getApiResponseJson(args: ITrackingArgs): Promise<SitcContainerTrackingApiResponseSchema> {
        return await this.request.sendRequestAndGetJson({
            url: `http://api.sitcline.com/ecm/cmcontainerhistory/movementSearchApp?containerNo=${args.number}`,
            method: "POST"
        })

    }

}

export class SitcInfoAboutMovingParser {
    public datetime: IDatetime;

    public constructor(datetime: IDatetime) {
        this.datetime = datetime
    }

    public getInfoAboutMoving(sitcApiResp: SitcContainerTrackingApiResponseSchema): OneTrackingEvent[] {
        let infoAboutMoving = []
        for (let item of sitcApiResp.data.list) {
            let oneEvent = {}
            try {
                oneEvent["time"] = this.datetime.strptime(item.eventDate, "YYYY-MM-DD HH:mm:ss").getTime()
            } catch (e) {
            }
            try {
                oneEvent["operationName"] = item.movementNameEn
            } catch (e) {
            }
            try {
                oneEvent["location"] = item.vesselCode
            } catch (e) {
            }
            try {
                oneEvent["vessel"] = item.eventPort
            } catch (e) {
            }
            if(Object.keys(oneEvent).length!==0){
                infoAboutMoving.push(oneEvent)

            }
        }
        return infoAboutMoving
    }
}

export class SitcContainer extends BaseTrackerByContainerNumber<fetchArgs> {
    protected request: SitcRequest;
    protected infoAboutMovingParser: SitcInfoAboutMovingParser;

    public constructor(args: BaseContainerConstructor<fetchArgs>) {
        super(args);
        this.request = new SitcRequest(args.requestSender)
        this.infoAboutMovingParser = new SitcInfoAboutMovingParser(args.datetime)
    }


    public async trackContainer(args: ITrackingArgs): Promise<TrackingContainerResponse> {
        let apiResp: SitcContainerTrackingApiResponseSchema = await this.request.getApiResponseJson(args);
        let infoAboutMoving: OneTrackingEvent[] = this.infoAboutMovingParser.getInfoAboutMoving(apiResp)
        if (!infoAboutMoving.length) {
            throw new NotThisShippingLineException()
        }
        //has no container size in tracking by container number ¯\_(ツ)_/¯
        return {container: args.number, containerSize: "", scac: "SITC", infoAboutMoving: infoAboutMoving}
    }
}