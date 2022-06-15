import {
    KmtuDataForInfoAboutMovingRequestCrawler,
    KmtuEtaParser,
    KmtuInfoAboutMovingParser
} from "../../src/trackTrace/TrackingByContainerNumber/kmtu/kmtu";
import {config} from "../classesConfigurator";

const path = require("path")
const fs = require("fs")


const assert = require("assert");
describe("KMTU container tracking Test", () => {
    it("KMTU info about moving parser test", () => {
        const buffetData = fs.readFileSync(path.resolve(__dirname, './kmtuInfoAboutMovingExample.html'))
        const infoAboutMovingParser = new KmtuInfoAboutMovingParser(config.DATETIME)
        const actualInfoAboutMoving = infoAboutMovingParser.getInfoAboutMoving(buffetData.toString("utf-8"))
        assert.ok(actualInfoAboutMoving.length > 1)
        assert.strictEqual(actualInfoAboutMoving.length,3)
        const expectedInfoAboutMoving = [
            {
                time: 1653400800000,
                operationName: 'container was picked',
                location: 'BUSAN,Hutchison Busan Container Terminal',
                vessel: ''
            },
            {
                time: 1653400800000,
                operationName: 'container was arrived',
                location: 'BUSAN,Busan Port Terminal',
                vessel: ''
            },
            {
                time: 1653400800000,
                operationName: 'container is onboard and is scheduled to arrive at transshipment',
                location: 'BUSAN,Busan Port Terminal',
                vessel: ''
            }
        ]
        assert.deepEqual(actualInfoAboutMoving,expectedInfoAboutMoving)
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