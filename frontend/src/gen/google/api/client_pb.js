// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// @generated by protoc-gen-es v2.0.0
// @generated from file google/api/client.proto (package google.api, syntax proto3)
/* eslint-disable */

import { enumDesc, extDesc, fileDesc, messageDesc, tsEnum } from "@bufbuild/protobuf/codegenv1";
import { file_google_api_launch_stage } from "./launch_stage_pb";
import { file_google_protobuf_descriptor, file_google_protobuf_duration } from "@bufbuild/protobuf/wkt";

/**
 * Describes the file google/api/client.proto.
 */
export const file_google_api_client = /*@__PURE__*/
  fileDesc("Chdnb29nbGUvYXBpL2NsaWVudC5wcm90bxIKZ29vZ2xlLmFwaSJ0ChZDb21tb25MYW5ndWFnZVNldHRpbmdzEh4KEnJlZmVyZW5jZV9kb2NzX3VyaRgBIAEoCUICGAESOgoMZGVzdGluYXRpb25zGAIgAygOMiQuZ29vZ2xlLmFwaS5DbGllbnRMaWJyYXJ5RGVzdGluYXRpb24i+wMKFUNsaWVudExpYnJhcnlTZXR0aW5ncxIPCgd2ZXJzaW9uGAEgASgJEi0KDGxhdW5jaF9zdGFnZRgCIAEoDjIXLmdvb2dsZS5hcGkuTGF1bmNoU3RhZ2USGgoScmVzdF9udW1lcmljX2VudW1zGAMgASgIEi8KDWphdmFfc2V0dGluZ3MYFSABKAsyGC5nb29nbGUuYXBpLkphdmFTZXR0aW5ncxItCgxjcHBfc2V0dGluZ3MYFiABKAsyFy5nb29nbGUuYXBpLkNwcFNldHRpbmdzEi0KDHBocF9zZXR0aW5ncxgXIAEoCzIXLmdvb2dsZS5hcGkuUGhwU2V0dGluZ3MSMwoPcHl0aG9uX3NldHRpbmdzGBggASgLMhouZ29vZ2xlLmFwaS5QeXRob25TZXR0aW5ncxIvCg1ub2RlX3NldHRpbmdzGBkgASgLMhguZ29vZ2xlLmFwaS5Ob2RlU2V0dGluZ3MSMwoPZG90bmV0X3NldHRpbmdzGBogASgLMhouZ29vZ2xlLmFwaS5Eb3RuZXRTZXR0aW5ncxIvCg1ydWJ5X3NldHRpbmdzGBsgASgLMhguZ29vZ2xlLmFwaS5SdWJ5U2V0dGluZ3MSKwoLZ29fc2V0dGluZ3MYHCABKAsyFi5nb29nbGUuYXBpLkdvU2V0dGluZ3Mi/gIKClB1Ymxpc2hpbmcSMwoPbWV0aG9kX3NldHRpbmdzGAIgAygLMhouZ29vZ2xlLmFwaS5NZXRob2RTZXR0aW5ncxIVCg1uZXdfaXNzdWVfdXJpGGUgASgJEhkKEWRvY3VtZW50YXRpb25fdXJpGGYgASgJEhYKDmFwaV9zaG9ydF9uYW1lGGcgASgJEhQKDGdpdGh1Yl9sYWJlbBhoIAEoCRIeChZjb2Rlb3duZXJfZ2l0aHViX3RlYW1zGGkgAygJEhYKDmRvY190YWdfcHJlZml4GGogASgJEjsKDG9yZ2FuaXphdGlvbhhrIAEoDjIlLmdvb2dsZS5hcGkuQ2xpZW50TGlicmFyeU9yZ2FuaXphdGlvbhI7ChBsaWJyYXJ5X3NldHRpbmdzGG0gAygLMiEuZ29vZ2xlLmFwaS5DbGllbnRMaWJyYXJ5U2V0dGluZ3MSKQohcHJvdG9fcmVmZXJlbmNlX2RvY3VtZW50YXRpb25fdXJpGG4gASgJIuMBCgxKYXZhU2V0dGluZ3MSFwoPbGlicmFyeV9wYWNrYWdlGAEgASgJEkwKE3NlcnZpY2VfY2xhc3NfbmFtZXMYAiADKAsyLy5nb29nbGUuYXBpLkphdmFTZXR0aW5ncy5TZXJ2aWNlQ2xhc3NOYW1lc0VudHJ5EjIKBmNvbW1vbhgDIAEoCzIiLmdvb2dsZS5hcGkuQ29tbW9uTGFuZ3VhZ2VTZXR0aW5ncxo4ChZTZXJ2aWNlQ2xhc3NOYW1lc0VudHJ5EgsKA2tleRgBIAEoCRINCgV2YWx1ZRgCIAEoCToCOAEiQQoLQ3BwU2V0dGluZ3MSMgoGY29tbW9uGAEgASgLMiIuZ29vZ2xlLmFwaS5Db21tb25MYW5ndWFnZVNldHRpbmdzIkEKC1BocFNldHRpbmdzEjIKBmNvbW1vbhgBIAEoCzIiLmdvb2dsZS5hcGkuQ29tbW9uTGFuZ3VhZ2VTZXR0aW5ncyJECg5QeXRob25TZXR0aW5ncxIyCgZjb21tb24YASABKAsyIi5nb29nbGUuYXBpLkNvbW1vbkxhbmd1YWdlU2V0dGluZ3MiQgoMTm9kZVNldHRpbmdzEjIKBmNvbW1vbhgBIAEoCzIiLmdvb2dsZS5hcGkuQ29tbW9uTGFuZ3VhZ2VTZXR0aW5ncyKqAwoORG90bmV0U2V0dGluZ3MSMgoGY29tbW9uGAEgASgLMiIuZ29vZ2xlLmFwaS5Db21tb25MYW5ndWFnZVNldHRpbmdzEkkKEHJlbmFtZWRfc2VydmljZXMYAiADKAsyLy5nb29nbGUuYXBpLkRvdG5ldFNldHRpbmdzLlJlbmFtZWRTZXJ2aWNlc0VudHJ5EksKEXJlbmFtZWRfcmVzb3VyY2VzGAMgAygLMjAuZ29vZ2xlLmFwaS5Eb3RuZXRTZXR0aW5ncy5SZW5hbWVkUmVzb3VyY2VzRW50cnkSGQoRaWdub3JlZF9yZXNvdXJjZXMYBCADKAkSIAoYZm9yY2VkX25hbWVzcGFjZV9hbGlhc2VzGAUgAygJEh4KFmhhbmR3cml0dGVuX3NpZ25hdHVyZXMYBiADKAkaNgoUUmVuYW1lZFNlcnZpY2VzRW50cnkSCwoDa2V5GAEgASgJEg0KBXZhbHVlGAIgASgJOgI4ARo3ChVSZW5hbWVkUmVzb3VyY2VzRW50cnkSCwoDa2V5GAEgASgJEg0KBXZhbHVlGAIgASgJOgI4ASJCCgxSdWJ5U2V0dGluZ3MSMgoGY29tbW9uGAEgASgLMiIuZ29vZ2xlLmFwaS5Db21tb25MYW5ndWFnZVNldHRpbmdzIkAKCkdvU2V0dGluZ3MSMgoGY29tbW9uGAEgASgLMiIuZ29vZ2xlLmFwaS5Db21tb25MYW5ndWFnZVNldHRpbmdzIs8CCg5NZXRob2RTZXR0aW5ncxIQCghzZWxlY3RvchgBIAEoCRI8Cgxsb25nX3J1bm5pbmcYAiABKAsyJi5nb29nbGUuYXBpLk1ldGhvZFNldHRpbmdzLkxvbmdSdW5uaW5nEh0KFWF1dG9fcG9wdWxhdGVkX2ZpZWxkcxgDIAMoCRrNAQoLTG9uZ1J1bm5pbmcSNQoSaW5pdGlhbF9wb2xsX2RlbGF5GAEgASgLMhkuZ29vZ2xlLnByb3RvYnVmLkR1cmF0aW9uEh0KFXBvbGxfZGVsYXlfbXVsdGlwbGllchgCIAEoAhIxCg5tYXhfcG9sbF9kZWxheRgDIAEoCzIZLmdvb2dsZS5wcm90b2J1Zi5EdXJhdGlvbhI1ChJ0b3RhbF9wb2xsX3RpbWVvdXQYBCABKAsyGS5nb29nbGUucHJvdG9idWYuRHVyYXRpb24qowEKGUNsaWVudExpYnJhcnlPcmdhbml6YXRpb24SKwonQ0xJRU5UX0xJQlJBUllfT1JHQU5JWkFUSU9OX1VOU1BFQ0lGSUVEEAASCQoFQ0xPVUQQARIHCgNBRFMQAhIKCgZQSE9UT1MQAxIPCgtTVFJFRVRfVklFVxAEEgwKCFNIT1BQSU5HEAUSBwoDR0VPEAYSEQoNR0VORVJBVElWRV9BSRAHKmcKGENsaWVudExpYnJhcnlEZXN0aW5hdGlvbhIqCiZDTElFTlRfTElCUkFSWV9ERVNUSU5BVElPTl9VTlNQRUNJRklFRBAAEgoKBkdJVEhVQhAKEhMKD1BBQ0tBR0VfTUFOQUdFUhAUOkoKEG1ldGhvZF9zaWduYXR1cmUSHi5nb29nbGUucHJvdG9idWYuTWV0aG9kT3B0aW9ucxibCCADKAlSD21ldGhvZFNpZ25hdHVyZTpDCgxkZWZhdWx0X2hvc3QSHy5nb29nbGUucHJvdG9idWYuU2VydmljZU9wdGlvbnMYmQggASgJUgtkZWZhdWx0SG9zdDpDCgxvYXV0aF9zY29wZXMSHy5nb29nbGUucHJvdG9idWYuU2VydmljZU9wdGlvbnMYmgggASgJUgtvYXV0aFNjb3Blc0JpCg5jb20uZ29vZ2xlLmFwaUILQ2xpZW50UHJvdG9QAVpBZ29vZ2xlLmdvbGFuZy5vcmcvZ2VucHJvdG8vZ29vZ2xlYXBpcy9hcGkvYW5ub3RhdGlvbnM7YW5ub3RhdGlvbnOiAgRHQVBJYgZwcm90bzM", [file_google_api_launch_stage, file_google_protobuf_descriptor, file_google_protobuf_duration]);

