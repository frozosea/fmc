import {SitcContainer, SitcInfoAboutMovingParser} from "../../src/trackTrace/TrackingByContainerNumber/sitc/sitc";
import {config} from "../classesConfigurator";
import {sitcExpectedResult} from "./sitcExpectedResult";
import {fetchArgs, IRequest} from "../../src/trackTrace/helpers/requestSender";

const assert = require("assert");


const requestMoch: IRequest<fetchArgs> = {
    async sendRequestAndGetJson(args: fetchArgs): Promise<any> {
        return sitcExpectedResult[args.url.split(`=`)[1]]
    },
    async sendRequestAndGetHtml(_: fetchArgs): Promise<string> {
        return ""
    }
}
describe("SITC Test", () => {
    const containers = ["SITU9130070", "UETU5790574"]
    it("SITC info about moving test", () => {
        let sitcInfoAboutMoving = new SitcInfoAboutMovingParser(config.DATETIME)
        for (let container of containers) {
            const result = sitcInfoAboutMoving.getInfoAboutMoving(sitcExpectedResult[container])
            for (let item in result) {
                assert.strictEqual(result[item].operationName, sitcExpectedResult[container].data.list[item].movementNameEn);
                assert.strictEqual(result[item].time, config.DATETIME.strptime(sitcExpectedResult[container].data.list[item].eventDate, "YYYY-MM-DD HH:mm:ss").getTime());
                assert.strictEqual(result[item].vessel, sitcExpectedResult[container].data.list[item].vesselCode);
                assert.strictEqual(result[item].location, sitcExpectedResult[container].data.list[item].eventPort);
            }
        }
    })
    it("SITC main class with moch test", async () => {
        const sitc = new SitcContainer({
            datetime: config.DATETIME,
            requestSender: requestMoch,
            UserAgentGenerator: config.USER_AGENT_GENERATOR
        })
        for (let container of containers) {
            const result = await sitc.trackContainer({container: container})
            for (let item = 0; item < result.infoAboutMoving.length; item++) {
                assert.strictEqual(result.infoAboutMoving[item].operationName, sitcExpectedResult[container].data.list[item].movementNameEn);
                assert.strictEqual(result.infoAboutMoving[item].time, config.DATETIME.strptime(sitcExpectedResult[container].data.list[item].eventDate, "YYYY-MM-DD HH:mm:ss").getTime());
                assert.strictEqual(result.infoAboutMoving[item].vessel, sitcExpectedResult[container].data.list[item].vesselCode);
                assert.strictEqual(result.infoAboutMoving[item].location, sitcExpectedResult[container].data.list[item].eventPort);
            }
        }
    })
})