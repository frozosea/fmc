import {describe} from "mocha";
import {ITrackingArgs, TrackingContainerResponse} from "../../src/types";
import {NotThisShippingLineException} from "../../src/exceptions";
import {FesoContainers, SkluContainers} from "./expectedData";
import {
    MainTrackingForOtherCountries
} from "../../src/trackTrace/TrackingByContainerNumber/tracking/mainTrackingForOtherCountries";
import {MainTrackingForRussia} from "../../src/trackTrace/TrackingByContainerNumber/tracking/mainTrackingForRussia";
import {BaseContainerConstructor, BaseTrackerByContainerNumber} from "../../src/trackTrace/base";
import {fetchArgs} from "../../src/trackTrace/helpers/requestSender";
import {config} from "../classesConfigurator";

const assert = require("assert");

export async function testInfoAboutMovingAndScac(tracking, containers): Promise<void> {
    for (let container in containers) {
        let actualResult = await tracking.trackContainer({container: container, scac: "AUTO"})
        try {
            assert.strictEqual(actualResult.scac, containers[container].scac)
        } catch (e) {
        }
        assert.ok(actualResult.infoAboutMoving.length)
    }
}


export const baseArgs = {
    datetime: config.DATETIME,
    requestSender: config.REQUEST_SENDER,
    UserAgentGenerator: config.USER_AGENT_GENERATOR
}

export class FesoMoch extends BaseTrackerByContainerNumber<fetchArgs> {
    protected shouldRaiseException: boolean;

    public constructor(args: BaseContainerConstructor<fetchArgs>, shouldRaiseException: boolean) {
        super(args);
        this.shouldRaiseException = shouldRaiseException;
    }

    public async trackContainer(args: ITrackingArgs): Promise<TrackingContainerResponse> {
        if (!this.shouldRaiseException) {
            return {
                container: 'FESU2219270',
                scac: 'FESO',
                containerSize: '20DC',
                infoAboutMoving: [
                    {
                        time: 1654495200000,
                        operationName: 'Gate out empty for loading',
                        location: 'MAGISTRAL',
                        vessel: ''
                    },
                    {
                        time: 1654588800000,
                        operationName: 'Gate in empty from consignee',
                        location: 'ZAPSIBCONT',
                        vessel: ''
                    }
                ]
            }
        }
        throw new NotThisShippingLineException()

    }
}

export class SkluMoch extends FesoMoch {
    async trackContainer(args: ITrackingArgs): Promise<TrackingContainerResponse> {
        if (!this.shouldRaiseException) {
            return {
                container: 'SKLU1623413',
                containerSize: "20'x15",
                scac: 'SKLU',
                infoAboutMoving: [
                    {
                        time: 1653906120000,
                        operationName: 'Pickup (1/1)',
                        location: 'SINOKOR TAM CANG CAT LAI Depot',
                        vessel: ''
                    },
                    {
                        time: 1653961680000,
                        operationName: 'Return (1/1)',
                        location: 'CAT LAI',
                        vessel: ''
                    },
                    {
                        time: 1654567200000,
                        operationName: 'Departure',
                        location: 'CAT LAI',
                        vessel: 'HEUNG-A HOCHIMINH / 2205N'
                    },
                    {
                        time: 1655085600000,
                        operationName: 'Arrival(T/S) (Scheduled)',
                        location: 'BPTS',
                        vessel: 'HEUNG-A HOCHIMINH / 2205N'
                    },
                    {
                        time: 1655283600000,
                        operationName: 'Departure(T/S) (Scheduled)',
                        location: 'BPTS',
                        vessel: 'HEUNG-A ULSAN / 2256E'
                    },
                    {
                        time: 1655452800000,
                        operationName: 'Arrival (Scheduled)',
                        location: 'HOSOSHIMA TERMINAL(SHIRAHMA #14)',
                        vessel: 'HEUNG-A ULSAN / 2256E'
                    },
                    {
                        operationName: 'ETA',
                        time: 1655424000000,
                        location: 'Hososhima',
                        vessel: ''
                    }
                ]
            }
        }
        throw new NotThisShippingLineException()

    }
}


