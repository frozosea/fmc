import ZhguApiResponseSchema from "../../src/trackTrace/trackingBybillNumber/zhgu/zhguApiResponseSchema";


export const zhguExampleResponseWithoutMistakes: ZhguApiResponseSchema= {
    "backCode": "200",
    "backMessage": "Request succeeded!",
    "object": [
        {
            "createdByUser": null,
            "createdOffice": null,
            "createdDtmLoc": null,
            "createdTimeZone": null,
            "updatedByUser": null,
            "updatedOffice": null,
            "updatedDtmLoc": null,
            "updatedTimeZone": null,
            "recordVersion": null,
            "rowStatus": 2,
            "principalGroupCode": null,
            "blNo": "ZGSHA0100001921",
            "tripNumber": "1",
            "lineType": "1",
            "vesselName": "ZHONG GU BO HAI",
            "voyage": "22004N",
            "portFrom": "CNSHA",
            "portFromName": "SHANGHAI",
            "portTo": "RUVYP",
            "portToName": "VOSTOCHNY",
            "etd": "2022-07-22",
            "atd": "2022-07-25",
            "eta": "2022-07-31",
            "ata": "",
            "departureTime": "ETD:2022-07-22\nATD:2022-07-25",
            "arrivalTime": "ETA:2022-07-31\nATA:"
        }
    ]
}
