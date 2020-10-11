/**
 * @fileoverview gRPC-Web generated client stub for places
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
proto.places = require('./placements_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.places.V1PlacesClient =
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
proto.places.V1PlacesPromiseClient =
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
 *   !proto.places.GetPlaceRequest,
 *   !proto.places.GetPlaceReply>}
 */
const methodDescriptor_V1Places_GetPlace = new grpc.web.MethodDescriptor(
  '/places.V1Places/GetPlace',
  grpc.web.MethodType.UNARY,
  proto.places.GetPlaceRequest,
  proto.places.GetPlaceReply,
  /**
   * @param {!proto.places.GetPlaceRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.places.GetPlaceReply.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.places.GetPlaceRequest,
 *   !proto.places.GetPlaceReply>}
 */
const methodInfo_V1Places_GetPlace = new grpc.web.AbstractClientBase.MethodInfo(
  proto.places.GetPlaceReply,
  /**
   * @param {!proto.places.GetPlaceRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.places.GetPlaceReply.deserializeBinary
);


/**
 * @param {!proto.places.GetPlaceRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.places.GetPlaceReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.places.GetPlaceReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.places.V1PlacesClient.prototype.getPlace =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/places.V1Places/GetPlace',
      request,
      metadata || {},
      methodDescriptor_V1Places_GetPlace,
      callback);
};


/**
 * @param {!proto.places.GetPlaceRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.places.GetPlaceReply>}
 *     Promise that resolves to the response
 */
proto.places.V1PlacesPromiseClient.prototype.getPlace =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/places.V1Places/GetPlace',
      request,
      metadata || {},
      methodDescriptor_V1Places_GetPlace);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.places.FindPlacesRequest,
 *   !proto.places.FindPlacesReply>}
 */
const methodDescriptor_V1Places_FindPlaces = new grpc.web.MethodDescriptor(
  '/places.V1Places/FindPlaces',
  grpc.web.MethodType.UNARY,
  proto.places.FindPlacesRequest,
  proto.places.FindPlacesReply,
  /**
   * @param {!proto.places.FindPlacesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.places.FindPlacesReply.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.places.FindPlacesRequest,
 *   !proto.places.FindPlacesReply>}
 */
const methodInfo_V1Places_FindPlaces = new grpc.web.AbstractClientBase.MethodInfo(
  proto.places.FindPlacesReply,
  /**
   * @param {!proto.places.FindPlacesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.places.FindPlacesReply.deserializeBinary
);


/**
 * @param {!proto.places.FindPlacesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.places.FindPlacesReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.places.FindPlacesReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.places.V1PlacesClient.prototype.findPlaces =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/places.V1Places/FindPlaces',
      request,
      metadata || {},
      methodDescriptor_V1Places_FindPlaces,
      callback);
};


/**
 * @param {!proto.places.FindPlacesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.places.FindPlacesReply>}
 *     Promise that resolves to the response
 */
proto.places.V1PlacesPromiseClient.prototype.findPlaces =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/places.V1Places/FindPlaces',
      request,
      metadata || {},
      methodDescriptor_V1Places_FindPlaces);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.places.CreatePlaceRequest,
 *   !proto.places.CreatePlaceReply>}
 */
const methodDescriptor_V1Places_CreatePlace = new grpc.web.MethodDescriptor(
  '/places.V1Places/CreatePlace',
  grpc.web.MethodType.UNARY,
  proto.places.CreatePlaceRequest,
  proto.places.CreatePlaceReply,
  /**
   * @param {!proto.places.CreatePlaceRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.places.CreatePlaceReply.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.places.CreatePlaceRequest,
 *   !proto.places.CreatePlaceReply>}
 */
const methodInfo_V1Places_CreatePlace = new grpc.web.AbstractClientBase.MethodInfo(
  proto.places.CreatePlaceReply,
  /**
   * @param {!proto.places.CreatePlaceRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.places.CreatePlaceReply.deserializeBinary
);


/**
 * @param {!proto.places.CreatePlaceRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.places.CreatePlaceReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.places.CreatePlaceReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.places.V1PlacesClient.prototype.createPlace =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/places.V1Places/CreatePlace',
      request,
      metadata || {},
      methodDescriptor_V1Places_CreatePlace,
      callback);
};


/**
 * @param {!proto.places.CreatePlaceRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.places.CreatePlaceReply>}
 *     Promise that resolves to the response
 */
proto.places.V1PlacesPromiseClient.prototype.createPlace =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/places.V1Places/CreatePlace',
      request,
      metadata || {},
      methodDescriptor_V1Places_CreatePlace);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.places.DestroyPlaceRequest,
 *   !proto.types.EmptyReply>}
 */
const methodDescriptor_V1Places_DestroyPlace = new grpc.web.MethodDescriptor(
  '/places.V1Places/DestroyPlace',
  grpc.web.MethodType.UNARY,
  proto.places.DestroyPlaceRequest,
  types_pb.EmptyReply,
  /**
   * @param {!proto.places.DestroyPlaceRequest} request
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
 *   !proto.places.DestroyPlaceRequest,
 *   !proto.types.EmptyReply>}
 */
const methodInfo_V1Places_DestroyPlace = new grpc.web.AbstractClientBase.MethodInfo(
  types_pb.EmptyReply,
  /**
   * @param {!proto.places.DestroyPlaceRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  types_pb.EmptyReply.deserializeBinary
);


/**
 * @param {!proto.places.DestroyPlaceRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.types.EmptyReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.types.EmptyReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.places.V1PlacesClient.prototype.destroyPlace =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/places.V1Places/DestroyPlace',
      request,
      metadata || {},
      methodDescriptor_V1Places_DestroyPlace,
      callback);
};


/**
 * @param {!proto.places.DestroyPlaceRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.types.EmptyReply>}
 *     Promise that resolves to the response
 */
proto.places.V1PlacesPromiseClient.prototype.destroyPlace =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/places.V1Places/DestroyPlace',
      request,
      metadata || {},
      methodDescriptor_V1Places_DestroyPlace);
};


module.exports = proto.places;

