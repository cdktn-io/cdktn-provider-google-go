// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkservicesagentgateway


type NetworkServicesAgentGatewayNetworkConfigEgress struct {
	// The URI of the Network Attachment resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/network_services_agent_gateway#network_attachment NetworkServicesAgentGateway#network_attachment}
	NetworkAttachment *string `field:"required" json:"networkAttachment" yaml:"networkAttachment"`
}

