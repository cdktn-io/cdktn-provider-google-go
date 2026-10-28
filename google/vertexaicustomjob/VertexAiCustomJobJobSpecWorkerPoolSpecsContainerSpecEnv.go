// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaicustomjob


type VertexAiCustomJobJobSpecWorkerPoolSpecsContainerSpecEnv struct {
	// Name of the environment variable. Must be a valid C identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_custom_job#name VertexAiCustomJob#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Variables that reference a $(VAR_NAME) are expanded using the previous defined environment variables in the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_custom_job#value VertexAiCustomJob#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

