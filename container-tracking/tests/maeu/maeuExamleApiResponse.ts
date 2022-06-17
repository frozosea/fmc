import {MaerskApiResponseSchema} from "../../src/trackTrace/TrackingByContainerNumber/maeu/maerskApiResponseSchema";

export const maeuExamleApiResponse: MaerskApiResponseSchema = {
    "isContainerSearch": true,
    "origin": {
        "terminal": "Laem Chabang Terminal PORT D1",
        "geo_site": "9QYVWASPIUQBZ",
        "city": "Laem Chabang",
        "state": "",
        "country": "Thailand",
        "country_code": "TH",
        "geoid_city": "0NTCCB4ENSFX8",
        "site_type": "TERMINAL"
    },
    "destination": {
        "terminal": "",
        "geo_site": "Unknown",
        "city": "Spartanburg",
        "state": "South Carolina",
        "country": "United States",
        "country_code": "US",
        "geoid_city": "09L62FUTQWHHB",
        "site_type": ""
    },
    "containers": [
        {
            "container_num": "MSKU6874333",
            "container_size": "40",
            "container_type": "Dry",
            "iso_code": "42G0",
            "operator": "MAEU",
            "locations": [
                {
                    "terminal": "Win Win Container Depot",
                    "geo_site": "1GP8G188VBR2G",
                    "city": "Laem Chabang",
                    "state": "",
                    "country": "Thailand",
                    "country_code": "TH",
                    "geoid_city": "0NTCCB4ENSFX8",
                    "site_type": "DEPOT",
                    "events": [
                        {
                            "activity": "GATE-OUT-EMPTY",
                            "stempty": true,
                            "actfor": "EXP",
                            "vessel_name": "MSC SVEVA",
                            "voyage_num": "204E",
                            "vessel_num": "Z77",
                            "actual_time": "2022-03-29T16:42:00.000",
                            "rkem_move": "GATE-OUT",
                            "is_cancelled": false,
                            "is_current": false
                        }
                    ]
                },
                {
                    "terminal": "Laem Chabang Terminal PORT D1",
                    "geo_site": "9QYVWASPIUQBZ",
                    "city": "Laem Chabang",
                    "state": "",
                    "country": "Thailand",
                    "country_code": "TH",
                    "geoid_city": "0NTCCB4ENSFX8",
                    "site_type": "TERMINAL",
                    "events": [
                        {
                            "activity": "GATE-IN",
                            "stempty": false,
                            "actfor": "EXP",
                            "vessel_name": "MSC SVEVA",
                            "voyage_num": "204E",
                            "vessel_num": "Z77",
                            "expected_time": "2022-04-16T14:00:00.000",
                            "actual_time": "2022-03-30T10:02:00.000",
                            "rkem_move": "GATE-IN",
                            "is_cancelled": false,
                            "is_current": false
                        },
                        {
                            "activity": "LOAD",
                            "stempty": false,
                            "actfor": "",
                            "vessel_name": "MSC SVEVA",
                            "voyage_num": "204E",
                            "vessel_num": "Z77",
                            "expected_time": "2022-04-16T14:00:00.000",
                            "actual_time": "2022-04-14T00:15:00.000",
                            "rkem_move": "LOAD",
                            "is_cancelled": false,
                            "is_current": false
                        }
                    ]
                },
                {
                    "terminal": "YANGSHAN SGH GUANDONG TERMINAL",
                    "geo_site": "37O5HQ17XCL3X",
                    "city": "Shanghai",
                    "state": "Shanghai",
                    "country": "China",
                    "country_code": "CN",
                    "geoid_city": "2IW9P6J7XAW72",
                    "site_type": "TERMINAL",
                    "events": [
                        {
                            "activity": "DISCHARG",
                            "stempty": false,
                            "actfor": "",
                            "vessel_name": "MSC SVEVA",
                            "voyage_num": "204E",
                            "vessel_num": "Z77",
                            "expected_time": "2022-04-25T08:00:00.000",
                            "actual_time": "2022-04-25T12:49:00.000",
                            "rkem_move": "DISCHARG",
                            "is_cancelled": false,
                            "is_current": false
                        },
                        {
                            "activity": "LOAD",
                            "stempty": false,
                            "actfor": "",
                            "vessel_name": "ZIM WILMINGTON",
                            "voyage_num": "006E",
                            "vessel_num": "U5T",
                            "expected_time": "2022-05-02T18:30:00.000",
                            "actual_time": "2022-05-02T01:11:00.000",
                            "rkem_move": "LOAD",
                            "is_cancelled": false,
                            "is_current": false
                        }
                    ]
                },
                {
                    "terminal": "Charleston Wando Welch terminal N59",
                    "geo_site": "1ML38I7Q8BBKU",
                    "city": "Charleston",
                    "state": "South Carolina",
                    "country": "United States",
                    "country_code": "US",
                    "geoid_city": "3RSB4DDP23AM7",
                    "site_type": "TERMINAL",
                    "events": [
                        {
                            "activity": "DISCHARG",
                            "stempty": false,
                            "actfor": "",
                            "vessel_name": "ZIM WILMINGTON",
                            "voyage_num": "006E",
                            "vessel_num": "U5T",
                            "expected_time": "2022-06-11T07:00:00.000",
                            "actual_time": "2022-06-11T13:54:00.000",
                            "rkem_move": "DISCHARG",
                            "is_cancelled": false,
                            "is_current": true
                        },
                        {
                            "activity": "GATE-OUT",
                            "stempty": false,
                            "actfor": "IMP",
                            "vessel_name": "",
                            "voyage_num": "",
                            "vessel_num": "",
                            "expected_time": "2022-06-14T10:00:00.000",
                            "is_current": false
                        }
                    ]
                }
            ],
            "eta_final_delivery": "2022-06-11T13:54:00.000",
            "latest": {
                "actual_time": "2022-06-11T13:54:00.000",
                "activity": "DISCHARG",
                "stempty": false,
                "actfor": "",
                "geo_site": "1ML38I7Q8BBKU",
                "city": "Charleston",
                "state": "South Carolina",
                "country": "United States",
                "country_code": "US"
            },
            "status": "IN-PROGRESS"
        }
    ]
}
export const MaeuinfoAboutMoving = [
    {
        "location": "Win Win Container Depot",
        "operationName": "GATE-OUT-EMPTY",
        "time": 1648572120000,
        "vessel": "MSC SVEVA"
    },
    {
        "location": "Laem Chabang Terminal PORT D1",
        "operationName": "GATE-IN",
        "time": 1648634520000,
        "vessel": "MSC SVEVA"
    },
    {
        "location": "Laem Chabang Terminal PORT D1",
        "operationName": "LOAD",
        "time": 1649895300000,
        "vessel": "MSC SVEVA"
    },
    {
        "location": "YANGSHAN SGH GUANDONG TERMINAL",
        "operationName": "DISCHARG",
        "time": 1650890940000,
        "vessel": "MSC SVEVA"
    },
    {
        "location": "YANGSHAN SGH GUANDONG TERMINAL",
        "operationName": "LOAD",
        "time": 1651453860000,
        "vessel": "ZIM WILMINGTON"
    },
    {
        "location": "Charleston Wando Welch terminal N59",
        "operationName": "DISCHARG",
        "time": 1654955640000,
        "vessel": "ZIM WILMINGTON"
    },
    {
        "location": "Charleston Wando Welch terminal N59",
        "operationName": "GATE-OUT",
        "time": 1655200800000,
        "vessel": ""
    }
]
export const expectedMaeuReadyObject = {
    'container': 'MSKU6874333',
    'containerSize': '40DRY',
    'scac': 'MAEU',
    'infoAboutMoving': [...MaeuinfoAboutMoving, {
        'time': 1654955640000,
        'operationName': 'ETA',
        'location': 'Spartanburg',
        'vessel': ''
    }]
}
