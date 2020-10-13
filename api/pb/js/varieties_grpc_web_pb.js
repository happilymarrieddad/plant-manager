/**
 * @fileoverview gRPC-Web generated client stub for varieties
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
proto.varieties = require('./varieties_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.varieties.V1VarietiesClient =
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
proto.varieties.V1VarietiesPromiseClient =
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
 *   !proto.varieties.GetVarietyRequest,
 *   !proto.varieties.GetVarietyReply>}
 */
const methodDescriptor_V1Varieties_GetVariety = new grpc.web.MethodDescriptor(
  '/varieties.V1Varieties/GetVariety',
  grpc.web.MethodType.UNARY,
  proto.varieties.GetVarietyRequest,
  proto.varieties.GetVarietyReply,
  /**
   * @param {!proto.varieties.GetVarietyRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.varieties.GetVarietyReply.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.varieties.GetVarietyRequest,
 *   !proto.varieties.GetVarietyReply>}
 */
const methodInfo_V1Varieties_GetVariety = new grpc.web.AbstractClientBase.MethodInfo(
  proto.varieties.GetVarietyReply,
  /**
   * @param {!proto.varieties.GetVarietyRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.varieties.GetVarietyReply.deserializeBinary
);


/**
 * @param {!proto.varieties.GetVarietyRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.varieties.GetVarietyReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.varieties.GetVarietyReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.varieties.V1VarietiesClient.prototype.getVariety =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/varieties.V1Varieties/GetVariety',
      request,
      metadata || {},
      methodDescriptor_V1Varieties_GetVariety,
      callback);
};


/**
 * @param {!proto.varieties.GetVarietyRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.varieties.GetVarietyReply>}
 *     Promise that resolves to the response
 */
proto.varieties.V1VarietiesPromiseClient.prototype.getVariety =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/varieties.V1Varieties/GetVariety',
      request,
      metadata || {},
      methodDescriptor_V1Varieties_GetVariety);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.varieties.FindVarietiesRequest,
 *   !proto.varieties.FindVarietiesReply>}
 */
const methodDescriptor_V1Varieties_FindVarieties = new grpc.web.MethodDescriptor(
  '/varieties.V1Varieties/FindVarieties',
  grpc.web.MethodType.UNARY,
  proto.varieties.FindVarietiesRequest,
  proto.varieties.FindVarietiesReply,
  /**
   * @param {!proto.varieties.FindVarietiesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.varieties.FindVarietiesReply.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.varieties.FindVarietiesRequest,
 *   !proto.varieties.FindVarietiesReply>}
 */
const methodInfo_V1Varieties_FindVarieties = new grpc.web.AbstractClientBase.MethodInfo(
  proto.varieties.FindVarietiesReply,
  /**
   * @param {!proto.varieties.FindVarietiesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.varieties.FindVarietiesReply.deserializeBinary
);


/**
 * @param {!proto.varieties.FindVarietiesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.varieties.FindVarietiesReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.varieties.FindVarietiesReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.varieties.V1VarietiesClient.prototype.findVarieties =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/varieties.V1Varieties/FindVarieties',
      request,
      metadata || {},
      methodDescriptor_V1Varieties_FindVarieties,
      callback);
};


/**
 * @param {!proto.varieties.FindVarietiesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.varieties.FindVarietiesReply>}
 *     Promise that resolves to the response
 */
proto.varieties.V1VarietiesPromiseClient.prototype.findVarieties =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/varieties.V1Varieties/FindVarieties',
      request,
      metadata || {},
      methodDescriptor_V1Varieties_FindVarieties);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.varieties.CreateVarietyRequest,
 *   !proto.varieties.CreateVarietyReply>}
 */
