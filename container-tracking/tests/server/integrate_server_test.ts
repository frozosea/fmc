import startServer, {server} from "../../src/server/server";
import {trackContainerByServer} from "../../src/server/client";
import {SCAC_TYPE} from "../../src/types";
import {Response} from "../../src/server/proto/server_pb";
import {TrackingServiceConverter} from "../../src/server/services/trackingService";
import {after} from "mocha";

const assert = require("assert")

const testTracking = (result: Response.AsObject, scac: SCAC_TYPE) => {
    try {
        assert.strictEqual(TrackingServiceConverter.convertEnumIntoScacType(result.scac), scac)
    } catch (e) {
        console.log("scac not equal")
    }
    if (!result.infoAboutMovingList.length) {
        throw new assert.AssertionError({message: "not info about moving len"})
    }
}
describe("grpc server test", () => {
    try {
        startServer()
    } catch (e) {
        server.tryShutdown(() => {
        })
        return
    }
    it("test FESO", async () => {
        let fesoResult = await trackContainerByServer("FESU2219270", "FESO", "RU")
        testTracking(fesoResult.toObject(), "FESO")
    }).timeout(10000)
    it("test MAEU", async () => {
        let result = await trackContainerByServer("MSKU6874333", "MAEU", "OTHER")
        testTracking(result.toObject(), "MAEU")

    }).timeout(10000)
    it("test MSCU", async () => {
        let result = await trackContainerByServer("MEDU3170580", "MSCU", "OTHER")
        testTracking(result.toObject(), "MSCU")
    }).timeout(10000)
    it("test SKLU", async () => {
        let result = await trackContainerByServer("SKLU1327134", "SKLU", "RU")
        testTracking(result.toObject(), "SKLU")
    }).timeout(10000)
    it("test SITC", async () => {
        let result = await trackContainerByServer("SITU9130070", "SITC", "RU")
        testTracking(result.toObject(), "SITC")
    }).timeout(10000)
    it("test KMTU", async () => {
        let result = await trackContainerByServer("KMTU7381545", "KMTU", "OTHER")
        testTracking(result.toObject(), "KMTU")
    }).timeout(10000000)

    it("test COSU", async () => {
        let result = await trackContainerByServer("CSNU6829160", "COSU", "OTHER")
        testTracking(result.toObject(), "COSU")
    }).timeout(10000)
    it("test ONEY", async () => {
        let result = await trackContainerByServer("GAOU6642924", "ONEY", "OTHER")
        testTracking(result.toObject(), "ONEY")
    }).timeout(10000)
    it("test AUTO ru", async () => {
        let result = await trackContainerByServer("FESU2219270", "AUTO", "RU")
        testTracking(result.toObject(), "FESO")
    }).timeout(10000)
    it("test AUTO other", async () => {
        let result = await trackContainerByServer("MSKU6874333", "AUTO", "OTHER")
        testTracking(result.toObject(), "MAEU")
    })
    after(() => {
        server.tryShutdown(() => {
            return
        })
        return
    })
})