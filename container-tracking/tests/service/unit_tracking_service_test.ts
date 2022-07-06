import {ICache} from "../../src/cache";
import ContainerTrackingController from "../../src/containerTrackingController";
import {IScacContainers} from "../../src/trackTrace/TrackingByContainerNumber/containerScacRepo";
import {COUNTRY_TYPE, SCAC_TYPE, TrackingContainerResponse} from "../../src/types";
import {IServiceLogger} from "../../src/logging";
import {
    trackingForOtherWorldForFeso, trackingForOtherWorldForKmtu,
    trackingForOtherWorldForMaeu,
    trackingForOtherWorldForMscu,
    trackingForOtherWorldForOney,
    trackingForOtherWorldForSitc,
    trackingForOtherWorldForSklu
} from "../tracking/unit_tracking_for_other_countries_test";

const assert = require("assert")
const moch = (key: string) => {
    return {
        container: key,
        scac: "FESO",
        containerSize: "40DHC",
        infoAboutMoving: [{time: 1, operationName: "moch operation", vessel: "", location: "moch loc"}]
    }
}
const cacheMoch: ICache = {
    async get(key: string): Promise<string> {
        if (key === "FESU2219270") return JSON.stringify(moch(key))
        else return null
    },
    async set<T>(key: string, value: T, ttl?: number): Promise<void> {
        return
    }

}
const scacTableMoch: IScacContainers = {
    async addContainer(container: string, _: SCAC_TYPE): Promise<void> {

    },
    async getScac(container: string): Promise<SCAC_TYPE | null> {
        return "SKLU"
    }
}

const loggerMoch: IServiceLogger = {
    containerSuccessLog(_: TrackingContainerResponse) {
    },
    containerNotFoundLog(_: string) {
    }
}

async function testNotAutoScac(service: ContainerTrackingController, container: string, scac: SCAC_TYPE, country: COUNTRY_TYPE): Promise<void> {
    let result = await service.trackContainer({number: container, scac: scac, country: country})
    assert.strictEqual(result.scac, scac)
    assert.strictEqual(result.container, container)
}

describe("service test with moch", () => {
    it("test return from cache", async () => {
        const testTrackContainer = (result) => {
            const mochRes = moch(fesoContainer)
            assert.strictEqual(result.container, fesoContainer)
            assert.strictEqual(result.scac, mochRes.scac)
            assert.strictEqual(result.containerSize, mochRes.containerSize)
            assert.deepEqual(result.infoAboutMoving, mochRes.infoAboutMoving)
        }
        let service = new ContainerTrackingController(trackingForOtherWorldForFeso, trackingForOtherWorldForFeso, scacTableMoch, cacheMoch, loggerMoch)
        const fesoContainer = "FESU2219270"
        let ruResult = await service.trackContainer({number: fesoContainer, scac: "AUTO", country: "RU"})
        testTrackContainer(ruResult)
        let otherWorldResult = await service.trackContainer({number: fesoContainer, scac: "AUTO", country: "OTHER"})
        testTrackContainer(otherWorldResult)
    })
    it("test get scac (auto scac) from db", async () => {
        const skluContainer = "SKLU1623413"
        const service = new ContainerTrackingController(trackingForOtherWorldForSklu, trackingForOtherWorldForSklu, scacTableMoch, cacheMoch, loggerMoch)
        await testNotAutoScac(service, skluContainer, "SKLU", "RU")
    })
    it("test sitc not auto scac", async () => {
        const sitcContainer = "SITU9130070"
        const service = new ContainerTrackingController(trackingForOtherWorldForSitc, trackingForOtherWorldForSitc, scacTableMoch, cacheMoch, loggerMoch)
        await testNotAutoScac(service, sitcContainer, "SITC", "RU")
    })
    it("test feso not auto scac", async () => {
        const service = new ContainerTrackingController(trackingForOtherWorldForFeso, trackingForOtherWorldForFeso, scacTableMoch, cacheMoch, loggerMoch)
        await testNotAutoScac(service, "FESU2219270", "FESO", "RU")
    })
    it("test maeu not auto scac", async () => {
        const service = new ContainerTrackingController(trackingForOtherWorldForMaeu, trackingForOtherWorldForMaeu, scacTableMoch, cacheMoch, loggerMoch)
        await testNotAutoScac(service, "MSKU6874333", "MAEU", "OTHER")
    }).timeout(100000)
    it("test mscu not auto scac", async () => {
        const service = new ContainerTrackingController(trackingForOtherWorldForMscu, trackingForOtherWorldForMscu, scacTableMoch, cacheMoch, loggerMoch)
        await testNotAutoScac(service, "MEDU3170580", "MSCU", "OTHER")
    }).timeout(100000)
    it("test oney not auto scac", async () => {
        const service = new ContainerTrackingController(trackingForOtherWorldForOney, trackingForOtherWorldForOney, scacTableMoch, cacheMoch, loggerMoch)
        await testNotAutoScac(service, "GAOU6642924", "ONEY", "OTHER")
    }).timeout(100000)

    it("test kmtu not auto scac", async () => {
        const service = new ContainerTrackingController(trackingForOtherWorldForKmtu, trackingForOtherWorldForKmtu, scacTableMoch, cacheMoch, loggerMoch)
        await testNotAutoScac(service, "TRHU3368865", "KMTU", "OTHER")
    }).timeout(100000)
})