const methodDescriptor_V1Varieties_CreateVariety = new grpc.web.MethodDescriptor(
  '/varieties.V1Varieties/CreateVariety',
  grpc.web.MethodType.UNARY,
  proto.varieties.CreateVarietyRequest,
  proto.varieties.CreateVarietyReply,
  /**
   * @param {!proto.varieties.CreateVarietyRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.varieties.CreateVarietyReply.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.varieties.CreateVarietyRequest,
 *   !proto.varieties.CreateVarietyReply>}
 */
const methodInfo_V1Varieties_CreateVariety = new grpc.web.AbstractClientBase.MethodInfo(
  proto.varieties.CreateVarietyReply,
  /**
   * @param {!proto.varieties.CreateVarietyRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.varieties.CreateVarietyReply.deserializeBinary
);


/**
 * @param {!proto.varieties.CreateVarietyRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.varieties.CreateVarietyReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.varieties.CreateVarietyReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.varieties.V1VarietiesClient.prototype.createVariety =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/varieties.V1Varieties/CreateVariety',
      request,
      metadata || {},
      methodDescriptor_V1Varieties_CreateVariety,
      callback);
};


/**
 * @param {!proto.varieties.CreateVarietyRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.varieties.CreateVarietyReply>}
 *     Promise that resolves to the response
 */
proto.varieties.V1VarietiesPromiseClient.prototype.createVariety =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/varieties.V1Varieties/CreateVariety',
      request,
      metadata || {},
      methodDescriptor_V1Varieties_CreateVariety);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.varieties.UpdateVarietyRequest,
 *   !proto.types.EmptyReply>}
 */
const methodDescriptor_V1Varieties_UpdateVariety = new grpc.web.MethodDescriptor(
  '/varieties.V1Varieties/UpdateVariety',
  grpc.web.MethodType.UNARY,
  proto.varieties.UpdateVarietyRequest,
  types_pb.EmptyReply,
  /**
   * @param {!proto.varieties.UpdateVarietyRequest} request
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
 *   !proto.varieties.UpdateVarietyRequest,
 *   !proto.types.EmptyReply>}
 */
const methodInfo_V1Varieties_UpdateVariety = new grpc.web.AbstractClientBase.MethodInfo(
  types_pb.EmptyReply,
  /**
   * @param {!proto.varieties.UpdateVarietyRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  types_pb.EmptyReply.deserializeBinary
);


/**
 * @param {!proto.varieties.UpdateVarietyRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.types.EmptyReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.types.EmptyReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.varieties.V1VarietiesClient.prototype.updateVariety =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/varieties.V1Varieties/UpdateVariety',
      request,
      metadata || {},
      methodDescriptor_V1Varieties_UpdateVariety,
      callback);
};


/**
 * @param {!proto.varieties.UpdateVarietyRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.types.EmptyReply>}
 *     Promise that resolves to the response
 */
proto.varieties.V1VarietiesPromiseClient.prototype.updateVariety =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/varieties.V1Varieties/UpdateVariety',
      request,
      metadata || {},
      methodDescriptor_V1Varieties_UpdateVariety);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.varieties.DestroyVarietyRequest,
 *   !proto.types.EmptyReply>}
 */
const methodDescriptor_V1Varieties_DestroyVariety = new grpc.web.MethodDescriptor(
  '/varieties.V1Varieties/DestroyVariety',
  grpc.web.MethodType.UNARY,
  proto.varieties.DestroyVarietyRequest,
  types_pb.EmptyReply,
  /**
   * @param {!proto.varieties.DestroyVarietyRequest} request
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
 *   !proto.varieties.DestroyVarietyRequest,
 *   !proto.types.EmptyReply>}
 */
const methodInfo_V1Varieties_DestroyVariety = new grpc.web.AbstractClientBase.MethodInfo(
  types_pb.EmptyReply,
  /**
   * @param {!proto.varieties.DestroyVarietyRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  types_pb.EmptyReply.deserializeBinary
);


/**
 * @param {!proto.varieties.DestroyVarietyRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.types.EmptyReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.types.EmptyReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.varieties.V1VarietiesClient.prototype.destroyVariety =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/varieties.V1Varieties/DestroyVariety',
      request,
      metadata || {},
      methodDescriptor_V1Varieties_DestroyVariety,
      callback);
};


/**
 * @param {!proto.varieties.DestroyVarietyRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.types.EmptyReply>}
 *     Promise that resolves to the response
 */
proto.varieties.V1VarietiesPromiseClient.prototype.destroyVariety =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/varieties.V1Varieties/DestroyVariety',
      request,
      metadata || {},
      methodDescriptor_V1Varieties_DestroyVariety);
};


module.exports = proto.varieties;

