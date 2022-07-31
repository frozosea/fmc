import {TimeInspector} from "../../TrackingByContainerNumber/tracking/mainTrackingForRussia";
import {IBillNumberTracker} from "../base";
import {TrackingArgsWithScac} from "../../base";
import {ContainerNotFoundException} from "../../../exceptions";
import {ITrackingByBillNumberResponse, SCAC_TYPE} from "../../../types";

export interface TrackingByBillNumberConstructorArgs {
    feso: IBillNumberTracker;
    sklu: IBillNumberTracker;
    halu: IBillNumberTracker;
    sitc: IBillNumberTracker;
    zhgu: IBillNumberTracker;
}

export interface BillNumberTrackingInRussia {
    FESO: IBillNumberTracker;
    SKLU: IBillNumberTracker;
    HALU: IBillNumberTracker;
    SITC: IBillNumberTracker;
    ZHGU: IBillNumberTracker;
}

export default class MainTrackingByBillNumberForRussia {
    public readonly feso: IBillNumberTracker;
    public readonly sklu: IBillNumberTracker;
    public readonly halu: IBillNumberTracker;
    public readonly sitc: IBillNumberTracker;
    public readonly zhgu: IBillNumberTracker;
    public readonly scacStruct: BillNumberTrackingInRussia;
    protected timeInspector: TimeInspector;

    public constructor(trackers: TrackingByBillNumberConstructorArgs) {
        this.feso = trackers.feso;
        this.sklu = trackers.sklu;
        this.halu = trackers.halu;
        this.sitc = trackers.sitc;
        this.zhgu = trackers.zhgu
        this.scacStruct = {
            FESO: this.feso,
            SKLU: this.sklu,
            HALU: this.halu,
            SITC: this.sitc,
            ZHGU: this.zhgu
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
                throw new ContainerNotFoundException()
            }
        }
    }
}