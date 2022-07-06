import {SkluContainer, SkluRequestSender} from "../sklu/sklu";
import {BaseContainerConstructor, ITrackingArgs, TrackingContainerResponse} from "../../base";
import {SinokorApiResponseSchema} from "../sklu/sinokorApiResponseSchema";
import {fetchArgs} from "../../helpers/requestSender";
import {IUnlocodesRepo} from "../sklu/unlocodesRepo";


export class HaluRequest extends SkluRequestSender {
    public async sendRequestToApi(args: ITrackingArgs): Promise<SinokorApiResponseSchema[]> {
        let todayYear = new Date().getFullYear()
        let requestHeaders = {
            "accept": "application/json, text/javascript, */*; q=0.01",
            "accept-language": "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6,zh-CN;q=0.5,zh;q=0.4",
            "x-requested-with": "XMLHttpRequest",
            "Referer": "http://ebiz.heung-a.com/Tracking",
            "Referrer-Policy": "strict-origin-when-cross-origin",
            "User-Agent": this.userAgentGenerator.generateUserAgent(),
        }
        let res: SinokorApiResponseSchema[] = await this.requestSender.sendRequestAndGetJson({
            url: `http://ebiz.heung-a.com/Tracking/GetBLList?cntrno=${args.number}&year=${todayYear}`,
            method: "GET",
            headers: requestHeaders,
            body: null
        })
        if (res !== []) {
            return res
        } else {
            res = await this.requestSender.sendRequestAndGetJson({
                url: `http://ebiz.heung-a.com/Tracking/GetBLList?cntrno=${args.number}&year=${todayYear - 1}`,
                method: "GET",
                headers: requestHeaders,
                body: null

            })
            if (res !== []) {
                return res

            }
        }
        return null
    }

    public async sendRequestAndGetInfoAboutMovingStringHtml(billNo: string, container?: string): Promise<string> {
        let url: string
        if (container) {
            url = `http://ebiz.heung-a.com/Tracking?blno=${billNo}&cntrno=${container}`
        } else {
            url = `http://ebiz.heung-a.com/Tracking?blno=${billNo}&cntrno=`
        }
        return await this.requestSender.sendRequestAndGetHtml({
                url: url,
                method: "GET",
                headers: {
                    "Accept": "application/json, text/javascript, */*; q=0.01",
                    "Accept-Encoding": "gzip, deflate",
                    "Accept-Language": "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6,zh;q=0.5",
                    "Connection": "keep-alive",
                    "x-requested-with": "XMLHttpRequest",
                    "Referer": "http://ebiz.heung-a.com/Tracking?blno=HASLC01220602499&cntrno=",
                    "Referrer-Policy": "strict-origin-when-cross-origin",
                    "User-Agent": this.userAgentGenerator.generateUserAgent(),
                }
            }
        )
    }
}


export class HaluContainer extends SkluContainer {
    public constructor(args: BaseContainerConstructor<fetchArgs>, unlocodesRepo: IUnlocodesRepo) {
        super(args, unlocodesRepo);
        this.skluRequest = new HaluRequest(this.requestSender, this.UserAgentGenerator)
    }

    public async trackContainer(args: ITrackingArgs): Promise<TrackingContainerResponse> {
        let result = await super.trackContainer(args);
        result.scac = "HALU"
        return result
    }
}