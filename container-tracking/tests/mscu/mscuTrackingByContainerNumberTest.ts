import {
    MscuContainer,
    MscuContainerSizeParser,
    MscuEtaParser, MscuInfoAboutMovingParser
} from "../../src/trackTrace/TrackingByContainerNumber/mscu/mscu";
import {config} from "../classesConfigurator";
import {mscuResponseExample} from "./mscuResponseExample";
import {OneTrackingEvent} from "../../src/types";

const assert = require("assert")

function testInfoAboutMoving(actualInfoAboutMoving: OneTrackingEvent[]): void {
    for (let event in actualInfoAboutMoving) {
        assert.strictEqual(actualInfoAboutMoving[event].location, mscuResponseExample.Data.BillOfLadings[0].ContainersInfo[0].Events[event].Location)
        assert.strictEqual(actualInfoAboutMoving[event].operationName, mscuResponseExample.Data.BillOfLadings[0].ContainersInfo[0].Events[event].Description)
    }
}

describe("MSCU tracking by container number test", () => {
    const expectedContainerSize = "20' DRY VAN"
    it("MSCU integration test", async () => {
        let mscu = new MscuContainer({
            datetime: config.DATETIME,
            requestSender: config.REQUEST_SENDER,
            UserAgentGenerator: config.USER_AGENT_GENERATOR
        })
        let actualInfoAboutMoving = await mscu.trackContainer({container: "MEDU3170580"})
        assert.strictEqual(actualInfoAboutMoving.containerSize, expectedContainerSize)
        assert.strictEqual(actualInfoAboutMoving.scac, "MSCU")
        if (!actualInfoAboutMoving.infoAboutMoving.length) {
            throw new assert.AssertionError()
        }
    })
    it("MSCU eta parser test", () => {
        let etaParser = new MscuEtaParser(config.DATETIME)
        const eta = 1654956000000
        assert.strictEqual(etaParser.getEta(mscuResponseExample).time, eta)
    })
    it("MSCU container size parser test", () => {
        let mscuContainerSizeParser = new MscuContainerSizeParser()
        assert.strictEqual(mscuContainerSizeParser.getContainerSize(mscuResponseExample), expectedContainerSize)
    })
    it("MSCU info about moving test", () => {
        let infoAboutMovingParser = new MscuInfoAboutMovingParser(config.DATETIME)
        let actualInfoAboutMoving = infoAboutMovingParser.getInfoAboutMoving(mscuResponseExample)
        testInfoAboutMoving(actualInfoAboutMoving)
    })
})