// @generated by protobuf-ts 2.9.0
// @generated from protobuf file "auth/auth.proto" (package "auth", syntax proto3)
// tslint:disable
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { AuthAPI } from "./auth";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { GetJWTResponse } from "./auth";
import type { GetJWTRequest } from "./auth";
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * AuthAPI is the main authentication API between the backend and the frontends.
 *
 * @generated from protobuf service auth.AuthAPI
 */
export interface IAuthAPIClient {
    /**
     * @generated from protobuf rpc: GetJWT(auth.GetJWTRequest) returns (auth.GetJWTResponse);
     */
    getJWT(input: GetJWTRequest, options?: RpcOptions): UnaryCall<GetJWTRequest, GetJWTResponse>;
}
/**
 * AuthAPI is the main authentication API between the backend and the frontends.
 *
 * @generated from protobuf service auth.AuthAPI
 */
export class AuthAPIClient implements IAuthAPIClient, ServiceInfo {
    typeName = AuthAPI.typeName;
    methods = AuthAPI.methods;
    options = AuthAPI.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * @generated from protobuf rpc: GetJWT(auth.GetJWTRequest) returns (auth.GetJWTResponse);
     */
    getJWT(input: GetJWTRequest, options?: RpcOptions): UnaryCall<GetJWTRequest, GetJWTResponse> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetJWTRequest, GetJWTResponse>("unary", this._transport, method, opt, input);
    }
}