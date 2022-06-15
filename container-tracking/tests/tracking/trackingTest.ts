import {
    containersForRussiaWithScac,
    FesoContainers, KmtuContainers,
    MaeuContainers,
    MscuContainers, SitcContainers,
    SkluContainers
} from "./expectedData";
import {TrackingForRussia} from "../../src/trackTrace/TrackingByContainerNumber/tracking/trackingForRussia";
import {FesoContainer} from "../../src/trackTrace/TrackingByContainerNumber/feso/feso";
import {config} from "../classesConfigurator";
import {SitcContainer} from "../../src/trackTrace/TrackingByContainerNumber/sitc/sitc";
import {SkluContainer} from "../../src/trackTrace/TrackingByContainerNumber/sklu/sklu";
import {UnlocodeObject} from "../../src/trackTrace/TrackingByContainerNumber/sklu/unlocodesRepo";
import {OneyContainer} from "../../src/trackTrace/TrackingByContainerNumber/oney/oney";
import {MaeuContainer} from "../../src/trackTrace/TrackingByContainerNumber/maeu/maeu";
import {MscuContainer} from "../../src/trackTrace/TrackingByContainerNumber/mscu/mscu";
import {CosuContainer} from "../../src/trackTrace/TrackingByContainerNumber/cosu/cosu";
import {
    TrackingForOtherCountries
} from "../../src/trackTrace/TrackingByContainerNumber/tracking/trackingForOtherCountries";
import {KmtuContainer} from "../../src/trackTrace/TrackingByContainerNumber/kmtu/kmtu";

const assert = require("assert")

async function testInfoAboutMovingAndScac(tracking, containers): Promise<void> {
    for (let container in containers) {
        let actualResult = await tracking.trackContainer({container: container, scac: "AUTO"})
        try {
            assert.strictEqual(actualResult.scac, containers[container])
        } catch (e) {
        }
        if (!actualResult.infoAboutMoving.length) {
            throw new assert.AssertionError({message: "not info about moving len"})
        }
    }


}

const repoMoch = {
    async getUnlocode(unlocode): Promise<string> {
        return ""
    },
    async addUnlocode(obj: UnlocodeObject): Promise<void> {

    }
}
const baseArgs = {
    datetime: config.DATETIME,
    requestSender: config.REQUEST_SENDER,
    UserAgentGenerator: config.USER_AGENT_GENERATOR
}
const feso = new FesoContainer(baseArgs);
const sitc = new SitcContainer(baseArgs)
const sklu = new SkluContainer(baseArgs, repoMoch)
const oney = new OneyContainer(baseArgs)
const maeu = new MaeuContainer(baseArgs)
const mscu = new MscuContainer(baseArgs)
const cosu = new CosuContainer(baseArgs)
const kmtu = new KmtuContainer(baseArgs)

export const trackingForRussia = new TrackingForRussia({
    fescoContainer: feso,
    sitcContainer: sitc,
    skluContainer: sklu
})
export const trackingForOtherWorld = new TrackingForOtherCountries({
    fescoContainer: feso,
    sitcContainer: sitc,
    skluContainer: sklu,
    oneyContainer: oney,
    maeuContainer: maeu,
    cosuContainer: cosu,
    mscuContainer: mscu,
    kmtuContainer: kmtu
})


describe("tracking test (this is class which parse all lines and try get information about moving)", () => {
    it("FESO test", async () => {
        return await testInfoAboutMovingAndScac(trackingForRussia, FesoContainers)
    }).timeout(10000)
    it("SKLU test", async () => {
        return await testInfoAboutMovingAndScac(trackingForRussia, SkluContainers)
    }).timeout(10000)
    it("SITC test", async () => {
        return await testInfoAboutMovingAndScac(trackingForRussia, SitcContainers)
    }).timeout(100000000000)
    it("MSCU test", async () => {
        return await testInfoAboutMovingAndScac(trackingForOtherWorld, MscuContainers)
    }).timeout(10000)
    it("MAEU test", async () => {
        return await testInfoAboutMovingAndScac(trackingForOtherWorld, MaeuContainers)
    }).timeout(10000)
    it("KMTU test", async () => {
        return await testInfoAboutMovingAndScac(trackingForOtherWorld, KmtuContainers)
    }).timeout(1000000000)

})