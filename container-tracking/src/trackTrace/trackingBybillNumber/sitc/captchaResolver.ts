import {fetchArgs, IRequest} from "../../helpers/requestSender";
import * as fs from "fs";

const FormData = require('form-data');

export interface IRandomStringGenerator {
    generate(): string
}

export interface ICaptchaGetter {
    get(randStr: string): Promise<any>;
}

export interface ICaptchaSolver {
    solve(image: Blob | Buffer | any): Promise<string>;
}

export interface ICaptcha {
    getSolvedCaptchaAndRandomString(): Promise<[string, string]>
}

export class RandomStringGenerator implements IRandomStringGenerator {
    public generate(): string {
        const desiredMaxLength = 17
        let randomNumber = '';
        for (let i = 0; i < desiredMaxLength; i++) {
            randomNumber += Math.floor(Math.random() * 10);
        }
        return randomNumber
    }
}

export class CaptchaGetter implements ICaptchaGetter {
    protected randomStringGenerator: IRandomStringGenerator;
    protected requestSender: IRequest<fetchArgs>;

    public constructor(generator: IRandomStringGenerator, requestSender: IRequest<fetchArgs>) {
        this.randomStringGenerator = generator;
        this.requestSender = requestSender;
    }

    public async get(randStr: string): Promise<any> {
        let url = `http://api.sitcline.com/code?randomStr=${randStr}`
        return await this.requestSender.sendRequestAndGetImage({
            url: url, headers: {
                "accept": "image/avif,image/webp,image/apng,image/svg+xml,image/*,*/*;q=0.8",
                "accept-language": "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6,zh-CN;q=0.5,zh;q=0.4",
                "Referer": "http://api.sitcline.com/app/cargoTrackSearch",
                "Referrer-Policy": "strict-origin-when-cross-origin"
            }, body: null, method: "GET"
        })
    }
}

interface CaptchaSolverGetIdResponse {
    status: string | null
    request: string | null
}

interface CaptchaSolverResponse {
    status: string | null | number
    text: string | null
    request?: string | null

}

//TODO test captcha solver
export class CaptchaSolver implements ICaptchaSolver {
    protected requestSender: IRequest<fetchArgs>;

    public constructor(requestSender: IRequest<fetchArgs>) {
        this.requestSender = requestSender;
    }

    protected async sendRequestToSolve(image: Blob): Promise<string> {
        let formData = new FormData();
        formData.append('key', process.env.CAPTCHA_SOLVER_SERVICE_KEY)
        formData.append('numeric', '1')
        formData.append('json', '1')
        formData.append('phrase', '0')
        formData.append('regsense', '0')
        formData.append('calc', '0')
        formData.append('min_len', '4')
        formData.append('max_len', '4')
        formData.append('language', '0')
        formData.append('file', image, {})
        let response: CaptchaSolverGetIdResponse = await this.requestSender.sendRequestAndGetJson({
            url: `http://2captcha.com/in.php`,
            method: "POST",
            body: formData,
        })
        return response.request
    }

    protected async sendRequestAndGetSolvedCaptcha(id: string): Promise<CaptchaSolverResponse> {
        return await this.requestSender.sendRequestAndGetJson({
            url: `http://2captcha.com/res.php?key=${process.env.CAPTCHA_SOLVER_SERVICE_KEY}&action=get&id=${id}&json=1`,
            method: "GET",
        })
    }

    public async solve(image: Blob): Promise<string> {
        let idOfRequest = await this.sendRequestToSolve(image)

        function sleep(ms): Promise<void> {
            return new Promise(resolve => setTimeout(resolve, ms));
        }

        do {
            await sleep(1000)
            let resp = await this.sendRequestAndGetSolvedCaptcha(idOfRequest)
            if (resp.request !== "CAPCHA_NOT_READY") return resp.request
        } while (true)
    }
}


export class Captcha implements ICaptcha {
    protected randomStringGenerator: IRandomStringGenerator;
    protected captchaGetter: ICaptchaGetter;
    protected captchaSolver: ICaptchaSolver;

    public constructor(randomStringGenerator: IRandomStringGenerator, captchaGetter: ICaptchaGetter, captchaSolver: ICaptchaSolver) {
        this.randomStringGenerator = randomStringGenerator;
        this.captchaGetter = captchaGetter;
        this.captchaSolver = captchaSolver;
    }

    public async getSolvedCaptchaAndRandomString(): Promise<[string, string]> {
        let randStr = this.randomStringGenerator.generate()
        let imageWithCaptcha = await this.captchaGetter.get(randStr)
        fs.writeFileSync(`${process.env.PWD}/${randStr}.jpeg`, Buffer.from(await imageWithCaptcha.arrayBuffer()))
        let response = await this.captchaSolver.solve(fs.createReadStream(`${process.env.PWD}/${randStr}.jpeg`))
        fs.unlink(`${process.env.PWD}/${randStr}.jpeg`, () => {
        })
        return [response, randStr]
    }
}