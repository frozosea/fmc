import {BaseTrackerByContainerNumber} from "../../base";
import {ContainersInTrackingForRussia, ScacStructForRussia, MainTrackingForRussia} from "./mainTrackingForRussia";
import {fetchArgs} from "../../helpers/requestSender";

export interface ContainersInTracking extends ContainersInTrackingForRussia {
    cosu: BaseTrackerByContainerNumber<fetchArgs>;
    maeu: BaseTrackerByContainerNumber<fetchArgs>;
    oney: BaseTrackerByContainerNumber<fetchArgs>;
    mscu: BaseTrackerByContainerNumber<fetchArgs>;
    kmtu: BaseTrackerByContainerNumber<fetchArgs>;
}

interface ScacStruct extends ScacStructForRussia {
    COSU: BaseTrackerByContainerNumber<fetchArgs>;
    MAEU: BaseTrackerByContainerNumber<fetchArgs>;
    ONEY: BaseTrackerByContainerNumber<fetchArgs>;
    MSCU: BaseTrackerByContainerNumber<fetchArgs>;
    KMTU: BaseTrackerByContainerNumber<fetchArgs>;
    HALU: BaseTrackerByContainerNumber<fetchArgs>;
}

export class MainTrackingForOtherCountries extends MainTrackingForRussia {
    public readonly cosu: BaseTrackerByContainerNumber<fetchArgs>;
    public readonly maeu: BaseTrackerByContainerNumber<fetchArgs>;
    public readonly oney: BaseTrackerByContainerNumber<fetchArgs>;
    public readonly mscu: BaseTrackerByContainerNumber<fetchArgs>;
    public readonly kmtu: BaseTrackerByContainerNumber<fetchArgs>;
    public readonly scacStruct: ScacStruct;

    public constructor(containers: ContainersInTracking) {
        super(containers)
        this.cosu = containers.cosu;
        this.maeu = containers.maeu;
        this.oney = containers.oney;
        this.mscu = containers.mscu;
        this.kmtu = containers.kmtu;
        this.scacStruct = {
            FESO: this.feso,
            COSU: this.cosu,
            MAEU: this.maeu,
            ONEY: this.oney,
            SITC: this.sitc,
            SKLU: this.sklu,
            MSCU: this.mscu,
            KMTU: this.kmtu,
            HALU: this.halu
        }
    }
}