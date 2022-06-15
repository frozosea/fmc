import {ICache} from "./cache";
import {TrackingArgsWithScac, TrackingContainerResponse} from "./types";
import {TrackingForRussia} from "./trackTrace/TrackingByContainerNumber/tracking/trackingForRussia";
import {TrackingForOtherCountries} from "./trackTrace/TrackingByContainerNumber/tracking/trackingForOtherCountries";
import {IScacContainers} from "./trackTrace/TrackingByContainerNumber/containerScacRepo";
import {IServiceLogger} from "./logging";
import {CONTAINER_TRACKING_RESULT_REDIS_TTL_SECONDS} from "../config.json"

export class CacheHandler {
    protected cache: ICache;

    public constructor(cache: ICache) {
        this.cache = cache

    }

    public async addTrackingResultToCache(container: string, trackingResult: TrackingContainerResponse, ttl?: number): Promise<void> {
        await this.cache.set(container, JSON.stringify(trackingResult), ttl)
    }

    public async getTrackingInfoFromCache(container: string): Promise<TrackingContainerResponse | null> {
        return JSON.parse(await this.cache.get(container))
    }
}

export default class TrackingController {
    protected trackingForRussia: TrackingForRussia;
    protected trackingForOtherCountries: TrackingForOtherCountries;
    protected scacContainersRepository: IScacContainers;
    protected cacheHandler: CacheHandler;
    protected logger: IServiceLogger;

    public constructor(trackingForRussia: TrackingForRussia,
                       trackingForOtherCountries: TrackingForOtherCountries,
                       scacContainersRepository: IScacContainers,
                       cache: ICache, logger: IServiceLogger) {
        this.trackingForRussia = trackingForRussia;
        this.trackingForOtherCountries = trackingForOtherCountries;
        this.scacContainersRepository = scacContainersRepository;
        this.cacheHandler = new CacheHandler(cache)
        this.logger = logger
    }

    protected async track(args: TrackingArgsWithScac, ttl?: number): Promise<TrackingContainerResponse> {
        const track = async (tracking: TrackingForRussia | TrackingForOtherCountries): Promise<TrackingContainerResponse> =>{
            let result = await tracking.trackContainer(args)
            await this.cacheHandler.addTrackingResultToCache(args.container, result, ttl)
            await this.scacContainersRepository.addContainer(result.container, result.scac)
            this.logger.containerSuccessLog(result)
            return result
        }
        switch (args.country) {
            case "RU":
                let result = await this.trackingForRussia.trackContainer(args)
                await this.cacheHandler.addTrackingResultToCache(args.container, result, ttl)
                await this.scacContainersRepository.addContainer(result.container, result.scac)
                this.logger.containerSuccessLog(result)
                return result
            case "OTHER":
                let res = await this.trackingForOtherCountries.trackContainer(args)
                await this.cacheHandler.addTrackingResultToCache(args.container, res, ttl)
                await this.scacContainersRepository.addContainer(res.container, res.scac)
                this.logger.containerSuccessLog(res)
                return res
            default:
                throw new Error("user only RU or OTHER")
        }
    }

    public async trackContainer(args: TrackingArgsWithScac): Promise<TrackingContainerResponse> {
        let result = await this.cacheHandler.getTrackingInfoFromCache(args.container)
        if (result !== null) return result
        if (args.scac === "AUTO") {
            let scacFromDb = await this.scacContainersRepository.getScac(args.container)
            if (scacFromDb !== null) {
                try {
                    return await this.track({
                        container: args.container,
                        country: args.country,
                        scac: scacFromDb
                    }, CONTAINER_TRACKING_RESULT_REDIS_TTL_SECONDS)
                } catch (e) {
                    this.logger.containerNotFoundLog(args.container)
                    return await this.track({
                        container: args.container,
                        country: args.country,
                        scac: "AUTO"
                    }, CONTAINER_TRACKING_RESULT_REDIS_TTL_SECONDS)
                }
            } else {
                try {
                    return await this.track(args, CONTAINER_TRACKING_RESULT_REDIS_TTL_SECONDS)
                } catch (e) {
                    this.logger.containerNotFoundLog(args.container)
                }
            }
        } else {
            try {
                return await this.track(args, CONTAINER_TRACKING_RESULT_REDIS_TTL_SECONDS)
            } catch (e) {
                this.logger.containerNotFoundLog(args.container)
            }
        }

    }
}

