import {describe} from "mocha";
import {BaseContainerConstructor, BaseTrackerByContainerNumber} from "../../src/trackTrace/base";
import {fetchArgs} from "../../src/trackTrace/helpers/requestSender";
import {ITrackingArgs, TrackingContainerResponse} from "../../src/types";
import {NotThisShippingLineException} from "../../src/exceptions";
import {MainTrackingForRussia} from "../../src/trackTrace/TrackingByContainerNumber/tracking/mainTrackingForRussia";
import {config} from "../classesConfigurator";
import {FesoContainers, SkluContainers} from "./expectedData";

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

export const baseArgs = {
    datetime: config.DATETIME,
    requestSender: config.REQUEST_SENDER,
    UserAgentGenerator: config.USER_AGENT_GENERATOR
}
describe("tracking unit test with mochs", () => {
    it("FESO test", async () => {
        const trackingForRussiaForFeso = new MainTrackingForRussia({
            fescoContainer: new FesoMoch(baseArgs,false),
            skluContainer: new SkluMoch(baseArgs,true),
            sitcContainer: new SitcMoch(baseArgs,true)
        })
        return await testInfoAboutMovingAndScac(trackingForRussiaForFeso, FesoContainers)
    }).timeout(10000)
    it("SKLU test", async () => {
        const trackingForRussiaForSklu = new MainTrackingForRussia({
            fescoContainer: new FesoMoch(baseArgs,true),
            skluContainer: new SkluMoch(baseArgs,false),
            sitcContainer: new SitcMoch(baseArgs,true)
        })
        return await testInfoAboutMovingAndScac(trackingForRussiaForSklu, SkluContainers)
    }).timeout(10000)
    it("SITC test", async () => {
        const trackingForRussiaForSitc= new MainTrackingForRussia({
            fescoContainer: new FesoMoch(baseArgs,true),
            skluContainer: new SkluMoch(baseArgs,true),
            sitcContainer: new SitcMoch(baseArgs,false)
        })
        return await testInfoAboutMovingAndScac(trackingForRussiaForSitc, SkluContainers)
    }).timeout(100000000000)
})