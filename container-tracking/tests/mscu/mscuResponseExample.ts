import {MscuApiResponseSchema} from "../../src/trackTrace/TrackingByContainerNumber/mscu/mscuApiResponseSchema";


export const mscuResponseExample: MscuApiResponseSchema = {
    "IsSuccess": true,
    "Data": {
        "TrackingType": "Container",
        "TrackingTitle": "CONTAINER NUMBER:",
        "TrackingNumber": "MEDU3170580",
        "CurrentDate": "12/06/2022",
        "PriceCalculationLabel": "* Price calculation date is indicative. Please contact your local MSC office to verify this information.",
        "TrackingResultsLabel": "Tracking results provided by MSC on 12.06.2022 at 10:00 Central Europe Standard Time",
        "BillOfLadings": [
            {
                "BillOfLadingNumber": "",
                "NumberOfContainers": 1,
                "GeneralTrackingInfo": {
                    "ShippedFrom": "",
                    "ShippedTo": "",
                    "PortOfLoad": "",
                    "PortOfDischarge": "",
                    "Transshipments": [],
                    "PriceCalculationDate": "",
                    "FinalPodEtaDate": "12/06/2022"
                },
                "ContainersInfo": [
                    {
                        "ContainerNumber": "MEDU3170580",
                        "PodEtaDate": "",
                        "ContainerType": "20' DRY VAN",
                        "LatestMove": "CHONGQING, CN",
                        "Events": [
                            {
                                "Order": 1,
                                "Date": "10/06/2022",
                                "Location": "CHONGQING, CN",
                                "Description": "Export at barge yard",
                                "Detail": [
                                    "LADEN"
                                ]
                            },
                            {
                                "Order": 0,
                                "Date": "09/06/2022",
                                "Location": "CHONGQING, CN",
                                "Description": "Empty to Shipper",
                                "Detail": [
                                    "EMPTY"
                                ]
                            }
                        ]
                    }
                ]
            }
        ]
    }
}