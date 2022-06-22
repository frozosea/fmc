import {BaseTrackerByContainerNumber} from "../../base";
import {ContainersInTrackingForRussia, ScacStructForRussia, MainTrackingForRussia} from "./mainTrackingForRussia";
import {fetchArgs} from "../../helpers/requestSender";

export interface ContainersInTracking extends ContainersInTrackingForRussia {
    cosuContainer: BaseTrackerByContainerNumber<fetchArgs>;
    maeuContainer: BaseTrackerByContainerNumber<fetchArgs>;
    oneyContainer: BaseTrackerByContainerNumber<fetchArgs>;
    mscuContainer: BaseTrackerByContainerNumber<fetchArgs>;
    kmtuContainer: BaseTrackerByContainerNumber<fetchArgs>;
}

interface ScacStruct extends ScacStructForRussia {
    COSU: BaseTrackerByContainerNumber<fetchArgs>;
    MAEU: BaseTrackerByContainerNumber<fetchArgs>;
    ONEY: BaseTrackerByContainerNumber<fetchArgs>;
    MSCU: BaseTrackerByContainerNumber<fetchArgs>;
    KMTU: BaseTrackerByContainerNumber<fetchArgs>;
}

export class MainTrackingForOtherCountries extends MainTrackingForRussia {
    public readonly cosuContainer: BaseTrackerByContainerNumber<fetchArgs>;
    public readonly maeuContainer: BaseTrackerByContainerNumber<fetchArgs>;
    public readonly oneyContainer: BaseTrackerByContainerNumber<fetchArgs>;
    public readonly mscuContainer: BaseTrackerByContainerNumber<fetchArgs>;
    public readonly kmtuContainer: BaseTrackerByContainerNumber<fetchArgs>;
    public readonly scacStruct: ScacStruct;

    public constructor(containers: ContainersInTracking) {
        super(containers)
        this.cosuContainer = containers.cosuContainer;
        this.maeuContainer = containers.maeuContainer;
        this.oneyContainer = containers.oneyContainer;
        this.mscuContainer = containers.mscuContainer;
        this.kmtuContainer = containers.kmtuContainer;
        this.scacStruct = {
            FESO: this.fescoContainer,
            COSU: this.cosuContainer,
            MAEU: this.maeuContainer,
            ONEY: this.oneyContainer,
            SITC: this.sitcContainer,
            SKLU: this.skluContainer,
            MSCU: this.mscuContainer,
            KMTU: this.kmtuContainer
        }
    }
}