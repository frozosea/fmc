// source: server.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

goog.exportSymbol('proto.tracking.Country', null, global);
goog.exportSymbol('proto.tracking.InfoAboutMoving', null, global);
goog.exportSymbol('proto.tracking.Request', null, global);
goog.exportSymbol('proto.tracking.Scac', null, global);
goog.exportSymbol('proto.tracking.TrackingByBillNumberResponse', null, global);
goog.exportSymbol('proto.tracking.TrackingByContainerNumberResponse', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.tracking.Request = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.tracking.Request, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.tracking.Request.displayName = 'proto.tracking.Request';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.tracking.InfoAboutMoving = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.tracking.InfoAboutMoving, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.tracking.InfoAboutMoving.displayName = 'proto.tracking.InfoAboutMoving';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.tracking.TrackingByContainerNumberResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.tracking.TrackingByContainerNumberResponse.repeatedFields_, null);
};
goog.inherits(proto.tracking.TrackingByContainerNumberResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.tracking.TrackingByContainerNumberResponse.displayName = 'proto.tracking.TrackingByContainerNumberResponse';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.tracking.TrackingByBillNumberResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.tracking.TrackingByBillNumberResponse.repeatedFields_, null);
};
goog.inherits(proto.tracking.TrackingByBillNumberResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.tracking.TrackingByBillNumberResponse.displayName = 'proto.tracking.TrackingByBillNumberResponse';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.tracking.Request.prototype.toObject = function(opt_includeInstance) {
  return proto.tracking.Request.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.tracking.Request} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tracking.Request.toObject = function(includeInstance, msg) {
  var f, obj = {
    number: jspb.Message.getFieldWithDefault(msg, 1, ""),
    scac: jspb.Message.getFieldWithDefault(msg, 2, 0),
    country: jspb.Message.getFieldWithDefault(msg, 3, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.tracking.Request}
 */
proto.tracking.Request.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.tracking.Request;
  return proto.tracking.Request.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.tracking.Request} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.tracking.Request}
 */
proto.tracking.Request.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setNumber(value);
      break;
    case 2:
      var value = /** @type {!proto.tracking.Scac} */ (reader.readEnum());
      msg.setScac(value);
      break;
    case 3:
      var value = /** @type {!proto.tracking.Country} */ (reader.readEnum());
      msg.setCountry(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.tracking.Request.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.tracking.Request.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.tracking.Request} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tracking.Request.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getNumber();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getScac();
  if (f !== 0.0) {
    writer.writeEnum(
      2,
      f
    );
  }
  f = message.getCountry();
  if (f !== 0.0) {
    writer.writeEnum(
      3,
      f
    );
  }
};


/**
 * optional string number = 1;
 * @return {string}
 */
proto.tracking.Request.prototype.getNumber = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.tracking.Request} returns this
 */
proto.tracking.Request.prototype.setNumber = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional Scac scac = 2;
 * @return {!proto.tracking.Scac}
 */
