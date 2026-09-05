// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkservicesagentgateway


type NetworkServicesAgentGatewayGoogleManaged struct {
	// Operating Mode of Agent Gateway. Possible values: ["AGENT_TO_ANYWHERE", "CLIENT_TO_AGENT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/network_services_agent_gateway#governed_access_path NetworkServicesAgentGateway#governed_access_path}
	GovernedAccessPath *string `field:"required" json:"governedAccessPath" yaml:"governedAccessPath"`
}

