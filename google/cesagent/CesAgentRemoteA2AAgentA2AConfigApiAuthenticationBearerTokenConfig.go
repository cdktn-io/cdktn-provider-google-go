// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cesagent


type CesAgentRemoteA2AAgentA2AConfigApiAuthenticationBearerTokenConfig struct {
	// The bearer token. Must be in the format '$context.variables.<name_of_variable>'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_agent#token CesAgent#token}
	Token *string `field:"required" json:"token" yaml:"token"`
}

