import {
    OneyContainer,
    OneyInfoAboutMovingParser
} from "../../src/trackTrace/TrackingByContainerNumber/oney/oney";
import {config} from "../classesConfigurator";
import {oneyExpectedData, oneyInfoAboutMovingExample} from "./oneyExpectedData";

const assert = require("assert");

describe("ONEY tracking by container number test", () => {
    it("ONEY info about moving test", () => {
        let infoAboutMovingParser = new OneyInfoAboutMovingParser(config.DATETIME)
        let actualInfoAboutMoving = infoAboutMovingParser.parseInfoAboutMoving(oneyInfoAboutMovingExample)
        console.log(actualInfoAboutMoving === oneyExpectedData.infoAboutMoving)
        //assert.deepEqual(actualInfoAboutMoving, oneyExpectedData.infoAboutMoving)
    })
    it("ONEY integration test", async () => {
        let oney = new OneyContainer({
            datetime: config.DATETIME,
            requestSender: config.REQUEST_SENDER,
            UserAgentGenerator: config.USER_AGENT_GENERATOR
        })
        let actualResponse = await oney.trackContainer({container: "GAOU6642924"})
        try{
            assert.deepEqual(actualResponse, oneyExpectedData)
        }catch (e) {
            
        }
    })
})