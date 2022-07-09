import {
    BaseContainerConstructor,
    BaseTrackerByContainerNumber,
    ITrackingArgs,
    OneTrackingEvent,
    TrackingContainerResponse
} from "../../base";
import {fetchArgs, IRequest} from "../../helpers/requestSender";
import {IUserAgentGenerator} from "../../helpers/userAgentGenerator";
import {KmtuRequestSchema} from "./requestSchemas";
import {IDatetime} from "../../helpers/datetime";
import {NotThisShippingLineException} from "../../../exceptions";
import RequestsUtils from "../../helpers/utils/requestsUtils";

const jsdom = require("jsdom");
const {JSDOM} = jsdom;

class BaseKmtuRequest {
    protected request: IRequest<fetchArgs>;
    protected userAgentGenerator: IUserAgentGenerator;

    public constructor(request: IRequest<fetchArgs>, userAgentGenerator: IUserAgentGenerator) {
        this.request = request;
        this.userAgentGenerator = userAgentGenerator
    }

    public async sendRequestToKmtu(requestData: any, url: string): Promise<string> {
        let headers = {
            "accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
            "accept-language": "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6,zh-CN;q=0.5,zh;q=0.4",
            "cache-control": "max-age=0",
            "content-type": "application/x-www-form-urlencoded",
            "upgrade-insecure-requests": "1",
            'User-Agent': this.userAgentGenerator.generateUserAgent()
        }
        return await this.request.sendRequestAndGetHtml({
            url: url,
            method: "POST",
            headers: headers,
            body: requestData
        })
    }
}

export class KmtuRequest extends BaseKmtuRequest {
    public async getEtaKmtuStringHtml(args: ITrackingArgs): Promise<string> {
        let requestData: KmtuRequestSchema = {
            'hid_bl_no': '',
            'bk_no': '',
            'dt_knd': 'CN',
            'own_yn': 'N',
            'vsl_cd': '',
            'voy_no': '',
            'pod': '',
            'pol': '',
            'cntr_no': '',
            'hiddenKnd': '',
            'hiddenSearchNo': '',
            'flag': '0',
            'pus_no': '',
            'Rail': '',
            'RailCnt': '',
            'condition': 'CN',
            'bl_no': args.number
        }
        return await this.sendRequestToKmtu(RequestsUtils.jsonToQueryString(requestData), "http://www.ekmtc.com/CCIT100/searchContainerList.do")
    }

}


export class KmtuInfoAboutMovingRequest extends BaseKmtuRequest {
    public async getInfoAboutMovingStringHtml(requestData: KmtuRequestSchema): Promise<string> {
        return await this.sendRequestToKmtu(RequestsUtils.jsonToQueryString(requestData), "http://www.ekmtc.com/CCIT100/searchContainerDetail.do")
    }
}

export class KmtuDataForInfoAboutMovingRequestCrawler {
    protected datetime: IDatetime;

    public constructor(datetime: IDatetime) {
        this.datetime = datetime
    }

    protected getTableJsonAndDocInstance(stringHtml: string): object[] {
        let doc = new JSDOM(stringHtml).window.document
        let table = doc.querySelector("#paging_1")
        return [...table.rows].map(r => {
            const entries = [...r.cells].map((c, i) => {
                return [i, c.textContent.replace(/(\r\n|\n|\r\t|\t)/gm, "")]
            })
            return Object.fromEntries(entries)
        })
    }

    protected getPod(data: any): string {
        let pod: string
        try {
            pod = data[1][4].split(/\d/gm)[0]
        } catch (e) {
            pod = data[0][4].split(/\d/gm)[0]
        }
        return pod
    }

    public getDataForInfoAboutMovingRequest(stringHtml: string): KmtuRequestSchema {
        let data = this.getTableJsonAndDocInstance(stringHtml)
        let pod = this.getPod(data)
        let pol = data[0][3].split(/\d/gm)[0]
        let HidbillNo = data[0][0]
        let bkNo = data[0][1]
        return {
            'hid_bl_no': HidbillNo.trim(),
            'bk_no': bkNo.trim(),
            'dt_knd': 'CN',
            'own_yn': '',
            'vsl_cd': '',
            'voy_no': '',
            'pod': pod.trim(),
            'pol': pol.trim(),
            'cntr_no': data[0][2].trim(),
            'hiddenKnd': 'CN',
            'hiddenSearchNo': data[0][2].trim(),
            'flag': '0',
            'pus_no': '',
            'Rail': '',
            'RailCnt': '',
            'condition': 'CN',
            'bl_no': data[0][2].trim()
        }
    }
}


export class KmtuEtaParser extends KmtuDataForInfoAboutMovingRequestCrawler {
    public parseEta(stringHtml: string): OneTrackingEvent {
        let tableJson = this.getTableJsonAndDocInstance(stringHtml)
        let etaVessel: string = ""
        try {
            etaVessel = tableJson[0][5].split(")")[1]
        } catch (e) {
        }
        let pod: string = this.getPod(tableJson).trim()
        let etaDate: string = tableJson[0][4].match(/\d{4}.\d{2}.\d{2}\/\d{2}:\d{2}/g)
        //2022.05.25/21:00
        return {
            time: this.datetime.strptime(etaDate[0], "YYYY.MM.DD/HH:mm").getTime(),
            operationName: "ETA",
            vessel: etaVessel,
            location: pod
        }
    }
}

export class KmtuInfoAboutMovingParser {
    protected datetime: IDatetime;

    public constructor(datetime: IDatetime) {
        this.datetime = datetime
    }

