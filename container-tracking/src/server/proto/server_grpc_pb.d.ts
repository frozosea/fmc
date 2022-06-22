// package: tracking
// file: services.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as server_pb from "./server_pb";

interface ITrackingByContainerNumberService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    track: ITrackingByContainerNumberService_ITrack;
}

interface ITrackingByContainerNumberService_ITrack extends grpc.MethodDefinition<server_pb.Request, server_pb.Response> {
    path: "/tracking.TrackingByContainerNumber/Track";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<server_pb.Request>;
    requestDeserialize: grpc.deserialize<server_pb.Request>;
    responseSerialize: grpc.serialize<server_pb.Response>;
    responseDeserialize: grpc.deserialize<server_pb.Response>;
}

export const TrackingByContainerNumberService: ITrackingByContainerNumberService;

export interface ITrackingByContainerNumberServer {
    track: grpc.handleUnaryCall<server_pb.Request, server_pb.Response>;
}

export interface ITrackingByContainerNumberClient {
    track(request: server_pb.Request, callback: (error: grpc.ServiceError | null, response: server_pb.Response) => void): grpc.ClientUnaryCall;
    track(request: server_pb.Request, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: server_pb.Response) => void): grpc.ClientUnaryCall;
    track(request: server_pb.Request, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: server_pb.Response) => void): grpc.ClientUnaryCall;
}

export class TrackingByContainerNumberClient extends grpc.Client implements ITrackingByContainerNumberClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public track(request: server_pb.Request, callback: (error: grpc.ServiceError | null, response: server_pb.Response) => void): grpc.ClientUnaryCall;
    public track(request: server_pb.Request, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: server_pb.Response) => void): grpc.ClientUnaryCall;
    public track(request: server_pb.Request, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: server_pb.Response) => void): grpc.ClientUnaryCall;
}
