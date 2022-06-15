import {ITrackingArgs} from "../../base";
import {fetchArgs, IRequest} from "../../helpers/requestSender";
import {JSDOM} from "jsdom";


const FESCO_VESSELS = {
    "FESCO DIOMID": 212232000,
    "FESCO DALNEGORSK": 209366000,
    "COLLETTE": 538006066,
    "CAROLINA TRADER": 248139000,
    "BUXCONTACT": 255806009,
    "A HOUOU": 371793000,
    "SONGA TIGER": 636020846,
    "FESCO SOFIA": 210315000,
    "VICTORY STAR": 352382000,
    "As Riccarda": 2,
    "ZEYA": 273347430,
    "AS RICCARDA": 1,
    "BAL BOAN": 372023000,
    "KAPITAN AFANASYEV": 209464000,
    "КАПИТАН АФАНАСЬЕВ": 209464000,
    "ORTOLAN EPSILON": 255806120
}
const CITIES_DICT = {
    "shekou": 36210,
    "yantian": 31004,
    "ninbo": 30979,
    "vladivostok": 31051,
    "vostochniy port": 31052,
    "syamen": 31068,
    "pusan": 31054,
    "shanghai": 30967,
    "qingdao": 31001,
    "dalyan": 30973
}


export interface IEtaWithAisSystem {
    getEta(args: ITrackingArgs): Promise<number[]>;
}

export interface IEtaWithParseSchedule {
    getEta(args: ITrackingArgs): Promise<number[]>;
}


class BaseEtaGetter {
    protected requestSender: IRequest<fetchArgs>;

    public constructor(requestSender: IRequest<fetchArgs>) {
        this.requestSender = requestSender;

    }

}

export class EtaGetterWithAisSystem extends BaseEtaGetter implements IEtaWithAisSystem {
    private parseAisSystemResponse(aisSystemResponse: object): number {
        return 1
    }

    private async sendRequestToAisSystem(vessel): Promise<object> {
        return {}
    }

    public async getEta(args: ITrackingArgs): Promise<number[]> {
        let outputArr: number[] = []
        return outputArr
    }
}

export class EtaGetterWithScheduleParser extends BaseEtaGetter implements IEtaWithParseSchedule {
    private getDocumentInstance(stringHtmlResponse: string): JSDOM {
        return new JSDOM(stringHtmlResponse);
    }

    private parseCountElementsInTableWithSchedule(doc: JSDOM): number {
        return 1
    }

    private parseDate(dateString: string): number {
        let splitDate = dateString.split(".")
        let day = splitDate[0]
        let month = splitDate[1]
        let year = `20${splitDate[2]}`
        let isoDateString = `${year}-${month}-${day}`
        return new Date(isoDateString).getTime();
    }

    private parseSchedule(stringHtmlResponse: string): number {
        let doc = this.getDocumentInstance(stringHtmlResponse);
        let countElementsInTable = this.parseCountElementsInTableWithSchedule(doc);
        for (let i = 2; i < countElementsInTable; i++) {
            let selector = `#schedule-table > div:nth-child(${i}) > div:nth-child(7) > div.schedule-table__td-content > span:nth-child(1)`
            let dateString = doc.querySelector(selector)
            if (dateString !== null) {
                let date = new Date(dateString).getDate()
            }
        }
        return countElementsInTable
    }

    private async sendRequestToScheduleAndGetStringHtml(fromCity: string, toCity: string): Promise<string> {
        return ""
    }

    public async getEta(args: ITrackingArgs): Promise<number[]> {
        let outputArr: number[] = []
        return outputArr
    }
}