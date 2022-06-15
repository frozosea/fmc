import {
    OneyContainer,
    OneyContainerSizeParser,
    OneyInfoAboutMovingParser
} from "../../src/trackTrace/TrackingByContainerNumber/oney/oney";
import {config} from "../classesConfigurator";
import {OneyContainerSizeExampleResponse, oneyExpectedData, oneyInfoAboutMovingExample} from "./oneyExpectedData";
import {fetchArgs, IRequest} from "../../src/trackTrace/helpers/requestSender";

const assert = require("assert");
const requestMoch: IRequest<fetchArgs> = {
    async sendRequestAndGetJson(args: fetchArgs): Promise<any> {
        return oneyInfoAboutMovingExample
    },
    async sendRequestAndGetHtml(args: fetchArgs): Promise<string> {
        return ""
    }
}
describe("ONEY tracking by container number test", () => {
    it("ONEY info about moving test", () => {
        let infoAboutMovingParser = new OneyInfoAboutMovingParser(config.DATETIME)
        let actualInfoAboutMoving = infoAboutMovingParser.parseInfoAboutMoving(oneyInfoAboutMovingExample)
        assert.deepEqual(actualInfoAboutMoving, oneyExpectedData.infoAboutMoving)
    })
    it("ONEY container size parser test", () => {
        const oneyContainerSizeParser = new OneyContainerSizeParser()
        assert.strictEqual(oneyContainerSizeParser.getContainerSize(OneyContainerSizeExampleResponse), oneyExpectedData.containerSize)
    })
    it("ONEY main class with moch test", async () => {
        const container = "GAOU6642924"
        let oney = new OneyContainer({
            datetime: config.DATETIME,
            requestSender: requestMoch,
            UserAgentGenerator: config.USER_AGENT_GENERATOR
        })
        let actualResponse = await oney.trackContainer({container: container})
        assert.strictEqual(actualResponse.container, container)
        assert.strictEqual(actualResponse.scac, "ONEY")
        assert.deepEqual(actualResponse.infoAboutMoving, oneyExpectedData.infoAboutMoving)
    })
})