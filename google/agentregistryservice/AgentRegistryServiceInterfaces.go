// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agentregistryservice


type AgentRegistryServiceInterfaces struct {
	// The protocol binding of the interface. Possible values: ["JSONRPC", "GRPC", "HTTP_JSON"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/agent_registry_service#protocol_binding AgentRegistryService#protocol_binding}
	ProtocolBinding *string `field:"required" json:"protocolBinding" yaml:"protocolBinding"`
	// The destination URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/agent_registry_service#url AgentRegistryService#url}
	Url *string `field:"required" json:"url" yaml:"url"`
}

