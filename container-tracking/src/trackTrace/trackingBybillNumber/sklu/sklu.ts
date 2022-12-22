import {NotThisShippingLineException} from "../../../exceptions";
import {SkluContainer} from "../../TrackingByContainerNumber/sklu/sklu";
import {BaseContainerConstructor, OneTrackingEvent} from "../../base";
import {fetchArgs} from "../../helpers/requestSender";
import {IUnlocodesRepo} from "../../TrackingByContainerNumber/sklu/unlocodesRepo";
import {IBillNumberTracker} from "../base";
import {ITrackingArgs, ITrackingByBillNumberResponse} from "../../../types";

const jsdom = require("jsdom");
const {JSDOM} = jsdom;

export class SkluContainerNumberParser {
    public getContainerNumberByStringHtml(stringHtml: string): string {
        let doc = new JSDOM(stringHtml).window.document
        let htmlElem = doc.querySelector("#wrapper > div > div:nth-child(5) > div.panel-body > div > div.form-group.table-responsive > div:nth-child(2) > table > tbody > tr:nth-child(1) > td:nth-child(1)")
        try {
            return htmlElem.textContent
        } catch (e) {
            throw new NotThisShippingLineException()
        }
    }
}


export class SkluBillNumber extends SkluContainer implements IBillNumberTracker {
    protected containerNumberParser: SkluContainerNumberParser

    public constructor(args: BaseContainerConstructor<fetchArgs>, unlocodesRepo: IUnlocodesRepo) {
        super(args, unlocodesRepo);
        this.containerNumberParser = new SkluContainerNumberParser()
    }

    public async trackByBillNumber(args: ITrackingArgs): Promise<ITrackingByBillNumberResponse> {
        let stringInfoAboutMovingHtml = await this.skluRequest.sendRequestAndGetInfoAboutMovingStringHtml(args.number)
        let containerNumber = this.containerNumberParser.getContainerNumberByStringHtml(stringInfoAboutMovingHtml)
        let apiResp = await this.skluRequest.sendRequestToApi({number: containerNumber})
        let eta: OneTrackingEvent = await this.etaParser.getEtaObject(this.apiParser.parseSinokorApiJson(apiResp))
        let infoAboutMoving: OneTrackingEvent[] = this.infoAboutMovingParser.parseInfoAboutMovingPage(await this.skluRequest.sendRequestAndGetInfoAboutMovingStringHtml(args.number, containerNumber), containerNumber)
        return {billNo: args.number, scac: "SKLU", infoAboutMoving: infoAboutMoving, etaFinalDelivery: eta.time}
    }
}
