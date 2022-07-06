import {FesoContainer} from "../../TrackingByContainerNumber/feso/feso";
import {IBillNumberTracker} from "../base";
import {ITrackingArgs, ITrackingByBillNumberResponse} from "../../../types";
import {FesoApiResponse} from "../../TrackingByContainerNumber/feso/fescoApiResponseSchemas";
import {fetchArgs} from "../../helpers/requestSender";
import {BaseContainerConstructor} from "../../base";
import {NotThisShippingLineException} from "../../../exceptions";
import {IDatetime} from "../../helpers/datetime";


export class FesoEtaParser {
    protected datetime: IDatetime;

    public constructor(datetime: IDatetime) {
        this.datetime = datetime;
    }

    public GetEta(fescoApiResponse: FesoApiResponse): number {
        let lastEvents = fescoApiResponse.lastEvents
        for (let item of lastEvents) {
            if (item.operationNameLatin === "ETA") {
                let time = item.time.split("Z")[0]
                return this.datetime.strptime(time,"YYYY-MM-DDTHH:mm:ss.SSS").getTime();
            }
        }
        return 0
    }
}


export class FesoBillNumber extends FesoContainer implements IBillNumberTracker {
    protected etaParser: FesoEtaParser

    public constructor(args: BaseContainerConstructor<fetchArgs>) {
        super(args);
        this.etaParser = new FesoEtaParser(args.datetime)
    }

    public async trackByBillNumber(args: ITrackingArgs): Promise<ITrackingByBillNumberResponse> {
        try {
            let response = await this.request.sendRequestToFescoGraphQlApiAndGetJsonResponse(args.number)
            let outputObject = this.apiParser.getOutputObjectAndGetEta(response)
            for (let event of outputObject.infoAboutMoving) {
                if (event.operationName === "ETA") {
                    outputObject.infoAboutMoving.splice(outputObject.infoAboutMoving.indexOf(event), 1)
                }
            }
            return {
                billNo: args.number,
                scac: "FESO",
                infoAboutMoving: outputObject.infoAboutMoving,
                etaFinalDelivery: this.etaParser.GetEta(response)
            }
        } catch (e) {
            throw new NotThisShippingLineException()
        }
    }
}