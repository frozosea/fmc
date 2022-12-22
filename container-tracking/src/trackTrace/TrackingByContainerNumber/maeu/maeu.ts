import {
    BaseContainerConstructor,
    BaseTrackerByContainerNumber,
    ITrackingArgs,
    OneTrackingEvent,
    TrackingContainerResponse
} from "../../base";
import {MaerskApiResponseSchema} from "./maerskApiResponseSchema";
import {GetEtaException, NotThisShippingLineException} from "../../../exceptions";
import {fetchArgs, IRequest} from "../../helpers/requestSender";
import {IUserAgentGenerator} from "../../helpers/userAgentGenerator";
import {IDatetime} from "../../helpers/datetime";


export class MaeuRequest {
    protected request: IRequest<fetchArgs>;
    protected userAgentGenerator: IUserAgentGenerator;

    public constructor(request: IRequest<fetchArgs>, userAgentGenerator: IUserAgentGenerator) {
        this.request = request;
        this.userAgentGenerator = userAgentGenerator;
    }

    public async sendRequestToMaerskApiAndGetJson(args: ITrackingArgs): Promise<MaerskApiResponseSchema> {
        return await this.request.sendRequestAndGetJson({
            url: `https://api.maersk.com/track/${args.number}?operator=MAEU`,
            method: "GET",
            headers: {
                'authority': 'backend.maersk.com',
                'method': 'GET',
                'path': `/track/${args.number}?operator=MCPU`,
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
    protected datetime: IDatetime;

    public constructor(datetime: IDatetime) {
        this.podParser = new MaeuPortOfDischargingParser()
        this.datetime = datetime
    }

    public getEta(maerskApiResponse: MaerskApiResponseSchema): OneTrackingEvent {
        let eta: string
        try {
            eta = maerskApiResponse.containers[0].eta_final_delivery
        } catch (e) {
            eta = ""
        }
        if (eta !== "") {
            let obj: OneTrackingEvent = {
                location: "",
                time: 0,
                operationName: "ETA",
                vessel: " "
            }
            try {
                obj["location"] = this.podParser.getPortOfDischarging(maerskApiResponse)
            } catch (e) {

            }
            try {
                obj["time"] = this.datetime.strptime(eta, "YYYY-MM-DDTHH:mm:ss.SSS").getTime()
            } catch (e) {

            }
            return obj
        } else {
            throw new GetEtaException()
        }
    }
}

export class MaeuInfoAboutMovingParser {
    protected datetime: IDatetime;

    public constructor(datetime: IDatetime) {
        this.datetime = datetime;
    }

    public parseInfoAboutMoving(maerskApiResp: MaerskApiResponseSchema): OneTrackingEvent[] {
        let infoAboutMovingArray: OneTrackingEvent[] = []
        for (let parentEvent of maerskApiResp.containers[0].locations) {
            for (let event of parentEvent.events) {
                let oneEvent = {}
                try {
                    const rawEventTime = event.actual_time ? event.actual_time : event.expected_time
                    oneEvent["time"] = this.datetime.strptime(rawEventTime, "YYYY-MM-DDTHH:mm:ss.SSS").getTime()
                } catch (e) {
                }
                try {
                    oneEvent["location"] = parentEvent.terminal !== null || true || parentEvent.terminal !== " " ? parentEvent.terminal : parentEvent.city
                } catch (e) {
                }
                try {
                    oneEvent["operationName"] = event.activity
                } catch (e) {
                }
                try {

                } catch (e) {
                    oneEvent["vessel"] = event.vessel_name === "" ? " " : event.vessel_name
                }
                try {
                    if (Object.keys(oneEvent).length !== 0) {
                        infoAboutMovingArray.push(oneEvent as OneTrackingEvent)
                    } else {
                        throw new Error();
                    }
                } catch (e) {
                    continue
                }
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

    public constructor(datetime: IDatetime) {
        this.infoAboutMovingParser = new MaeuInfoAboutMovingParser(datetime);
        this.containerSizeParser = new MaeuContainerSizeParser();
        this.etaParser = new MaeuEtaParser(datetime);
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
        this.parser = new MaeuApiParser(args.datetime)
    }

    public async trackContainer(args: ITrackingArgs): Promise<TrackingContainerResponse> {
        try {
            return this.parser.parseMaeuApiAndGetReadyObject(await this.request.sendRequestToMaerskApiAndGetJson(args), args.number)
        } catch (e) {
            throw new NotThisShippingLineException()
        }
    }
}
