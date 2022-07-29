import {
    BaseContainerConstructor,
    BaseTrackerByContainerNumber,
    ITrackingArgs,
    OneTrackingEvent,
    TrackingContainerResponse
} from "../../base";
import {IUnlocodesRepo, UnlocodeObject} from "./unlocodesRepo";
import {fetchArgs, IRequest} from "../../helpers/requestSender";
import {SinokorApiResponseSchema} from "./sinokorApiResponseSchema";
import {NotThisShippingLineException} from "../../../exceptions";
import {IUserAgentGenerator} from "../../helpers/userAgentGenerator";
import {IDatetime} from "../../helpers/datetime";
import {SliceArray} from "slice";

const jsdom = require("jsdom");
const {JSDOM} = jsdom;

interface _NextRequestDataResp {
    billNo: string,
    eta: number,
    containerSize: string,
    unlocode: string
}


export class SkluRequestSender {
    protected requestSender: IRequest<fetchArgs>
    protected userAgentGenerator: IUserAgentGenerator

    public constructor(requestSender: IRequest<fetchArgs>, userAgentGenerator: IUserAgentGenerator) {
        this.requestSender = requestSender;
        this.userAgentGenerator = userAgentGenerator
    }

    public async sendRequestToApi(args: ITrackingArgs): Promise<SinokorApiResponseSchema[]> {
        let todayYear = new Date().getFullYear()
        let headers = {
            "Accept": "application/json, text/javascript, */*; q=0.01",
            "Accept-Encoding": "gzip, deflate",
            "Accept-Language": "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6,zh;q=0.5",
            "Connection": "keep-alive",
            "Host": "ebiz.sinokor.co.kr",
            "Referer": "http://ebiz.sinokor.co.kr/Tracking",
            "User-Agent": this.userAgentGenerator.generateUserAgent(),
            "X-Requested-With": "XMLHttpRequest"
        }
        let res: SinokorApiResponseSchema[] = await this.requestSender.sendRequestAndGetJson({
            url: `http://ebiz.sinokor.co.kr/Tracking/GetBLList?cntrno=${args.number}&year=${todayYear}`,
            method: "GET",
            headers: headers
        })
        if (res !== []) {
            return res
        } else {
            res = await this.requestSender.sendRequestAndGetJson({
                url: `http://ebiz.sinokor.co.kr/Tracking/GetBLList?cntrno=${args.number}&year=${todayYear - 1}`,
                method: "GET",
                headers: headers
            })
            if (res !== []) {
                return res

            }
        }
        return null
    }

    async sendRequestAndGetInfoAboutMovingStringHtml(billNo: string, container?: string): Promise<string> {
        let url: string
        if (container) {
            url = `http://ebiz.sinokor.co.kr/Tracking?blno=${billNo}&cntrno=${container}`
        } else {
            url = `http://ebiz.sinokor.co.kr/Tracking?blno=${billNo}&cntrno=`
        }
        return await this.requestSender.sendRequestAndGetHtml({
                url: url,
                method: "GET",
                headers: {
                    "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
                    "Accept-Encoding": "gzip, deflate",
                    "Accept-Language": "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6,zh;q=0.5",
                    "Cache-Control": "max-age=0",
                    "Connection": "keep-alive",
                    "Host": "ebiz.sinokor.co.kr",
                    "Referer": "http://ebiz.sinokor.co.kr/Tracking?blno=SNKO026210400554&cntrno=",
                    "Upgrade-Insecure-Requests": "1",
                    "User-Agent": this.userAgentGenerator.generateUserAgent()
                }
            }
        )
    }
}


export class SkluApiParser {
    protected datetime: IDatetime;

    public constructor(datetime: IDatetime) {
        this.datetime = datetime
    }

    public parseSinokorApiJson(sinokorApiJson: SinokorApiResponseSchema[]): _NextRequestDataResp {
        let lastBillNumber: string = sinokorApiJson[0].BKNO
        let eta: number = this.datetime.strptime(sinokorApiJson[0].ETA, "YYYY-MM-DD").getTime()
        let containerSize: string = sinokorApiJson[0].CNTR
        return {billNo: lastBillNumber, eta: eta, containerSize: containerSize, unlocode: sinokorApiJson[0].POD}
    }
}

export class SkluInfoAboutMovingParser {
    protected datetime: IDatetime;

    public constructor(datetime: IDatetime) {
        this.datetime = datetime
    }

    protected parseAndConvertTime(time: string): number {
        let splitTime = time.split(" ")
        let dayOfWeek = splitTime[1].toLowerCase().capitalizeFirstLetter()
        return this.datetime.strptime(`${splitTime[0]} ${dayOfWeek} ${splitTime[2]}`, "YYYY-MM-DD ddd HH:mm").getTime()
    }

