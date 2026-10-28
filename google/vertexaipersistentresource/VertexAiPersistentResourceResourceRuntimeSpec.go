// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaipersistentresource


type VertexAiPersistentResourceResourceRuntimeSpec struct {
	// ray_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_persistent_resource#ray_spec VertexAiPersistentResource#ray_spec}
	RaySpec *VertexAiPersistentResourceResourceRuntimeSpecRaySpec `field:"optional" json:"raySpec" yaml:"raySpec"`
	// service_account_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_persistent_resource#service_account_spec VertexAiPersistentResource#service_account_spec}
	ServiceAccountSpec *VertexAiPersistentResourceResourceRuntimeSpecServiceAccountSpec `field:"optional" json:"serviceAccountSpec" yaml:"serviceAccountSpec"`
}

