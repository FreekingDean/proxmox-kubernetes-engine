// @generated by protoc-gen-es v1.8.0
// @generated from file proxmox_kubernetes_engine/v1/machine_pools.proto (package proxmox_kubernetes_engine.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { FieldMask, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message proxmox_kubernetes_engine.v1.GetMachinePoolRequest
 */
export const GetMachinePoolRequest = /*@__PURE__*/ proto3.makeMessageType(
  "proxmox_kubernetes_engine.v1.GetMachinePoolRequest",
  () => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message proxmox_kubernetes_engine.v1.ListMachinePoolsRequest
 */
export const ListMachinePoolsRequest = /*@__PURE__*/ proto3.makeMessageType(
  "proxmox_kubernetes_engine.v1.ListMachinePoolsRequest",
  () => [
    { no: 1, name: "page_size", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "page_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message proxmox_kubernetes_engine.v1.ListMachinePoolsResponse
 */
export const ListMachinePoolsResponse = /*@__PURE__*/ proto3.makeMessageType(
  "proxmox_kubernetes_engine.v1.ListMachinePoolsResponse",
  () => [
    { no: 1, name: "machine_pools", kind: "message", T: MachinePool, repeated: true },
    { no: 2, name: "next_page_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message proxmox_kubernetes_engine.v1.CreateMachinePoolRequest
 */
export const CreateMachinePoolRequest = /*@__PURE__*/ proto3.makeMessageType(
  "proxmox_kubernetes_engine.v1.CreateMachinePoolRequest",
  () => [
    { no: 1, name: "machine_pool_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "machine_pool", kind: "message", T: MachinePool },
  ],
);

/**
 * @generated from message proxmox_kubernetes_engine.v1.UpdateMachinePoolRequest
 */
export const UpdateMachinePoolRequest = /*@__PURE__*/ proto3.makeMessageType(
  "proxmox_kubernetes_engine.v1.UpdateMachinePoolRequest",
  () => [
    { no: 1, name: "machine_pool", kind: "message", T: MachinePool },
    { no: 2, name: "update_mask", kind: "message", T: FieldMask },
  ],
);

/**
 * @generated from message proxmox_kubernetes_engine.v1.DeleteMachinePoolRequest
 */
export const DeleteMachinePoolRequest = /*@__PURE__*/ proto3.makeMessageType(
  "proxmox_kubernetes_engine.v1.DeleteMachinePoolRequest",
  () => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message proxmox_kubernetes_engine.v1.MachinePool
 */
export const MachinePool = /*@__PURE__*/ proto3.makeMessageType(
  "proxmox_kubernetes_engine.v1.MachinePool",
  () => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "image", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "memory", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "cpus", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "group", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "display_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

