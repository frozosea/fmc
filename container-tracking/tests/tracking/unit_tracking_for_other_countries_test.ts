import {describe} from "mocha";
import {baseArgs, FesoMoch, SitcMoch, SkluMoch, testInfoAboutMovingAndScac} from "./unit_tracking_for_russia_test";
import {ITrackingArgs, TrackingContainerResponse} from "../../src/types";
import {NotThisShippingLineException} from "../../src/exceptions";
import {FesoContainers, SkluContainers} from "./expectedData";
import {
    TrackingForOtherCountries
} from "../../src/trackTrace/TrackingByContainerNumber/tracking/trackingForOtherCountries";


export class OneyMoch extends FesoMoch {
    async trackContainer(args: ITrackingArgs): Promise<TrackingContainerResponse> {
        if (!this.shouldRaiseException) {
            return {
                container: 'GAOU6642924',
                containerSize: undefined,
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
                container: args.container, scac: "COSU", containerSize: "", infoAboutMoving: [
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
                container: args.container, scac: "KMTU", containerSize: "", infoAboutMoving: [
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

export const trackingForOtherWorldForSklu = new TrackingForOtherCountries({
    fescoContainer: new FesoMoch(baseArgs, true),
    skluContainer: new SkluMoch(baseArgs, false),
    sitcContainer: new SitcMoch(baseArgs, true),
    maeuContainer: new MaeuMoch(baseArgs, true),
    mscuContainer: new MscuMoch(baseArgs, true),
    kmtuContainer: new KmtuMoch(baseArgs, true),
    oneyContainer: new OneyMoch(baseArgs, true),
    cosuContainer: new CosuMoch(baseArgs, true)
})

export const trackingForOtherWorldForFeso = new TrackingForOtherCountries({
    fescoContainer: new FesoMoch(baseArgs, false),
    skluContainer: new SkluMoch(baseArgs, true),
    sitcContainer: new SitcMoch(baseArgs, true),
    maeuContainer: new MaeuMoch(baseArgs, true),
    mscuContainer: new MscuMoch(baseArgs, true),
    kmtuContainer: new KmtuMoch(baseArgs, true),
    oneyContainer: new OneyMoch(baseArgs, true),
    cosuContainer: new CosuMoch(baseArgs, true)
})
export const trackingForOtherWorldForSitc = new TrackingForOtherCountries({
    fescoContainer: new FesoMoch(baseArgs, true),
    skluContainer: new SkluMoch(baseArgs, true),
    sitcContainer: new SitcMoch(baseArgs, false),
    maeuContainer: new MaeuMoch(baseArgs, true),
    mscuContainer: new MscuMoch(baseArgs, true),
    kmtuContainer: new KmtuMoch(baseArgs, true),
    oneyContainer: new OneyMoch(baseArgs, true),
    cosuContainer: new CosuMoch(baseArgs, true)
})
export const trackingForOtherWorldForMscu = new TrackingForOtherCountries({
    fescoContainer: new FesoMoch(baseArgs, true),
    skluContainer: new SkluMoch(baseArgs, true),
    sitcContainer: new SitcMoch(baseArgs, true),
    maeuContainer: new MaeuMoch(baseArgs, true),
    mscuContainer: new MscuMoch(baseArgs, false),
    kmtuContainer: new KmtuMoch(baseArgs, true),
    oneyContainer: new OneyMoch(baseArgs, true),
    cosuContainer: new CosuMoch(baseArgs, true)
})
export const trackingForOtherWorldForMaeu = new TrackingForOtherCountries({
    fescoContainer: new FesoMoch(baseArgs, true),
    skluContainer: new SkluMoch(baseArgs, true),
    sitcContainer: new SitcMoch(baseArgs, true),
    maeuContainer: new MaeuMoch(baseArgs, false),
    mscuContainer: new MscuMoch(baseArgs, true),
    kmtuContainer: new KmtuMoch(baseArgs, true),
    oneyContainer: new OneyMoch(baseArgs, true),
    cosuContainer: new CosuMoch(baseArgs, true)
})
export const trackingForOtherWorldForKmtu = new TrackingForOtherCountries({
    fescoContainer: new FesoMoch(baseArgs, true),
    skluContainer: new SkluMoch(baseArgs, true),
    sitcContainer: new SitcMoch(baseArgs, true),
    maeuContainer: new MaeuMoch(baseArgs, true),
    mscuContainer: new MscuMoch(baseArgs, true),
    kmtuContainer: new KmtuMoch(baseArgs, false),
    oneyContainer: new OneyMoch(baseArgs, true),
    cosuContainer: new CosuMoch(baseArgs, true)
})
export const trackingForOtherWorldForOney = new TrackingForOtherCountries({
    fescoContainer: new FesoMoch(baseArgs, true),
    skluContainer: new SkluMoch(baseArgs, true),
    sitcContainer: new SitcMoch(baseArgs, true),
    maeuContainer: new MaeuMoch(baseArgs, true),
    mscuContainer: new MscuMoch(baseArgs, true),
    kmtuContainer: new KmtuMoch(baseArgs, true),
    oneyContainer: new OneyMoch(baseArgs, false),
    cosuContainer: new CosuMoch(baseArgs, true)
})
export const trackingForOtherWorldForCosu = new TrackingForOtherCountries({
    fescoContainer: new FesoMoch(baseArgs, true),
    skluContainer: new SkluMoch(baseArgs, true),
    sitcContainer: new SitcMoch(baseArgs, true),
    maeuContainer: new MaeuMoch(baseArgs, true),
    mscuContainer: new MscuMoch(baseArgs, true),
    kmtuContainer: new KmtuMoch(baseArgs, true),
    oneyContainer: new OneyMoch(baseArgs, true),
    cosuContainer: new CosuMoch(baseArgs, false)
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
    it("ONEY test",async()=>{
        return await testInfoAboutMovingAndScac(trackingForOtherWorldForMscu,SkluContainers)
    })
    it("COSU test",async()=>{
        return await testInfoAboutMovingAndScac(trackingForOtherWorldForCosu,SkluContainers)
    })
})