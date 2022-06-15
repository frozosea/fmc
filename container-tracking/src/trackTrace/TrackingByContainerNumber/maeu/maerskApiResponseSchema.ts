interface _BaseMaerskApiStruct {
    actual_time?: string
    activity?: string,
    stempty?: boolean | string,
    terminal?: string,
    geo_site: string,
    actfor?: string
    city: string,
    state: string,
    country: string,
    country_code: string,
    geoid_city?: string,
    site_type?: string
}

interface EvenetsMaerskStruct {
    activity?: string
    stempty?: boolean
    actfor?: string
    vessel_name: string
    voyage_num: string
    vessel_num: string
    actual_time?: string
    expected_time?: string
    rkem_move?: string
    is_cancelled?: boolean
    is_current: boolean
}

interface locationsStruct extends _BaseMaerskApiStruct {
    events: EvenetsMaerskStruct[]
}

export interface MaerskApiResponseSchema {
    isContainerSearch: boolean,
    origin: _BaseMaerskApiStruct,
    destination: _BaseMaerskApiStruct,
    containers: [
        {
            container_num: string,
            container_size: string,
            container_type: string,
            iso_code: string,
            operator: "MAEU"
            locations: locationsStruct[]
            eta_final_delivery: string
            latest: _BaseMaerskApiStruct
            status: string
        }
    ]

}