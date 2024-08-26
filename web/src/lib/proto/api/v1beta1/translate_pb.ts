// @generated by protoc-gen-es v1.10.0 with parameter "target=ts"
// @generated from file api/v1beta1/translate.proto (package api.v1beta1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum api.v1beta1.Language
 */
export enum Language {
  /**
   * @generated from enum value: UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: ENGLISH = 1;
   */
  ENGLISH = 1,

  /**
   * @generated from enum value: SPANISH = 2;
   */
  SPANISH = 2,

  /**
   * @generated from enum value: CHINESE = 3;
   */
  CHINESE = 3,

  /**
   * @generated from enum value: KOREAN = 4;
   */
  KOREAN = 4,

  /**
   * @generated from enum value: JAPANESE = 5;
   */
  JAPANESE = 5,

  /**
   * @generated from enum value: GERMAN = 6;
   */
  GERMAN = 6,

  /**
   * @generated from enum value: RUSSIAN = 7;
   */
  RUSSIAN = 7,

  /**
   * @generated from enum value: FRENCH = 8;
   */
  FRENCH = 8,

  /**
   * @generated from enum value: DUTCH = 9;
   */
  DUTCH = 9,

  /**
   * @generated from enum value: ITALIAN = 10;
   */
  ITALIAN = 10,

  /**
   * @generated from enum value: INDONESIAN = 11;
   */
  INDONESIAN = 11,

  /**
   * @generated from enum value: PORTUGUESE = 12;
   */
  PORTUGUESE = 12,

  /**
   * @generated from enum value: TAIWANESE = 13;
   */
  TAIWANESE = 13,
}
// Retrieve enum metadata with: proto3.getEnumType(Language)
proto3.util.setEnumType(Language, "api.v1beta1.Language", [
  { no: 0, name: "UNSPECIFIED" },
  { no: 1, name: "ENGLISH" },
  { no: 2, name: "SPANISH" },
  { no: 3, name: "CHINESE" },
  { no: 4, name: "KOREAN" },
  { no: 5, name: "JAPANESE" },
  { no: 6, name: "GERMAN" },
  { no: 7, name: "RUSSIAN" },
  { no: 8, name: "FRENCH" },
  { no: 9, name: "DUTCH" },
  { no: 10, name: "ITALIAN" },
  { no: 11, name: "INDONESIAN" },
  { no: 12, name: "PORTUGUESE" },
  { no: 13, name: "TAIWANESE" },
]);

/**
 * @generated from message api.v1beta1.HealthzRequest
 */
export class HealthzRequest extends Message<HealthzRequest> {
  constructor(data?: PartialMessage<HealthzRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "api.v1beta1.HealthzRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): HealthzRequest {
    return new HealthzRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): HealthzRequest {
    return new HealthzRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): HealthzRequest {
    return new HealthzRequest().fromJsonString(jsonString, options);
  }

  static equals(a: HealthzRequest | PlainMessage<HealthzRequest> | undefined, b: HealthzRequest | PlainMessage<HealthzRequest> | undefined): boolean {
    return proto3.util.equals(HealthzRequest, a, b);
  }
}

/**
 * @generated from message api.v1beta1.HealthzResponse
 */
export class HealthzResponse extends Message<HealthzResponse> {
  /**
   * @generated from field: bool healthy = 1;
   */
  healthy = false;

  /**
   * @generated from field: string message = 2;
   */
  message = "";

  /**
   * @generated from field: string version = 3;
   */
  version = "";

  constructor(data?: PartialMessage<HealthzResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "api.v1beta1.HealthzResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "healthy", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
    { no: 2, name: "message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "version", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): HealthzResponse {
    return new HealthzResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): HealthzResponse {
    return new HealthzResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): HealthzResponse {
    return new HealthzResponse().fromJsonString(jsonString, options);
  }

  static equals(a: HealthzResponse | PlainMessage<HealthzResponse> | undefined, b: HealthzResponse | PlainMessage<HealthzResponse> | undefined): boolean {
    return proto3.util.equals(HealthzResponse, a, b);
  }
}

/**
 * @generated from message api.v1beta1.LLMCredentials
 */
export class LLMCredentials extends Message<LLMCredentials> {
  /**
   * @generated from field: string api_key = 1;
   */
  apiKey = "";

