// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agentidentityauthprovider


type AgentIdentityAuthProviderAuthProviderTypeParams struct {
	// api_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/agent_identity_auth_provider#api_key AgentIdentityAuthProvider#api_key}
	ApiKey *AgentIdentityAuthProviderAuthProviderTypeParamsApiKey `field:"optional" json:"apiKey" yaml:"apiKey"`
	// three_legged_oauth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/agent_identity_auth_provider#three_legged_oauth AgentIdentityAuthProvider#three_legged_oauth}
	ThreeLeggedOauth *AgentIdentityAuthProviderAuthProviderTypeParamsThreeLeggedOauth `field:"optional" json:"threeLeggedOauth" yaml:"threeLeggedOauth"`
	// two_legged_oauth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/agent_identity_auth_provider#two_legged_oauth AgentIdentityAuthProvider#two_legged_oauth}
	TwoLeggedOauth *AgentIdentityAuthProviderAuthProviderTypeParamsTwoLeggedOauth `field:"optional" json:"twoLeggedOauth" yaml:"twoLeggedOauth"`
}