/**
 * Describes the message google.api.CommonLanguageSettings.
 * Use `create(CommonLanguageSettingsSchema)` to create a new message.
 */
export const CommonLanguageSettingsSchema = /*@__PURE__*/
  messageDesc(file_google_api_client, 0);

/**
 * Describes the message google.api.ClientLibrarySettings.
 * Use `create(ClientLibrarySettingsSchema)` to create a new message.
 */
export const ClientLibrarySettingsSchema = /*@__PURE__*/
  messageDesc(file_google_api_client, 1);

/**
 * Describes the message google.api.Publishing.
 * Use `create(PublishingSchema)` to create a new message.
 */
export const PublishingSchema = /*@__PURE__*/
  messageDesc(file_google_api_client, 2);

/**
 * Describes the message google.api.JavaSettings.
 * Use `create(JavaSettingsSchema)` to create a new message.
 */
export const JavaSettingsSchema = /*@__PURE__*/
  messageDesc(file_google_api_client, 3);

/**
 * Describes the message google.api.CppSettings.
 * Use `create(CppSettingsSchema)` to create a new message.
 */
export const CppSettingsSchema = /*@__PURE__*/
  messageDesc(file_google_api_client, 4);

/**
 * Describes the message google.api.PhpSettings.
 * Use `create(PhpSettingsSchema)` to create a new message.
 */