    protected* zip(...iterables) {
        let iterators = iterables.map(i => i[Symbol.iterator]())
        while (true) {
            let results = iterators.map(iter => iter.next())
            if (results.some(res => res.done)) return
            else yield results.map(res => res.value)
        }
    }

    protected parseTable(stringHtml: string): [string[], typeof JSDOM.window.document] {
        let doc = new JSDOM(stringHtml).window.document
        let outputArray: string[] = []
        let table = doc.querySelector("#wrapper > div > div:nth-child(6) > div.panel-body > div > table").querySelectorAll("td")
        for (let item of table) {
            outputArray.push(item.textContent)
        }
        return [outputArray, doc]
    }

    protected parseAllOperations(doc: typeof JSDOM.window.document): string[] {
        let allOperations: string[] = []
        let allTexts = doc.getElementsByClassName("splitTable")[0].getElementsByClassName("firstTh")
        for (let item of allTexts) {
            allOperations.push(item.textContent)
        }
        return allOperations
    }

    protected getInfoAboutMoving(stringHtml: string, container: string): OneTrackingEvent[] {
        let outputArr: OneTrackingEvent[] = []
        let [text, doc] = this.parseTable(stringHtml)
        let allEvents = this.parseAllOperations(doc)
        let sliceArr = SliceArray.from(text)
        for (let [times, locations, vessels, operation] of this.zip(sliceArr[[2, , 3]], sliceArr[[1, , 3]], sliceArr[[, , 3]], allEvents)) {
            let oneEvent = {
                time: this.parseAndConvertTime(times),
                operationName: operation.trim(),
                location: locations.trim(),
                vessel: vessels.trim()
            }
            outputArr.push(oneEvent)
        }
        for (let item of outputArr) {
            if (item.vessel === container) {
                item.vessel = ""
            }
        }
        return outputArr
    }

    public parseInfoAboutMovingPage(infoAboutMovingString: string, container: string): OneTrackingEvent[] {
        return this.getInfoAboutMoving(infoAboutMovingString, container)
    }
}

export class SkluEtaParser {
    private repo: IUnlocodesRepo

    public constructor(repo: IUnlocodesRepo) {
        this.repo = repo;
    }

    async getEtaObject(data: _NextRequestDataResp): Promise<OneTrackingEvent> {
        let etaPortFullName: string = await this.repo.getUnlocode(data.unlocode)
        return {operationName: "ETA", time: data.eta, location: etaPortFullName, vessel: ""}
    }
}

export class SkluContainer extends BaseTrackerByContainerNumber<fetchArgs> {
    protected skluRequest: SkluRequestSender;
    protected infoAboutMovingParser: SkluInfoAboutMovingParser;
    protected apiParser: SkluApiParser;
    protected etaParser: SkluEtaParser;

    public constructor(args: BaseContainerConstructor<fetchArgs>, repo: IUnlocodesRepo) {
        super(args);
        this.skluRequest = new SkluRequestSender(args.requestSender, args.UserAgentGenerator);
        this.infoAboutMovingParser = new SkluInfoAboutMovingParser(args.datetime);
        this.apiParser = new SkluApiParser(args.datetime);
        this.etaParser = new SkluEtaParser(repo);
    }

    public async trackContainer(args: ITrackingArgs): Promise<TrackingContainerResponse> {
        try {
            let apiResponse: SinokorApiResponseSchema[] = await this.skluRequest.sendRequestToApi(args);
            if (apiResponse !== null) {
                let nextRequestDataObject = await this.apiParser.parseSinokorApiJson(apiResponse)
                let eta: OneTrackingEvent = await this.etaParser.getEtaObject(nextRequestDataObject);
                let infoAboutMovingStringHtml: string = await this.skluRequest.sendRequestAndGetInfoAboutMovingStringHtml(nextRequestDataObject.billNo, args.number);
                let infoAboutMoving: OneTrackingEvent[] = this.infoAboutMovingParser.parseInfoAboutMovingPage(infoAboutMovingStringHtml, args.number);
                infoAboutMoving.push(eta)
                console.log(eta)
                return {
                    container: args.number,
                    containerSize: nextRequestDataObject.containerSize,
                    scac: "SKLU",
                    infoAboutMoving: infoAboutMoving
                }
            } else {
                throw new NotThisShippingLineException()
            }
        } catch (e) {
            throw new NotThisShippingLineException()
        }


    }
}

export class UnlocodesRepoMoch implements IUnlocodesRepo {
    public async getUnlocode(unlocode: string): Promise<string> {
        return "Hososhima"
    }

    public async addUnlocode(obj: UnlocodeObject): Promise<void> {
    }
}
