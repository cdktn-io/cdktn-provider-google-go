// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vectorsearchindex


type VectorSearchIndexDedicatedInfrastructure struct {
	// autoscaling_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/vector_search_index#autoscaling_spec VectorSearchIndex#autoscaling_spec}
	AutoscalingSpec *VectorSearchIndexDedicatedInfrastructureAutoscalingSpec `field:"optional" json:"autoscalingSpec" yaml:"autoscalingSpec"`
	// Mode of the dedicated infrastructure. Defaults to 'PERFORMANCE_OPTIMIZED'. Possible values: ["MODE_UNSPECIFIED", "STORAGE_OPTIMIZED", "PERFORMANCE_OPTIMIZED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/vector_search_index#mode VectorSearchIndex#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