export const PhpSettingsSchema = /*@__PURE__*/
  messageDesc(file_google_api_client, 5);

/**
 * Describes the message google.api.PythonSettings.
 * Use `create(PythonSettingsSchema)` to create a new message.
 */
export const PythonSettingsSchema = /*@__PURE__*/
  messageDesc(file_google_api_client, 6);

/**
 * Describes the message google.api.NodeSettings.
 * Use `create(NodeSettingsSchema)` to create a new message.
 */
export const NodeSettingsSchema = /*@__PURE__*/
  messageDesc(file_google_api_client, 7);

/**
 * Describes the message google.api.DotnetSettings.
 * Use `create(DotnetSettingsSchema)` to create a new message.
 */
export const DotnetSettingsSchema = /*@__PURE__*/
  messageDesc(file_google_api_client, 8);

/**
 * Describes the message google.api.RubySettings.
 * Use `create(RubySettingsSchema)` to create a new message.
 */
export const RubySettingsSchema = /*@__PURE__*/
  messageDesc(file_google_api_client, 9);

/**
 * Describes the message google.api.GoSettings.
 * Use `create(GoSettingsSchema)` to create a new message.
 */
export const GoSettingsSchema = /*@__PURE__*/
  messageDesc(file_google_api_client, 10);

/**
 * Describes the message google.api.MethodSettings.
 * Use `create(MethodSettingsSchema)` to create a new message.
 */
