// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaiendpoint


type VertexAiEndpointPrivateServiceConnectConfigPscAutomationConfigs struct {
	// The full name of the Google Compute Engine [network](https://cloud.google.com/compute/docs/networks-and-firewalls#networks). [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/get): projects/{project}/global/networks/{network}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/vertex_ai_endpoint#network VertexAiEndpoint#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// Project id used to create forwarding rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/vertex_ai_endpoint#project_id VertexAiEndpoint#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
}

