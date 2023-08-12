// @generated by protoc-gen-es v1.3.0
// @generated from file kotobako/v1/kotobako.proto (package kotobako.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

// buf:lint:ignore PACKAGE_DIRECTORY_MATCH

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message kotobako.v1.HealthRequest
 */
export declare class HealthRequest extends Message<HealthRequest> {
  constructor(data?: PartialMessage<HealthRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "kotobako.v1.HealthRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): HealthRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): HealthRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): HealthRequest;

  static equals(a: HealthRequest | PlainMessage<HealthRequest> | undefined, b: HealthRequest | PlainMessage<HealthRequest> | undefined): boolean;
}

/**
 * @generated from message kotobako.v1.HealthResponse
 */
export declare class HealthResponse extends Message<HealthResponse> {
  /**
   * @generated from field: string status = 1;
   */
  status: string;

  constructor(data?: PartialMessage<HealthResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "kotobako.v1.HealthResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): HealthResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): HealthResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): HealthResponse;

  static equals(a: HealthResponse | PlainMessage<HealthResponse> | undefined, b: HealthResponse | PlainMessage<HealthResponse> | undefined): boolean;
}

/**
 * @generated from message kotobako.v1.ListQuotesRequest
 */
export declare class ListQuotesRequest extends Message<ListQuotesRequest> {
  constructor(data?: PartialMessage<ListQuotesRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "kotobako.v1.ListQuotesRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListQuotesRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListQuotesRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListQuotesRequest;

  static equals(a: ListQuotesRequest | PlainMessage<ListQuotesRequest> | undefined, b: ListQuotesRequest | PlainMessage<ListQuotesRequest> | undefined): boolean;
}

/**
 * @generated from message kotobako.v1.ListQuotesResponse
 */
export declare class ListQuotesResponse extends Message<ListQuotesResponse> {
  /**
   * @generated from field: repeated kotobako.v1.Quote quotes = 1;
   */
  quotes: Quote[];

  constructor(data?: PartialMessage<ListQuotesResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "kotobako.v1.ListQuotesResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListQuotesResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListQuotesResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListQuotesResponse;

  static equals(a: ListQuotesResponse | PlainMessage<ListQuotesResponse> | undefined, b: ListQuotesResponse | PlainMessage<ListQuotesResponse> | undefined): boolean;
}

/**
 * @generated from message kotobako.v1.GetQuoteRequest
 */
export declare class GetQuoteRequest extends Message<GetQuoteRequest> {
  /**
   * @generated from field: string quoteId = 1;
   */
  quoteId: string;

  constructor(data?: PartialMessage<GetQuoteRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "kotobako.v1.GetQuoteRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetQuoteRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetQuoteRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetQuoteRequest;

  static equals(a: GetQuoteRequest | PlainMessage<GetQuoteRequest> | undefined, b: GetQuoteRequest | PlainMessage<GetQuoteRequest> | undefined): boolean;
}

/**
 * @generated from message kotobako.v1.GetQuoteResponse
 */
export declare class GetQuoteResponse extends Message<GetQuoteResponse> {
  /**
   * @generated from field: string quoteId = 1;
   */
  quoteId: string;

  /**
   * @generated from field: string authorName = 2;
   */
  authorName: string;

  /**
   * @generated from field: string quoteMediaType = 3;
   */
  quoteMediaType: string;

  /**
   * @generated from field: string quoteSourceName = 4;
   */
  quoteSourceName: string;

  /**
   * @generated from field: string sentence = 5;
   */
  sentence: string;

  constructor(data?: PartialMessage<GetQuoteResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "kotobako.v1.GetQuoteResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetQuoteResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetQuoteResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetQuoteResponse;

  static equals(a: GetQuoteResponse | PlainMessage<GetQuoteResponse> | undefined, b: GetQuoteResponse | PlainMessage<GetQuoteResponse> | undefined): boolean;
}

/**
 * @generated from message kotobako.v1.Quote
 */
export declare class Quote extends Message<Quote> {
  /**
   * @generated from field: string quoteId = 1;
   */
  quoteId: string;

  /**
   * @generated from field: string authorName = 2;
   */
  authorName: string;

  /**
   * @generated from field: string quoteMediaType = 3;
   */
  quoteMediaType: string;

  /**
   * @generated from field: string quoteSourceName = 4;
   */
  quoteSourceName: string;

  /**
   * @generated from field: string sentence = 5;
   */
  sentence: string;

  constructor(data?: PartialMessage<Quote>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "kotobako.v1.Quote";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Quote;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Quote;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Quote;

  static equals(a: Quote | PlainMessage<Quote> | undefined, b: Quote | PlainMessage<Quote> | undefined): boolean;
}

