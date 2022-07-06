import {
    MscuContainer,
    MscuContainerSizeParser,
    MscuEtaParser,
    MscuInfoAboutMovingParser
} from "../../src/trackTrace/TrackingByContainerNumber/mscu/mscu";
import {config} from "../classesConfigurator";
import {mscuResponseExample} from "./mscuResponseExample";
import {OneTrackingEvent} from "../../src/types";
import {fetchArgs, IRequest} from "../../src/trackTrace/helpers/requestSender";

const assert = require("assert")
function testInfoAboutMoving(actualInfoAboutMoving: OneTrackingEvent[]): void {
    for (let event in mscuResponseExample.Data.BillOfLadings[0].ContainersInfo[0].Events.reverse()) {
        assert.strictEqual(actualInfoAboutMoving[event].time, new Date(new Date(actualInfoAboutMoving[event].time).toUTCString()).getTime())
        assert.strictEqual(actualInfoAboutMoving[event].location, mscuResponseExample.Data.BillOfLadings[0].ContainersInfo[0].Events[event].Location)
        assert.strictEqual(actualInfoAboutMoving[event].operationName, mscuResponseExample.Data.BillOfLadings[0].ContainersInfo[0].Events[event].Description)
    }
}

const requestMoch: IRequest<fetchArgs> = {
    async sendRequestAndGetJson(_: fetchArgs): Promise<any> {
        return mscuResponseExample
    },
    async sendRequestAndGetHtml(_: fetchArgs): Promise<string> {
        return ""
    }
}
describe("MSCU tracking by container number test", () => {
    const expectedContainerSize = "20' DRY VAN"
    it("MSCU main class test", async () => {
        let mscu = new MscuContainer({
            datetime: config.DATETIME,
            requestSender: requestMoch,
            UserAgentGenerator: config.USER_AGENT_GENERATOR
        })
        let actualInfoAboutMoving = await mscu.trackContainer({number: "MEDU3170580"})
        assert.strictEqual(actualInfoAboutMoving.containerSize, expectedContainerSize)
        assert.strictEqual(actualInfoAboutMoving.scac, "MSCU")
        testInfoAboutMoving([actualInfoAboutMoving.infoAboutMoving[0], actualInfoAboutMoving.infoAboutMoving[1]])
        assert.strictEqual(actualInfoAboutMoving.infoAboutMoving[actualInfoAboutMoving.infoAboutMoving.length - 1].operationName, "ETA")

    })
    it("MSCU eta parser test", () => {
        let etaParser = new MscuEtaParser(config.DATETIME)
        const eta = 1654992000000
        assert.strictEqual(etaParser.getEta(mscuResponseExample).time, new Date(new Date(eta).toUTCString()).getTime())
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