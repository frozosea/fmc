import {
    OneyGetContainerSizeSchema,
    OneyInfoAboutMovingSchema
} from "../../src/trackTrace/TrackingByContainerNumber/oney/oneyApiResponseSchemas";
import {TrackingContainerResponse} from "../../src/types";

export const oneyExpectedData: TrackingContainerResponse = {
    'container': 'GAOU6642924',
    'scac': 'ONEY',
    'containerSize': "40'DRY HC.",
    'infoAboutMoving': [{
        'time': 1648084740000,
        'operationName': 'Empty Container Release to Shipper',
        'location': 'PUSAN, KOREA REPUBLIC OF',
        'vessel': ''
    }, {
        'time': 1649119200000,
        'operationName': 'Gate In to Outbound Terminal',
        'location': 'PUSAN, KOREA REPUBLIC OF',
        'vessel': ''
    }, {
        'time': 1649297340000,
        'operationName': "Loaded on 'HYUNDAI SINGAPORE 126E' at Port of Loading",
        'location': 'PUSAN, KOREA REPUBLIC OF',
        'vessel': 'HYUNDAI SINGAPORE'
    }, {
        'time': 1649329200000,
        'operationName': "'HYUNDAI SINGAPORE 126E' Departure from Port of Loading",
        'location': 'PUSAN, KOREA REPUBLIC OF',
        'vessel': 'HYUNDAI SINGAPORE'
    }, {
        'time': 1650261600000,
        'operationName': "'HYUNDAI SINGAPORE 126E' Arrival at Port of Discharging",
        'location': 'VANCOUVER, BC, CANADA',
        'vessel': 'HYUNDAI SINGAPORE'
    }, {
        'time': 1653671640000,
        'operationName': "'HYUNDAI SINGAPORE 126E' POD Berthing Destination",
        'location': 'VANCOUVER, BC, CANADA',
        'vessel': 'HYUNDAI SINGAPORE'
    }, {
        'time': 1653997620000,
        'operationName': "Unloaded from 'HYUNDAI SINGAPORE 126E' at Port of Discharging",
        'location': 'VANCOUVER, BC, CANADA',
        'vessel': 'HYUNDAI SINGAPORE'
    }, {
        'time': 1654179840000,
        'operationName': 'Loaded on rail at inbound rail origin',
        'location': 'VANCOUVER, BC, CANADA',
        'vessel': ''
    }, {
        'time': 1654204140000,
        'operationName': 'Inbound Rail Departure',
        'location': 'VANCOUVER, BC, CANADA',
        'vessel': ''
    }, {
        'time': 1654763640000,
        'operationName': 'Inbound Rail Arrival',
        'location': 'DETROIT, MI, UNITED STATES',
        'vessel': ''
    }, {
        'time': 1654772520000,
        'operationName': 'Unloaded from rail at inbound rail destination',
        'location': 'DETROIT, MI, UNITED STATES',
        'vessel': ''
    }, {
        'time': 1654821480000,
        'operationName': 'Gate Out from Inbound CY for Delivery to Consignee',
        'location': 'DETROIT, MI, UNITED STATES',
        'vessel': ''
    }, {
        'time': 1654827840000,
        'operationName': 'Empty Container Returned from Customer',
        'location': 'DETROIT, MI, UNITED STATES',
        'vessel': ''
    }]
}

