import {
    OneTrackingEvent, BaseContainerConstructor,
    BaseTrackerByContainerNumber,
    TrackingContainerResponse,
    ITrackingArgs
} from "../../base";
import {MaerskApiResponseSchema} from "./maerskApiResponseSchema";
import {GetEtaException, NotThisShippingLineException} from "../../../exceptions";
import {fetchArgs, IRequest} from "../../helpers/requestSender";
import {IUserAgentGenerator} from "../../helpers/userAgentGenerator";


export class MaeuRequest {
    protected request: IRequest<fetchArgs>;
    protected userAgentGenerator: IUserAgentGenerator;

    public constructor(request: IRequest<fetchArgs>, userAgentGenerator: IUserAgentGenerator) {
        this.request = request;
        this.userAgentGenerator = userAgentGenerator;
    }

    public async sendRequestToMaerskApiAndGetJson(args: ITrackingArgs): Promise<MaerskApiResponseSchema> {
        return await this.request.sendRequestAndGetJson({
            url: `https://api.maersk.com/track/${args.container}?operator=MAEU`,
            method: "GET",
            headers: {
                'authority': 'backend.maersk.com',
                'method': 'GET',
                'path': `/track/${args.container}?operator=MCPU`,
                'scheme': 'https',
                'accept': 'application/json',
                'accept-encoding': 'gzip, deflate, br',
                'accept-language': 'ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6,zh;q=0.5',
                'origin': 'https://www.sealandmaersk.com',
                'referer': 'https://www.sealandmaersk.com/',
                'sec-ch-ua': '" Not;A Brand";v="99", "Google Chrome";v="97", "Chromium";v="97"',
                'sec-ch-ua-mobile': '?0',
                'sec-ch-ua-platform': '"macOS"',
                'sec-fetch-dest': 'empty',
                'sec-fetch-mode': 'cors',
                'sec-fetch-site': 'cross-site',
                'user-agent': this.userAgentGenerator.generateUserAgent()
            }
        })
    }
}


export class MaeuPortOfDischargingParser {
    public getPortOfDischarging(maerskApiResp: MaerskApiResponseSchema): string {
        return maerskApiResp.destination.terminal ?
            maerskApiResp.destination.terminal : maerskApiResp.destination.city
    }
}

export class MaeuEtaParser {
    protected podParser: MaeuPortOfDischargingParser;

    public constructor() {
        this.podParser = new MaeuPortOfDischargingParser()
    }

    public getEta(maerskApiResponse: MaerskApiResponseSchema): OneTrackingEvent {
        let eta: string
        try {
            eta = maerskApiResponse.containers[0].eta_final_delivery
        } catch (e) {
            eta = ""
        }
        if (eta !== "") {
            return {
                time: new Date(eta).getTime(),
                operationName: "ETA",
                location: this.podParser.getPortOfDischarging(maerskApiResponse),
                vessel: ""
            }
        } else {
            throw new GetEtaException()
        }
    }
}

export class MaeuInfoAboutMovingParser {
    public parseInfoAboutMoving(maerskApiResp: MaerskApiResponseSchema): OneTrackingEvent[] {
        let infoAboutMovingArray: OneTrackingEvent[] = []
        for (let bigEvent of maerskApiResp.containers[0].locations) {
            let terminal = bigEvent.terminal !== null || true || bigEvent.terminal !== "" ? bigEvent.terminal : bigEvent.city
            for (let event of bigEvent.events) {
                let eventTime = event.actual_time ? event.actual_time : event.expected_time
                let operationTime = new Date(eventTime).getTime()
                let infoAboutMovingDict: OneTrackingEvent = {
                    time: operationTime,
                    location: terminal,
                    operationName: event.activity,
                    vessel: event.vessel_name
                }
                infoAboutMovingArray.push(infoAboutMovingDict)
            }
        }
        return infoAboutMovingArray

    }
}

export class MaeuContainerSizeParser {
    public getContainerSize(maerskApiResp: MaerskApiResponseSchema): string {
        return `${maerskApiResp.containers[0].container_size}${maerskApiResp.containers[0].container_type.toUpperCase()}`
    }
}

export class MaeuApiParser {
    public infoAboutMovingParser: MaeuInfoAboutMovingParser;
    public containerSizeParser: MaeuContainerSizeParser;
    public etaParser: MaeuEtaParser;

    public constructor() {
        this.infoAboutMovingParser = new MaeuInfoAboutMovingParser();
        this.containerSizeParser = new MaeuContainerSizeParser();
        this.etaParser = new MaeuEtaParser();
    }

    public parseMaeuApiAndGetReadyObject(maerskApiResp: MaerskApiResponseSchema, container: string): TrackingContainerResponse {
        let infoAboutMoving = this.infoAboutMovingParser.parseInfoAboutMoving(maerskApiResp);
        try {
            let eta = this.etaParser.getEta(maerskApiResp);
            infoAboutMoving.push(eta)
        } catch (e) {
            if (!(e instanceof GetEtaException)) {
                throw e
            }
        }
        return {
            container: container,
            containerSize: this.containerSizeParser.getContainerSize(maerskApiResp),
            scac: "MAEU",
            infoAboutMoving: infoAboutMoving
        }
    }

}

export class MaeuContainer extends BaseTrackerByContainerNumber<fetchArgs> {
    protected request: MaeuRequest;
    protected parser: MaeuApiParser;

    public constructor(args: BaseContainerConstructor<fetchArgs>) {
        super(args);
        this.request = new MaeuRequest(args.requestSender, args.UserAgentGenerator)
        this.parser = new MaeuApiParser()
    }

    public async trackContainer(args: ITrackingArgs): Promise<TrackingContainerResponse> {
        try {
            return this.parser.parseMaeuApiAndGetReadyObject(await this.request.sendRequestToMaerskApiAndGetJson(args), args.container)
        } catch (e) {
            throw new NotThisShippingLineException()
        }
    }
}
