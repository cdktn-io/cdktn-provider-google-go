// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dialogflowtool


type DialogflowToolOpenApiSpecAuthenticationServiceAgentAuthConfig struct {
	// Optional.
	//
	// Indicate the auth token type generated from the Dialogflow service agent.
	// The generated token is sent in the Authorization header. Possible values: ["ID_TOKEN", "ACCESS_TOKEN"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/dialogflow_tool#service_agent_auth DialogflowTool#service_agent_auth}
	ServiceAgentAuth *string `field:"optional" json:"serviceAgentAuth" yaml:"serviceAgentAuth"`
}

