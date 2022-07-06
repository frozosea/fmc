import {fetchArgs, IRequest} from "../../src/trackTrace/helpers/requestSender"
import {expectedInfoAboutMoving, HaluExampleApiResponse} from "./exampleApiResponse";
import {HaluContainer} from "../../src/trackTrace/TrackingByContainerNumber/halu/halu";
import {config} from "../classesConfigurator";
import {UnlocodesRepoMoch} from "../sklu/skluApiResponseExample";

const assert = require("assert")
const path = require("path")
const fs = require("fs")
export const requestMoch: IRequest<fetchArgs> = {
    async sendRequestAndGetJson(_: fetchArgs): Promise<any> {
        return HaluExampleApiResponse
    },
    async sendRequestAndGetHtml(_: fetchArgs): Promise<string> {
        return fs.readFileSync(path.resolve(__dirname, './haluExampleInfoAboutMoving.txt')).toString("utf-8")
    }
}

describe("HALU tracking by container number test", () => {
    let unlocodesRepoMoch = new UnlocodesRepoMoch()
    const container = "TKRU3089090"
    it("HALU main class with moch test", async () => {
        const haluContainer = new HaluContainer({
            requestSender: requestMoch,
            datetime: config.DATETIME,
            UserAgentGenerator: config.USER_AGENT_GENERATOR
        }, unlocodesRepoMoch)
        let result = await haluContainer.trackContainer({number: container})
        assert.deepEqual(result.infoAboutMoving, expectedInfoAboutMoving)
        assert.strictEqual(result.containerSize, "20'x4")
    })
})