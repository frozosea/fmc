const asyncRedis = require("async-redis");
import {REDIS_URL} from "../config.json"

export interface ICache {
    get(key: string): Promise<string>

    set<T>(key: string, value: T, ttl?: number): Promise<void>
}

export class Cache implements ICache {
    protected client

    public constructor() {
        this.client = asyncRedis.createClient({
            url: REDIS_URL
        })
    }

    public async get(key: string): Promise<string> {
        return await this.client.get(key)
    }
    public async set<T>(key: string, value: T, ttl?: number): Promise<void> {
        return await this.client.set(key, value)
    }
}
