import {
    MaeuContainer,
    MaeuEtaParser,
    MaeuInfoAboutMovingParser,
    MaeuPortOfDischargingParser
} from "../../src/trackTrace/TrackingByContainerNumber/maeu/maeu";
import {expectedMaeuReadyObject, maeuExamleApiResponse, MaeuinfoAboutMoving} from "./maeuExamleApiResponse";
import {OneTrackingEvent} from "../../src/types";
import {config} from "../classesConfigurator";
import {fetchArgs, IRequest} from "../../src/trackTrace/helpers/requestSender";

const assert = require("assert")

const requestMoch: IRequest<fetchArgs> = {
    async sendRequestAndGetJson(args: fetchArgs): Promise<any> {
        return maeuExamleApiResponse
    },
    async sendRequestAndGetHtml(args: fetchArgs): Promise<string> {
        return ""
    }
}


describe("MAEU Tracking by container number test", () => {
    it("MAEU pod test", () => {
        let maeuPodParser = new MaeuPortOfDischargingParser()
        let actualPod = maeuPodParser.getPortOfDischarging(maeuExamleApiResponse)
        assert.strictEqual(actualPod, maeuExamleApiResponse.destination.city)
    })
    it("MAEU get eta test", () => {
        let maeuEtaParser = new MaeuEtaParser()
        const expectedEtaObject: OneTrackingEvent = {
            time: new Date("2022-06-11T13:54:00.000").getTime(),
            operationName: "ETA",
            location: "Spartanburg",
            vessel: ""
        }
        let acutalEtaObject: OneTrackingEvent = maeuEtaParser.getEta(maeuExamleApiResponse);
        assert.deepEqual(acutalEtaObject, expectedEtaObject)
    })
    it("MAEU info about moving parser test", () => {
        let infoAboutMovingParser = new MaeuInfoAboutMovingParser()
        let actualInfoAboutMoving = infoAboutMovingParser.parseInfoAboutMoving(maeuExamleApiResponse)
        assert.deepEqual(actualInfoAboutMoving, MaeuinfoAboutMoving)
    })
    it("MAEU main class with moch test", () => {
        let maeu = new MaeuContainer({
            datetime: config.DATETIME,
            requestSender: requestMoch,
            UserAgentGenerator: config.USER_AGENT_GENERATOR
        });
        return (async () => {
            let actualResponse = await maeu.trackContainer({container: "MSKU6874333"})
            assert.deepEqual(actualResponse, expectedMaeuReadyObject)
        })()
    })
})