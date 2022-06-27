import {config} from "../classesConfigurator";
import {expectedInfoAboutMoving} from "./exampleApiResponse";
import {requestMoch} from "./unit_halu_tracking_by_container_number_test";
import {UnlocodesRepoMoch} from "../sklu/skluApiResponseExample";
import {HaluBillNumber} from "../../src/trackTrace/trackingBybillNumber/halu/halu";

const assert = require("assert")

describe("HALU tracking by bill number test", () => {
    let unlocodesRepoMoch = new UnlocodesRepoMoch()
    const billNo = "HASLC01220602499"
    it("HALU main class with moch test", async () => {
        const haluContainer = new HaluBillNumber({
            requestSender: requestMoch,
            datetime: config.DATETIME,
            UserAgentGenerator: config.USER_AGENT_GENERATOR
        }, unlocodesRepoMoch)
        let result = await haluContainer.trackByBillNumber({number: billNo})
        expectedInfoAboutMoving.pop()
        assert.deepEqual(result.infoAboutMoving, expectedInfoAboutMoving)
        assert.strictEqual(result.billNo,billNo)
        assert.strictEqual(result.etaFinalDelivery,1656892800000)
    })
})