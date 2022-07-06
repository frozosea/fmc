interface list1Event {
    blNo: string
    delen: string
    polen: string
    pol: string
    del: string
}

interface OneSitcTrackingEvent {
    rowNo: string
    totalCount: string
    containerNo: string
    sealNo: string
    voyageNo: string
    cntrType: string
    quanity: string
    cntrSize: string
    weight: string
    currentport: string
    movementname: string
    movementnameen: string
}

interface list2Event {
    vesselName: string | null
    voyageNo: string | null
    voyageLeg: string | null
    portFrom: string | null
    portFromName: string | null
    portTo: string | null
    portToName: string | null
    eta: string | null
    cta: string | null
    ata: string | null
    ccta: string | null
    etd: string | null
    cctd: string | null
    ctd: string | null
    atd: string | null
    cctb: string | null
    agtb: string | null
    ctb: string | null
    atb: string | null
    agtd: string | null
    agta: string | null
}

interface movementCode {
    movementStatus: string
}

export default interface SitcBillNumberApiResponseSchema {
    code: number
    msg: string
    data: {
        list1: list1Event[]
        containerNo: string | null
        blNo: string
        list3: OneSitcTrackingEvent[]
        list2: list2Event[]
        movementcode: movementCode[]
    }

}