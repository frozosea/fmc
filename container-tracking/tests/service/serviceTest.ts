import {ICache} from "../../src/cache";
import {trackingForOtherWorld, trackingForRussia} from "../tracking/trackingTest";
import TrackingController from "../../src/trackingController";
import {IScacContainers} from "../../src/trackTrace/TrackingByContainerNumber/containerScacRepo";
import {COUNTRY_TYPE, SCAC_TYPE, TrackingContainerResponse} from "../../src/types";
import {IServiceLogger} from "../../src/logging";

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
    async addContainer(container: string, scac: SCAC_TYPE): Promise<void> {

    },
    async getScac(container: string): Promise<SCAC_TYPE | null> {
        return "SKLU"
    }
}

const loggerMoch: IServiceLogger = {
    containerSuccessLog(result: TrackingContainerResponse) {
    },
    containerNotFoundLog(container: string) {
    }
}

async function testNotAutoScac(service: TrackingController, container: string, scac: SCAC_TYPE, country: COUNTRY_TYPE): Promise<void> {
    let result = await service.trackContainer({container: container, scac: scac, country: country})
    assert.strictEqual(result.scac, scac)
    assert.strictEqual(result.container, container)
}

describe("service test with moch", () => {
    let service = new TrackingController(trackingForRussia, trackingForOtherWorld, scacTableMoch, cacheMoch, loggerMoch)
    const fesoContainer = "FESU2219270"
    it("test return from cache", async () => {
        const testTrackContainer = (result) => {
            const mochRes = moch(fesoContainer)
            assert.strictEqual(result.container, fesoContainer)
            assert.strictEqual(result.scac, mochRes.scac)
            assert.strictEqual(result.containerSize, mochRes.containerSize)
            assert.deepEqual(result.infoAboutMoving, mochRes.infoAboutMoving)
        }
        let ruResult = await service.trackContainer({container: fesoContainer, scac: "AUTO", country: "RU"})
        testTrackContainer(ruResult)
        let otherWorldResult = await service.trackContainer({container: fesoContainer, scac: "AUTO", country: "OTHER"})
        testTrackContainer(otherWorldResult)
    })
    it("test get scac (auto scac) from db", async () => {
        const skluContainer = "SKLU1623413"
        await testNotAutoScac(service, skluContainer, "SKLU", "RU")
    }).timeout(100000)
    it("test sitc not auto scac", async () => {
        const sitcContainer = "SITU9130070"
        await testNotAutoScac(service, sitcContainer, "SITC", "RU")
    }).timeout(100000)
    it("test feso not auto scac", async () => {
        await testNotAutoScac(service, fesoContainer, "FESO", "RU")
    }).timeout(100000)
    it("test maeu not auto scac", async () => {
        await testNotAutoScac(service, "MSKU6874333", "MAEU", "OTHER")
    }).timeout(100000)
    it("test mscu not auto scac", async () => {
        await testNotAutoScac(service, "MEDU3170580", "MSCU", "OTHER")
    }).timeout(100000)
    it("test oney not auto scac", async () => {
        await testNotAutoScac(service, "GAOU6642924", "ONEY", "OTHER")
    }).timeout(100000)
    it("test kmtu not auto scac", async () => {
        await testNotAutoScac(service, "TRHU3368865", "KMTU", "OTHER")
    }).timeout(100000)
    it("should be exception test", async () => {
        try {
            await service.trackContainer({container: "this container does not exist", scac: "AUTO", country: "RU"})
        } catch (e) {
            console.log("TEST SUCCESS")
        }
    })
})