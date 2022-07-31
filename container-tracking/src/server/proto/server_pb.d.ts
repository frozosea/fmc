// package: tracking
// file: server.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class Request extends jspb.Message { 
    getNumber(): string;
    setNumber(value: string): Request;
    getScac(): Scac;
    setScac(value: Scac): Request;
    getCountry(): Country;
    setCountry(value: Country): Request;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Request.AsObject;
    static toObject(includeInstance: boolean, msg: Request): Request.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Request, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Request;
    static deserializeBinaryFromReader(message: Request, reader: jspb.BinaryReader): Request;
}

export namespace Request {
    export type AsObject = {
        number: string,
        scac: Scac,
        country: Country,
    }
}

export class InfoAboutMoving extends jspb.Message { 
    getTime(): number;
    setTime(value: number): InfoAboutMoving;
    getOperationName(): string;
    setOperationName(value: string): InfoAboutMoving;
    getLocation(): string;
    setLocation(value: string): InfoAboutMoving;
    getVessel(): string;
    setVessel(value: string): InfoAboutMoving;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): InfoAboutMoving.AsObject;
    static toObject(includeInstance: boolean, msg: InfoAboutMoving): InfoAboutMoving.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: InfoAboutMoving, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): InfoAboutMoving;
    static deserializeBinaryFromReader(message: InfoAboutMoving, reader: jspb.BinaryReader): InfoAboutMoving;
}

export namespace InfoAboutMoving {
    export type AsObject = {
        time: number,
        operationName: string,
        location: string,
        vessel: string,
    }
}

export class TrackingByContainerNumberResponse extends jspb.Message { 
    getContainer(): string;
    setContainer(value: string): TrackingByContainerNumberResponse;
    getContainerSize(): string;
    setContainerSize(value: string): TrackingByContainerNumberResponse;
    getScac(): Scac;
    setScac(value: Scac): TrackingByContainerNumberResponse;
    clearInfoAboutMovingList(): void;
    getInfoAboutMovingList(): Array<InfoAboutMoving>;
    setInfoAboutMovingList(value: Array<InfoAboutMoving>): TrackingByContainerNumberResponse;
    addInfoAboutMoving(value?: InfoAboutMoving, index?: number): InfoAboutMoving;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): TrackingByContainerNumberResponse.AsObject;
    static toObject(includeInstance: boolean, msg: TrackingByContainerNumberResponse): TrackingByContainerNumberResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: TrackingByContainerNumberResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): TrackingByContainerNumberResponse;
    static deserializeBinaryFromReader(message: TrackingByContainerNumberResponse, reader: jspb.BinaryReader): TrackingByContainerNumberResponse;
}

export namespace TrackingByContainerNumberResponse {
    export type AsObject = {
        container: string,
        containerSize: string,
        scac: Scac,
        infoAboutMovingList: Array<InfoAboutMoving.AsObject>,
    }
}

export class TrackingByBillNumberResponse extends jspb.Message { 
    getBillno(): string;
    setBillno(value: string): TrackingByBillNumberResponse;
    getScac(): Scac;
    setScac(value: Scac): TrackingByBillNumberResponse;
    clearInfoAboutMovingList(): void;
    getInfoAboutMovingList(): Array<InfoAboutMoving>;
    setInfoAboutMovingList(value: Array<InfoAboutMoving>): TrackingByBillNumberResponse;
    addInfoAboutMoving(value?: InfoAboutMoving, index?: number): InfoAboutMoving;
    getEtaFinalDelivery(): number;
    setEtaFinalDelivery(value: number): TrackingByBillNumberResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): TrackingByBillNumberResponse.AsObject;
    static toObject(includeInstance: boolean, msg: TrackingByBillNumberResponse): TrackingByBillNumberResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: TrackingByBillNumberResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): TrackingByBillNumberResponse;
    static deserializeBinaryFromReader(message: TrackingByBillNumberResponse, reader: jspb.BinaryReader): TrackingByBillNumberResponse;
}

export namespace TrackingByBillNumberResponse {
    export type AsObject = {
        billno: string,
        scac: Scac,
        infoAboutMovingList: Array<InfoAboutMoving.AsObject>,
        etaFinalDelivery: number,
    }
}

export enum Scac {
    FESO = 0,
    SKLU = 1,
    MAEU = 2,
    COSU = 3,
    KMTU = 4,
    ONEY = 5,
    SITC = 6,
    MSCU = 7,
    HALU = 8,
    ZHGU = 9,
    AUTO = 10,
}

export enum Country {
    RU = 0,
    OTHER = 1,
}