export const oneyInfoAboutMovingExample: OneyInfoAboutMovingSchema = {
    "TRANS_RESULT_KEY": "S",
    "Exception": "",
    "count": "13",
    "list": [
        {
            "maxRows": 0,
            "models": [],
            "vslCd": "",
            "no": "1",
            "copNo": "CSEL2303645419",
            "eventDt": "2022-03-24 11:19",
            "vslEngNm": "",
            "placeNm": "PUSAN, KOREA REPUBLIC OF",
            "skdVoyNo": "",
            "yardNm": "HANJIN BUSAN NEW PORT COMPANY(HJNC)",
            "copDtlSeq": "1011",
            "skdDirCd": "",
            "actTpCd": "A",
            "statusNm": "Empty Container Release to Shipper",
            "statusCd": "MOTYDO",
            "nodCd": "KRPUS14",
            "vvd": "",
            "lloydNo": "",
            "hashColumns": [
                [
                    "no",
                    "1"
                ],
                [
                    "bl_no",
                    {}
                ],
                [
                    "cntr_no",
                    {}
                ],
                [
                    "cntr_tpsz_nm",
                    {}
                ],
                [
                    "ibflag",
                    {}
                ],
                [
                    "event_dt",
                    "2022-03-24 11:19"
                ],
                [
                    "lloyd_no",
                    ""
                ],
                [
                    "cop_dtl_seq",
                    "1011"
                ],
                [
                    "yard_cd",
                    {}
                ],
                [
                    "seal_no",
                    {}
                ],
                [
                    "vgm_rcv",
                    {}
                ],
                [
                    "dsp_bkg_no",
                    {}
                ],
                [
                    "place_cd",
                    {}
                ],
                [
                    "status_cd",
                    "MOTYDO"
                ],
                [
                    "cop_no",
                    "CSEL2303645419"
                ],
                [
                    "nod_cd",
                    "KRPUS14"
                ],
                [
                    "act_tp_cd",
                    "A"
                ],
                [
                    "po_no",
                    {}
                ],
                [
                    "soc_flg",
                    {}
                ],
                [
                    "pagerows",
                    {}
                ],
                [
                    "cntr_tpsz_cd",
                    {}
                ],
                [
                    "enbl_flag",
                    {}
                ],
                [
                    "vvd",
                    ""
                ],
                [
                    "weight",
                    {}
                ],
                [
                    "yard_nm",
                    "HANJIN BUSAN NEW PORT COMPANY(HJNC)"
                ],
                [
                    "cntr_flg",
                    {}
                ],
                [
                    "bkg_no",
                    {}
                ],
                [
                    "vsl_eng_nm",
                    ""
                ],
                [
                    "vsl_cd",
                    ""
                ],
                [
                    "cop_sts_cd",
                    {}
                ],
                [
                    "piece",
                    {}
                ],
                [
                    "mvmt_sts_cd",
                    {}
                ],
                [
                    "status_nm",
                    "Empty Container Release to Shipper"
                ],
                [
                    "skd_dir_cd",
                    ""
                ],
                [
                    "place_nm",
                    "PUSAN, KOREA REPUBLIC OF"
                ],
                [
                    "skd_voy_no",
                    ""
                ]
            ],
            "hashFields": []
        },
        {
            "maxRows": 0,
            "models": [],
            "vslCd": "",
            "no": "2",
            "copNo": "CSEL2303645419",
            "eventDt": "2022-04-05 10:40",
            "vslEngNm": "",
            "placeNm": "PUSAN, KOREA REPUBLIC OF",
            "skdVoyNo": "",
            "yardNm": "HANJIN BUSAN NEW PORT COMPANY(HJNC)",
            "copDtlSeq": "1031",
            "skdDirCd": "",
            "actTpCd": "A",
            "statusNm": "Gate In to Outbound Terminal",
            "statusCd": "FOTMAD",
            "nodCd": "KRPUS14",
            "vvd": "",
            "lloydNo": "",
            "hashColumns": [],
            "hashFields": []
        },
        {
            "maxRows": 0,
            "models": [],
            "vslCd": "HHVT",
            "no": "3",
            "copNo": "CSEL2303645419",
            "eventDt": "2022-04-07 12:09",
            "vslEngNm": "HYUNDAI SINGAPORE",
            "placeNm": "PUSAN, KOREA REPUBLIC OF",
            "skdVoyNo": "0126",
            "yardNm": "HANJIN BUSAN NEW PORT COMPANY(HJNC)",
            "copDtlSeq": "1032",
            "skdDirCd": "E",
            "actTpCd": "A",
            "statusNm": "Loaded on 'HYUNDAI SINGAPORE 126E' at Port of Loading",
            "statusCd": "FLVMLO",
            "nodCd": "KRPUS14",
            "vvd": "HYUNDAI SINGAPORE 126E",
            "lloydNo": "9305685",
            "hashColumns": [],
            "hashFields": []
        },
        {
            "maxRows": 0,
            "models": [],
            "vslCd": "HHVT",
            "no": "4",
            "copNo": "CSEL2303645419",
            "eventDt": "2022-04-07 21:00",
            "vslEngNm": "HYUNDAI SINGAPORE",
            "placeNm": "PUSAN, KOREA REPUBLIC OF",
            "skdVoyNo": "0126",
            "yardNm": "HANJIN BUSAN NEW PORT COMPANY(HJNC)",
            "copDtlSeq": "4033",
            "skdDirCd": "E",
            "actTpCd": "A",
            "statusNm": "'HYUNDAI SINGAPORE 126E' Departure from Port of Loading",
            "statusCd": "FLVMDO",
            "nodCd": "KRPUS14",
            "vvd": "HYUNDAI SINGAPORE 126E",
            "lloydNo": "9305685",
            "hashColumns": [],
            "hashFields": []
        },
        {
            "maxRows": 0,
            "models": [],
            "vslCd": "HHVT",
            "no": "5",
            "copNo": "CSEL2303645419",
            "eventDt": "2022-04-18 16:00",
            "vslEngNm": "HYUNDAI SINGAPORE",
            "placeNm": "VANCOUVER, BC, CANADA",
            "skdVoyNo": "0126",
            "yardNm": "3891 DELTAPORT GCT",
            "copDtlSeq": "4051",
            "skdDirCd": "E",
            "actTpCd": "A",
            "statusNm": "'HYUNDAI SINGAPORE 126E' Arrival at Port of Discharging",
            "statusCd": "FUVMAD",
            "nodCd": "CAVAN01",
            "vvd": "HYUNDAI SINGAPORE 126E",
            "lloydNo": "9305685",
            "hashColumns": [],
            "hashFields": []
        },
        {
            "maxRows": 0,
            "models": [],
            "vslCd": "HHVT",
            "no": "6",
            "copNo": "CSEL2303645419",
            "eventDt": "2022-05-28 03:14",
            "vslEngNm": "HYUNDAI SINGAPORE",
            "placeNm": "VANCOUVER, BC, CANADA",
            "skdVoyNo": "0126",
            "yardNm": "3891 DELTAPORT GCT",
            "copDtlSeq": "4052",
            "skdDirCd": "E",
            "actTpCd": "A",
            "statusNm": "'HYUNDAI SINGAPORE 126E' POD Berthing Destination",
            "statusCd": "FUVMBD",
            "nodCd": "CAVAN01",
            "vvd": "HYUNDAI SINGAPORE 126E",
            "lloydNo": "9305685",
            "hashColumns": [],
            "hashFields": []
        },
        {
            "maxRows": 0,
            "models": [],
            "vslCd": "HHVT",
            "no": "7",
            "copNo": "CSEL2303645419",
            "eventDt": "2022-05-31 21:47",
            "vslEngNm": "HYUNDAI SINGAPORE",
            "placeNm": "VANCOUVER, BC, CANADA",
            "skdVoyNo": "0126",
            "yardNm": "3891 DELTAPORT GCT",
            "copDtlSeq": "6053",
            "skdDirCd": "E",
            "actTpCd": "A",
            "statusNm": "Unloaded from 'HYUNDAI SINGAPORE 126E' at Port of Discharging",
            "statusCd": "FUVMUD",
            "nodCd": "CAVAN01",
            "vvd": "HYUNDAI SINGAPORE 126E",
            "lloydNo": "9305685",
            "hashColumns": [],
            "hashFields": []
        },
        {
            "maxRows": 0,
            "models": [],
            "vslCd": "",
            "no": "8",
            "copNo": "CSEL2303645419",
            "eventDt": "2022-06-03 00:24",
            "vslEngNm": "",
            "placeNm": "VANCOUVER, BC, CANADA",
            "skdVoyNo": "",
            "yardNm": "3891 DELTAPORT GCT",
            "copDtlSeq": "6054",
            "skdDirCd": "",
            "actTpCd": "A",
            "statusNm": "Loaded on rail at inbound rail origin",
            "statusCd": "FIRRLO",
            "nodCd": "CAVAN01",
            "vvd": "",
            "lloydNo": "",
            "hashColumns": [],
            "hashFields": []
        },
        {
            "maxRows": 0,
            "models": [],
            "vslCd": "",
            "no": "9",
            "copNo": "CSEL2303645419",
            "eventDt": "2022-06-03 07:09",
            "vslEngNm": "",
            "placeNm": "VANCOUVER, BC, CANADA",
            "skdVoyNo": "",
            "yardNm": "3891 DELTAPORT GCT",
            "copDtlSeq": "6055",
            "skdDirCd": "",
            "actTpCd": "A",
            "statusNm": "Inbound Rail Departure",
            "statusCd": "FIRRDO",
            "nodCd": "CAVAN01",
            "vvd": "",
            "lloydNo": "",
            "hashColumns": [],
            "hashFields": []
        },
        {
            "maxRows": 0,
            "models": [],
            "vslCd": "",
            "no": "10",
            "copNo": "CSEL2303645419",
            "eventDt": "2022-06-09 18:34",
            "vslEngNm": "",
            "placeNm": "DETROIT, MI, UNITED STATES",
            "skdVoyNo": "",
            "yardNm": "CN RAIL - DETROIT",
            "copDtlSeq": "6071",
            "skdDirCd": "",
            "actTpCd": "A",
            "statusNm": "Inbound Rail Arrival",
            "statusCd": "FIRRAD",
            "nodCd": "USDET65",
            "vvd": "",
            "lloydNo": "",
            "hashColumns": [],
            "hashFields": []
        },
        {
            "maxRows": 0,
            "models": [],
            "vslCd": "",
            "no": "11",
            "copNo": "CSEL2303645419",
            "eventDt": "2022-06-09 21:02",
            "vslEngNm": "",
            "placeNm": "DETROIT, MI, UNITED STATES",
            "skdVoyNo": "",
            "yardNm": "CN RAIL - DETROIT",
            "copDtlSeq": "6072",
            "skdDirCd": "",
            "actTpCd": "A",
            "statusNm": "Unloaded from rail at inbound rail destination",
            "statusCd": "FIRRUD",
            "nodCd": "USDET65",
            "vvd": "",
            "lloydNo": "",
            "hashColumns": [],
            "hashFields": []
        },
        {
            "maxRows": 0,
            "models": [],
            "vslCd": "",
            "no": "12",
            "copNo": "CSEL2303645419",
            "eventDt": "2022-06-10 10:38",
            "vslEngNm": "",
            "placeNm": "DETROIT, MI, UNITED STATES",
            "skdVoyNo": "",
            "yardNm": "CN RAIL - DETROIT",
            "copDtlSeq": "6073",
            "skdDirCd": "",
            "actTpCd": "A",
            "statusNm": "Gate Out from Inbound CY for Delivery to Consignee",
            "statusCd": "FITRDO",
            "nodCd": "USDET65",
            "vvd": "",
            "lloydNo": "",
            "hashColumns": [],
            "hashFields": []
        },
        {
            "maxRows": 0,
            "models": [],
            "vslCd": "",
            "no": "13",
            "copNo": "CSEL2303645419",
            "eventDt": "2022-06-10 12:24",
            "vslEngNm": "",
            "placeNm": "DETROIT, MI, UNITED STATES",
            "skdVoyNo": "",
            "yardNm": "UNIVERSAL INTERMODAL - DETROIT (DEPOT)",
            "copDtlSeq": "6091",
            "skdDirCd": "",
            "actTpCd": "A",
            "statusNm": "Empty Container Returned from Customer",
            "statusCd": "MITYAD",
            "nodCd": "USDET31",
            "vvd": "",
            "lloydNo": "",
            "hashColumns": [],
            "hashFields": []
        }
    ]
}
export const OneyContainerSizeExampleResponse:OneyGetContainerSizeSchema = {
    "TRANS_RESULT_KEY": "S",
    "Exception": "",
    "count": "1",
    "list": [
        {
            "maxRows": 0,
            "models": [],
            "weight": "",
            "copNo": "CCHI0C23914256",
            "blNo": "CHIU06661600",
            "eventDt": "2022-06-14 15:01",
            "cntrTpszCd": "D5",
            "copStsCd": "T",
            "piece": "",
            "sealNo": "",
            "placeNm": "DETROIT, MI, UNITED STATES",
            "yardNm": "UNIVERSAL INTERMODAL - DETROIT (DEPOT)",
            "statusNm": "Empty Container Release to Shipper",
            "bkgNo": "CHIU06661600",
            "poNo": "",
            "yardCd": "USDET31",
            "statusCd": "MOTYDO",
            "cntrNo": "GAOU6642924",
            "cntrTpszNm": "40'DRY HC.",
            "dspbkgNo": "",
            "socFlg": "N",
            "mvmtStsCd": "OP",
            "hashColumns": [
                [
                    "no",
                    {}
                ],
                [
                    "bl_no",
                    "CHIU06661600"
                ],
                [
                    "cntr_no",
                    "GAOU6642924"
                ],
                [
                    "cntr_tpsz_nm",
                    "40'DRY HC."
                ],
                [
                    "ibflag",
                    {}
                ],
                [
                    "event_dt",
                    "2022-06-14 15:01"
                ],
                [
                    "lloyd_no",
                    {}
                ],
                [
                    "cop_dtl_seq",
                    {}
                ],
                [
                    "yard_cd",
                    "USDET31"
                ],
                [
                    "seal_no",
                    ""
                ],
                [
                    "vgm_rcv",
                    {}
                ],
                [
                    "dsp_bkg_no",
                    ""
                ],
                [
                    "place_cd",
                    {}
                ],
                [
                    "status_cd",
                    "MOTYDO"
                ],
                [
                    "cop_no",
                    "CCHI0C23914256"
                ],
                [
                    "nod_cd",
                    {}
                ],
                [
                    "act_tp_cd",
                    {}
                ],
                [
                    "po_no",
                    ""
                ],
                [
                    "soc_flg",
                    "N"
                ],
                [
                    "pagerows",
                    {}
                ],
                [
                    "cntr_tpsz_cd",
                    "D5"
                ],
                [
                    "enbl_flag",
                    {}
                ],
                [
                    "vvd",
                    {}
                ],
                [
                    "weight",
                    ""
                ],
                [
                    "yard_nm",
                    "UNIVERSAL INTERMODAL - DETROIT (DEPOT)"
                ],
                [
                    "cntr_flg",
                    {}
                ],
                [
                    "bkg_no",
                    "CHIU06661600"
                ],
                [
                    "vsl_eng_nm",
                    {}
                ],
                [
                    "vsl_cd",
                    {}
                ],
                [
                    "cop_sts_cd",
                    "T"
                ],
                [
                    "piece",
                    ""
                ],
                [
                    "mvmt_sts_cd",
                    "OP"
                ],
                [
                    "status_nm",
                    "Empty Container Release to Shipper"
                ],
                [
                    "skd_dir_cd",
                    {}
                ],
                [
                    "place_nm",
                    "DETROIT, MI, UNITED STATES"
                ],
                [
                    "skd_voy_no",
                    {}
                ]
            ],
            "hashFields": []
        }
    ]
}