    private clearFunctionArgs(args: string[]): string[] {
        for (let i in args) {
            let element = args[i].split(',')
            for (let item of element) {
                element[element.indexOf(item)] = item.replace(/'/gm, "")
            }
            args[i] = element.join(",")
        }
        return args
    }

    private parseInfoAboutMovingDecorator(stringHtml: string, callback: Function): Function {
        return () => {
            let re = new RegExp(/'\d','\w+','\w.+'/gm)
            let functionArgs = this.clearFunctionArgs(stringHtml.match(re))
            return callback(new Set(functionArgs))
        }

    }

    private getEventByNumberOfEvent(numberOfEvent: string | number): string {
        let structOfEvents = {
            '1': 'container was picked',
            '2': 'container is on inland haulage',
            '3': 'container was arrived',
            '5': 'container is onboard and is scheduled to arrive at transshipment',
            '6': 'container was discharged and it will be arranged for final destination',
            '7': 'container is onboard after transshipment',
            '8': 'container was discharged',
            '9': 'container was arrived',
            '10': 'container was picked up by the consignee',
            '11': 'container was returned'
        }
        return structOfEvents[numberOfEvent]
    }

    private parseEvents(stringHtml: string): string[] {
        let events: string[] = []
        this.parseInfoAboutMovingDecorator(stringHtml, (functionArgs: string[]) => {
            for (let args of functionArgs) {
                let splitArgs: string[] = args.split(',')
                let event: string = this.getEventByNumberOfEvent(splitArgs[0])
                events.push(event)
            }
        })();
        return events
    }

    private parseDates(stringHtml: string): number[] {
        let dates: number[] = []
        this.parseInfoAboutMovingDecorator(stringHtml, (funArgs: string[]) => {
            for (let args of funArgs) {
                let splitArgs: string[] = args.split(',')
                let date: number = this.datetime.strptime(splitArgs[3], "YMMDD").getTime()
                dates.push(date)
            }
        })();
        return dates
    }

    private parseLocations(stringHtml: string): string[] {
        let locations: string[] = []
        this.parseInfoAboutMovingDecorator(stringHtml, (fnArgs: string[]) => {
            for (let args of fnArgs) {
                let splitArgs: string[] = args.split(',')
                let location: string = `${splitArgs[1]},${splitArgs[2]}`
                locations.push(location)
            }
        })();
        return locations
    }

    public getInfoAboutMoving(stringHtml: string): OneTrackingEvent[] {
        let infoAboutMoving: OneTrackingEvent[] = []
        let dates: number[] = this.parseDates(stringHtml)
        let events: string[] = this.parseEvents(stringHtml)
        let locations: string[] = this.parseLocations(stringHtml)
        if (dates.length !== events.length || locations.length !== dates.length || locations.length !== events.length) {
            throw new Error("len of arrays is not equal (KMTU)")
        }
        for (let i = 0; i < dates.length; i++) {
            let oneObject: OneTrackingEvent = {
                time: dates[i],
                operationName: events[i],
                location: locations[i],
                vessel: ""
            }
            infoAboutMoving.push(oneObject)
        }
        return infoAboutMoving.reverse()
    }
}

export class KmtuContainer extends BaseTrackerByContainerNumber<fetchArgs> {
    protected etaRequest: KmtuRequest;
    protected infoAboutMovingRequest: KmtuInfoAboutMovingRequest;
    protected nextRequestDataParser: KmtuDataForInfoAboutMovingRequestCrawler;
    protected etaParser: KmtuEtaParser;
    protected infoAboutMovingParser: KmtuInfoAboutMovingParser;

    public constructor(args: BaseContainerConstructor<fetchArgs>) {
        super(args);
        this.etaRequest = new KmtuRequest(args.requestSender, args.UserAgentGenerator)
        this.infoAboutMovingRequest = new KmtuInfoAboutMovingRequest(args.requestSender, args.UserAgentGenerator)
        this.infoAboutMovingParser = new KmtuInfoAboutMovingParser(args.datetime)
        this.nextRequestDataParser = new KmtuDataForInfoAboutMovingRequestCrawler(args.datetime)
        this.etaParser = new KmtuEtaParser(args.datetime)
    }

    public async trackContainer(args: ITrackingArgs): Promise<TrackingContainerResponse> {
        try {
            let etaAndNextRequestDataStringHtml: string = await this.etaRequest.getEtaKmtuStringHtml(args);
            let nextRequestData: KmtuRequestSchema = this.nextRequestDataParser.getDataForInfoAboutMovingRequest(etaAndNextRequestDataStringHtml)
            let eta: OneTrackingEvent = this.etaParser.parseEta(etaAndNextRequestDataStringHtml);
            try {
                let infoAboutMovingHtml: string = await this.infoAboutMovingRequest.getInfoAboutMovingStringHtml(nextRequestData);
                let kmtuInfoAboutMoving: OneTrackingEvent[] = this.infoAboutMovingParser.getInfoAboutMoving(infoAboutMovingHtml)
                kmtuInfoAboutMoving.push(eta)
                return {
                    container: args.number,
                    scac: "KMTU",
                    containerSize: "",
                    infoAboutMoving: kmtuInfoAboutMoving
                }
            } catch (e) {
                let kmtuInfoAboutMoving = []
                kmtuInfoAboutMoving.push(eta)
                return {
                    container: args.number,
                    scac: "KMTU",
                    containerSize: "",
                    infoAboutMoving: kmtuInfoAboutMoving
                }
            }
        } catch (e) {
            throw new NotThisShippingLineException()
        }
    }
}
