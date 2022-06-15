import {
    FesoApiParser,
    FesoContainer,
    FesoContainerSizeParser,
    FesoInfoAboutMovingParser
} from "../../src/trackTrace/TrackingByContainerNumber/feso/feso";
import {FESO, OneTrackingEvent} from "../../src/types";
import {
    FesoApiFullResponseSchema,
    FesoApiResponse,
    FesoLastEventsSchema
} from "../../src/trackTrace/TrackingByContainerNumber/feso/fescoApiResponseSchemas"
import {config} from "../classesConfigurator";
import {fetchArgs, IRequest} from "../../src/trackTrace/helpers/requestSender";

const assert = require('assert');

function parseFesoResp(fesoApiResp: FesoApiFullResponseSchema): FesoApiResponse {
    return JSON.parse(fesoApiResp.data.tracking.data.containers[0])
}

const RawFesoApiResp: FesoApiFullResponseSchema = {
    "data": {
        "tracking": {
            "data": {
                "requestKey": "Ors2jy_ihL",
                "containers": [
                    "{\"container\":\"FESU2219270\",\"time\":\"2022-06-11T04:51:35.626Z\",\"containerCTCode\":\"20DC\",\"containerOwner\":\"COC\",\"latLng\":null,\"lastEvents\":[{\"time\":\"2022-06-06T16:00:00\",\"operation\":\"GATE-OUT\",\"operationName\":\"Вывоз с терминала\",\"operationNameLatin\":\"Gate out empty for loading\",\"locId\":43765,\"locName\":\"МАГИСТРАЛЬ\",\"locNameLatin\":\"MAGISTRAL\",\"locIdTo\":33427,\"locNameTo\":\"СКЛАД ГРУЗОВЛАДЕЛЬЦА\",\"locNameLatinTo\":\"sklad gruzovladel'сa\",\"etCode\":null,\"transportType\":null},{\"time\":\"2022-06-07T18:00:00\",\"operation\":\"GATE-IN\",\"operationName\":\"Прибытие на терминал\",\"operationNameLatin\":\"Gate in empty from consignee\",\"locId\":33378,\"locName\":\"Трансгарант\",\"locNameLatin\":\"ZAPSIBCONT\",\"locIdTo\":33378,\"locNameTo\":\"Трансгарант\",\"locNameLatinTo\":\"ZAPSIBCONT\",\"etCode\":\"T\",\"transportType\":\"Автомобиль\",\"vessel\":\"\",\"location\":{\"id\":43765,\"text\":\"МАГИСТРАЛЬ\",\"textLatin\":\"MAGISTRAL\",\"parentText\":\"Новосибирск\",\"parentTextLatin\":\"Novosibirsk\",\"country\":\"Россия\",\"countryLatin\":\"Russia\",\"ltCode\":\"T\",\"softshipCode\":\"MAGIST\",\"code\":null}}]}"
                ],
                "missing": [],
                "__typename": "tracking_screenResult"
            },
            "__typename": "trackingQueries"
        }
    }
}

const ExampleFesoInfoAboutMovingApiResponse: FesoApiResponse = parseFesoResp(RawFesoApiResp)


function FesoInfoAboutMovingTest(fesoInfoAboutMovingParser: FesoInfoAboutMovingParser) {
    let expectedInfoAboutMoving: FesoLastEventsSchema[] = ExampleFesoInfoAboutMovingApiResponse.lastEvents
    let actualInfoAboutMoving: OneTrackingEvent[] = fesoInfoAboutMovingParser.getInfoAboutMoving(ExampleFesoInfoAboutMovingApiResponse)
    for (let event in actualInfoAboutMoving) {
        //locNameLatin
        assert.strictEqual(actualInfoAboutMoving[event].time, new Date(expectedInfoAboutMoving[event].time).getTime())
        assert.strictEqual(actualInfoAboutMoving[event].operationName, expectedInfoAboutMoving[event].operationNameLatin)
        assert.strictEqual(actualInfoAboutMoving[event].location, expectedInfoAboutMoving[event].locNameLatin)
    }
}

function FesoContainerSizeTest(containerSizeParser: FesoContainerSizeParser) {
    let expectedContainerSize = ExampleFesoInfoAboutMovingApiResponse.containerCTCode
    let actualContainerSize = containerSizeParser.getContainerSize(ExampleFesoInfoAboutMovingApiResponse)
    assert.strictEqual(actualContainerSize, expectedContainerSize)
}

function FesoApiParserTest(apiParser: FesoApiParser, infoAboutMovingParser: FesoInfoAboutMovingParser, containerSizeParser: FesoContainerSizeParser, container: string) {
    let expectedReadyObject = {
        container: container,
        scac: "FESO",
        containerSize: containerSizeParser.getContainerSize(ExampleFesoInfoAboutMovingApiResponse),
        infoAboutMoving: infoAboutMovingParser.getInfoAboutMoving(ExampleFesoInfoAboutMovingApiResponse)
    }
    let actualReadyObject = apiParser.getOutputObjectAndGetEta(ExampleFesoInfoAboutMovingApiResponse)
    assert.deepEqual(actualReadyObject, expectedReadyObject)
}


const requestMoch: IRequest<fetchArgs> ={
    async sendRequestAndGetJson(args: fetchArgs): Promise<any> {
        return RawFesoApiResp
    },
    async sendRequestAndGetHtml(args: fetchArgs): Promise<string> {
        return ""
    }
}

describe("FESO container tracking Test", () => {
    const container = "FESU2219270"

    const fesoContainerSizeParser = new FesoContainerSizeParser()
    const fesoInfoAboutMovingParser = new FesoInfoAboutMovingParser()
    const fesoApiParser = new FesoApiParser()

    it("FESO container size parser test", () => {
        FesoContainerSizeTest(fesoContainerSizeParser)
    })

    it("FESO info about moving parser test", () => {
        FesoInfoAboutMovingTest(fesoInfoAboutMovingParser)
    })
    it("FESO full api parser test", () => {
        FesoApiParserTest(fesoApiParser, fesoInfoAboutMovingParser, fesoContainerSizeParser, container)
    })
    it("FESO main class with moch test", async () => {
        let feso = new FesoContainer({
            datetime: config.DATETIME,
            requestSender: requestMoch,
            UserAgentGenerator: config.USER_AGENT_GENERATOR
        })
        let actualResult = await feso.trackContainer({container: "FESU2219270"})
        let expectedInfoAboutMoving: FesoLastEventsSchema[] = ExampleFesoInfoAboutMovingApiResponse.lastEvents
        for (let event in actualResult.infoAboutMoving) {
            //locNameLatin
            assert.strictEqual(actualResult.infoAboutMoving[event].time, new Date(expectedInfoAboutMoving[event].time).getTime())
            assert.strictEqual(actualResult.infoAboutMoving[event].operationName, expectedInfoAboutMoving[event].operationNameLatin)
            assert.strictEqual(actualResult.infoAboutMoving[event].location, expectedInfoAboutMoving[event].locNameLatin)
        }
    })
})