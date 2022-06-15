interface _SitcInfoAboutMoving {
    containerNo: string
    movementName: string
    movementCode: string
    movementNameEn: string
    eventPort: string
    eventDate: string
    vesselCode: string
    voyageNo: string
}

export interface SitcContainerTrackingApiResponseSchema {
    code: number
    msg: string
    data: {
        list: _SitcInfoAboutMoving[]
    }
}