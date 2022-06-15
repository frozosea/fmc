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
export const expectedInfoAboutMoving = [{
    'time': 1653870120000,
    'operationName': 'Pickup (1/1)',
    'location': 'SINOKOR TAM CANG CAT LAI Depot',
    'vessel': ''
}, {
    'time': 1653925680000,
    'operationName': 'Return (1/1)',
    'location': 'CAT LAI',
    'vessel': ''
},
    {
        'time': 1654531200000,
        'operationName': 'Departure',
        'location': 'CAT LAI',
        'vessel': 'HEUNG-A HOCHIMINH / 2205N'
    }, {
        'time': 1655049600000,
        'operationName': 'Arrival(T/S) (Scheduled)',
        'location': 'BPTS',
        'vessel': 'HEUNG-A HOCHIMINH / 2205N'
    }, {
        'time': 1655247600000,
        'operationName': 'Departure(T/S) (Scheduled)',
        'location': 'BPTS',
        'vessel': 'HEUNG-A ULSAN / 2256E'
    }, {
        'time': 1655416800000,
        'operationName': 'Arrival (Scheduled)',
        'location': 'HOSOSHIMA TERMINAL(SHIRAHMA #14)',
        'vessel': 'HEUNG-A ULSAN / 2256E'
    }]
export const expectedResult = {
    'container': 'TEMU2094051',
    'containerSize': "20'x15",
    'scac': 'SKLU',
    'infoAboutMoving': [...expectedInfoAboutMoving, {
        "location": "Hososhima",
        "operationName": "ETA",
        "time": 1655424000000,
        "vessel": ""
    }]
}


export class UnlocodesRepoMoch implements IUnlocodesRepo {
    public async getUnlocode(unlocode: string): Promise<string> {
        return "Hososhima"
    }

    public async addUnlocode(obj: UnlocodeObject): Promise<void> {
    }
}