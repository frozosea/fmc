import {IUnlocodesRepo, UnlocodeObject} from "../../src/trackTrace/TrackingByContainerNumber/sklu/unlocodesRepo";

export const skluApiResponseExample = [
    {
        "BKNO": "SNKO101220501450",
        "CNTR": "20'x15",
        "POR": null,
        "POL": "VNSGN",
        "POD": "JPHSM",
        "DLV": "JPHSM",
        "VSL": "HAHM",
        "VYG": "2205N",
        "ETD": "2022-06-07",
        "ETA": "2022-06-17"
    },
    {
        "BKNO": "SNKO190220401055",
        "CNTR": "20'x11",
        "POR": "THLCH",
        "POL": "THLCH",
        "POD": "VNSGN",
        "DLV": "VNSGN",
        "VSL": "SWSG",
        "VYG": "0184N",
        "ETD": "2022-04-29",
        "ETA": "2022-05-03"
    },
    {
        "BKNO": "SNKO011220302714",
        "CNTR": "20'x3",
        "POR": null,
        "POL": "KRPUS",
        "POD": "THLCH",
        "DLV": "THLCH",
        "VSL": "PDN4",
        "VYG": "2203S",
        "ETD": "2022-03-28",
        "ETA": "2022-04-07"
    }
]
export const expectedInfoAboutMoving = [
    {
        "location": "SINOKOR TAM CANG CAT LAI Depot",
        "operationName": "Pickup (1/1)",
        "time": 1653906120000,
        "vessel": " "
    },
    {
        "location": "CAT LAI",
        "operationName": "Return (1/1)",
        "time": 1653961680000,
        "vessel": " "
    },
    {
        "location": "CAT LAI",
        "operationName": "Departure",
        "time": 1654567200000,
        "vessel": "HEUNG-A HOCHIMINH / 2205N"
    },
    {
        "location": "BPTS",
        "operationName": "Arrival(T/S) (Scheduled)",
        "time": 1655085600000,
        "vessel": "HEUNG-A HOCHIMINH / 2205N"
    },
    {
        "location": "BPTS",
        "operationName": "Departure(T/S) (Scheduled)",
        "time": 1655283600000,
        "vessel": "HEUNG-A ULSAN / 2256E"
    },
    {
        "location": "HOSOSHIMA TERMINAL(SHIRAHMA #14)",
        "operationName": "Arrival (Scheduled)",
        "time": 1655452800000,
        "vessel": "HEUNG-A ULSAN / 2256E"
    }
]

export const expectedResult = {
    'container': 'TEMU2094051',
    'containerSize': "20'x15",
    'scac': 'SKLU',
    'infoAboutMoving': [...expectedInfoAboutMoving, {
        "location": "Hososhima",
        "operationName": "ETA",
        "time": 1655424000000,
        "vessel": " "
    }]
}


export class UnlocodesRepoMoch implements IUnlocodesRepo {
    public async getUnlocode(unlocode: string): Promise<string> {
        return "Hososhima"
    }

    public async addUnlocode(obj: UnlocodeObject): Promise<void> {
    }
}