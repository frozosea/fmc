// package: tracking
// file: server.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class Request extends jspb.Message { 
    getContainer(): string;
    setContainer(value: string): Request;
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
        container: string,
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

export class Response extends jspb.Message { 
    getContainer(): string;
    setContainer(value: string): Response;
    getContainerSize(): string;
    setContainerSize(value: string): Response;
    getScac(): Scac;
    setScac(value: Scac): Response;
    clearInfoAboutMovingList(): void;
    getInfoAboutMovingList(): Array<InfoAboutMoving>;
    setInfoAboutMovingList(value: Array<InfoAboutMoving>): Response;
    addInfoAboutMoving(value?: InfoAboutMoving, index?: number): InfoAboutMoving;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Response.AsObject;
    static toObject(includeInstance: boolean, msg: Response): Response.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Response, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Response;
    static deserializeBinaryFromReader(message: Response, reader: jspb.BinaryReader): Response;
}

export namespace Response {
    export type AsObject = {
        container: string,
        containerSize: string,
        scac: Scac,
        infoAboutMovingList: Array<InfoAboutMoving.AsObject>,
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
    AUTO = 8,
}

export enum Country {
    RU = 0,
    OTHER = 1,
}
