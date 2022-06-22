import {FesoContainer} from "./trackTrace/TrackingByContainerNumber/feso/feso";
import {SitcContainer} from "./trackTrace/TrackingByContainerNumber/sitc/sitc";
import {SkluContainer} from "./trackTrace/TrackingByContainerNumber/sklu/sklu";
import {OneyContainer} from "./trackTrace/TrackingByContainerNumber/oney/oney";
import {MaeuContainer} from "./trackTrace/TrackingByContainerNumber/maeu/maeu";
import {MscuContainer} from "./trackTrace/TrackingByContainerNumber/mscu/mscu";
import {CosuContainer} from "./trackTrace/TrackingByContainerNumber/cosu/cosu";
import {KmtuContainer} from "./trackTrace/TrackingByContainerNumber/kmtu/kmtu";
import {MainTrackingForRussia} from "./trackTrace/TrackingByContainerNumber/tracking/mainTrackingForRussia";
import {MainTrackingForOtherCountries} from "./trackTrace/TrackingByContainerNumber/tracking/mainTrackingForOtherCountries";
import {UnlocodesRepo} from "./trackTrace/TrackingByContainerNumber/sklu/unlocodesRepo";
import TrackingController from "./trackingController";
import {ScacRepository} from "./trackTrace/TrackingByContainerNumber/containerScacRepo";
import {Cache} from "./cache";
import {Logger, ServiceLogger} from "./logging";
import {UserAgentGenerator} from "./trackTrace/helpers/userAgentGenerator";
import {RequestSender} from "./trackTrace/helpers/requestSender";
import {Datetime} from "./trackTrace/helpers/datetime";
import {AppDataSource} from "./db/data-source";
import {TrackingService} from "./server/services/trackingService";

// container.register<>("",{})
const baseArgs = {
    datetime: new Datetime(),
    requestSender: new RequestSender(),
    UserAgentGenerator: new UserAgentGenerator()
}
const feso = new FesoContainer(baseArgs);
const sitc = new SitcContainer(baseArgs)
const sklu = new SkluContainer(baseArgs, new UnlocodesRepo())
const oney = new OneyContainer(baseArgs)
const maeu = new MaeuContainer(baseArgs)
const mscu = new MscuContainer(baseArgs)
const cosu = new CosuContainer(baseArgs)
const kmtu = new KmtuContainer(baseArgs)

export const trackingForRussia = new MainTrackingForRussia({
    fescoContainer: feso,
    sitcContainer: sitc,
    skluContainer: sklu
})
export const trackingForOtherWorld = new MainTrackingForOtherCountries({
    fescoContainer: feso,
    sitcContainer: sitc,
    skluContainer: sklu,
    oneyContainer: oney,
    maeuContainer: maeu,
    cosuContainer: cosu,
    mscuContainer: mscu,
    kmtuContainer: kmtu
})
const scacRepo = new ScacRepository()
const cache = new Cache()
const serviceLogger = new ServiceLogger()
export const baseLogger = new Logger()
export const service = new TrackingController(trackingForRussia, trackingForOtherWorld, scacRepo, cache, serviceLogger)
export const grpcService = new TrackingService(service, baseLogger);

AppDataSource.initialize()
    .then(async (source) => {

    })
    .catch((error) => console.log("Error: ", error))