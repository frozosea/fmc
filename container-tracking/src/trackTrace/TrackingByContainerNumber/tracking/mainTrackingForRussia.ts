import {BaseTrackerByContainerNumber, TrackingContainerResponse, TrackingArgsWithScac} from "../../base";
import {ContainerNotFoundException} from "../../../exceptions";
import {fetchArgs} from "../../helpers/requestSender";
import {SCAC_TYPE} from "../../../types";

export interface ContainersInTrackingForRussia {
    fescoContainer: BaseTrackerByContainerNumber<fetchArgs>;
    sitcContainer: BaseTrackerByContainerNumber<fetchArgs>;
    skluContainer: BaseTrackerByContainerNumber<fetchArgs>;
}

export interface ScacStructForRussia {
    FESO: BaseTrackerByContainerNumber<fetchArgs>;
    SITC: BaseTrackerByContainerNumber<fetchArgs>;
    SKLU: BaseTrackerByContainerNumber<fetchArgs>;
}

export class TimeInspector {
    protected getDifferenceBetweenDatesInMonths(end: Date, start: Date): number {
        var timeDiff = Math.abs(end.getTime() - start.getTime());
        return Math.round(timeDiff / (2e3 * 3600 * 365.25));
    }

    public inspectTime(response: TrackingContainerResponse): boolean {
        const lastDateInInfoAboutMoving = new Date(response.infoAboutMoving[response.infoAboutMoving.length - 1].time);
        const today = new Date();
        let diffBetweenDates = this.getDifferenceBetweenDatesInMonths(today, lastDateInInfoAboutMoving);
        return diffBetweenDates <= 3;
    }
}

export class MainTrackingForRussia {
    public readonly fescoContainer: BaseTrackerByContainerNumber<fetchArgs>;
    public readonly sitcContainer: BaseTrackerByContainerNumber<fetchArgs>;
    public readonly skluContainer: BaseTrackerByContainerNumber<fetchArgs>;
    public readonly scacStruct: ScacStructForRussia;
    protected timeInspector: TimeInspector;

    public constructor(containers: ContainersInTrackingForRussia) {
        this.fescoContainer = containers.fescoContainer;
        this.skluContainer = containers.skluContainer;
        this.sitcContainer = containers.sitcContainer;
        this.scacStruct = {
            FESO: this.fescoContainer,
            SITC: this.sitcContainer,
            SKLU: this.skluContainer
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
                    let res = await task.trackContainer(args);
                    if (res !== undefined) {
                        if (this.timeInspector.inspectTime(res)) {
                            return res
                        }
                    }
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