  constructor(data?: PartialMessage<LLMCredentials>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "api.v1beta1.LLMCredentials";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "api_key", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LLMCredentials {
    return new LLMCredentials().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LLMCredentials {
    return new LLMCredentials().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LLMCredentials {
    return new LLMCredentials().fromJsonString(jsonString, options);
  }

  static equals(a: LLMCredentials | PlainMessage<LLMCredentials> | undefined, b: LLMCredentials | PlainMessage<LLMCredentials> | undefined): boolean {
    return proto3.util.equals(LLMCredentials, a, b);
  }
}

/**
 * @generated from message api.v1beta1.LLMOptions
 */
export class LLMOptions extends Message<LLMOptions> {
  /**
   * @generated from field: api.v1beta1.LLMOptions.Provider provider = 10;
   */
  provider = LLMOptions_Provider.UNSPECIFIED;

  /**
   * @generated from field: optional api.v1beta1.LLMCredentials credentials = 20;
   */
  credentials?: LLMCredentials;

  constructor(data?: PartialMessage<LLMOptions>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "api.v1beta1.LLMOptions";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 10, name: "provider", kind: "enum", T: proto3.getEnumType(LLMOptions_Provider) },
    { no: 20, name: "credentials", kind: "message", T: LLMCredentials, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LLMOptions {
    return new LLMOptions().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LLMOptions {
    return new LLMOptions().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LLMOptions {
    return new LLMOptions().fromJsonString(jsonString, options);
  }

  static equals(a: LLMOptions | PlainMessage<LLMOptions> | undefined, b: LLMOptions | PlainMessage<LLMOptions> | undefined): boolean {
    return proto3.util.equals(LLMOptions, a, b);
  }
}

/**
 * @generated from enum api.v1beta1.LLMOptions.Provider
 */
export enum LLMOptions_Provider {
  /**
   * @generated from enum value: UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: SERVER = 1;
   */
  SERVER = 1,

  /**
   * @generated from enum value: AISTUDIO = 2;
   */
  AISTUDIO = 2,

  /**
   * @generated from enum value: OPENAI = 3;
   */
  OPENAI = 3,

  /**
   * @generated from enum value: ANTHROPIC = 4;
   */
  ANTHROPIC = 4,
}
// Retrieve enum metadata with: proto3.getEnumType(LLMOptions_Provider)
proto3.util.setEnumType(LLMOptions_Provider, "api.v1beta1.LLMOptions.Provider", [
  { no: 0, name: "UNSPECIFIED" },
  { no: 1, name: "SERVER" },
  { no: 2, name: "AISTUDIO" },
  { no: 3, name: "OPENAI" },
  { no: 4, name: "ANTHROPIC" },
]);

/**
 * @generated from message api.v1beta1.TranslateRequest
 */
export class TranslateRequest extends Message<TranslateRequest> {
  /**
   * @generated from field: string text = 10;
   */
  text = "";

  /**
   * @generated from field: api.v1beta1.Language target_language = 20;
   */
  targetLanguage = Language.UNSPECIFIED;

  /**
   * @generated from field: api.v1beta1.LLMOptions options = 30;
   */
  options?: LLMOptions;

  constructor(data?: PartialMessage<TranslateRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "api.v1beta1.TranslateRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 10, name: "text", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 20, name: "target_language", kind: "enum", T: proto3.getEnumType(Language) },
    { no: 30, name: "options", kind: "message", T: LLMOptions },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): TranslateRequest {
    return new TranslateRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): TranslateRequest {
    return new TranslateRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): TranslateRequest {
    return new TranslateRequest().fromJsonString(jsonString, options);
  }

  static equals(a: TranslateRequest | PlainMessage<TranslateRequest> | undefined, b: TranslateRequest | PlainMessage<TranslateRequest> | undefined): boolean {
    return proto3.util.equals(TranslateRequest, a, b);
  }
}

