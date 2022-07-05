import {FesoInfoAboutMovingParser} from "../../src/trackTrace/TrackingByContainerNumber/feso/feso";
import {
    FesoApiFullResponseSchema,
    FesoApiResponse,
    FesoLastEventsSchema
} from "../../src/trackTrace/TrackingByContainerNumber/feso/fescoApiResponseSchemas";
import {FesoInfoAboutMovingTest, parseFesoResp} from "./unit_feso_tracking_by_container_number_test";
import {FesoBillNumber, FesoEtaParser} from "../../src/trackTrace/trackingBybillNumber/feso/feso";
import {fetchArgs, IRequest} from "../../src/trackTrace/helpers/requestSender";
import {config} from "../classesConfigurator";

const assert = require("assert")


const RawFesoApiResp: FesoApiFullResponseSchema = {
    "data": {
        "tracking": {
            "data": {
                "requestKey": "ELMzn-Gvy",
                "containers": [
                    "{\"container\":\"TLLU8400756\",\"time\":\"2022-06-27T07:54:21.339Z\",\"containerCTCode\":\"20DC\",\"containerOwner\":\"COC\",\"latLng\":null,\"lastEvents\":[{\"time\":\"2022-06-10T02:55:00\",\"operation\":\"GATE-OUT\",\"operationName\":\"Вывоз с терминала\",\"operationNameLatin\":\"Gate out empty for loading\",\"locId\":42522,\"locName\":\"XIN BA DA\",\"locNameLatin\":\"XIN BA DA\",\"locIdTo\":33427,\"locNameTo\":\"СКЛАД ГРУЗОВЛАДЕЛЬЦА\",\"locNameLatinTo\":\"sklad gruzovladel'сa\",\"etCode\":null,\"transportType\":null},{\"time\":\"2022-06-23T01:55:00\",\"operation\":\"GATE-IN\",\"operationName\":\"Прибытие на терминал\",\"operationNameLatin\":\"Gate in empty from consignee\",\"locId\":38727,\"locName\":\"Honghai\",\"locNameLatin\":\"Honghai\",\"locIdTo\":38727,\"locNameTo\":\"Honghai\",\"locNameLatinTo\":\"Honghai\",\"etCode\":null,\"transportType\":null},{\"time\":\"2022-06-25T00:05:00\",\"operationName\":\"ETS\",\"operationNameLatin\":\"ETS\",\"vessel\":\"FESCO DALNEGORSK\",\"locId\":\"30967\",\"locName\":\"Шанхай\",\"locNameLatin\":\"Shanghai\"},{\"time\":\"2022-06-25T11:28:00\",\"operation\":\"LOAD\",\"operationName\":\"Погрузка на борт судна\",\"operationNameLatin\":\"Load on board\",\"locId\":38727,\"locName\":\"Honghai\",\"locNameLatin\":\"Honghai\",\"locIdTo\":31052,\"locNameTo\":\"Восточный\",\"locNameLatinTo\":\"Vostochny\",\"etCode\":\"S\",\"transportType\":\"Судно\",\"vessel\":\"FESCO DALNEGORSK\"},{\"time\":\"2022-07-05T00:00:00\",\"operationName\":\"ETA\",\"operationNameLatin\":\"ETA\",\"vessel\":\"FESCO DALNEGORSK\",\"locId\":\"31052\",\"locName\":\"Восточный\",\"locNameLatin\":\"Vostochny\",\"location\":{\"id\":42522,\"text\":\"XIN BA DA\",\"textLatin\":\"XIN BA DA\",\"parentText\":\"Шанхай\",\"parentTextLatin\":\"Shanghai\",\"country\":\"Китай\",\"countryLatin\":\"China\",\"ltCode\":\"T\",\"softshipCode\":\"CNXBD\",\"code\":null}}]}",
                    "{\"container\":\"TRHU1717783\",\"time\":\"2022-06-27T07:54:21.084Z\",\"containerCTCode\":\"20DC\",\"containerOwner\":\"COC\",\"latLng\":null,\"lastEvents\":[{\"time\":\"2022-06-10T03:26:00\",\"operation\":\"GATE-OUT\",\"operationName\":\"Вывоз с терминала\",\"operationNameLatin\":\"Gate out empty for loading\",\"locId\":42522,\"locName\":\"XIN BA DA\",\"locNameLatin\":\"XIN BA DA\",\"locIdTo\":33427,\"locNameTo\":\"СКЛАД ГРУЗОВЛАДЕЛЬЦА\",\"locNameLatinTo\":\"sklad gruzovladel'сa\",\"etCode\":null,\"transportType\":null},{\"time\":\"2022-06-23T01:54:00\",\"operation\":\"GATE-IN\",\"operationName\":\"Прибытие на терминал\",\"operationNameLatin\":\"Gate in empty from consignee\",\"locId\":38727,\"locName\":\"Honghai\",\"locNameLatin\":\"Honghai\",\"locIdTo\":38727,\"locNameTo\":\"Honghai\",\"locNameLatinTo\":\"Honghai\",\"etCode\":null,\"transportType\":null},{\"time\":\"2022-06-25T00:05:00\",\"operationName\":\"ETS\",\"operationNameLatin\":\"ETS\",\"vessel\":\"FESCO DALNEGORSK\",\"locId\":\"30967\",\"locName\":\"Шанхай\",\"locNameLatin\":\"Shanghai\"},{\"time\":\"2022-06-25T11:16:00\",\"operation\":\"LOAD\",\"operationName\":\"Погрузка на борт судна\",\"operationNameLatin\":\"Load on board\",\"locId\":38727,\"locName\":\"Honghai\",\"locNameLatin\":\"Honghai\",\"locIdTo\":31052,\"locNameTo\":\"Восточный\",\"locNameLatinTo\":\"Vostochny\",\"etCode\":\"S\",\"transportType\":\"Судно\",\"vessel\":\"FESCO DALNEGORSK\"},{\"time\":\"2022-07-05T00:00:00\",\"operationName\":\"ETA\",\"operationNameLatin\":\"ETA\",\"vessel\":\"FESCO DALNEGORSK\",\"locId\":\"31052\",\"locName\":\"Восточный\",\"locNameLatin\":\"Vostochny\",\"location\":{\"id\":42522,\"text\":\"XIN BA DA\",\"textLatin\":\"XIN BA DA\",\"parentText\":\"Шанхай\",\"parentTextLatin\":\"Shanghai\",\"country\":\"Китай\",\"countryLatin\":\"China\",\"ltCode\":\"T\",\"softshipCode\":\"CNXBD\",\"code\":null}}]}",
                    "{\"container\":\"TRHU2046355\",\"time\":\"2022-06-27T07:54:21.268Z\",\"containerCTCode\":\"20DC\",\"containerOwner\":\"COC\",\"requestDate\":\"2021-12-13T11:17:46+03:00\",\"latLng\":null,\"locations\":[{\"id\":27829,\"text\":\"Владивосток (эксп.)\",\"textLatin\":\"Vladivost.e\",\"parentText\":\"Владивосток\",\"parentTextLatin\":\"Vladivostok\",\"country\":\"Россия\",\"countryLatin\":\"Russia\",\"ltCode\":\"S\",\"softshipCode\":\"VVOEX\",\"code\":\"98020\",\"type\":\"RR\",\"info\":[{\"type\":\"departure\",\"transportType\":\"train\",\"date\":\"2022-01-03T06:53:39.000Z\"}],\"to\":\"train\",\"passed\":true},{\"id\":26433,\"text\":\"Новосибирск-Восточный\",\"textLatin\":\"Novosibirsk-Vostochnyj\",\"parentText\":\"Новосибирск\",\"parentTextLatin\":\"Novosibirsk\",\"country\":\"Россия\",\"countryLatin\":\"Russia\",\"ltCode\":\"S\",\"softshipCode\":\"RUNVV\",\"code\":\"85150\",\"type\":\"end\",\"info\":[{\"type\":\"arrival\",\"transportType\":\"train\",\"date\":\"2022-01-10T22:15:00.000Z\"}],\"from\":\"train\",\"here\":true}],\"lastEvents\":[{\"time\":\"2022-06-10T03:26:00\",\"operation\":\"GATE-OUT\",\"operationName\":\"Вывоз с терминала\",\"operationNameLatin\":\"Gate out empty for loading\",\"locId\":42522,\"locName\":\"XIN BA DA\",\"locNameLatin\":\"XIN BA DA\",\"locIdTo\":33427,\"locNameTo\":\"СКЛАД ГРУЗОВЛАДЕЛЬЦА\",\"locNameLatinTo\":\"sklad gruzovladel'сa\",\"etCode\":null,\"transportType\":null},{\"time\":\"2022-06-22T23:01:00\",\"operation\":\"GATE-IN\",\"operationName\":\"Прибытие на терминал\",\"operationNameLatin\":\"Gate in empty from consignee\",\"locId\":38727,\"locName\":\"Honghai\",\"locNameLatin\":\"Honghai\",\"locIdTo\":38727,\"locNameTo\":\"Honghai\",\"locNameLatinTo\":\"Honghai\",\"etCode\":null,\"transportType\":null},{\"time\":\"2022-06-25T00:05:00\",\"operationName\":\"ETS\",\"operationNameLatin\":\"ETS\",\"vessel\":\"FESCO DALNEGORSK\",\"locId\":\"30967\",\"locName\":\"Шанхай\",\"locNameLatin\":\"Shanghai\"},{\"time\":\"2022-06-25T11:18:00\",\"operation\":\"LOAD\",\"operationName\":\"Погрузка на борт судна\",\"operationNameLatin\":\"Load on board\",\"locId\":38727,\"locName\":\"Honghai\",\"locNameLatin\":\"Honghai\",\"locIdTo\":31052,\"locNameTo\":\"Восточный\",\"locNameLatinTo\":\"Vostochny\",\"etCode\":\"S\",\"transportType\":\"Судно\",\"vessel\":\"FESCO DALNEGORSK\"},{\"time\":\"2022-07-05T00:00:00\",\"operationName\":\"ETA\",\"operationNameLatin\":\"ETA\",\"vessel\":\"FESCO DALNEGORSK\",\"locId\":\"31052\",\"locName\":\"Восточный\",\"locNameLatin\":\"Vostochny\"}]}",
                    "{\"container\":\"FESU2122594\",\"time\":\"2022-06-27T07:54:22.020Z\",\"containerCTCode\":\"20DC\",\"containerOwner\":\"COC\",\"latLng\":null,\"lastEvents\":[{\"time\":\"2022-06-09T23:21:00\",\"operation\":\"GATE-OUT\",\"operationName\":\"Вывоз с терминала\",\"operationNameLatin\":\"Gate out empty for loading\",\"locId\":42522,\"locName\":\"XIN BA DA\",\"locNameLatin\":\"XIN BA DA\",\"locIdTo\":33427,\"locNameTo\":\"СКЛАД ГРУЗОВЛАДЕЛЬЦА\",\"locNameLatinTo\":\"sklad gruzovladel'сa\",\"etCode\":null,\"transportType\":null},{\"time\":\"2022-06-22T23:01:00\",\"operation\":\"GATE-IN\",\"operationName\":\"Прибытие на терминал\",\"operationNameLatin\":\"Gate in empty from consignee\",\"locId\":38727,\"locName\":\"Honghai\",\"locNameLatin\":\"Honghai\",\"locIdTo\":38727,\"locNameTo\":\"Honghai\",\"locNameLatinTo\":\"Honghai\",\"etCode\":null,\"transportType\":null},{\"time\":\"2022-06-25T00:05:00\",\"operationName\":\"ETS\",\"operationNameLatin\":\"ETS\",\"vessel\":\"FESCO DALNEGORSK\",\"locId\":\"30967\",\"locName\":\"Шанхай\",\"locNameLatin\":\"Shanghai\"},{\"time\":\"2022-06-25T11:42:00\",\"operation\":\"LOAD\",\"operationName\":\"Погрузка на борт судна\",\"operationNameLatin\":\"Load on board\",\"locId\":38727,\"locName\":\"Honghai\",\"locNameLatin\":\"Honghai\",\"locIdTo\":31052,\"locNameTo\":\"Восточный\",\"locNameLatinTo\":\"Vostochny\",\"etCode\":\"S\",\"transportType\":\"Судно\",\"vessel\":\"FESCO DALNEGORSK\"},{\"time\":\"2022-07-05T00:00:00\",\"operationName\":\"ETA\",\"operationNameLatin\":\"ETA\",\"vessel\":\"FESCO DALNEGORSK\",\"locId\":\"31052\",\"locName\":\"Восточный\",\"locNameLatin\":\"Vostochny\",\"location\":{\"id\":42522,\"text\":\"XIN BA DA\",\"textLatin\":\"XIN BA DA\",\"parentText\":\"Шанхай\",\"parentTextLatin\":\"Shanghai\",\"country\":\"Китай\",\"countryLatin\":\"China\",\"ltCode\":\"T\",\"softshipCode\":\"CNXBD\",\"code\":null}}]}",
                    "{\"container\":\"TCLU2451737\",\"time\":\"2022-06-27T07:54:21.894Z\",\"containerCTCode\":\"20DC\",\"containerOwner\":\"COC\",\"latLng\":null,\"lastEvents\":[{\"time\":\"2022-06-10T02:55:00\",\"operation\":\"GATE-OUT\",\"operationName\":\"Вывоз с терминала\",\"operationNameLatin\":\"Gate out empty for loading\",\"locId\":42522,\"locName\":\"XIN BA DA\",\"locNameLatin\":\"XIN BA DA\",\"locIdTo\":33427,\"locNameTo\":\"СКЛАД ГРУЗОВЛАДЕЛЬЦА\",\"locNameLatinTo\":\"sklad gruzovladel'сa\",\"etCode\":null,\"transportType\":null},{\"time\":\"2022-06-23T02:24:00\",\"operation\":\"GATE-IN\",\"operationName\":\"Прибытие на терминал\",\"operationNameLatin\":\"Gate in empty from consignee\",\"locId\":38727,\"locName\":\"Honghai\",\"locNameLatin\":\"Honghai\",\"locIdTo\":38727,\"locNameTo\":\"Honghai\",\"locNameLatinTo\":\"Honghai\",\"etCode\":null,\"transportType\":null},{\"time\":\"2022-06-25T00:05:00\",\"operationName\":\"ETS\",\"operationNameLatin\":\"ETS\",\"vessel\":\"FESCO DALNEGORSK\",\"locId\":\"30967\",\"locName\":\"Шанхай\",\"locNameLatin\":\"Shanghai\"},{\"time\":\"2022-06-25T11:23:00\",\"operation\":\"LOAD\",\"operationName\":\"Погрузка на борт судна\",\"operationNameLatin\":\"Load on board\",\"locId\":38727,\"locName\":\"Honghai\",\"locNameLatin\":\"Honghai\",\"locIdTo\":31052,\"locNameTo\":\"Восточный\",\"locNameLatinTo\":\"Vostochny\",\"etCode\":\"S\",\"transportType\":\"Судно\",\"vessel\":\"FESCO DALNEGORSK\"},{\"time\":\"2022-07-05T00:00:00\",\"operationName\":\"ETA\",\"operationNameLatin\":\"ETA\",\"vessel\":\"FESCO DALNEGORSK\",\"locId\":\"31052\",\"locName\":\"Восточный\",\"locNameLatin\":\"Vostochny\",\"location\":{\"id\":42522,\"text\":\"XIN BA DA\",\"textLatin\":\"XIN BA DA\",\"parentText\":\"Шанхай\",\"parentTextLatin\":\"Shanghai\",\"country\":\"Китай\",\"countryLatin\":\"China\",\"ltCode\":\"T\",\"softshipCode\":\"CNXBD\",\"code\":null}}]}"
                ],
                "missing": [],
                "__typename": "tracking_screenResult"
            },
            "__typename": "trackingQueries"
        }
    }
}
const ExampleFesoInfoAboutMovingApiResponse: FesoApiResponse = parseFesoResp(RawFesoApiResp)