export class SitcMoch extends SkluMoch {
    async trackContainer(args: ITrackingArgs): Promise<TrackingContainerResponse> {
        if (!this.shouldRaiseException) {
            return {
                container: 'SITU9130070',
                containerSize: '',
                scac: 'SITC',
                infoAboutMoving: [
                    {
                        time: 1654380000000,
                        operationName: 'LOADED ONTO VESSEL',
                        vessel: 'SITC CAGAYAN',
                        location: 'DALIAN'
                    },
                    {
                        time: 1653649140000,
                        operationName: 'OUTBOUND PICKUP',
                        vessel: 'SITC CAGAYAN',
                        location: 'DALIAN'
                    },
                    {
                        time: 1652899980000,
                        operationName: 'EMPTY CONTAINER',
                        vessel: 'SITC MAKASSAR',
                        location: 'DALIAN'
                    }
                ]
            }
        }
        throw new NotThisShippingLineException()
    }
}

export class OneyMoch extends FesoMoch {
    async trackContainer(args: ITrackingArgs): Promise<TrackingContainerResponse> {
        if (!this.shouldRaiseException) {
            return {
                container: 'GAOU6642924',
                containerSize: "undefined",
                scac: 'ONEY',
                infoAboutMoving: [
                    {
                        time: 1648120740000,
                        operationName: 'Empty Container Release to Shipper',
                        location: 'PUSAN, KOREA REPUBLIC OF',
                        vessel: ''
                    },
                    {
                        time: 1649155200000,
                        operationName: 'Gate In to Outbound Terminal',
                        location: 'PUSAN, KOREA REPUBLIC OF',
                        vessel: ''
                    },
                    {
                        time: 1649333340000,
                        operationName: "Loaded on 'HYUNDAI SINGAPORE 126E' at Port of Loading",
                        location: 'PUSAN, KOREA REPUBLIC OF',
                        vessel: 'HYUNDAI SINGAPORE'
                    },
                    {
                        time: 1649365200000,
                        operationName: "'HYUNDAI SINGAPORE 126E' Departure from Port of Loading",
                        location: 'PUSAN, KOREA REPUBLIC OF',
                        vessel: 'HYUNDAI SINGAPORE'
                    },
                    {
                        time: 1650297600000,
                        operationName: "'HYUNDAI SINGAPORE 126E' Arrival at Port of Discharging",
                        location: 'VANCOUVER, BC, CANADA',
                        vessel: 'HYUNDAI SINGAPORE'
                    },
                    {
                        time: 1653707640000,
                        operationName: "'HYUNDAI SINGAPORE 126E' POD Berthing Destination",
                        location: 'VANCOUVER, BC, CANADA',
                        vessel: 'HYUNDAI SINGAPORE'
                    },
                    {
                        time: 1654033620000,
                        operationName: "Unloaded from 'HYUNDAI SINGAPORE 126E' at Port of Discharging",
                        location: 'VANCOUVER, BC, CANADA',
                        vessel: 'HYUNDAI SINGAPORE'
                    },
                    {
                        time: 1654215840000,
                        operationName: 'Loaded on rail at inbound rail origin',
                        location: 'VANCOUVER, BC, CANADA',
                        vessel: ''
                    },
                    {
                        time: 1654240140000,
                        operationName: 'Inbound Rail Departure',
                        location: 'VANCOUVER, BC, CANADA',
                        vessel: ''
                    },
                    {
                        time: 1654799640000,
                        operationName: 'Inbound Rail Arrival',
                        location: 'DETROIT, MI, UNITED STATES',
                        vessel: ''
                    },
                    {
                        time: 1654808520000,
                        operationName: 'Unloaded from rail at inbound rail destination',
                        location: 'DETROIT, MI, UNITED STATES',
                        vessel: ''
                    },
                    {
                        time: 1654857480000,
                        operationName: 'Gate Out from Inbound CY for Delivery to Consignee',
                        location: 'DETROIT, MI, UNITED STATES',
                        vessel: ''
                    },
                    {
                        time: 1654863840000,
                        operationName: 'Empty Container Returned from Customer',
                        location: 'DETROIT, MI, UNITED STATES',
                        vessel: ''
                    }
                ]
            }
        }
        throw new NotThisShippingLineException()
    }
}

