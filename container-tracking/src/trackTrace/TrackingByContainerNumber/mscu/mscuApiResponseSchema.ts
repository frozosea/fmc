interface _Events {
    Order?: number
    Date: string
    Location: string
    Description: string
    Detail: string[]
}

interface _ContainerInfo {
    Order?: number
    ContainerNumber?: string
    PodEtaDate?: string
    ContainerType: string
    LatestMove: string
    Events: _Events[]
}

interface _BillOfLandings {
    BillOfLadingNumber: string
    NumberOfContainers: number
    GeneralTrackingInfo: {
        ShippedFrom: string
        ShippedTo: string
        PortOfLoad: string
        PortOfDischarge: string
        Transshipments: string[]
        PriceCalculationDate: string
        FinalPodEtaDate: string
    }
    ContainersInfo: _ContainerInfo[]
}

export interface MscuApiResponseSchema {
    IsSuccess: boolean
    Data: {
        TrackingType: string
        TrackingTitle: string
        TrackingNumber: string
        CurrentDate: string
        PriceCalculationLabel: string
        TrackingResultsLabel: string
        BillOfLadings: _BillOfLandings[]
    }
}