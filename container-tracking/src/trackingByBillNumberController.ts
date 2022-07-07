import {IScacContainers} from "./trackTrace/TrackingByContainerNumber/containerScacRepo";
import {IServiceLogger} from "./logging";
import {ICache} from "./cache";
import {ITrackingByBillNumberResponse, TrackingArgsWithScac} from "./types";
import {CacheHandler} from "./containerTrackingController";
import MainTrackingByBillNumberForRussia
    from "./trackTrace/trackingBybillNumber/tracking/mainTrackingByBillNumberForRussia";


export default class BillNumberTrackingController {
    protected trackingForRussia: MainTrackingByBillNumberForRussia;
    protected scacContainersRepository: IScacContainers;
    protected cacheHandler: CacheHandler;
    protected logger: IServiceLogger;

    public constructor(trackingForRussia: MainTrackingByBillNumberForRussia,
                       scacContainersRepository: IScacContainers,
                       cache: ICache, logger: IServiceLogger) {
        this.trackingForRussia = trackingForRussia;
        this.scacContainersRepository = scacContainersRepository;
        this.cacheHandler = new CacheHandler(cache)
        this.logger = logger
    }

    protected async track(args: TrackingArgsWithScac, ttl?: number): Promise<ITrackingByBillNumberResponse> {
        switch (args.country) {
            case "RU":
                let result = await this.trackingForRussia.trackByBillNumber(args)
                await this.cacheHandler.addTrackingResultToCache(args.number, result, ttl)
                await this.scacContainersRepository.addContainer(result.billNo, result.scac)
                this.logger.containerSuccessLog(result)
                return result
            default:
                throw new Error("use only RU")
        }
    }

    public async trackByBillNumber(args: TrackingArgsWithScac): Promise<ITrackingByBillNumberResponse> {
        let result = await this.cacheHandler.getTrackingInfoFromCache<ITrackingByBillNumberResponse>(args.number)
        if (result !== null) return result
        if (args.scac === "AUTO") {
            let scacFromDb = await this.scacContainersRepository.getScac(args.number)
            if (scacFromDb !== null) {
                try {
                    return await this.track({
                        number: args.number,
                        country: args.country,
                        scac: scacFromDb
                    }, Number(process.env.CONTAINER_TRACKING_RESULT_REDIS_TTL_SECONDS))
                } catch (e) {
                    this.logger.containerNotFoundLog(args.number)
                    return await this.track({
                        number: args.number,
                        country: args.country,
                        scac: "AUTO"
                    }, Number(process.env.CONTAINER_TRACKING_RESULT_REDIS_TTL_SECONDS))
                }
            } else {
                try {
                    return await this.track(args, Number(process.env.CONTAINER_TRACKING_RESULT_REDIS_TTL_SECONDS))
                } catch (e) {
                    this.logger.containerNotFoundLog(args.number)
                }
            }
        } else {
            try {
                return await this.track(args, Number(process.env.CONTAINER_TRACKING_RESULT_REDIS_TTL_SECONDS))
            } catch (e) {
                this.logger.containerNotFoundLog(args.number)
            }
        }

    }
}