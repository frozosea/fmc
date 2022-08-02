interface zhguEventSchema {
    createdByUser: null | string
    createdOffice: null | string
    createdDtmLoc: null | string
    createdTimeZone: null | string
    updatedByUser: null | string
    updatedOffice: null | string
    updatedDtmLoc: null | string
    updatedTimeZone: null | string
    recordVersion: null | string
    rowStatus: 2
    principalGroupCode: null | string
    blNo: string
    tripNumber: string
    lineType: string
    vesselName: string
    voyage: string
    portFrom: string
    portFromName: string
    portTo: string
    portToName: string
    etd: string
    atd: string
    eta: string
    ata: string
    departureTime: string
    arrivalTime: string
}

export default interface ZhguApiResponseSchema {
    "backCode": string,
    "backMessage": string,
    "object": zhguEventSchema[]
}

export interface ZhguCheckBookApiResp{
    backCode: string
    backMessage: string
    object: any[]
}