proto.tracking.Request.prototype.getScac = function() {
  return /** @type {!proto.tracking.Scac} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {!proto.tracking.Scac} value
 * @return {!proto.tracking.Request} returns this
 */
proto.tracking.Request.prototype.setScac = function(value) {
  return jspb.Message.setProto3EnumField(this, 2, value);
};


/**
 * optional Country country = 3;
 * @return {!proto.tracking.Country}
 */
proto.tracking.Request.prototype.getCountry = function() {
  return /** @type {!proto.tracking.Country} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {!proto.tracking.Country} value
 * @return {!proto.tracking.Request} returns this
 */
proto.tracking.Request.prototype.setCountry = function(value) {
  return jspb.Message.setProto3EnumField(this, 3, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.tracking.InfoAboutMoving.prototype.toObject = function(opt_includeInstance) {
  return proto.tracking.InfoAboutMoving.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.tracking.InfoAboutMoving} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tracking.InfoAboutMoving.toObject = function(includeInstance, msg) {
  var f, obj = {
    time: jspb.Message.getFieldWithDefault(msg, 1, 0),
    operationName: jspb.Message.getFieldWithDefault(msg, 2, ""),
    location: jspb.Message.getFieldWithDefault(msg, 3, ""),
    vessel: jspb.Message.getFieldWithDefault(msg, 4, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.tracking.InfoAboutMoving}
 */
proto.tracking.InfoAboutMoving.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.tracking.InfoAboutMoving;
  return proto.tracking.InfoAboutMoving.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.tracking.InfoAboutMoving} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.tracking.InfoAboutMoving}
 */
proto.tracking.InfoAboutMoving.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setTime(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setOperationName(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setLocation(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setVessel(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.tracking.InfoAboutMoving.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.tracking.InfoAboutMoving.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.tracking.InfoAboutMoving} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tracking.InfoAboutMoving.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getTime();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getOperationName();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getLocation();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getVessel();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
};


/**
 * optional int64 time = 1;
 * @return {number}
 */
proto.tracking.InfoAboutMoving.prototype.getTime = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/**
 * @param {number} value
 * @return {!proto.tracking.InfoAboutMoving} returns this
 */
proto.tracking.InfoAboutMoving.prototype.setTime = function(value) {
  return jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional string operation_name = 2;
 * @return {string}
 */
proto.tracking.InfoAboutMoving.prototype.getOperationName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.tracking.InfoAboutMoving} returns this
 */
proto.tracking.InfoAboutMoving.prototype.setOperationName = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string location = 3;
 * @return {string}
 */
proto.tracking.InfoAboutMoving.prototype.getLocation = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.tracking.InfoAboutMoving} returns this
 */
proto.tracking.InfoAboutMoving.prototype.setLocation = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string vessel = 4;
 * @return {string}
 */
proto.tracking.InfoAboutMoving.prototype.getVessel = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.tracking.InfoAboutMoving} returns this
 */
proto.tracking.InfoAboutMoving.prototype.setVessel = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.tracking.TrackingByContainerNumberResponse.repeatedFields_ = [4];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.tracking.TrackingByContainerNumberResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.tracking.TrackingByContainerNumberResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.tracking.TrackingByContainerNumberResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tracking.TrackingByContainerNumberResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    container: jspb.Message.getFieldWithDefault(msg, 1, ""),
    containerSize: jspb.Message.getFieldWithDefault(msg, 2, ""),
    scac: jspb.Message.getFieldWithDefault(msg, 3, 0),
    infoAboutMovingList: jspb.Message.toObjectList(msg.getInfoAboutMovingList(),
    proto.tracking.InfoAboutMoving.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.tracking.TrackingByContainerNumberResponse}
 */
proto.tracking.TrackingByContainerNumberResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.tracking.TrackingByContainerNumberResponse;
  return proto.tracking.TrackingByContainerNumberResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.tracking.TrackingByContainerNumberResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.tracking.TrackingByContainerNumberResponse}
 */
proto.tracking.TrackingByContainerNumberResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setContainer(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setContainerSize(value);
      break;
    case 3:
      var value = /** @type {!proto.tracking.Scac} */ (reader.readEnum());
      msg.setScac(value);
      break;
    case 4:
      var value = new proto.tracking.InfoAboutMoving;
      reader.readMessage(value,proto.tracking.InfoAboutMoving.deserializeBinaryFromReader);
      msg.addInfoAboutMoving(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.tracking.TrackingByContainerNumberResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.tracking.TrackingByContainerNumberResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.tracking.TrackingByContainerNumberResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tracking.TrackingByContainerNumberResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getContainer();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getContainerSize();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getScac();
  if (f !== 0.0) {
    writer.writeEnum(
      3,
      f
    );
  }
  f = message.getInfoAboutMovingList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      4,
      f,
      proto.tracking.InfoAboutMoving.serializeBinaryToWriter
    );
  }
};


/**
 * optional string container = 1;
 * @return {string}
 */
proto.tracking.TrackingByContainerNumberResponse.prototype.getContainer = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.tracking.TrackingByContainerNumberResponse} returns this
 */
proto.tracking.TrackingByContainerNumberResponse.prototype.setContainer = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string container_size = 2;
 * @return {string}
 */
proto.tracking.TrackingByContainerNumberResponse.prototype.getContainerSize = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.tracking.TrackingByContainerNumberResponse} returns this
 */
proto.tracking.TrackingByContainerNumberResponse.prototype.setContainerSize = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional Scac scac = 3;
 * @return {!proto.tracking.Scac}
 */
proto.tracking.TrackingByContainerNumberResponse.prototype.getScac = function() {
  return /** @type {!proto.tracking.Scac} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {!proto.tracking.Scac} value
 * @return {!proto.tracking.TrackingByContainerNumberResponse} returns this
 */
proto.tracking.TrackingByContainerNumberResponse.prototype.setScac = function(value) {
  return jspb.Message.setProto3EnumField(this, 3, value);
};


/**
 * repeated InfoAboutMoving info_about_moving = 4;
 * @return {!Array<!proto.tracking.InfoAboutMoving>}
 */
proto.tracking.TrackingByContainerNumberResponse.prototype.getInfoAboutMovingList = function() {
  return /** @type{!Array<!proto.tracking.InfoAboutMoving>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.tracking.InfoAboutMoving, 4));
};


/**
 * @param {!Array<!proto.tracking.InfoAboutMoving>} value
 * @return {!proto.tracking.TrackingByContainerNumberResponse} returns this
*/
proto.tracking.TrackingByContainerNumberResponse.prototype.setInfoAboutMovingList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 4, value);
};


/**
 * @param {!proto.tracking.InfoAboutMoving=} opt_value
 * @param {number=} opt_index
 * @return {!proto.tracking.InfoAboutMoving}
 */
proto.tracking.TrackingByContainerNumberResponse.prototype.addInfoAboutMoving = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 4, opt_value, proto.tracking.InfoAboutMoving, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.tracking.TrackingByContainerNumberResponse} returns this
 */
