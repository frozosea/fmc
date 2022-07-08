import {RandomStringGenerator} from "../../src/trackTrace/trackingBybillNumber/sitc/captchaResolver";

const assert = require("assert");
describe("captcha solver test", () => {
    it("random string generator test", () => {
        let generator = new RandomStringGenerator();
        for (let i = 0; i < 20; i++) {
            let randString = generator.generate()
            assert.equal(randString.length, 17)
        }
    });

})