export interface FesoLastEventsSchema {
    time: string,
    operation?: string
    operationName: string,
    operationNameLatin: string,
    vessel?: string,
    locId: number,
    locName: string,
    locNameLatin: string,
    etCode?: string | null,
    transportType?: string | null
}

export interface FesoApiResponse {
    container: string,
    time: string,
    containerCTCode: string,
    containerOwner: string,
    latLng: string,
    lastEvents: FesoLastEventsSchema[]
}



export interface FesoApiFullResponseSchema {
    data: {
        tracking: {
            data: {
                requestKey: string,
                containers: FesoApiResponse | string[]
                missing: [],
                __typename: string
            },
            __typename: string
        }
    }
}