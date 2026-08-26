// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaipersistentresource


type VertexAiPersistentResourceResourceRuntimeSpec struct {
	// service_account_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/vertex_ai_persistent_resource#service_account_spec VertexAiPersistentResource#service_account_spec}
	ServiceAccountSpec *VertexAiPersistentResourceResourceRuntimeSpecServiceAccountSpec `field:"optional" json:"serviceAccountSpec" yaml:"serviceAccountSpec"`
}

