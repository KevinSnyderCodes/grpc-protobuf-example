/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

goog.provide('proto.Bounds');

goog.require('jspb.BinaryReader');
goog.require('jspb.BinaryWriter');
goog.require('jspb.Message');
goog.require('proto.Coordinates');

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
proto.Bounds = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.Bounds, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.Bounds.displayName = 'proto.Bounds';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.Bounds.prototype.toObject = function(opt_includeInstance) {
  return proto.Bounds.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.Bounds} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.Bounds.toObject = function(includeInstance, msg) {
  var f, obj = {
    min: (f = msg.getMin()) && proto.Coordinates.toObject(includeInstance, f),
    max: (f = msg.getMax()) && proto.Coordinates.toObject(includeInstance, f)
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
 * @return {!proto.Bounds}
 */
proto.Bounds.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.Bounds;
  return proto.Bounds.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.Bounds} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.Bounds}
 */
proto.Bounds.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.Coordinates;
      reader.readMessage(value,proto.Coordinates.deserializeBinaryFromReader);
      msg.setMin(value);
      break;
    case 2:
      var value = new proto.Coordinates;
      reader.readMessage(value,proto.Coordinates.deserializeBinaryFromReader);
      msg.setMax(value);
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
proto.Bounds.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.Bounds.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.Bounds} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.Bounds.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getMin();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.Coordinates.serializeBinaryToWriter
    );
  }
  f = message.getMax();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      proto.Coordinates.serializeBinaryToWriter
    );
  }
};


/**
 * optional Coordinates min = 1;
 * @return {?proto.Coordinates}
 */
proto.Bounds.prototype.getMin = function() {
  return /** @type{?proto.Coordinates} */ (
    jspb.Message.getWrapperField(this, proto.Coordinates, 1));
};


/** @param {?proto.Coordinates|undefined} value */
proto.Bounds.prototype.setMin = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.Bounds.prototype.clearMin = function() {
  this.setMin(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.Bounds.prototype.hasMin = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional Coordinates max = 2;
 * @return {?proto.Coordinates}
 */
proto.Bounds.prototype.getMax = function() {
  return /** @type{?proto.Coordinates} */ (
    jspb.Message.getWrapperField(this, proto.Coordinates, 2));
};


/** @param {?proto.Coordinates|undefined} value */
proto.Bounds.prototype.setMax = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


proto.Bounds.prototype.clearMax = function() {
  this.setMax(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.Bounds.prototype.hasMax = function() {
  return jspb.Message.getField(this, 2) != null;
};


