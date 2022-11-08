// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var tracking_pb = require('./tracking_pb.js');
var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js');

function serialize_google_protobuf_Empty(arg) {
  if (!(arg instanceof google_protobuf_empty_pb.Empty)) {
    throw new Error('Expected argument of type google.protobuf.Empty');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_google_protobuf_Empty(buffer_arg) {
  return google_protobuf_empty_pb.Empty.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_tracking_GetAllScacResponse(arg) {
  if (!(arg instanceof tracking_pb.GetAllScacResponse)) {
    throw new Error('Expected argument of type tracking.GetAllScacResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_tracking_GetAllScacResponse(buffer_arg) {
  return tracking_pb.GetAllScacResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_tracking_Request(arg) {
  if (!(arg instanceof tracking_pb.Request)) {
    throw new Error('Expected argument of type tracking.Request');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_tracking_Request(buffer_arg) {
  return tracking_pb.Request.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_tracking_TrackingByBillNumberResponse(arg) {
  if (!(arg instanceof tracking_pb.TrackingByBillNumberResponse)) {
    throw new Error('Expected argument of type tracking.TrackingByBillNumberResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_tracking_TrackingByBillNumberResponse(buffer_arg) {
  return tracking_pb.TrackingByBillNumberResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_tracking_TrackingByContainerNumberResponse(arg) {
  if (!(arg instanceof tracking_pb.TrackingByContainerNumberResponse)) {
    throw new Error('Expected argument of type tracking.TrackingByContainerNumberResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_tracking_TrackingByContainerNumberResponse(buffer_arg) {
  return tracking_pb.TrackingByContainerNumberResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


var TrackingByContainerNumberService = exports.TrackingByContainerNumberService = {
  trackByContainerNumber: {
    path: '/tracking.TrackingByContainerNumber/TrackByContainerNumber',
    requestStream: false,
    responseStream: false,
    requestType: tracking_pb.Request,
    responseType: tracking_pb.TrackingByContainerNumberResponse,
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
    requestType: tracking_pb.Request,
    responseType: tracking_pb.TrackingByBillNumberResponse,
    requestSerialize: serialize_tracking_Request,
    requestDeserialize: deserialize_tracking_Request,
    responseSerialize: serialize_tracking_TrackingByBillNumberResponse,
    responseDeserialize: deserialize_tracking_TrackingByBillNumberResponse,
  },
};

exports.TrackingByBillNumberClient = grpc.makeGenericClientConstructor(TrackingByBillNumberService);
var ScacServiceService = exports.ScacServiceService = {
  getAll: {
    path: '/tracking.ScacService/GetAll',
    requestStream: false,
    responseStream: false,
    requestType: google_protobuf_empty_pb.Empty,
    responseType: tracking_pb.GetAllScacResponse,
    requestSerialize: serialize_google_protobuf_Empty,
    requestDeserialize: deserialize_google_protobuf_Empty,
    responseSerialize: serialize_tracking_GetAllScacResponse,
    responseDeserialize: deserialize_tracking_GetAllScacResponse,
  },
};

exports.ScacServiceClient = grpc.makeGenericClientConstructor(ScacServiceService);
