// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkservicesagentgateway


type NetworkServicesAgentGatewayNetworkConfig struct {
	// egress block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/network_services_agent_gateway#egress NetworkServicesAgentGateway#egress}
	Egress *NetworkServicesAgentGatewayNetworkConfigEgress `field:"required" json:"egress" yaml:"egress"`
	// dns_peering_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/network_services_agent_gateway#dns_peering_config NetworkServicesAgentGateway#dns_peering_config}
	DnsPeeringConfig *NetworkServicesAgentGatewayNetworkConfigDnsPeeringConfig `field:"optional" json:"dnsPeeringConfig" yaml:"dnsPeeringConfig"`
}

