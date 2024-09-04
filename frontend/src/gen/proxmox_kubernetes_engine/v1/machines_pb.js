// @generated by protoc-gen-es v1.8.0
// @generated from file proxmox_kubernetes_engine/v1/machines.proto (package proxmox_kubernetes_engine.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * @generated from enum proxmox_kubernetes_engine.v1.State
 */
export const State = /*@__PURE__*/ proto3.makeEnum(
  "proxmox_kubernetes_engine.v1.State",
  [
    {no: 0, name: "UNKNOWN"},
    {no: 1, name: "STARTING"},
    {no: 2, name: "RUNNING"},
    {no: 3, name: "STOPPING"},
    {no: 4, name: "STOPPED"},
    {no: 5, name: "ERROR"},
    {no: 6, name: "CREATING"},
    {no: 7, name: "CREATED"},
    {no: 8, name: "DESTROYING"},
    {no: 9, name: "DESTROYED"},
    {no: 10, name: "INITIALIZING"},
    {no: 11, name: "TO_DESTROY"},
  ],
);

/**
 * @generated from message proxmox_kubernetes_engine.v1.GetMachineRequest
 */
export const GetMachineRequest = /*@__PURE__*/ proto3.makeMessageType(
  "proxmox_kubernetes_engine.v1.GetMachineRequest",
  () => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message proxmox_kubernetes_engine.v1.ListMachinesRequest
 */
export const ListMachinesRequest = /*@__PURE__*/ proto3.makeMessageType(
  "proxmox_kubernetes_engine.v1.ListMachinesRequest",
  () => [
    { no: 1, name: "parent", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "page_size", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "page_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message proxmox_kubernetes_engine.v1.ListMachinesResponse
 */
export const ListMachinesResponse = /*@__PURE__*/ proto3.makeMessageType(
  "proxmox_kubernetes_engine.v1.ListMachinesResponse",
  () => [
    { no: 1, name: "machines", kind: "message", T: Machine, repeated: true },
    { no: 2, name: "next_page_token", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message proxmox_kubernetes_engine.v1.Machine
 */
export const Machine = /*@__PURE__*/ proto3.makeMessageType(
  "proxmox_kubernetes_engine.v1.Machine",
  () => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "image", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "memory", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "cpus", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 5, name: "state", kind: "enum", T: proto3.getEnumType(State) },
    { no: 6, name: "group", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "role", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 8, name: "node", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 9, name: "vmid", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 10, name: "created_at", kind: "message", T: Timestamp },
    { no: 11, name: "updated_at", kind: "message", T: Timestamp },
    { no: 12, name: "display_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);