const deleteEta = (infoAboutMoving) => {
    const copy = JSON.parse(JSON.stringify(infoAboutMoving))
    for (let item of copy) {
        if (item.operationNameLatin === "ETA") {
            copy.splice(copy.indexOf(item))
        }
    }
    return copy
}

const requestMoch: IRequest<fetchArgs> = {
    async sendRequestAndGetJson(): Promise<any> {
        return RawFesoApiResp
    },
    async sendRequestAndGetHtml(): Promise<string> {
        return ""
    }
}
describe("FESO tracking by bill number test", () => {
    const fesoInfoAboutMovingParser = new FesoInfoAboutMovingParser(config.DATETIME)
    const etaParser = new FesoEtaParser(config.DATETIME)
    it("FESO eta parser test", () => {
        const eta = etaParser.GetEta(ExampleFesoInfoAboutMovingApiResponse)
        assert.strictEqual(eta, 1656979200000) //1656943200000
    })
    it("FESO main class with moch test", async () => {
        const fesoBillNumber = new FesoBillNumber({
            requestSender: requestMoch,
            UserAgentGenerator: config.USER_AGENT_GENERATOR,
            datetime: config.DATETIME
        })
        const result = await fesoBillNumber.trackByBillNumber({number: ""})
        assert.strictEqual(result.etaFinalDelivery, 1656979200000)
        let expectedInfoAboutMoving: FesoLastEventsSchema[] = ExampleFesoInfoAboutMovingApiResponse.lastEvents
        FesoInfoAboutMovingTest(fesoInfoAboutMovingParser, result.infoAboutMoving, deleteEta(expectedInfoAboutMoving))
    })
})