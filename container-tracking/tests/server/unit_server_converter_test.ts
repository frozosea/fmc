import {TrackingServiceConverter} from "../../src/server/services/trackingByContainerNumberService";
import {Country, Scac} from "../../src/server/proto/server_pb";

const assert = require("assert")
describe("converters test", () => {
    const FesoScac = Scac.FESO
    const SkluScac = Scac.SKLU
    const SitcScac = Scac.SITC
    const MaeuScac = Scac.MAEU
    const MscuScac = Scac.MSCU
    const CosuScac = Scac.COSU
    const OneyScac = Scac.ONEY
    it("test scac enum into scac type converter", () => {
        assert.strictEqual(TrackingServiceConverter.convertEnumIntoScacType(FesoScac), "FESO")
        assert.strictEqual(TrackingServiceConverter.convertEnumIntoScacType(SkluScac), "SKLU")
        assert.strictEqual(TrackingServiceConverter.convertEnumIntoScacType(MaeuScac), "MAEU")
        assert.strictEqual(TrackingServiceConverter.convertEnumIntoScacType(MscuScac), "MSCU")
        assert.strictEqual(TrackingServiceConverter.convertEnumIntoScacType(CosuScac), "COSU")
        assert.strictEqual(TrackingServiceConverter.convertEnumIntoScacType(OneyScac), "ONEY")
        assert.strictEqual(TrackingServiceConverter.convertEnumIntoScacType(SitcScac), "SITC")
    })
    it("test country enum to country type converter", () => {
        const ru = Country.RU
        const other = Country.OTHER
        assert.strictEqual(TrackingServiceConverter.convertEnumCountryIntoCountryType(ru), "RU")
        assert.strictEqual(TrackingServiceConverter.convertEnumCountryIntoCountryType(other), "OTHER")
    })
    it("test scac type into enum scac", () => {
        assert.strictEqual(TrackingServiceConverter.convertScacIntoEnum("FESO"), FesoScac)
        assert.strictEqual(TrackingServiceConverter.convertScacIntoEnum("SKLU"), SkluScac)
        assert.strictEqual(TrackingServiceConverter.convertScacIntoEnum("MAEU"), MaeuScac)
        assert.strictEqual(TrackingServiceConverter.convertScacIntoEnum("MSCU"), MscuScac)
        assert.strictEqual(TrackingServiceConverter.convertScacIntoEnum("COSU"), CosuScac)
        assert.strictEqual(TrackingServiceConverter.convertScacIntoEnum("ONEY"), OneyScac)
        assert.strictEqual(TrackingServiceConverter.convertScacIntoEnum("SITC"), SitcScac)
    })
    it("test country type into enum converter",()=>{
        assert.strictEqual(TrackingServiceConverter.convertCountryTypeIntoEnum("RU"), Country.RU)
        assert.strictEqual(TrackingServiceConverter.convertCountryTypeIntoEnum("OTHER"), Country.OTHER)
    })
})