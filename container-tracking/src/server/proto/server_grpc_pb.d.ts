// package: tracking
// file: server.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "grpc";
import * as server_pb from "./server_pb";

interface ITrackingByContainerNumberService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    trackByContainerNumber: ITrackingByContainerNumberService_ITrackByContainerNumber;
}

interface ITrackingByContainerNumberService_ITrackByContainerNumber extends grpc.MethodDefinition<server_pb.Request, server_pb.TrackingByContainerNumberResponse> {
    path: "/tracking.TrackingByContainerNumber/TrackByContainerNumber";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<server_pb.Request>;
    requestDeserialize: grpc.deserialize<server_pb.Request>;
    responseSerialize: grpc.serialize<server_pb.TrackingByContainerNumberResponse>;
    responseDeserialize: grpc.deserialize<server_pb.TrackingByContainerNumberResponse>;
}

export const TrackingByContainerNumberService: ITrackingByContainerNumberService;

export interface ITrackingByContainerNumberServer {
    trackByContainerNumber: grpc.handleUnaryCall<server_pb.Request, server_pb.TrackingByContainerNumberResponse>;
}

export interface ITrackingByContainerNumberClient {
    trackByContainerNumber(request: server_pb.Request, callback: (error: grpc.ServiceError | null, response: server_pb.TrackingByContainerNumberResponse) => void): grpc.ClientUnaryCall;
    trackByContainerNumber(request: server_pb.Request, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: server_pb.TrackingByContainerNumberResponse) => void): grpc.ClientUnaryCall;
    trackByContainerNumber(request: server_pb.Request, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: server_pb.TrackingByContainerNumberResponse) => void): grpc.ClientUnaryCall;
}

export class TrackingByContainerNumberClient extends grpc.Client implements ITrackingByContainerNumberClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public trackByContainerNumber(request: server_pb.Request, callback: (error: grpc.ServiceError | null, response: server_pb.TrackingByContainerNumberResponse) => void): grpc.ClientUnaryCall;
    public trackByContainerNumber(request: server_pb.Request, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: server_pb.TrackingByContainerNumberResponse) => void): grpc.ClientUnaryCall;
    public trackByContainerNumber(request: server_pb.Request, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: server_pb.TrackingByContainerNumberResponse) => void): grpc.ClientUnaryCall;
}

interface ITrackingByBillNumberService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    trackByBillNumber: ITrackingByBillNumberService_ITrackByBillNumber;
}

interface ITrackingByBillNumberService_ITrackByBillNumber extends grpc.MethodDefinition<server_pb.Request, server_pb.TrackingByBillNumberResponse> {
    path: "/tracking.TrackingByBillNumber/TrackByBillNumber";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<server_pb.Request>;
    requestDeserialize: grpc.deserialize<server_pb.Request>;
    responseSerialize: grpc.serialize<server_pb.TrackingByBillNumberResponse>;
    responseDeserialize: grpc.deserialize<server_pb.TrackingByBillNumberResponse>;
}

export const TrackingByBillNumberService: ITrackingByBillNumberService;

export interface ITrackingByBillNumberServer {
    trackByBillNumber: grpc.handleUnaryCall<server_pb.Request, server_pb.TrackingByBillNumberResponse>;
}

export interface ITrackingByBillNumberClient {
    trackByBillNumber(request: server_pb.Request, callback: (error: grpc.ServiceError | null, response: server_pb.TrackingByBillNumberResponse) => void): grpc.ClientUnaryCall;
    trackByBillNumber(request: server_pb.Request, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: server_pb.TrackingByBillNumberResponse) => void): grpc.ClientUnaryCall;
    trackByBillNumber(request: server_pb.Request, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: server_pb.TrackingByBillNumberResponse) => void): grpc.ClientUnaryCall;
}

export class TrackingByBillNumberClient extends grpc.Client implements ITrackingByBillNumberClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public trackByBillNumber(request: server_pb.Request, callback: (error: grpc.ServiceError | null, response: server_pb.TrackingByBillNumberResponse) => void): grpc.ClientUnaryCall;
    public trackByBillNumber(request: server_pb.Request, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: server_pb.TrackingByBillNumberResponse) => void): grpc.ClientUnaryCall;
    public trackByBillNumber(request: server_pb.Request, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: server_pb.TrackingByBillNumberResponse) => void): grpc.ClientUnaryCall;
}