export const MethodSettingsSchema = /*@__PURE__*/
  messageDesc(file_google_api_client, 11);

/**
 * Describes the message google.api.MethodSettings.LongRunning.
 * Use `create(MethodSettings_LongRunningSchema)` to create a new message.
 */
export const MethodSettings_LongRunningSchema = /*@__PURE__*/
  messageDesc(file_google_api_client, 11, 0);

/**
 * Describes the enum google.api.ClientLibraryOrganization.
 */
export const ClientLibraryOrganizationSchema = /*@__PURE__*/
  enumDesc(file_google_api_client, 0);

/**
 * The organization for which the client libraries are being published.
 * Affects the url where generated docs are published, etc.
 *
 * @generated from enum google.api.ClientLibraryOrganization
 */
export const ClientLibraryOrganization = /*@__PURE__*/
  tsEnum(ClientLibraryOrganizationSchema);

/**
 * Describes the enum google.api.ClientLibraryDestination.
 */
export const ClientLibraryDestinationSchema = /*@__PURE__*/
  enumDesc(file_google_api_client, 1);

/**
 * To where should client libraries be published?
 *
 * @generated from enum google.api.ClientLibraryDestination
 */
export const ClientLibraryDestination = /*@__PURE__*/
  tsEnum(ClientLibraryDestinationSchema);

/**
 * A definition of a client library method signature.
 *
 * In client libraries, each proto RPC corresponds to one or more methods
 * which the end user is able to call, and calls the underlying RPC.
 * Normally, this method receives a single argument (a struct or instance
 * corresponding to the RPC request object). Defining this field will
 * add one or more overloads providing flattened or simpler method signatures
 * in some languages.
 *
 * The fields on the method signature are provided as a comma-separated
 * string.
 *
 * For example, the proto RPC and annotation:
 *
 *   rpc CreateSubscription(CreateSubscriptionRequest)
 *       returns (Subscription) {
 *     option (google.api.method_signature) = "name,topic";
 *   }
 *
 * Would add the following Java overload (in addition to the method accepting
 * the request object):
 *
 *   public final Subscription createSubscription(String name, String topic)
 *
 * The following backwards-compatibility guidelines apply:
 *
 *   * Adding this annotation to an unannotated method is backwards
 *     compatible.
 *   * Adding this annotation to a method which already has existing
 *     method signature annotations is backwards compatible if and only if
 *     the new method signature annotation is last in the sequence.
 *   * Modifying or removing an existing method signature annotation is
 *     a breaking change.
 *   * Re-ordering existing method signature annotations is a breaking
 *     change.
 *
 * @generated from extension: repeated string method_signature = 1051;
 */
export const method_signature = /*@__PURE__*/
  extDesc(file_google_api_client, 0);

/**
 * The hostname for this service.
 * This should be specified with no prefix or protocol.
 *
 * Example:
 *
 *   service Foo {
 *     option (google.api.default_host) = "foo.googleapi.com";
 *     ...
 *   }
 *
 * @generated from extension: string default_host = 1049;
 */
export const default_host = /*@__PURE__*/
  extDesc(file_google_api_client, 1);

/**
 * OAuth scopes needed for the client.
 *
 * Example:
 *
 *   service Foo {
 *     option (google.api.oauth_scopes) = \
 *       "https://www.googleapis.com/auth/cloud-platform";
 *     ...
 *   }
 *
 * If there is more than one scope, use a comma-separated string:
 *
 * Example:
 *
 *   service Foo {
 *     option (google.api.oauth_scopes) = \
 *       "https://www.googleapis.com/auth/cloud-platform,"
 *       "https://www.googleapis.com/auth/monitoring";
 *     ...
 *   }
 *
 * @generated from extension: string oauth_scopes = 1050;
 */
export const oauth_scopes = /*@__PURE__*/
  extDesc(file_google_api_client, 2);

