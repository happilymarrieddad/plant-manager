/**
 * @fileoverview gRPC-Web generated client stub for plant_types
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
proto.plant_types = require('./plant_types_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.plant_types.V1PlantTypesClient =
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
proto.plant_types.V1PlantTypesPromiseClient =
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
 *   !proto.plant_types.GetPlantTypeRequest,
 *   !proto.plant_types.GetPlantTypeReply>}
 */
const methodDescriptor_V1PlantTypes_GetPlantType = new grpc.web.MethodDescriptor(
  '/plant_types.V1PlantTypes/GetPlantType',
  grpc.web.MethodType.UNARY,
  proto.plant_types.GetPlantTypeRequest,
  proto.plant_types.GetPlantTypeReply,
  /**
   * @param {!proto.plant_types.GetPlantTypeRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.plant_types.GetPlantTypeReply.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.plant_types.GetPlantTypeRequest,
 *   !proto.plant_types.GetPlantTypeReply>}
 */
const methodInfo_V1PlantTypes_GetPlantType = new grpc.web.AbstractClientBase.MethodInfo(
  proto.plant_types.GetPlantTypeReply,
  /**
   * @param {!proto.plant_types.GetPlantTypeRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.plant_types.GetPlantTypeReply.deserializeBinary
);


/**
 * @param {!proto.plant_types.GetPlantTypeRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.plant_types.GetPlantTypeReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.plant_types.GetPlantTypeReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.plant_types.V1PlantTypesClient.prototype.getPlantType =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/plant_types.V1PlantTypes/GetPlantType',
      request,
      metadata || {},
      methodDescriptor_V1PlantTypes_GetPlantType,
      callback);
};


/**
 * @param {!proto.plant_types.GetPlantTypeRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.plant_types.GetPlantTypeReply>}
 *     Promise that resolves to the response
 */
proto.plant_types.V1PlantTypesPromiseClient.prototype.getPlantType =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/plant_types.V1PlantTypes/GetPlantType',
      request,
      metadata || {},
      methodDescriptor_V1PlantTypes_GetPlantType);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.plant_types.FindPlantTypesRequest,
 *   !proto.plant_types.FindPlantTypesReply>}
 */
const methodDescriptor_V1PlantTypes_FindPlantTypes = new grpc.web.MethodDescriptor(
  '/plant_types.V1PlantTypes/FindPlantTypes',
  grpc.web.MethodType.UNARY,
  proto.plant_types.FindPlantTypesRequest,
  proto.plant_types.FindPlantTypesReply,
  /**
   * @param {!proto.plant_types.FindPlantTypesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.plant_types.FindPlantTypesReply.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.plant_types.FindPlantTypesRequest,
 *   !proto.plant_types.FindPlantTypesReply>}
 */
const methodInfo_V1PlantTypes_FindPlantTypes = new grpc.web.AbstractClientBase.MethodInfo(
  proto.plant_types.FindPlantTypesReply,
  /**
   * @param {!proto.plant_types.FindPlantTypesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.plant_types.FindPlantTypesReply.deserializeBinary
);


/**
 * @param {!proto.plant_types.FindPlantTypesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.plant_types.FindPlantTypesReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.plant_types.FindPlantTypesReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.plant_types.V1PlantTypesClient.prototype.findPlantTypes =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/plant_types.V1PlantTypes/FindPlantTypes',
      request,
      metadata || {},
      methodDescriptor_V1PlantTypes_FindPlantTypes,
      callback);
};


/**
 * @param {!proto.plant_types.FindPlantTypesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.plant_types.FindPlantTypesReply>}
 *     Promise that resolves to the response
 */
proto.plant_types.V1PlantTypesPromiseClient.prototype.findPlantTypes =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/plant_types.V1PlantTypes/FindPlantTypes',
      request,
      metadata || {},
      methodDescriptor_V1PlantTypes_FindPlantTypes);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.plant_types.CreatePlantTypeRequest,
 *   !proto.plant_types.CreatePlantTypeReply>}
 */
