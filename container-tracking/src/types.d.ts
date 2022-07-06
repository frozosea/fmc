/// <reference types="node" />

interface OneTrackingEvent {
    time: number
    location: string
    operationName: string
    vessel: string

}

export interface TrackingContainerResponse {
    container: string
    containerSize: string
    scac: SCAC_TYPE
    infoAboutMoving: OneTrackingEvent[]
}

export interface ITrackingArgs {
    number: string
}


declare type FESO = "FESO"
declare type SKLU = "SKLU"
declare type MAEU = "MAEU"
declare type COSU = "COSU"
declare type KMTU = "KMTU"
declare type ONEY = "ONEY"
declare type SITC = "SITC"
declare type MSCU = "MSCU"
declare type AUTO = "AUTO"
declare type HALU = "HALU"
declare type SCAC_TYPE = FESO | SKLU | MAEU | COSU | KMTU | ONEY | SITC | MSCU | HALU | AUTO


export interface TrackingArgsWithScac extends ITrackingArgs {
    scac: SCAC_TYPE,
    country: COUNTRY_TYPE
}

declare global {
    interface String {
        capitalizeFirstLetter(): string
    }
}
type BkgNo = string
type CopNo = string

type COUNTRY_TYPE = "RU" | "OTHER"

export interface ITrackingByBillNumberResponse {
    billNo: string
    scac: SCAC_TYPE
    infoAboutMoving: OneTrackingEvent[]
    etaFinalDelivery: number
}