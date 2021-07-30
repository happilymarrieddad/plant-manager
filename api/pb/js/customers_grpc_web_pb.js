/**
 * @fileoverview gRPC-Web generated client stub for customers
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
proto.customers = require('./customers_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.customers.V1CustomersClient =
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
proto.customers.V1CustomersPromiseClient =
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
 *   !proto.customers.GetCustomerRequest,
 *   !proto.customers.GetCustomerReply>}
 */
const methodDescriptor_V1Customers_GetCustomer = new grpc.web.MethodDescriptor(
  '/customers.V1Customers/GetCustomer',
  grpc.web.MethodType.UNARY,
  proto.customers.GetCustomerRequest,
  proto.customers.GetCustomerReply,
  /**
   * @param {!proto.customers.GetCustomerRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.customers.GetCustomerReply.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.customers.GetCustomerRequest,
 *   !proto.customers.GetCustomerReply>}
 */
const methodInfo_V1Customers_GetCustomer = new grpc.web.AbstractClientBase.MethodInfo(
  proto.customers.GetCustomerReply,
  /**
   * @param {!proto.customers.GetCustomerRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.customers.GetCustomerReply.deserializeBinary
);


/**
 * @param {!proto.customers.GetCustomerRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.customers.GetCustomerReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.customers.GetCustomerReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.customers.V1CustomersClient.prototype.getCustomer =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/customers.V1Customers/GetCustomer',
      request,
      metadata || {},
      methodDescriptor_V1Customers_GetCustomer,
      callback);
};


/**
 * @param {!proto.customers.GetCustomerRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.customers.GetCustomerReply>}
 *     Promise that resolves to the response
 */
proto.customers.V1CustomersPromiseClient.prototype.getCustomer =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/customers.V1Customers/GetCustomer',
      request,
      metadata || {},
      methodDescriptor_V1Customers_GetCustomer);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.customers.FindCustomersRequest,
 *   !proto.customers.FindCustomersReply>}
 */
const methodDescriptor_V1Customers_FindCustomers = new grpc.web.MethodDescriptor(
  '/customers.V1Customers/FindCustomers',
  grpc.web.MethodType.UNARY,
  proto.customers.FindCustomersRequest,
  proto.customers.FindCustomersReply,
  /**
   * @param {!proto.customers.FindCustomersRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.customers.FindCustomersReply.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.customers.FindCustomersRequest,
 *   !proto.customers.FindCustomersReply>}
 */
const methodInfo_V1Customers_FindCustomers = new grpc.web.AbstractClientBase.MethodInfo(
  proto.customers.FindCustomersReply,
  /**
   * @param {!proto.customers.FindCustomersRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.customers.FindCustomersReply.deserializeBinary
);


/**
 * @param {!proto.customers.FindCustomersRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.customers.FindCustomersReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.customers.FindCustomersReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.customers.V1CustomersClient.prototype.findCustomers =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/customers.V1Customers/FindCustomers',
      request,
      metadata || {},
      methodDescriptor_V1Customers_FindCustomers,
      callback);
};


/**
 * @param {!proto.customers.FindCustomersRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.customers.FindCustomersReply>}
 *     Promise that resolves to the response
 */
proto.customers.V1CustomersPromiseClient.prototype.findCustomers =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/customers.V1Customers/FindCustomers',
      request,
      metadata || {},
      methodDescriptor_V1Customers_FindCustomers);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.customers.CreateCustomerRequest,
 *   !proto.customers.CreateCustomerReply>}
 */
const methodDescriptor_V1Customers_CreateCustomer = new grpc.web.MethodDescriptor(
  '/customers.V1Customers/CreateCustomer',
  grpc.web.MethodType.UNARY,
  proto.customers.CreateCustomerRequest,
  proto.customers.CreateCustomerReply,
  /**
   * @param {!proto.customers.CreateCustomerRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.customers.CreateCustomerReply.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.customers.CreateCustomerRequest,
 *   !proto.customers.CreateCustomerReply>}
 */
const methodInfo_V1Customers_CreateCustomer = new grpc.web.AbstractClientBase.MethodInfo(
  proto.customers.CreateCustomerReply,
  /**
   * @param {!proto.customers.CreateCustomerRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.customers.CreateCustomerReply.deserializeBinary
);


/**
 * @param {!proto.customers.CreateCustomerRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.customers.CreateCustomerReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.customers.CreateCustomerReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.customers.V1CustomersClient.prototype.createCustomer =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/customers.V1Customers/CreateCustomer',
      request,
      metadata || {},
      methodDescriptor_V1Customers_CreateCustomer,
      callback);
};


/**
 * @param {!proto.customers.CreateCustomerRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.customers.CreateCustomerReply>}
 *     Promise that resolves to the response
 */
proto.customers.V1CustomersPromiseClient.prototype.createCustomer =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/customers.V1Customers/CreateCustomer',
      request,
      metadata || {},
      methodDescriptor_V1Customers_CreateCustomer);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.customers.DestroyCustomerRequest,
 *   !proto.types.EmptyReply>}
 */
const methodDescriptor_V1Customers_DestroyCustomer = new grpc.web.MethodDescriptor(
  '/customers.V1Customers/DestroyCustomer',
  grpc.web.MethodType.UNARY,
  proto.customers.DestroyCustomerRequest,
  types_pb.EmptyReply,
  /**
   * @param {!proto.customers.DestroyCustomerRequest} request
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
 *   !proto.customers.DestroyCustomerRequest,
 *   !proto.types.EmptyReply>}
 */
const methodInfo_V1Customers_DestroyCustomer = new grpc.web.AbstractClientBase.MethodInfo(
  types_pb.EmptyReply,
  /**
   * @param {!proto.customers.DestroyCustomerRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  types_pb.EmptyReply.deserializeBinary
);


/**
 * @param {!proto.customers.DestroyCustomerRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.types.EmptyReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.types.EmptyReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.customers.V1CustomersClient.prototype.destroyCustomer =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/customers.V1Customers/DestroyCustomer',
      request,
      metadata || {},
      methodDescriptor_V1Customers_DestroyCustomer,
      callback);
};


/**
 * @param {!proto.customers.DestroyCustomerRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.types.EmptyReply>}
 *     Promise that resolves to the response
 */
proto.customers.V1CustomersPromiseClient.prototype.destroyCustomer =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/customers.V1Customers/DestroyCustomer',
      request,
      metadata || {},
      methodDescriptor_V1Customers_DestroyCustomer);
};


module.exports = proto.customers;

