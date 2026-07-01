// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkservicesagentgateway


type NetworkServicesAgentGatewaySelfManaged struct {
	// A supported Google Cloud networking proxy in the Project and Location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/network_services_agent_gateway#resource_uri NetworkServicesAgentGateway#resource_uri}
	ResourceUri *string `field:"required" json:"resourceUri" yaml:"resourceUri"`
}

