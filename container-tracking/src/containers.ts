import {FesoContainer} from "./trackTrace/TrackingByContainerNumber/feso/feso";
import {SitcContainer} from "./trackTrace/TrackingByContainerNumber/sitc/sitc";
import {SkluContainer} from "./trackTrace/TrackingByContainerNumber/sklu/sklu";
import {OneyContainer} from "./trackTrace/TrackingByContainerNumber/oney/oney";
import {MaeuContainer} from "./trackTrace/TrackingByContainerNumber/maeu/maeu";
import {MscuContainer} from "./trackTrace/TrackingByContainerNumber/mscu/mscu";
import {CosuContainer} from "./trackTrace/TrackingByContainerNumber/cosu/cosu";
import {KmtuContainer} from "./trackTrace/TrackingByContainerNumber/kmtu/kmtu";
import {MainTrackingForRussia} from "./trackTrace/TrackingByContainerNumber/tracking/mainTrackingForRussia";
import {
    MainTrackingForOtherCountries
} from "./trackTrace/TrackingByContainerNumber/tracking/mainTrackingForOtherCountries";
import {UnlocodesRepo} from "./trackTrace/TrackingByContainerNumber/sklu/unlocodesRepo";
import ContainerTrackingController from "./containerTrackingController";
import {ScacRepository} from "./trackTrace/TrackingByContainerNumber/containerScacRepo";
import {Cache} from "./cache";
import {Logger, ServiceLogger} from "./logging";
import {UserAgentGenerator} from "./trackTrace/helpers/userAgentGenerator";
import {RequestSender} from "./trackTrace/helpers/requestSender";
import {Datetime} from "./trackTrace/helpers/datetime";
import {TrackingByContainerNumberService} from "./server/services/trackingByContainerNumberService";
import {HaluContainer} from "./trackTrace/TrackingByContainerNumber/halu/halu";
import {TrackingBybillNumberService} from "./server/services/trackingByBillNumberService";
import BillNumberTrackingController from "./trackingByBillNumberController";
import {FesoBillNumber} from "./trackTrace/trackingBybillNumber/feso/feso";
import {SkluBillNumber} from "./trackTrace/trackingBybillNumber/sklu/sklu";
import {HaluBillNumber} from "./trackTrace/trackingBybillNumber/halu/halu";
import MainTrackingByBillNumberForRussia
    from "./trackTrace/trackingBybillNumber/tracking/mainTrackingByBillNumberForRussia";
import {
    Captcha,
    CaptchaGetter,
    CaptchaSolver,
    RandomStringGenerator
} from "./trackTrace/trackingBybillNumber/sitc/captchaResolver";
import {SitcBillNumber, SitcBillNumberRequest} from "./trackTrace/trackingBybillNumber/sitc/sitc";
import {ZhguBillNumber} from "./trackTrace/trackingBybillNumber/zhgu/zhgu";
import {AppDataSource} from "./db/data-source";

// container.register<>("",{})
const baseArgs = {
    datetime: new Datetime(),
    requestSender: new RequestSender(),
    UserAgentGenerator: new UserAgentGenerator()
}
export const unlocodesRepo = new UnlocodesRepo()
export const feso = new FesoContainer(baseArgs);
export const sitc = new SitcContainer(baseArgs)
export const sklu = new SkluContainer(baseArgs, unlocodesRepo)
export const oney = new OneyContainer(baseArgs)
export const maeu = new MaeuContainer(baseArgs)
export const mscu = new MscuContainer(baseArgs)
export const cosu = new CosuContainer(baseArgs)
export const kmtu = new KmtuContainer(baseArgs)
export const halu = new HaluContainer(baseArgs, unlocodesRepo)

export const fesoBill = new FesoBillNumber(baseArgs)
export const skluBill = new SkluBillNumber(baseArgs, unlocodesRepo)
export const haluBill = new HaluBillNumber(baseArgs, unlocodesRepo)
export const zhguBill = new ZhguBillNumber(baseArgs)
export const randomStringGenerator = new RandomStringGenerator()
export const captchaGetter = new CaptchaGetter(randomStringGenerator, baseArgs.requestSender)
export const captchaSolver = new CaptchaSolver(baseArgs.requestSender)
export const captchaController = new Captcha(randomStringGenerator, captchaGetter, captchaSolver)

export const sitcbill = new SitcBillNumber(baseArgs, captchaController, new SitcBillNumberRequest(baseArgs.requestSender))
export const trackingByContainerNumberForRussia = new MainTrackingForRussia({
    feso: feso,
    sitc: sitc,
    sklu: sklu,
    halu: halu
})
export const trackingByContainerNumberForOtherWorld = new MainTrackingForOtherCountries({
    feso: feso,
    sitc: sitc,
    sklu: sklu,
    halu: halu,
    oney: oney,
    maeu: maeu,
    cosu: cosu,
    mscu: mscu,
    kmtu: kmtu
})
export const mainTrackingByBillNumberForRussia = new MainTrackingByBillNumberForRussia({
    feso: fesoBill,
    sklu: skluBill,
    halu: haluBill,
    sitc: sitcbill,
    zhgu: zhguBill
})
const scacRepo = new ScacRepository()
const cache = new Cache()
const serviceLogger = new ServiceLogger()
export const baseLogger = new Logger()
export const trackingByContainerNumberGrpcService = new ContainerTrackingController(trackingByContainerNumberForRussia, trackingByContainerNumberForOtherWorld, scacRepo, cache, serviceLogger)
export const trackingByContainerNumberService = new TrackingByContainerNumberService(trackingByContainerNumberGrpcService, baseLogger);
export const billNumberTrackingController = new BillNumberTrackingController(mainTrackingByBillNumberForRussia, scacRepo, cache, serviceLogger)
export const trackingByBillNumberService = new TrackingBybillNumberService(billNumberTrackingController, baseLogger)