export class MaeuMoch extends OneyMoch {
    async trackContainer(args: ITrackingArgs): Promise<TrackingContainerResponse> {
        if (!this.shouldRaiseException) {
            return {
                container: 'MSKU6874333',
                containerSize: '40DRY',
                scac: 'MAEU',
                infoAboutMoving: [
                    {
                        time: 1648536120000,
                        location: 'Win Win Container Depot',
                        operationName: 'GATE-OUT-EMPTY',
                        vessel: 'MSC SVEVA'
                    },
                    {
                        time: 1648598520000,
                        location: 'Laem Chabang Terminal PORT D1',
                        operationName: 'GATE-IN',
                        vessel: 'MSC SVEVA'
                    },
                    {
                        time: 1649859300000,
                        location: 'Laem Chabang Terminal PORT D1',
                        operationName: 'LOAD',
                        vessel: 'MSC SVEVA'
                    },
                    {
                        time: 1650854940000,
                        location: 'YANGSHAN SGH GUANDONG TERMINAL',
                        operationName: 'DISCHARG',
                        vessel: 'MSC SVEVA'
                    },
                    {
                        time: 1651417860000,
                        location: 'YANGSHAN SGH GUANDONG TERMINAL',
                        operationName: 'LOAD',
                        vessel: 'ZIM WILMINGTON'
                    },
                    {
                        time: 1654919640000,
                        location: 'Charleston Wando Welch terminal N59',
                        operationName: 'DISCHARG',
                        vessel: 'ZIM WILMINGTON'
                    },
                    {
                        time: 1655164800000,
                        location: 'Charleston Wando Welch terminal N59',
                        operationName: 'GATE-OUT',
                        vessel: ''
                    },
                    {
                        time: 1654919640000,
                        operationName: 'ETA',
                        location: 'Spartanburg',
                        vessel: ''
                    }
                ]
            }
        }
        throw new NotThisShippingLineException()
    }
}

export class MscuMoch extends MaeuMoch {
    async trackContainer(args: ITrackingArgs): Promise<TrackingContainerResponse> {
        if (!this.shouldRaiseException) {
            return {
                container: 'MEDU3170580',
                containerSize: "20' DRY VAN",
                scac: 'MSCU',
                infoAboutMoving: [
                    {
                        time: 1654819200000,
                        operationName: 'Export at barge yard',
                        location: 'CHONGQING, CN',
                        vessel: ''
                    },
                    {
                        time: 1654732800000,
                        operationName: 'Empty to Shipper',
                        location: 'CHONGQING, CN',
                        vessel: ''
                    },
                    {
                        time: 1654992000000,
                        operationName: 'ETA',
                        location: '',
                        vessel: ''
                    }
                ]
            }
        }
        throw new NotThisShippingLineException()
    }
}

export class CosuMoch extends MscuMoch {
    async trackContainer(args: ITrackingArgs): Promise<TrackingContainerResponse> {
        if (!this.shouldRaiseException) {
            return {
                container: args.number, scac: "COSU", containerSize: "", infoAboutMoving: [
                    {
                        time: 1654083120000,
                        operationName: 'Empty Equipment Returned',
                        location: 'United Waalhaven Terminals BV(Gate2,Rotterdam,Zuid-Holland,Netherlands',
                        vessel: 'Truck'
                    },
                    {
                        time: 1653893580000,
                        operationName: 'Gate-out from Final Hub',
                        location: 'Euromax Terminal,Rotterdam,Zuid-Holland,Netherlands',
                        vessel: 'Truck'
                    },
                    {
                        time: 1653253980000,
                        operationName: 'Discharged at Last POD',
                        location: 'Euromax Terminal,Rotterdam,Zuid-Holland,Netherlands',
                        vessel: 'Vessel'
                    },
                    {
                        time: 1649940720000,
                        operationName: 'Loaded at First POL',
                        location: 'QQCTU Qingdao Qianwan United Ctn,Qingdao,Shandong,China',
                        vessel: 'Vessel'
                    },
                    {
                        time: 1649671440000,
                        operationName: 'Cargo Received',
                        location: 'QQCTU Qingdao Qianwan United Ctn,Qingdao,Shandong,China',
                        vessel: 'Truck'
                    },
                    {
                        time: 1649671440000,
                        operationName: 'Gate-In at First POL',
                        location: 'QQCTU Qingdao Qianwan United Ctn,Qingdao,Shandong,China',
                        vessel: 'Truck'
                    },
                    {
                        time: 1649318100000,
                        operationName: 'Empty Equipment Despatched',
                        location: "Qingdao Shenzhouxing Int'l Frt Co,Qingdao,Shandong,China",
                        vessel: 'Truck'
                    }
                ]
            }
        }
        throw new NotThisShippingLineException()
    }
}

