// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agentregistryservice


type AgentRegistryServiceEndpointSpec struct {
	// The type of the Endpoint spec content. Possible values: ["NO_SPEC"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/agent_registry_service#type AgentRegistryService#type}
	Type *string `field:"required" json:"type" yaml:"type"`
}

