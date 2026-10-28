// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaicustomjob


type VertexAiCustomJobJobSpecWorkerPoolSpecsContainerSpec struct {
	// The URI of a container image in the Container Registry that is to be run on each worker replica.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_custom_job#image_uri VertexAiCustomJob#image_uri}
	ImageUri *string `field:"required" json:"imageUri" yaml:"imageUri"`
	// The arguments to be passed when starting the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_custom_job#args VertexAiCustomJob#args}
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// The command to be invoked when the container is started. It overrides the entrypoint instruction in Dockerfile when provided.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_custom_job#command VertexAiCustomJob#command}
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// env block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_custom_job#env VertexAiCustomJob#env}
	Env interface{} `field:"optional" json:"env" yaml:"env"`
}

