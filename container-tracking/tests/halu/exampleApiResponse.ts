export const HaluExampleApiResponse = [
    {
        "BKNO": "HASLC01220602499",
        "CNTR": "20'x4",
        "POR": null,
        "POL": "CNSHA",
        "POD": "RUVYP",
        "DLV": "RUVYP",
        "VSL": "SWCG",
        "VYG": "9999N",
        "ETD": "2022-06-19",
        "ETA": "2022-07-04"
    }
]
export const expectedInfoAboutMoving = [
    {
        time: 1655645400000,
        operationName: 'Departure',
        location: 'WAIGAOQIAO PIER #5',
        vessel: 'SURABAYA VOYAGER / 2231W'
    },
    {
        time: 1655645400000,
        operationName: 'Arrival(T/S)',
        location: 'WAIGAOQIAO PIER #5',
        vessel: 'SAWASDEE CHITTAGONG / 9999N'
    },
    {
        time: 1655848800000,
        operationName: 'Departure(T/S) (Scheduled)',
        location: 'BUSAN PORT TERMINAL',
        vessel: 'SAWASDEE CHITTAGONG / 9999N'
    },
    {
        time: 1656759600000,
        operationName: 'Arrival (Scheduled)',
        location: 'PUSAN INTERNATIONAL TERMINAL',
        vessel: 'SURABAYA VOYAGER / 2231W'
    },
    {
        operationName: 'ETA',
        time: 1656892800000,
        location: 'Hososhima',
        vessel: ''
    }
]