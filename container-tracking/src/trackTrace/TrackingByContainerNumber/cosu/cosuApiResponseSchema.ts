interface _ContainerInfo {
    containerUuid: string
    containerNumber: string
    containerType: string
    grossWeight: string
    piecesNumber: number
    label: string
    sealNumber: string
    location: string
    locationDateTime: string
    transportation: string
    flag: string | number
    railRef: string | null
    inlandMvId: string | number
    containerLocation: string | number
    isShow: boolean
    polEtd: string | null
    polAtd: string | null
    podEta: string | null
    podAta: string | null
    transportId: string | null | number
    pol: string | null
    pod: string | null
    hsCode: any
    isNorthAmericaRails: boolean
}

export interface CosuInfoAboutMoving {
    uuid: string | null
    containerNumber: string | null
    containerNumberStatus: string | null
    location: string | null
    timeOfIssue: string | null
    transportation: string | null
    polEtd: string | null
    polAtd: string | null
    podEta: string | null
    podAta: string | null
    transportId: string | null
    pol: string | null
    pod: string | null
}


export interface CosuApiResponseSchema {
    code: number | string
    message: string
    data: {
        content: {
            containers: [{
                container: _ContainerInfo
                containerCircleStatus: CosuInfoAboutMoving[]
                containerHistorys: CosuInfoAboutMoving[]
            }?]
            notFound: string
        }
    }
}

export interface EtaResponseSchema {
    code: string
    message: string
    data: {
        content: string
    }
}
