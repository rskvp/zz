// @generated by protobuf-ts 2.9.0
// @generated from protobuf file "auth/auth.proto" (package "auth", syntax proto3)
// tslint:disable
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import { MessageType } from "@protobuf-ts/runtime";
/**
 * Account are the credentials for authentication API.
 *
 * @generated from protobuf message auth.Account
 */
export interface Account {
    /**
     * *
     * ID of the provider used for this account.
     *
     * Based on the ID of the provider, it will check the access token on the OAuth/OIDC Provider.
     *
     * @generated from protobuf field: string provider = 1;
     */
    provider: string;
    /**
     * * Provider's type for this account, oauth or oidc.
     *
     * @generated from protobuf field: string type = 2;
     */
    type: string;
    /**
     * * The provider account ID.
     *
     * @generated from protobuf field: string provider_account_id = 3;
     */
    providerAccountId: string;
    /**
     * * The provider access_token.
     *
     * @generated from protobuf field: string access_token = 4;
     */
    accessToken: string;
}
/**
 * @generated from protobuf message auth.GetJWTRequest
 */
export interface GetJWTRequest {
    /**
     * @generated from protobuf field: auth.Account account = 1;
     */
    account?: Account;
}
/**
 * @generated from protobuf message auth.GetJWTResponse
 */
export interface GetJWTResponse {
    /**
     * @generated from protobuf field: string token = 1;
     */
    token: string;
}
// @generated message type with reflection information, may provide speed optimized methods
class Account$Type extends MessageType<Account> {
    constructor() {
        super("auth.Account", [
            { no: 1, name: "provider", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "type", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "provider_account_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "access_token", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message auth.Account
 */
export const Account = new Account$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetJWTRequest$Type extends MessageType<GetJWTRequest> {
    constructor() {
        super("auth.GetJWTRequest", [
            { no: 1, name: "account", kind: "message", T: () => Account }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message auth.GetJWTRequest
 */
export const GetJWTRequest = new GetJWTRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetJWTResponse$Type extends MessageType<GetJWTResponse> {
    constructor() {
        super("auth.GetJWTResponse", [
            { no: 1, name: "token", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message auth.GetJWTResponse
 */
export const GetJWTResponse = new GetJWTResponse$Type();
/**
 * @generated ServiceType for protobuf service auth.AuthAPI
 */
export const AuthAPI = new ServiceType("auth.AuthAPI", [
    { name: "GetJWT", options: {}, I: GetJWTRequest, O: GetJWTResponse }
]);
