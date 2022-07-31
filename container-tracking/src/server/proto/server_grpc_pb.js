// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
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

function serialize_tracking_TrackingByBillNumberResponse(arg) {
  if (!(arg instanceof server_pb.TrackingByBillNumberResponse)) {
    throw new Error('Expected argument of type tracking.TrackingByBillNumberResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_tracking_TrackingByBillNumberResponse(buffer_arg) {
  return server_pb.TrackingByBillNumberResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_tracking_TrackingByContainerNumberResponse(arg) {
  if (!(arg instanceof server_pb.TrackingByContainerNumberResponse)) {
    throw new Error('Expected argument of type tracking.TrackingByContainerNumberResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_tracking_TrackingByContainerNumberResponse(buffer_arg) {
  return server_pb.TrackingByContainerNumberResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var TrackingByContainerNumberService = exports.TrackingByContainerNumberService = {
  trackByContainerNumber: {
    path: '/tracking.TrackingByContainerNumber/TrackByContainerNumber',
    requestStream: false,
    responseStream: false,
    requestType: server_pb.Request,
    responseType: server_pb.TrackingByContainerNumberResponse,
    requestSerialize: serialize_tracking_Request,
    requestDeserialize: deserialize_tracking_Request,
    responseSerialize: serialize_tracking_TrackingByContainerNumberResponse,
    responseDeserialize: deserialize_tracking_TrackingByContainerNumberResponse,
  },
};

exports.TrackingByContainerNumberClient = grpc.makeGenericClientConstructor(TrackingByContainerNumberService);
var TrackingByBillNumberService = exports.TrackingByBillNumberService = {
  trackByBillNumber: {
    path: '/tracking.TrackingByBillNumber/TrackByBillNumber',
    requestStream: false,
    responseStream: false,
    requestType: server_pb.Request,
    responseType: server_pb.TrackingByBillNumberResponse,
    requestSerialize: serialize_tracking_Request,
    requestDeserialize: deserialize_tracking_Request,
    responseSerialize: serialize_tracking_TrackingByBillNumberResponse,
    responseDeserialize: deserialize_tracking_TrackingByBillNumberResponse,
  },
};

exports.TrackingByBillNumberClient = grpc.makeGenericClientConstructor(TrackingByBillNumberService);