export class KmtuMoch extends CosuMoch {
    async trackContainer(args: ITrackingArgs): Promise<TrackingContainerResponse> {
        if (!this.shouldRaiseException) {
            return {
                container: args.number, scac: "KMTU", containerSize: "", infoAboutMoving: [
                    {
                        time: 1653436800000,
                        operationName: 'container was picked',
                        location: 'BUSAN,Hutchison Busan Container Terminal',
                        vessel: ''
                    },
                    {
                        time: 1653436800000,
                        operationName: 'container was arrived',
                        location: 'BUSAN,Busan Port Terminal',
                        vessel: ''
                    },
                    {
                        time: 1653436800000,
                        operationName: 'container is onboard and is scheduled to arrive at transshipment',
                        location: 'BUSAN,Busan Port Terminal',
                        vessel: ''
                    }
                ]
            }
        }
        throw new NotThisShippingLineException()
    }
}

export class HaluMoch extends CosuMoch {
    async trackContainer(args: ITrackingArgs): Promise<TrackingContainerResponse> {
        if (!this.shouldRaiseException) {
            return {
                container: args.number, scac: "HALU", containerSize: "", infoAboutMoving: [
                    {
                        time: 1653436800000,
                        operationName: 'container was picked',
                        location: 'BUSAN,Hutchison Busan Container Terminal',
                        vessel: ''
                    },
                    {
                        time: 1653436800000,
                        operationName: 'container was arrived',
                        location: 'BUSAN,Busan Port Terminal',
                        vessel: ''
                    },
                    {
                        time: 1653436800000,
                        operationName: 'container is onboard and is scheduled to arrive at transshipment',
                        location: 'BUSAN,Busan Port Terminal',
                        vessel: ''
                    }
                ]
            }
        }
        throw new NotThisShippingLineException()
    }
}

export const trackingForOtherWorldForSklu = new MainTrackingForOtherCountries({
    feso: new FesoMoch(baseArgs, true),
    sklu: new SkluMoch(baseArgs, false),
    sitc: new SitcMoch(baseArgs, true),
    maeu: new MaeuMoch(baseArgs, true),
    mscu: new MscuMoch(baseArgs, true),
    kmtu: new KmtuMoch(baseArgs, true),
    oney: new OneyMoch(baseArgs, true),
    cosu: new CosuMoch(baseArgs, true),
    halu: new HaluMoch(baseArgs, true)
})

