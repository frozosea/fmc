import {BaseTrackerByContainerNumber, TrackingArgsWithScac, TrackingContainerResponse} from "../../base";
import {ContainerNotFoundException} from "../../../exceptions";
import {fetchArgs} from "../../helpers/requestSender";
import {ITrackingByBillNumberResponse, SCAC_TYPE} from "../../../types";

export interface ContainersInTrackingForRussia {
    feso: BaseTrackerByContainerNumber<fetchArgs>;
    sitc: BaseTrackerByContainerNumber<fetchArgs>;
    sklu: BaseTrackerByContainerNumber<fetchArgs>;
    halu: BaseTrackerByContainerNumber<fetchArgs>;
}

export interface ScacStructForRussia {
    FESO: BaseTrackerByContainerNumber<fetchArgs>;
    SITC: BaseTrackerByContainerNumber<fetchArgs>;
    SKLU: BaseTrackerByContainerNumber<fetchArgs>;
    HALU: BaseTrackerByContainerNumber<fetchArgs>;
}

export class TimeInspector {
    protected getDifferenceBetweenDatesInMonths(end: Date, start: Date): number {
        var timeDiff = Math.abs(end.getTime() - start.getTime());
        return Math.round(timeDiff / (2e3 * 3600 * 365.25));
    }

    public inspectTime(response: TrackingContainerResponse | ITrackingByBillNumberResponse): boolean {
        const lastDateInInfoAboutMoving = new Date(response.infoAboutMoving[response.infoAboutMoving.length - 1].time);
        const today = new Date();
        let diffBetweenDates = this.getDifferenceBetweenDatesInMonths(today, lastDateInInfoAboutMoving);
        return diffBetweenDates <= 3;
    }
}

export class MainTrackingForRussia {
    public readonly feso: BaseTrackerByContainerNumber<fetchArgs>;
    public readonly sitc: BaseTrackerByContainerNumber<fetchArgs>;
    public readonly sklu: BaseTrackerByContainerNumber<fetchArgs>;
    public readonly halu: BaseTrackerByContainerNumber<fetchArgs>;
    public readonly scacStruct: ScacStructForRussia;
    protected timeInspector: TimeInspector;

    public constructor(containers: ContainersInTrackingForRussia) {
        this.feso = containers.feso;
        this.sklu = containers.sklu;
        this.sitc = containers.sitc;
        this.halu = containers.halu;
        this.scacStruct = {
            FESO: this.feso,
            SITC: this.sitc,
            SKLU: this.sklu,
            HALU: this.halu,
        };
        this.timeInspector = new TimeInspector();
    }

    protected getContainerByScac(scac: SCAC_TYPE): BaseTrackerByContainerNumber<fetchArgs> {
        return this.scacStruct[scac]
    }

    public async trackContainer(args: TrackingArgsWithScac): Promise<TrackingContainerResponse> {
        let tasks: BaseTrackerByContainerNumber<fetchArgs>[] = Object.values(this.scacStruct)
        if (args.scac === "AUTO") {
            for (let task of tasks) {
                try {
                    for (let i = 0; i < 3; i++) {
                        try {
                            let res = await task.trackContainer(args);
                            if (res) {
                                if (this.timeInspector.inspectTime(res)) {
                                    return res
                                }
                            }
                        } catch (e) {
                            continue
                        }
                    }
                    throw new Error("");

                } catch (e) {
                    continue
                }
            }
            throw new ContainerNotFoundException()
        } else {
            try {
                return await this.getContainerByScac(args.scac).trackContainer(args);
            } catch (e) {
                throw new ContainerNotFoundException()
            }
        }
    }
}