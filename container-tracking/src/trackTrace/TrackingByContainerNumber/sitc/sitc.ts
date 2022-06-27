import {
    BaseTrackerByContainerNumber,
    ITrackingArgs,
    TrackingContainerResponse,
    OneTrackingEvent, BaseContainerConstructor
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
        let infoAboutMoving: OneTrackingEvent[] = []
        for (let item of sitcApiResp.data.list) {
            let infoAboutMovingOneEvent: OneTrackingEvent = {
                time: this.datetime.strptime(item.eventDate, "YYYY-MM-DD HH:mm:ss").getTime(),
                operationName: item.movementNameEn,
                vessel: item.vesselCode,
                location: item.eventPort
            }
            infoAboutMoving.push(infoAboutMovingOneEvent)
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