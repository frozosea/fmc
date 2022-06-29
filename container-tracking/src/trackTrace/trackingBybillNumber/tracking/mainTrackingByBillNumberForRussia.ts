import {TimeInspector} from "../../TrackingByContainerNumber/tracking/mainTrackingForRussia";
import {IBillNumberTracker} from "../base";
import {TrackingArgsWithScac} from "../../base";
import {ContainerNotFoundException} from "../../../exceptions";
import {ITrackingByBillNumberResponse, SCAC_TYPE} from "../../../types";

export interface TrackingByBillNumberConstructorArgs {
    feso: IBillNumberTracker;
    sklu: IBillNumberTracker;
    halu: IBillNumberTracker;
}

export interface BillNumberTrackingInRussia {
    FESO: IBillNumberTracker;
    SKLU: IBillNumberTracker;
    HALU: IBillNumberTracker;
}

export default class MainTrackingByBillNumberForRussia {
    public readonly feso: IBillNumberTracker;
    public readonly sklu: IBillNumberTracker;
    public readonly halu: IBillNumberTracker;
    public readonly scacStruct: BillNumberTrackingInRussia;
    protected timeInspector: TimeInspector;

    public constructor(trackers: TrackingByBillNumberConstructorArgs) {
        this.feso = trackers.feso;
        this.sklu = trackers.sklu;
        this.halu = trackers.halu
        this.scacStruct = {
            FESO: this.feso,
            SKLU: this.sklu,
            HALU: this.halu,
        };
        this.timeInspector = new TimeInspector();
    }

    protected getContainerByScac(scac: SCAC_TYPE): IBillNumberTracker {
        return this.scacStruct[scac]
    }

    public async trackByBillNumber(args: TrackingArgsWithScac): Promise<ITrackingByBillNumberResponse> {
        let tasks: IBillNumberTracker[] = Object.values(this.scacStruct)
        if (args.scac === "AUTO") {
            for (let task of tasks) {
                try {
                    let res = await task.trackByBillNumber(args);
                    // console.log(res)
                    if (res !== undefined) {
                        if (this.timeInspector.inspectTime(res)) {
                            return res
                        }
                    }
                } catch (e) {

                }

            }
            throw new ContainerNotFoundException()
        } else {
            try {
                return await this.getContainerByScac(args.scac).trackByBillNumber(args);
            } catch (e) {
                console.log(e)
                // throw new ContainerNotFoundException()
            }
        }
    }
}