import {ICache} from "./cache";
import {ITrackingByBillNumberResponse, TrackingArgsWithScac, TrackingContainerResponse} from "./types";
import {MainTrackingForRussia} from "./trackTrace/TrackingByContainerNumber/tracking/mainTrackingForRussia";
import {
    MainTrackingForOtherCountries
} from "./trackTrace/TrackingByContainerNumber/tracking/mainTrackingForOtherCountries";
import {IScacContainers} from "./trackTrace/TrackingByContainerNumber/containerScacRepo";
import {IServiceLogger} from "./logging";

export class CacheHandler {
    protected cache: ICache;

    public constructor(cache: ICache) {
        this.cache = cache

    }

    public async addTrackingResultToCache(container: string, trackingResult: TrackingContainerResponse | ITrackingByBillNumberResponse, ttl?: number): Promise<void> {
        await this.cache.set(container, JSON.stringify(trackingResult), ttl)
    }

    public async getTrackingInfoFromCache<T>(container: string): Promise<T | null> {
        return JSON.parse(await this.cache.get(container))
    }
}

export default class ContainerTrackingController {
    protected trackingForRussia: MainTrackingForRussia;
    protected trackingForOtherCountries: MainTrackingForOtherCountries;
    protected scacContainersRepository: IScacContainers;
    protected cacheHandler: CacheHandler;
    protected logger: IServiceLogger;

    public constructor(trackingForRussia: MainTrackingForRussia,
                       trackingForOtherCountries: MainTrackingForOtherCountries,
                       scacContainersRepository: IScacContainers,
                       cache: ICache, logger: IServiceLogger) {
        this.trackingForRussia = trackingForRussia;
        this.trackingForOtherCountries = trackingForOtherCountries;
        this.scacContainersRepository = scacContainersRepository;
        this.cacheHandler = new CacheHandler(cache)
        this.logger = logger
    }

    protected async track(args: TrackingArgsWithScac, ttl?: number): Promise<TrackingContainerResponse> {
        switch (args.country) {
            case "RU":
                let result = await this.trackingForRussia.trackContainer(args)
                await this.cacheHandler.addTrackingResultToCache(args.number, result, ttl)
                await this.scacContainersRepository.addContainer(result.container, result.scac)
                this.logger.containerSuccessLog(result)
                return result
            case "OTHER":
                let res = await this.trackingForOtherCountries.trackContainer(args)
                await this.cacheHandler.addTrackingResultToCache(args.number, res, ttl)
                await this.scacContainersRepository.addContainer(res.container, res.scac)
                this.logger.containerSuccessLog(res)
                return res
            default:
                throw new Error("user-pb only RU or OTHER")
        }
    }

    public async trackContainer(args: TrackingArgsWithScac): Promise<TrackingContainerResponse> {
        let result = await this.cacheHandler.getTrackingInfoFromCache<TrackingContainerResponse>(args.number)
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