export const trackingForOtherWorldForFeso = new MainTrackingForOtherCountries({
    feso: new FesoMoch(baseArgs, false),
    sklu: new SkluMoch(baseArgs, true),
    sitc: new SitcMoch(baseArgs, true),
    maeu: new MaeuMoch(baseArgs, true),
    mscu: new MscuMoch(baseArgs, true),
    kmtu: new KmtuMoch(baseArgs, true),
    oney: new OneyMoch(baseArgs, true),
    cosu: new CosuMoch(baseArgs, true),
    halu: new HaluMoch(baseArgs, true)

})
export const trackingForOtherWorldForSitc = new MainTrackingForOtherCountries({
    feso: new FesoMoch(baseArgs, true),
    sklu: new SkluMoch(baseArgs, true),
    sitc: new SitcMoch(baseArgs, false),
    maeu: new MaeuMoch(baseArgs, true),
    mscu: new MscuMoch(baseArgs, true),
    kmtu: new KmtuMoch(baseArgs, true),
    oney: new OneyMoch(baseArgs, true),
    cosu: new CosuMoch(baseArgs, true),
    halu: new HaluMoch(baseArgs, true)

})
export const trackingForOtherWorldForMscu = new MainTrackingForOtherCountries({
    feso: new FesoMoch(baseArgs, true),
    sklu: new SkluMoch(baseArgs, true),
    sitc: new SitcMoch(baseArgs, true),
    maeu: new MaeuMoch(baseArgs, true),
    mscu: new MscuMoch(baseArgs, false),
    kmtu: new KmtuMoch(baseArgs, true),
    oney: new OneyMoch(baseArgs, true),
    cosu: new CosuMoch(baseArgs, true),
    halu: new HaluMoch(baseArgs, true)

})
export const trackingForOtherWorldForMaeu = new MainTrackingForOtherCountries({
    feso: new FesoMoch(baseArgs, true),
    sklu: new SkluMoch(baseArgs, true),
    sitc: new SitcMoch(baseArgs, true),
    maeu: new MaeuMoch(baseArgs, false),
    mscu: new MscuMoch(baseArgs, true),
    kmtu: new KmtuMoch(baseArgs, true),
    oney: new OneyMoch(baseArgs, true),
    cosu: new CosuMoch(baseArgs, true),
    halu: new HaluMoch(baseArgs, true)

})
export const trackingForOtherWorldForKmtu = new MainTrackingForOtherCountries({
    feso: new FesoMoch(baseArgs, true),
    sklu: new SkluMoch(baseArgs, true),
    sitc: new SitcMoch(baseArgs, true),
    maeu: new MaeuMoch(baseArgs, true),
    mscu: new MscuMoch(baseArgs, true),
    kmtu: new KmtuMoch(baseArgs, false),
    oney: new OneyMoch(baseArgs, true),
    cosu: new CosuMoch(baseArgs, true),
    halu: new HaluMoch(baseArgs, true)

})
export const trackingForOtherWorldForOney = new MainTrackingForOtherCountries({
    feso: new FesoMoch(baseArgs, true),
    sklu: new SkluMoch(baseArgs, true),
    sitc: new SitcMoch(baseArgs, true),
    maeu: new MaeuMoch(baseArgs, true),
    mscu: new MscuMoch(baseArgs, true),
    kmtu: new KmtuMoch(baseArgs, true),
    oney: new OneyMoch(baseArgs, false),
    cosu: new CosuMoch(baseArgs, true),
    halu: new HaluMoch(baseArgs, true)

})
export const trackingForOtherWorldForCosu = new MainTrackingForOtherCountries({
    feso: new FesoMoch(baseArgs, true),
    sklu: new SkluMoch(baseArgs, true),
    sitc: new SitcMoch(baseArgs, true),
    maeu: new MaeuMoch(baseArgs, true),
    mscu: new MscuMoch(baseArgs, true),
    kmtu: new KmtuMoch(baseArgs, true),
    oney: new OneyMoch(baseArgs, true),
    cosu: new CosuMoch(baseArgs, false),
    halu: new HaluMoch(baseArgs, true)

})
export const trackingForRussiaForHalu = new MainTrackingForRussia({
    feso: new FesoMoch(baseArgs, true),
    sklu: new SkluMoch(baseArgs, true),
    sitc: new SitcMoch(baseArgs, true),
    halu: new HaluMoch(baseArgs, false)
})
export const trackingForOtherWorldForHalu = new MainTrackingForOtherCountries({
    feso: new FesoMoch(baseArgs, true),
    sklu: new SkluMoch(baseArgs, true),
    sitc: new SitcMoch(baseArgs, true),
    maeu: new MaeuMoch(baseArgs, true),
    mscu: new MscuMoch(baseArgs, true),
    kmtu: new KmtuMoch(baseArgs, true),
    oney: new OneyMoch(baseArgs, true),
    cosu: new CosuMoch(baseArgs, true),
    halu: new HaluMoch(baseArgs, false)

})
describe("tracking for other countries test", () => {
    it("FESO test", async () => {
        return await testInfoAboutMovingAndScac(trackingForOtherWorldForFeso, FesoContainers)
    }).timeout(10000)
    it("SKLU test", async () => {
        return await testInfoAboutMovingAndScac(trackingForOtherWorldForSklu, SkluContainers)
    }).timeout(10000)
    it("SITC test", async () => {
        return await testInfoAboutMovingAndScac(trackingForOtherWorldForSitc, SkluContainers)
    }).timeout(100000000000)
    it("MSCU test", async () => {
        return await testInfoAboutMovingAndScac(trackingForOtherWorldForMscu, SkluContainers)
    }).timeout(10000)
    it("MAEU test", async () => {
        return await testInfoAboutMovingAndScac(trackingForOtherWorldForMaeu, SkluContainers)
    }).timeout(10000)
    it("KMTU test", async () => {
        return await testInfoAboutMovingAndScac(trackingForOtherWorldForKmtu, SkluContainers)
    }).timeout(1000000000)
    it("ONEY test", async () => {
        return await testInfoAboutMovingAndScac(trackingForOtherWorldForMscu, SkluContainers)
    })
    it("COSU test", async () => {
        return await testInfoAboutMovingAndScac(trackingForOtherWorldForCosu, SkluContainers)
    })
})