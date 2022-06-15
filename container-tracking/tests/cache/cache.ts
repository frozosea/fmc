import {Cache, ICache} from "../../src/cache";

const assert = require("assert")

describe("redis test", () => {
    const cache: ICache = new Cache();
    const key = "testKey"
    const value = "testValue"
    const nonStringObj = {
        key: "key",
        value: "value"
    }
    it("put into cache test", async () => {
        await cache.set(key, value)
    })
    it("get from cache test", async () => {
        await cache.set(key, value)
        let result = await cache.get(key)
        assert.strictEqual(result, value)
    })
    it("put non string object into cache and parse it", async () => {
        const keyForNonString = "nonStringObjKey"
        await cache.set(keyForNonString, JSON.stringify(nonStringObj))
        let result = await cache.get(keyForNonString)
        assert.deepEqual(JSON.parse(result), nonStringObj)
    })
    after(() => {
        return
    })
})