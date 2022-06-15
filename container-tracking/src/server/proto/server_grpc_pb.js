// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var server_pb = require('./server_pb.js');

function serialize_tracking_Request(arg) {
  if (!(arg instanceof server_pb.Request)) {
    throw new Error('Expected argument of type tracking.Request');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_tracking_Request(buffer_arg) {
  return server_pb.Request.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_tracking_Response(arg) {
  if (!(arg instanceof server_pb.Response)) {
    throw new Error('Expected argument of type tracking.Response');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_tracking_Response(buffer_arg) {
  return server_pb.Response.deserializeBinary(new Uint8Array(buffer_arg));
}


var TrackingByContainerNumberService = exports.TrackingByContainerNumberService = {
  track: {
    path: '/tracking.TrackingByContainerNumber/Track',
    requestStream: false,
    responseStream: false,
    requestType: server_pb.Request,
    responseType: server_pb.Response,
    requestSerialize: serialize_tracking_Request,
    requestDeserialize: deserialize_tracking_Request,
    responseSerialize: serialize_tracking_Response,
    responseDeserialize: deserialize_tracking_Response,
  },
};

exports.TrackingByContainerNumberClient = grpc.makeGenericClientConstructor(TrackingByContainerNumberService);
