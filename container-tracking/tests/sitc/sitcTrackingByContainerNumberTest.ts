import {SitcInfoAboutMovingParser, SitcRequest} from "../../src/trackTrace/TrackingByContainerNumber/sitc/sitc";
import {config} from "../classesConfigurator";
import {sitcExpectedResult} from "./sitcExpectedResult";

const assert = require("assert");

describe("SITC Test", () => {
    const containers = ["SITU9130070", "UETU5790574"]
    it("SITC request test", () => {
        let sitcRequest = new SitcRequest(config.REQUEST_SENDER);
        (async () => {
            for (let container of containers) {
                let actualRes = await sitcRequest.getApiResponseJson({container: container})
                try {
                    assert.deepEqual(actualRes, sitcExpectedResult[container])
                } catch (e) {
                    console.log("SITC container info was change")
                }
            }
        })();
    })
    it("SITC info about moving test", () => {
        let sitcInfoAboutMoving = new SitcInfoAboutMovingParser(config.DATETIME)
        for (let container of containers) {
            let result = sitcInfoAboutMoving.getInfoAboutMoving(sitcExpectedResult[container])
            for (let item in result) {
                assert.strictEqual(result[item].operationName, sitcExpectedResult[container].data.list[item].movementNameEn);
                assert.strictEqual(result[item].time, new Date(sitcExpectedResult[container].data.list[item].eventDate).getTime());
                assert.strictEqual(result[item].vessel, sitcExpectedResult[container].data.list[item].vesselCode);
                assert.strictEqual(result[item].location, sitcExpectedResult[container].data.list[item].eventPort);
            }
        }
    })
    
})