proto.tracking.TrackingByContainerNumberResponse.prototype.clearInfoAboutMovingList = function() {
  return this.setInfoAboutMovingList([]);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.tracking.TrackingByBillNumberResponse.repeatedFields_ = [4];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.tracking.TrackingByBillNumberResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.tracking.TrackingByBillNumberResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.tracking.TrackingByBillNumberResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tracking.TrackingByBillNumberResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    billno: jspb.Message.getFieldWithDefault(msg, 1, ""),
    scac: jspb.Message.getFieldWithDefault(msg, 3, 0),
    infoAboutMovingList: jspb.Message.toObjectList(msg.getInfoAboutMovingList(),
    proto.tracking.InfoAboutMoving.toObject, includeInstance),
    etaFinalDelivery: jspb.Message.getFieldWithDefault(msg, 5, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.tracking.TrackingByBillNumberResponse}
 */
proto.tracking.TrackingByBillNumberResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.tracking.TrackingByBillNumberResponse;
  return proto.tracking.TrackingByBillNumberResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.tracking.TrackingByBillNumberResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.tracking.TrackingByBillNumberResponse}
 */
proto.tracking.TrackingByBillNumberResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setBillno(value);
      break;
    case 3:
      var value = /** @type {!proto.tracking.Scac} */ (reader.readEnum());
      msg.setScac(value);
      break;
    case 4:
      var value = new proto.tracking.InfoAboutMoving;
      reader.readMessage(value,proto.tracking.InfoAboutMoving.deserializeBinaryFromReader);
      msg.addInfoAboutMoving(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setEtaFinalDelivery(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.tracking.TrackingByBillNumberResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.tracking.TrackingByBillNumberResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.tracking.TrackingByBillNumberResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.tracking.TrackingByBillNumberResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getBillno();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getScac();
  if (f !== 0.0) {
    writer.writeEnum(
      3,
      f
    );
  }
  f = message.getInfoAboutMovingList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      4,
      f,
      proto.tracking.InfoAboutMoving.serializeBinaryToWriter
    );
  }
  f = message.getEtaFinalDelivery();
  if (f !== 0) {
    writer.writeInt64(
      5,
      f
    );
  }
};


/**
 * optional string billNo = 1;
 * @return {string}
 */
proto.tracking.TrackingByBillNumberResponse.prototype.getBillno = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.tracking.TrackingByBillNumberResponse} returns this
 */
proto.tracking.TrackingByBillNumberResponse.prototype.setBillno = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional Scac scac = 3;
 * @return {!proto.tracking.Scac}
 */
proto.tracking.TrackingByBillNumberResponse.prototype.getScac = function() {
  return /** @type {!proto.tracking.Scac} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {!proto.tracking.Scac} value
 * @return {!proto.tracking.TrackingByBillNumberResponse} returns this
 */
proto.tracking.TrackingByBillNumberResponse.prototype.setScac = function(value) {
  return jspb.Message.setProto3EnumField(this, 3, value);
};


/**
 * repeated InfoAboutMoving info_about_moving = 4;
 * @return {!Array<!proto.tracking.InfoAboutMoving>}
 */
proto.tracking.TrackingByBillNumberResponse.prototype.getInfoAboutMovingList = function() {
  return /** @type{!Array<!proto.tracking.InfoAboutMoving>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.tracking.InfoAboutMoving, 4));
};


/**
 * @param {!Array<!proto.tracking.InfoAboutMoving>} value
 * @return {!proto.tracking.TrackingByBillNumberResponse} returns this
*/
proto.tracking.TrackingByBillNumberResponse.prototype.setInfoAboutMovingList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 4, value);
};


/**
 * @param {!proto.tracking.InfoAboutMoving=} opt_value
 * @param {number=} opt_index
 * @return {!proto.tracking.InfoAboutMoving}
 */
proto.tracking.TrackingByBillNumberResponse.prototype.addInfoAboutMoving = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 4, opt_value, proto.tracking.InfoAboutMoving, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.tracking.TrackingByBillNumberResponse} returns this
 */
proto.tracking.TrackingByBillNumberResponse.prototype.clearInfoAboutMovingList = function() {
  return this.setInfoAboutMovingList([]);
};


/**
 * optional int64 eta_final_delivery = 5;
 * @return {number}
 */
proto.tracking.TrackingByBillNumberResponse.prototype.getEtaFinalDelivery = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/**
 * @param {number} value
 * @return {!proto.tracking.TrackingByBillNumberResponse} returns this
 */
proto.tracking.TrackingByBillNumberResponse.prototype.setEtaFinalDelivery = function(value) {
  return jspb.Message.setProto3IntField(this, 5, value);
};


/**
 * @enum {number}
 */
proto.tracking.Scac = {
  FESO: 0,
  SKLU: 1,
  MAEU: 2,
  COSU: 3,
  KMTU: 4,
  ONEY: 5,
  SITC: 6,
  MSCU: 7,
  HALU: 8,
  AUTO: 9
};

/**
 * @enum {number}
 */
proto.tracking.Country = {
  RU: 0,
  OTHER: 1
};

goog.object.extend(exports, proto.tracking);
