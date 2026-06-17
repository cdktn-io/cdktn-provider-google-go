// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkservicesagentgateway


type NetworkServicesAgentGatewayNetworkConfig struct {
	// egress block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.37.0/docs/resources/network_services_agent_gateway#egress NetworkServicesAgentGateway#egress}
	Egress *NetworkServicesAgentGatewayNetworkConfigEgress `field:"required" json:"egress" yaml:"egress"`
}

