/**
 * @fileoverview gRPC-Web generated client stub for users
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
proto.users = require('./users_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.users.V1UsersClient =
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
proto.users.V1UsersPromiseClient =
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
 *   !proto.users.LoginRequest,
 *   !proto.users.LoginReply>}
 */
const methodDescriptor_V1Users_Login = new grpc.web.MethodDescriptor(
  '/users.V1Users/Login',
  grpc.web.MethodType.UNARY,
  proto.users.LoginRequest,
  proto.users.LoginReply,
  /**
   * @param {!proto.users.LoginRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.users.LoginReply.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.users.LoginRequest,
 *   !proto.users.LoginReply>}
 */
const methodInfo_V1Users_Login = new grpc.web.AbstractClientBase.MethodInfo(
  proto.users.LoginReply,
  /**
   * @param {!proto.users.LoginRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.users.LoginReply.deserializeBinary
);


/**
 * @param {!proto.users.LoginRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.users.LoginReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.users.LoginReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.users.V1UsersClient.prototype.login =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/users.V1Users/Login',
      request,
      metadata || {},
      methodDescriptor_V1Users_Login,
      callback);
};


/**
 * @param {!proto.users.LoginRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.users.LoginReply>}
 *     Promise that resolves to the response
 */
proto.users.V1UsersPromiseClient.prototype.login =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/users.V1Users/Login',
      request,
      metadata || {},
      methodDescriptor_V1Users_Login);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.users.VerifyJWTRequest,
 *   !proto.users.VerifyJWTReply>}
 */
const methodDescriptor_V1Users_VerifyJWT = new grpc.web.MethodDescriptor(
  '/users.V1Users/VerifyJWT',
  grpc.web.MethodType.UNARY,
  proto.users.VerifyJWTRequest,
  proto.users.VerifyJWTReply,
  /**
   * @param {!proto.users.VerifyJWTRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.users.VerifyJWTReply.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.users.VerifyJWTRequest,
 *   !proto.users.VerifyJWTReply>}
 */
const methodInfo_V1Users_VerifyJWT = new grpc.web.AbstractClientBase.MethodInfo(
  proto.users.VerifyJWTReply,
  /**
   * @param {!proto.users.VerifyJWTRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.users.VerifyJWTReply.deserializeBinary
);


/**
 * @param {!proto.users.VerifyJWTRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.users.VerifyJWTReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.users.VerifyJWTReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.users.V1UsersClient.prototype.verifyJWT =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/users.V1Users/VerifyJWT',
      request,
      metadata || {},
      methodDescriptor_V1Users_VerifyJWT,
      callback);
};


/**
 * @param {!proto.users.VerifyJWTRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.users.VerifyJWTReply>}
 *     Promise that resolves to the response
 */
proto.users.V1UsersPromiseClient.prototype.verifyJWT =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/users.V1Users/VerifyJWT',
      request,
      metadata || {},
      methodDescriptor_V1Users_VerifyJWT);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.users.AddUserRequest,
 *   !proto.users.AddUserReply>}
 */
const methodDescriptor_V1Users_AddUser = new grpc.web.MethodDescriptor(
  '/users.V1Users/AddUser',
  grpc.web.MethodType.UNARY,
  proto.users.AddUserRequest,
  proto.users.AddUserReply,
  /**
   * @param {!proto.users.AddUserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.users.AddUserReply.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.users.AddUserRequest,
 *   !proto.users.AddUserReply>}
 */
const methodInfo_V1Users_AddUser = new grpc.web.AbstractClientBase.MethodInfo(
  proto.users.AddUserReply,
  /**
   * @param {!proto.users.AddUserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.users.AddUserReply.deserializeBinary
);


/**
 * @param {!proto.users.AddUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.users.AddUserReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.users.AddUserReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.users.V1UsersClient.prototype.addUser =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/users.V1Users/AddUser',
      request,
      metadata || {},
      methodDescriptor_V1Users_AddUser,
      callback);
};


/**
 * @param {!proto.users.AddUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.users.AddUserReply>}
 *     Promise that resolves to the response
 */
proto.users.V1UsersPromiseClient.prototype.addUser =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/users.V1Users/AddUser',
      request,
      metadata || {},
      methodDescriptor_V1Users_AddUser);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.users.GetUserRequest,
 *   !proto.users.GetUserReply>}
 */
const methodDescriptor_V1Users_GetUser = new grpc.web.MethodDescriptor(
  '/users.V1Users/GetUser',
  grpc.web.MethodType.UNARY,
  proto.users.GetUserRequest,
  proto.users.GetUserReply,
  /**
   * @param {!proto.users.GetUserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.users.GetUserReply.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.users.GetUserRequest,
 *   !proto.users.GetUserReply>}
 */
const methodInfo_V1Users_GetUser = new grpc.web.AbstractClientBase.MethodInfo(
  proto.users.GetUserReply,
  /**
   * @param {!proto.users.GetUserRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.users.GetUserReply.deserializeBinary
);


/**
 * @param {!proto.users.GetUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.users.GetUserReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.users.GetUserReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.users.V1UsersClient.prototype.getUser =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/users.V1Users/GetUser',
      request,
      metadata || {},
      methodDescriptor_V1Users_GetUser,
      callback);
};


/**
 * @param {!proto.users.GetUserRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.users.GetUserReply>}
 *     Promise that resolves to the response
 */
proto.users.V1UsersPromiseClient.prototype.getUser =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/users.V1Users/GetUser',
      request,
      metadata || {},
      methodDescriptor_V1Users_GetUser);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.users.FindUsersRequest,
 *   !proto.users.FindUsersReply>}
 */
const methodDescriptor_V1Users_FindUsers = new grpc.web.MethodDescriptor(
  '/users.V1Users/FindUsers',
  grpc.web.MethodType.UNARY,
  proto.users.FindUsersRequest,
  proto.users.FindUsersReply,
  /**
   * @param {!proto.users.FindUsersRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.users.FindUsersReply.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.users.FindUsersRequest,
 *   !proto.users.FindUsersReply>}
 */
const methodInfo_V1Users_FindUsers = new grpc.web.AbstractClientBase.MethodInfo(
  proto.users.FindUsersReply,
  /**
   * @param {!proto.users.FindUsersRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.users.FindUsersReply.deserializeBinary
);


/**
 * @param {!proto.users.FindUsersRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.users.FindUsersReply)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.users.FindUsersReply>|undefined}
 *     The XHR Node Readable Stream
 */
proto.users.V1UsersClient.prototype.findUsers =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/users.V1Users/FindUsers',
      request,
      metadata || {},
      methodDescriptor_V1Users_FindUsers,
      callback);
};


/**
 * @param {!proto.users.FindUsersRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.users.FindUsersReply>}
 *     Promise that resolves to the response
 */
proto.users.V1UsersPromiseClient.prototype.findUsers =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/users.V1Users/FindUsers',
      request,
      metadata || {},
      methodDescriptor_V1Users_FindUsers);
};


module.exports = proto.users;

