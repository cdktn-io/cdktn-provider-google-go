// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaicustomjob


type VertexAiCustomJobJobSpecBaseOutputDirectory struct {
	// Cloud Storage folder URL to store the output.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_custom_job#output_uri_prefix VertexAiCustomJob#output_uri_prefix}
	OutputUriPrefix *string `field:"required" json:"outputUriPrefix" yaml:"outputUriPrefix"`
}

