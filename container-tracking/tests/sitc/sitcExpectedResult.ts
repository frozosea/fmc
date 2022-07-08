import SitcBillNumberApiResponseSchema from "../../src/trackTrace/trackingBybillNumber/sitc/sitcApiResponseSchema";

export const sitcExpectedResult = {
    "SITU9130070": {
        "code": 1,
        "msg": "success",
        "data": {
            "list": [
                {
                    "containerNo": "SITU9130070",
                    "movementName": "出口装船",
                    "movementCode": "VL",
                    "movementNameEn": "LOADED ONTO VESSEL",
                    "eventPort": "DALIAN",
                    "eventDate": "2022-06-04 22:00:00",
                    "vesselCode": "SITC CAGAYAN",
                    "voyageNo": "2212S"
                },
                {
                    "containerNo": "SITU9130070",
                    "movementName": "客户提空箱",
                    "movementCode": "OP",
                    "movementNameEn": "OUTBOUND PICKUP",
                    "eventPort": "DALIAN",
                    "eventDate": "2022-05-27 10:59:00",
                    "vesselCode": "SITC CAGAYAN",
                    "voyageNo": "2212S"
                },
                {
                    "containerNo": "SITU9130070",
                    "movementName": "空箱入场",
                    "movementCode": "MT",
                    "movementNameEn": "EMPTY CONTAINER",
                    "eventPort": "DALIAN",
                    "eventDate": "2022-05-18 18:53:00",
                    "vesselCode": "SITC MAKASSAR",
                    "voyageNo": "2210S"
                }
            ]
        }
    },
    "UETU5790574": {
        "code": 1,
        "msg": "success",
        "data": {
            "list": [
                {
                    "containerNo": "UETU5790574",
                    "movementName": "出口装船",
                    "movementCode": "VL",
                    "movementNameEn": "LOADED ONTO VESSEL",
                    "eventPort": "DALIAN",
                    "eventDate": "2022-06-04 22:00:00",
                    "vesselCode": "SITC CAGAYAN",
                    "voyageNo": "2212S"
                },
                {
                    "containerNo": "UETU5790574",
                    "movementName": "客户提空箱",
                    "movementCode": "OP",
                    "movementNameEn": "OUTBOUND PICKUP",
                    "eventPort": "DALIAN",
                    "eventDate": "2022-05-27 10:47:00",
                    "vesselCode": "SITC CAGAYAN",
                    "voyageNo": "2212S"
                },
                {
                    "containerNo": "UETU5790574",
                    "movementName": "空箱入场",
                    "movementCode": "MT",
                    "movementNameEn": "EMPTY CONTAINER",
                    "eventPort": "DALIAN",
                    "eventDate": "2022-05-18 18:52:00",
                    "vesselCode": "SITC MAKASSAR",
                    "voyageNo": "2210S"
                }
            ]
        }
    }
}

export const SitcBillNumberResponse: SitcBillNumberApiResponseSchema = {
    "code": 1,
    "msg": "success",
    "data": {
        "list1": [
            {
                "blNo": "SITDLVK222G951",
                "polen": "DALIAN",
                "del": "海参崴",
                "delen": "VLADIVOSTOK COMMERCIAL PORT",
                "pol": "大连"
            }
        ],
        "containerNo": null,
        "blNo": "SITDLVK222G951",
        "list3": [
            {
                "rowNo": "1",
                "totalCount": "2",
                "containerNo": "SITU9130070",
                "sealNo": "SITW962404",
                "voyageNo": "STCG2212S",
                "cntrType": "40HC",
                "quantity": "333",
                "cntrSize": "66.76",
                "weight": "5028.5",
                "currentport": "SHANGHAI",
                "movementname": "出口装船",
                "movementnameen": "LOADED ONTO VESSEL"
            },
            {
                "rowNo": "2",
                "totalCount": "2",
                "containerNo": "UETU5790574",
                "sealNo": "SITW962403",
                "voyageNo": "STCG2212S",
                "cntrType": "40HC",
                "quantity": "432",
                "cntrSize": "61.72",
                "weight": "4760.5",
                "currentport": "SHANGHAI",
                "movementname": "出口装船",
                "movementnameen": "LOADED ONTO VESSEL"
            }
        ],
        "list2": [
            {
                "vesselName": "___",
                "voyageNo": "2212",
                "voyageLeg": "S",
                "portFrom": "CNDLC",
                "portFromName": "DALIAN",
                "portTo": "CNSHA",
                "portToName": "SHANGHAI",
                "eta": "2022-06-05 12:00:00",
                "etd": "2022-05-31 12:00:00",
                "atd": "2022-06-04 23:00",
                "agtb": "2022-06-13 23:00",
                "cctd": null,
                "cta": null,
                "ata": "2022-06-11 11:36",
                "ctd": null,
                "agta": "2022-06-11 01:00",
                "ccta": null,
                "atb": "2022-06-13 22:42",
                "agtd": "2022-06-04 23:00",
                "ctb": null,
                "cctb": null
            },
            {
                "vesselName": "HF FORTUNE",
                "voyageNo": "2229",
                "voyageLeg": "N",
                "portFrom": "CNSHA",
                "portFromName": "SHANGHAI",
                "portTo": "RUVVO",
                "portToName": "VLADIVOSTOK COMMERCIAL PORT",
                "eta": "2022-07-05 21:00:00",
                "etd": "2022-06-26 00:00:00",
                "atd": "2022-07-03 15:30",
                "agtb": null,
                "cctd": null,
                "cta": null,
                "ata": null,
                "ctd": null,
                "agta": null,
                "ccta": "2022-07-13 02:00",
                "atb": null,
                "agtd": "2022-07-03 15:30",
                "ctb": null,
                "cctb": "2022-07-13 04:00"
            }
        ],
        "movementcode": [
            {
                "movementStatus": "3"
            }
        ]
    }
}