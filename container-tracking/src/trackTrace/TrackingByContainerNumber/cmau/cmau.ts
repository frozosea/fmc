import {
    BaseTrackerByContainerNumber,
    ITrackingArgs,
    TrackingContainerResponse,
    OneTrackingEvent, BaseContainerConstructor
} from "../../base";
import {IRequest} from "../../helpers/requestSender";
import {IBrowserArgs} from "../../helpers/browser"
import {GetContainerSizeException, GetEtaException} from "../../../exceptions";
import {WAIT_SELECTOR_TIMEOUT,REQUEST_TIMEOUT_MS} from "../../../../config.json"
const jsdom = require("jsdom");
const {JSDOM} = jsdom;


export class CmauRequest {
    protected requestSender: IRequest<IBrowserArgs>;

    public constructor(requestSender: IRequest<IBrowserArgs>) {
        this.requestSender = requestSender;
    }

    public async getCmaCgmResponse(args: ITrackingArgs): Promise<string> {
        return await this.requestSender.sendRequestAndGetHtml({
            url: "https://www.cma-cgm.com/ebusiness/tracking",
            TypeToIntoField: {
                selector: "#Reference",
                text: args.number,
                clickAfterType: {
                    locatorSelector: "#btnTracking",
                    key: "Enter",
                    waitingSelector: {selector: "#trackingsearchsection", timeout: WAIT_SELECTOR_TIMEOUT}
                }
            }, requestTimeOut: REQUEST_TIMEOUT_MS
        })
    }
}

//TODO create cmau parser
export class CmauParser {
    public getContainerSize(doc: typeof JSDOM): string {
        let htmlStrongObj = doc.window.document.querySelector("#trackingsearchsection > header > div > div > div > ul > li.ico-container > strong")
        if (htmlStrongObj !== null || htmlStrongObj !== "undefined") {
            return htmlStrongObj.textContent;
        }
        throw new GetContainerSizeException()
    }

    public getInfoAboutMoving(doc: typeof JSDOM): OneTrackingEvent {
        let table: HTMLTableElement = doc.window.document.querySelector("#gridTrackingDetails > div.k-grid-content > table")
        return
    }

    public getTimeInInfoAboutMoving(doc: typeof JSDOM): number[] {
        let time: number[] = []
        return time
    }

    //TODO CMA-CGM create method which parsing info about moving table
    public GetCmau(htmlString: string): OneTrackingEvent {
        let doc = new JSDOM(htmlString)

        return
    }

    //TODO CMA-CGM create method which parsing eta
    public getEta(doc: typeof JSDOM): number {
        throw new GetEtaException()
    }
}

export class CmauContainer extends BaseTrackerByContainerNumber<IBrowserArgs> {
    protected parser: CmauParser;
    protected request: CmauRequest;

    public constructor(args: BaseContainerConstructor<IBrowserArgs>) {
        super(args);
        this.parser = new CmauParser()
        this.request = new CmauRequest(args.requestSender)
    }

    public trackContainer(args: ITrackingArgs): Promise<TrackingContainerResponse> {
        return Promise.resolve(undefined);
    }
}