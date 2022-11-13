// package: tracking
// file: tracking.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as tracking_pb from "./tracking_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";

interface ITrackingByContainerNumberService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    trackByContainerNumber: ITrackingByContainerNumberService_ITrackByContainerNumber;
}

interface ITrackingByContainerNumberService_ITrackByContainerNumber extends grpc.MethodDefinition<tracking_pb.Request, tracking_pb.TrackingByContainerNumberResponse> {
    path: "/tracking.TrackingByContainerNumber/TrackByContainerNumber";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<tracking_pb.Request>;
    requestDeserialize: grpc.deserialize<tracking_pb.Request>;
    responseSerialize: grpc.serialize<tracking_pb.TrackingByContainerNumberResponse>;
    responseDeserialize: grpc.deserialize<tracking_pb.TrackingByContainerNumberResponse>;
}

export const TrackingByContainerNumberService: ITrackingByContainerNumberService;

export interface ITrackingByContainerNumberServer {
    trackByContainerNumber: grpc.handleUnaryCall<tracking_pb.Request, tracking_pb.TrackingByContainerNumberResponse>;
}

export interface ITrackingByContainerNumberClient {
    trackByContainerNumber(request: tracking_pb.Request, callback: (error: grpc.ServiceError | null, response: tracking_pb.TrackingByContainerNumberResponse) => void): grpc.ClientUnaryCall;
    trackByContainerNumber(request: tracking_pb.Request, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: tracking_pb.TrackingByContainerNumberResponse) => void): grpc.ClientUnaryCall;
    trackByContainerNumber(request: tracking_pb.Request, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: tracking_pb.TrackingByContainerNumberResponse) => void): grpc.ClientUnaryCall;
}

export class TrackingByContainerNumberClient extends grpc.Client implements ITrackingByContainerNumberClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public trackByContainerNumber(request: tracking_pb.Request, callback: (error: grpc.ServiceError | null, response: tracking_pb.TrackingByContainerNumberResponse) => void): grpc.ClientUnaryCall;
    public trackByContainerNumber(request: tracking_pb.Request, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: tracking_pb.TrackingByContainerNumberResponse) => void): grpc.ClientUnaryCall;
    public trackByContainerNumber(request: tracking_pb.Request, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: tracking_pb.TrackingByContainerNumberResponse) => void): grpc.ClientUnaryCall;
}

interface ITrackingByBillNumberService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    trackByBillNumber: ITrackingByBillNumberService_ITrackByBillNumber;
}

interface ITrackingByBillNumberService_ITrackByBillNumber extends grpc.MethodDefinition<tracking_pb.Request, tracking_pb.TrackingByBillNumberResponse> {
    path: "/tracking.TrackingByBillNumber/TrackByBillNumber";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<tracking_pb.Request>;
    requestDeserialize: grpc.deserialize<tracking_pb.Request>;
    responseSerialize: grpc.serialize<tracking_pb.TrackingByBillNumberResponse>;
    responseDeserialize: grpc.deserialize<tracking_pb.TrackingByBillNumberResponse>;
}

export const TrackingByBillNumberService: ITrackingByBillNumberService;

export interface ITrackingByBillNumberServer {
    trackByBillNumber: grpc.handleUnaryCall<tracking_pb.Request, tracking_pb.TrackingByBillNumberResponse>;
}

export interface ITrackingByBillNumberClient {
    trackByBillNumber(request: tracking_pb.Request, callback: (error: grpc.ServiceError | null, response: tracking_pb.TrackingByBillNumberResponse) => void): grpc.ClientUnaryCall;
    trackByBillNumber(request: tracking_pb.Request, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: tracking_pb.TrackingByBillNumberResponse) => void): grpc.ClientUnaryCall;
    trackByBillNumber(request: tracking_pb.Request, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: tracking_pb.TrackingByBillNumberResponse) => void): grpc.ClientUnaryCall;
}

export class TrackingByBillNumberClient extends grpc.Client implements ITrackingByBillNumberClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public trackByBillNumber(request: tracking_pb.Request, callback: (error: grpc.ServiceError | null, response: tracking_pb.TrackingByBillNumberResponse) => void): grpc.ClientUnaryCall;
    public trackByBillNumber(request: tracking_pb.Request, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: tracking_pb.TrackingByBillNumberResponse) => void): grpc.ClientUnaryCall;
    public trackByBillNumber(request: tracking_pb.Request, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: tracking_pb.TrackingByBillNumberResponse) => void): grpc.ClientUnaryCall;
}

interface IScacServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    getAll: IScacServiceService_IGetAll;
}

interface IScacServiceService_IGetAll extends grpc.MethodDefinition<google_protobuf_empty_pb.Empty, tracking_pb.GetAllScacResponse> {
    path: "/tracking.ScacService/GetAll";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<google_protobuf_empty_pb.Empty>;
    requestDeserialize: grpc.deserialize<google_protobuf_empty_pb.Empty>;
    responseSerialize: grpc.serialize<tracking_pb.GetAllScacResponse>;
    responseDeserialize: grpc.deserialize<tracking_pb.GetAllScacResponse>;
}

export const ScacServiceService: IScacServiceService;

export interface IScacServiceServer {
    getAll: grpc.handleUnaryCall<google_protobuf_empty_pb.Empty, tracking_pb.GetAllScacResponse>;
}

export interface IScacServiceClient {
    getAll(request: google_protobuf_empty_pb.Empty, callback: (error: grpc.ServiceError | null, response: tracking_pb.GetAllScacResponse) => void): grpc.ClientUnaryCall;
    getAll(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: tracking_pb.GetAllScacResponse) => void): grpc.ClientUnaryCall;
    getAll(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: tracking_pb.GetAllScacResponse) => void): grpc.ClientUnaryCall;
}

export class ScacServiceClient extends grpc.Client implements IScacServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public getAll(request: google_protobuf_empty_pb.Empty, callback: (error: grpc.ServiceError | null, response: tracking_pb.GetAllScacResponse) => void): grpc.ClientUnaryCall;
    public getAll(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: tracking_pb.GetAllScacResponse) => void): grpc.ClientUnaryCall;
    public getAll(request: google_protobuf_empty_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: tracking_pb.GetAllScacResponse) => void): grpc.ClientUnaryCall;
}
