/**
 * @fileoverview gRPC-Web generated client stub for place_slots
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var types_pb = require('./types_pb.js')
const proto = {};
proto.place_slots = require('./place_slots_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.place_slots.V1PlaceSlotsClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.place_slots.V1PlaceSlotsPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.place_slots.GetPlaceSlotRequest,
 *   !proto.place_slots.GetPlaceSlotReply>}
 */
const methodDescriptor_V1PlaceSlots_GetPlaceSlot = new grpc.web.MethodDescriptor(
  '/place_slots.V1PlaceSlots/GetPlaceSlot',
  grpc.web.MethodType.UNARY,
  proto.place_slots.GetPlaceSlotRequest,
  proto.place_slots.GetPlaceSlotReply,
  /**
   * @param {!proto.place_slots.GetPlaceSlotRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.place_slots.GetPlaceSlotReply.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.place_slots.GetPlaceSlotRequest,
 *   !proto.place_slots.GetPlaceSlotReply>}
 */
const methodInfo_V1PlaceSlots_GetPlaceSlot = new grpc.web.AbstractClientBase.MethodInfo(
  proto.place_slots.GetPlaceSlotReply,
  /**
   * @param {!proto.place_slots.GetPlaceSlotRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.place_slots.GetPlaceSlotReply.deserializeBinary
);


/**
 * @param {!proto.place_slots.GetPlaceSlotRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.place_slots.GetPlaceSlotReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.place_slots.GetPlaceSlotReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.place_slots.V1PlaceSlotsClient.prototype.getPlaceSlot =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/place_slots.V1PlaceSlots/GetPlaceSlot',
      request,
      metadata || {},
      methodDescriptor_V1PlaceSlots_GetPlaceSlot,
      callback);
};


/**
 * @param {!proto.place_slots.GetPlaceSlotRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.place_slots.GetPlaceSlotReply>}
 *     Promise that resolves to the response
 */
proto.place_slots.V1PlaceSlotsPromiseClient.prototype.getPlaceSlot =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/place_slots.V1PlaceSlots/GetPlaceSlot',
      request,
      metadata || {},
      methodDescriptor_V1PlaceSlots_GetPlaceSlot);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.place_slots.UpdatePlaceSlotRequest,
 *   !proto.types.EmptyReply>}
 */
const methodDescriptor_V1PlaceSlots_UpdatePlaceSlot = new grpc.web.MethodDescriptor(
  '/place_slots.V1PlaceSlots/UpdatePlaceSlot',
  grpc.web.MethodType.UNARY,
  proto.place_slots.UpdatePlaceSlotRequest,
  types_pb.EmptyReply,
  /**
   * @param {!proto.place_slots.UpdatePlaceSlotRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  types_pb.EmptyReply.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.place_slots.UpdatePlaceSlotRequest,
 *   !proto.types.EmptyReply>}
 */
const methodInfo_V1PlaceSlots_UpdatePlaceSlot = new grpc.web.AbstractClientBase.MethodInfo(
  types_pb.EmptyReply,
  /**
   * @param {!proto.place_slots.UpdatePlaceSlotRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  types_pb.EmptyReply.deserializeBinary
);


/**
 * @param {!proto.place_slots.UpdatePlaceSlotRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.types.EmptyReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.types.EmptyReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.place_slots.V1PlaceSlotsClient.prototype.updatePlaceSlot =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/place_slots.V1PlaceSlots/UpdatePlaceSlot',
      request,
      metadata || {},
      methodDescriptor_V1PlaceSlots_UpdatePlaceSlot,
      callback);
};


/**
 * @param {!proto.place_slots.UpdatePlaceSlotRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.types.EmptyReply>}
 *     Promise that resolves to the response
 */
proto.place_slots.V1PlaceSlotsPromiseClient.prototype.updatePlaceSlot =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/place_slots.V1PlaceSlots/UpdatePlaceSlot',
      request,
      metadata || {},
      methodDescriptor_V1PlaceSlots_UpdatePlaceSlot);
};


module.exports = proto.place_slots;

