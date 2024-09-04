// @generated by protoc-gen-es v1.8.0
// @generated from file proxmox_kubernetes_engine/v1/machines.proto (package proxmox_kubernetes_engine.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage, Timestamp } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from enum proxmox_kubernetes_engine.v1.State
 */
export declare enum State {
  /**
   * @generated from enum value: UNKNOWN = 0;
   */
  UNKNOWN = 0,

  /**
   * @generated from enum value: STARTING = 1;
   */
  STARTING = 1,

  /**
   * @generated from enum value: RUNNING = 2;
   */
  RUNNING = 2,

  /**
   * @generated from enum value: STOPPING = 3;
   */
  STOPPING = 3,

  /**
   * @generated from enum value: STOPPED = 4;
   */
  STOPPED = 4,

  /**
   * Terminal
   *
   * @generated from enum value: ERROR = 5;
   */
  ERROR = 5,

  /**
   * The machine reconciler has picked up the machine and is attempting
   * to submit the machine to proxmox to create
   *
   * @generated from enum value: CREATING = 6;
   */
  CREATING = 6,

  /**
   * @generated from enum value: CREATED = 7;
   */
  CREATED = 7,

  /**
   * @generated from enum value: DESTROYING = 8;
   */
  DESTROYING = 8,

  /**
   * Terminal
   *
   * @generated from enum value: DESTROYED = 9;
   */
  DESTROYED = 9,

  /**
   * Initializing is a status given when configuration is still
   * being reconciled on a machine but has not been requested to
   * create on proxmox yet.
   *
   * @generated from enum value: INITIALIZING = 10;
   */
  INITIALIZING = 10,

  /**
   * @generated from enum value: TO_DESTROY = 11;
   */
  TO_DESTROY = 11,
}

/**
 * @generated from message proxmox_kubernetes_engine.v1.GetMachineRequest
 */
export declare class GetMachineRequest extends Message<GetMachineRequest> {
  /**
   * @generated from field: string name = 1;
   */
  name: string;

  constructor(data?: PartialMessage<GetMachineRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "proxmox_kubernetes_engine.v1.GetMachineRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetMachineRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetMachineRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetMachineRequest;

  static equals(a: GetMachineRequest | PlainMessage<GetMachineRequest> | undefined, b: GetMachineRequest | PlainMessage<GetMachineRequest> | undefined): boolean;
}

/**
 * @generated from message proxmox_kubernetes_engine.v1.ListMachinesRequest
 */
export declare class ListMachinesRequest extends Message<ListMachinesRequest> {
  /**
   * @generated from field: string parent = 1;
   */
  parent: string;

  /**
   * @generated from field: int32 page_size = 2;
   */
  pageSize: number;

  /**
   * @generated from field: string page_token = 3;
   */
  pageToken: string;

  constructor(data?: PartialMessage<ListMachinesRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "proxmox_kubernetes_engine.v1.ListMachinesRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListMachinesRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListMachinesRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListMachinesRequest;

  static equals(a: ListMachinesRequest | PlainMessage<ListMachinesRequest> | undefined, b: ListMachinesRequest | PlainMessage<ListMachinesRequest> | undefined): boolean;
}

/**
 * @generated from message proxmox_kubernetes_engine.v1.ListMachinesResponse
 */
export declare class ListMachinesResponse extends Message<ListMachinesResponse> {
  /**
   * @generated from field: repeated proxmox_kubernetes_engine.v1.Machine machines = 1;
   */
  machines: Machine[];

  /**
   * @generated from field: string next_page_token = 2;
   */
  nextPageToken: string;

  constructor(data?: PartialMessage<ListMachinesResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "proxmox_kubernetes_engine.v1.ListMachinesResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListMachinesResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListMachinesResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListMachinesResponse;

  static equals(a: ListMachinesResponse | PlainMessage<ListMachinesResponse> | undefined, b: ListMachinesResponse | PlainMessage<ListMachinesResponse> | undefined): boolean;
}

/**
 * @generated from message proxmox_kubernetes_engine.v1.Machine
 */
export declare class Machine extends Message<Machine> {
  /**
   * @generated from field: string name = 1;
   */
  name: string;

  /**
   * @generated from field: string image = 2;
   */
  image: string;

  /**
   * @generated from field: int32 memory = 3;
   */
  memory: number;

  /**
   * @generated from field: int32 cpus = 4;
   */
  cpus: number;

  /**
   * @generated from field: proxmox_kubernetes_engine.v1.State state = 5;
   */
  state: State;

  /**
   * @generated from field: string group = 6;
   */
  group: string;

  /**
   * @generated from field: string role = 7;
   */
  role: string;

  /**
   * @generated from field: string node = 8;
   */
  node: string;

  /**
   * @generated from field: int32 vmid = 9;
   */
  vmid: number;

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 10;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp updated_at = 11;
   */
  updatedAt?: Timestamp;

  /**
   * @generated from field: string display_name = 12;
   */
  displayName: string;

  constructor(data?: PartialMessage<Machine>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "proxmox_kubernetes_engine.v1.Machine";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Machine;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Machine;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Machine;

  static equals(a: Machine | PlainMessage<Machine> | undefined, b: Machine | PlainMessage<Machine> | undefined): boolean;
}

