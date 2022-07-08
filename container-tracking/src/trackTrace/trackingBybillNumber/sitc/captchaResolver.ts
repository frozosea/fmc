import {fetchArgs, IRequest} from "../../helpers/requestSender";

export interface IRandomStringGenerator {
    generate(): string
}

export interface ICaptchaGetter {
    get(): Promise<any>;
}

export interface ICaptchaSolver {
    solve(image): Promise<string>;
}
export interface ICaptcha{
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

    public async get(): Promise<any> {
        let randStr = this.randomStringGenerator.generate();
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

//TODO captcha solver
export class CaptchaSolver implements ICaptchaSolver {

    public async solve(image): Promise<string> {
        return ""
    }
}


export class Captcha implements ICaptcha{
    protected randomStringGenerator: IRandomStringGenerator;
    protected captchaGetter: ICaptchaGetter;
    protected captchaSolver: ICaptchaSolver;

    public constructor(randomStringGenerator: IRandomStringGenerator, captchaGetter: ICaptchaGetter, captchaSolver: ICaptchaSolver) {
        this.randomStringGenerator = randomStringGenerator;
        this.captchaGetter = captchaGetter;
        this.captchaSolver = captchaSolver;
    }

    public async getSolvedCaptchaAndRandomString(): Promise<[string, string]> {
        return ["", ""]
    }
}