const methodDescriptor_V1PlantTypes_CreatePlantType = new grpc.web.MethodDescriptor(
  '/plant_types.V1PlantTypes/CreatePlantType',
  grpc.web.MethodType.UNARY,
  proto.plant_types.CreatePlantTypeRequest,
  proto.plant_types.CreatePlantTypeReply,
  /**
   * @param {!proto.plant_types.CreatePlantTypeRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.plant_types.CreatePlantTypeReply.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.plant_types.CreatePlantTypeRequest,
 *   !proto.plant_types.CreatePlantTypeReply>}
 */
const methodInfo_V1PlantTypes_CreatePlantType = new grpc.web.AbstractClientBase.MethodInfo(
  proto.plant_types.CreatePlantTypeReply,
  /**
   * @param {!proto.plant_types.CreatePlantTypeRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.plant_types.CreatePlantTypeReply.deserializeBinary
);


/**
 * @param {!proto.plant_types.CreatePlantTypeRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.plant_types.CreatePlantTypeReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.plant_types.CreatePlantTypeReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.plant_types.V1PlantTypesClient.prototype.createPlantType =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/plant_types.V1PlantTypes/CreatePlantType',
      request,
      metadata || {},
      methodDescriptor_V1PlantTypes_CreatePlantType,
      callback);
};


/**
 * @param {!proto.plant_types.CreatePlantTypeRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.plant_types.CreatePlantTypeReply>}
 *     Promise that resolves to the response
 */
proto.plant_types.V1PlantTypesPromiseClient.prototype.createPlantType =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/plant_types.V1PlantTypes/CreatePlantType',
      request,
      metadata || {},
      methodDescriptor_V1PlantTypes_CreatePlantType);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.plant_types.UpdatePlantTypeRequest,
 *   !proto.types.EmptyReply>}
 */
const methodDescriptor_V1PlantTypes_UpdatePlantType = new grpc.web.MethodDescriptor(
  '/plant_types.V1PlantTypes/UpdatePlantType',
  grpc.web.MethodType.UNARY,
  proto.plant_types.UpdatePlantTypeRequest,
  types_pb.EmptyReply,
  /**
   * @param {!proto.plant_types.UpdatePlantTypeRequest} request
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
 *   !proto.plant_types.UpdatePlantTypeRequest,
 *   !proto.types.EmptyReply>}
 */
const methodInfo_V1PlantTypes_UpdatePlantType = new grpc.web.AbstractClientBase.MethodInfo(
  types_pb.EmptyReply,
  /**
   * @param {!proto.plant_types.UpdatePlantTypeRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  types_pb.EmptyReply.deserializeBinary
);


/**
 * @param {!proto.plant_types.UpdatePlantTypeRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.types.EmptyReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.types.EmptyReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.plant_types.V1PlantTypesClient.prototype.updatePlantType =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/plant_types.V1PlantTypes/UpdatePlantType',
      request,
      metadata || {},
      methodDescriptor_V1PlantTypes_UpdatePlantType,
      callback);
};


/**
 * @param {!proto.plant_types.UpdatePlantTypeRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.types.EmptyReply>}
 *     Promise that resolves to the response
 */
proto.plant_types.V1PlantTypesPromiseClient.prototype.updatePlantType =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/plant_types.V1PlantTypes/UpdatePlantType',
      request,
      metadata || {},
      methodDescriptor_V1PlantTypes_UpdatePlantType);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.plant_types.DestroyPlantTypeRequest,
 *   !proto.types.EmptyReply>}
 */
const methodDescriptor_V1PlantTypes_DestroyPlantType = new grpc.web.MethodDescriptor(
  '/plant_types.V1PlantTypes/DestroyPlantType',
  grpc.web.MethodType.UNARY,
  proto.plant_types.DestroyPlantTypeRequest,
  types_pb.EmptyReply,
  /**
   * @param {!proto.plant_types.DestroyPlantTypeRequest} request
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
 *   !proto.plant_types.DestroyPlantTypeRequest,
 *   !proto.types.EmptyReply>}
 */
const methodInfo_V1PlantTypes_DestroyPlantType = new grpc.web.AbstractClientBase.MethodInfo(
  types_pb.EmptyReply,
  /**
   * @param {!proto.plant_types.DestroyPlantTypeRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  types_pb.EmptyReply.deserializeBinary
);


/**
 * @param {!proto.plant_types.DestroyPlantTypeRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.types.EmptyReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.types.EmptyReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.plant_types.V1PlantTypesClient.prototype.destroyPlantType =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/plant_types.V1PlantTypes/DestroyPlantType',
      request,
      metadata || {},
      methodDescriptor_V1PlantTypes_DestroyPlantType,
      callback);
};


/**
 * @param {!proto.plant_types.DestroyPlantTypeRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.types.EmptyReply>}
 *     Promise that resolves to the response
 */
proto.plant_types.V1PlantTypesPromiseClient.prototype.destroyPlantType =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/plant_types.V1PlantTypes/DestroyPlantType',
      request,
      metadata || {},
      methodDescriptor_V1PlantTypes_DestroyPlantType);
};


module.exports = proto.plant_types;

