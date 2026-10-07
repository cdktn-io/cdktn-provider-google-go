// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkservicesagentconnectivitytemplate


type NetworkServicesAgentConnectivityTemplateEgressNetworkConfig struct {
	// dns_peering_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/network_services_agent_connectivity_template#dns_peering_config NetworkServicesAgentConnectivityTemplate#dns_peering_config}
	DnsPeeringConfig *NetworkServicesAgentConnectivityTemplateEgressNetworkConfigDnsPeeringConfig `field:"optional" json:"dnsPeeringConfig" yaml:"dnsPeeringConfig"`
	// The network attachment resource name. Format: projects/{project}/regions/{region}/networkAttachments/{network_attachment_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/network_services_agent_connectivity_template#network_attachment NetworkServicesAgentConnectivityTemplate#network_attachment}
	NetworkAttachment *string `field:"optional" json:"networkAttachment" yaml:"networkAttachment"`
	// tls_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/network_services_agent_connectivity_template#tls_config NetworkServicesAgentConnectivityTemplate#tls_config}
	TlsConfig *NetworkServicesAgentConnectivityTemplateEgressNetworkConfigTlsConfig `field:"optional" json:"tlsConfig" yaml:"tlsConfig"`
	// The VPC egress setting. Possible values: ["ALL_TRAFFIC", "PRIVATE_RANGES_ONLY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/network_services_agent_connectivity_template#vpc_egress NetworkServicesAgentConnectivityTemplate#vpc_egress}
	VpcEgress *string `field:"optional" json:"vpcEgress" yaml:"vpcEgress"`
}

