import {
    KmtuContainer,
    KmtuDataForInfoAboutMovingRequestCrawler, KmtuEtaParser, KmtuInfoAboutMovingParser
} from "../../src/trackTrace/TrackingByContainerNumber/kmtu/kmtu";
import {config} from "../classesConfigurator";
const path = require("path")
const fs = require("fs")


const assert = require("assert");
describe("KMTU container tracking Test", () => {
    it("KMTU integration test", () => {
        return (async () => {
            let kmtu = new KmtuContainer({
                datetime: config.DATETIME,
                requestSender: config.REQUEST_SENDER,
                UserAgentGenerator: config.USER_AGENT_GENERATOR
            })
            let firstContainer = await kmtu.trackContainer({container: "KMTU7381545"})
            let secondContainer = await kmtu.trackContainer({container: "TRHU3368865"})
            assert.strictEqual(firstContainer.infoAboutMoving[firstContainer.infoAboutMoving.length - 1].operationName, "ETA")
            if (!secondContainer.infoAboutMoving.length) {
                throw new assert.AssertionError("not len of KMTU")
            }
            assert.strictEqual(secondContainer.infoAboutMoving[secondContainer.infoAboutMoving.length - 1].operationName, "ETA")
        })();

    })
    it("KMTU next request data parser test", (done) => {
        const hid_bl_no = "JKT4165031"
        const bk_no = "ID00784886"
        const pod = "HONG KONG"
        const pol = "JAKARTA"
        let kmtuNextRequestDataParser = new KmtuDataForInfoAboutMovingRequestCrawler(config.DATETIME)
        let data = fs.readFileSync(path.resolve(__dirname, './kmtuEtaHtmlExample.html'))
        let requestData = kmtuNextRequestDataParser.getDataForInfoAboutMovingRequest(data.toString("utf-8"))
        assert.strictEqual(requestData.pod, pod)
        assert.strictEqual(requestData.pol, pol)
        assert.strictEqual(requestData.hid_bl_no, hid_bl_no)
        assert.strictEqual(requestData.bk_no, bk_no)
        done()

    })
    it("KMTU eta parser test", () => {
        let etaParser = new KmtuEtaParser(config.DATETIME)
        let data = fs.readFileSync(path.resolve(__dirname, './kmtuEtaHtmlExample.html'))
        let eta = etaParser.parseEta(data.toString("utf-8"))
        assert.strictEqual(eta.operationName, "ETA")
        assert.strictEqual(eta.time, 1655652600000)
        assert.strictEqual(eta.location, "HONG KONG")
        assert.strictEqual(eta.vessel, "KMTC SEOUL/2204N")
    })
})