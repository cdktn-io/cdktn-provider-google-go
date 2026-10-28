// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaicustomjob


type VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpec struct {
	// The URI of a container image in Artifact Registry that will run the provided Python package.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_custom_job#executor_image_uri VertexAiCustomJob#executor_image_uri}
	ExecutorImageUri *string `field:"required" json:"executorImageUri" yaml:"executorImageUri"`
	// The Google Cloud Storage location of the Python package files which are the training program and its dependent packages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_custom_job#package_uris VertexAiCustomJob#package_uris}
	PackageUris *[]*string `field:"required" json:"packageUris" yaml:"packageUris"`
	// The Python module name to run after installing the packages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_custom_job#python_module VertexAiCustomJob#python_module}
	PythonModule *string `field:"required" json:"pythonModule" yaml:"pythonModule"`
	// Command line arguments to be passed to the Python task.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_custom_job#args VertexAiCustomJob#args}
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// env block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_custom_job#env VertexAiCustomJob#env}
	Env interface{} `field:"optional" json:"env" yaml:"env"`
}

