


import {
    CosuContainerSizeParser,
    CosuEtaParser,
    CosuInfoAboutMovingParser,
    CosuPodParser,
    CosuRequest
} from "../../src/trackTrace/TrackingByContainerNumber/cosu/cosu";
import {
    CosuApiResponseSchema,
    EtaResponseSchema
} from "../../src/trackTrace/TrackingByContainerNumber/cosu/cosuApiResponseSchema";
import {OneTrackingEvent} from "../../src/trackTrace/base";
import {config} from "../classesConfigurator";
import 'mocha';
import {cosuResponse} from "./cosuResponseExample";

const assert = require('assert');


let etaApiResp = {"code": "200", "message": "", "data": {"content": "2022-05-22 23:00"}}

function containerSizeParserTest(containerSizeParser: CosuContainerSizeParser, rawInfoAboutMoving: CosuApiResponseSchema) {
    let actualContainerSize = containerSizeParser.getContainerSize(rawInfoAboutMoving)
    let expectedContainerSize = cosuResponse.data.content.containers[0].container.containerType
    assert.strictEqual(actualContainerSize, expectedContainerSize)
}

function etaParserTest(etaParser: CosuEtaParser, rawEtaResp: EtaResponseSchema, pod: string) {
    let etaObj: OneTrackingEvent = etaParser.getEtaObject(rawEtaResp, pod)
    assert.strictEqual(etaObj.operationName, "ETA")
    assert.strictEqual(etaObj.location, pod)
    let expectedEta = new Date(etaApiResp.data.content).getTime()
    assert.strictEqual(etaObj.time, expectedEta)
}

function podParserTest(podParser: CosuPodParser, rawInfoAboutMoving: CosuApiResponseSchema) {
    let expectedPod = cosuResponse.data.content.containers[0].container.pod
    let actualPod = podParser.getPod(rawInfoAboutMoving)
    assert.notEqual(actualPod, expectedPod)
}

function infoAboutMovingTest(infoAboutMovingParser: CosuInfoAboutMovingParser, rawInfoAboutMoving: CosuApiResponseSchema) {
    let infoAboutMoving: OneTrackingEvent[] = infoAboutMovingParser.getInfoAboutMoving(rawInfoAboutMoving)
    let containerExpectedHistory = cosuResponse.data.content.containers[0].containerHistorys
    assert.strictEqual(infoAboutMoving.length, containerExpectedHistory.length)
    for (let event in infoAboutMoving) {
        let actualEvent = infoAboutMoving[event]
        assert.strictEqual(actualEvent.location, containerExpectedHistory[event].location)
        assert.strictEqual(actualEvent.operationName, containerExpectedHistory[event].containerNumberStatus)
        assert.strictEqual(actualEvent.time, new Date(containerExpectedHistory[event].timeOfIssue).getTime())
        assert.strictEqual(actualEvent.vessel, containerExpectedHistory[event].transportation)
    }
}



describe('COSU container tracking test', function () {
    it("COSU pod parser test", () => {
        let podParser = new CosuPodParser()
        podParserTest(podParser, cosuResponse)
    })
    it("COSU eta parser test", () => {
        let etaParser = new CosuEtaParser(config.DATETIME)
        etaParserTest(etaParser, etaApiResp, cosuResponse.data.content.containers[0].container.pod)
    })
    it("COSU container size parser test", () => {
        let containerSizeParser = new CosuContainerSizeParser()
        containerSizeParserTest(containerSizeParser, cosuResponse)

    })
    it("COSU info about moving test", () => {
        let infoAboutMovingParser = new CosuInfoAboutMovingParser(config.DATETIME)
        infoAboutMovingTest(infoAboutMovingParser, cosuResponse)
    })
});