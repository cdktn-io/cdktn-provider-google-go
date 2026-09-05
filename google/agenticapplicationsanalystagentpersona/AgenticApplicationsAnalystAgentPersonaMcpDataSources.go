// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package agenticapplicationsanalystagentpersona


type AgenticApplicationsAnalystAgentPersonaMcpDataSources struct {
	// The description of the MCP agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/agentic_applications_analyst_agent_persona#description AgenticApplicationsAnalystAgentPersona#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// The display name of the MCP server.
	//
	// Must be no longer than 63 characters
	// and can only contain letters, numbers, spaces, underscores, and hyphens.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/agentic_applications_analyst_agent_persona#display_name AgenticApplicationsAnalystAgentPersona#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Whether this external data source is enabled for the current analysis.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/agentic_applications_analyst_agent_persona#enabled AgenticApplicationsAnalystAgentPersona#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The URL of the MCP server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/agentic_applications_analyst_agent_persona#server_url AgenticApplicationsAnalystAgentPersona#server_url}
	ServerUrl *string `field:"required" json:"serverUrl" yaml:"serverUrl"`
	// Input only. The API key of the MCP server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/agentic_applications_analyst_agent_persona#api_key AgenticApplicationsAnalystAgentPersona#api_key}
	ApiKey *string `field:"optional" json:"apiKey" yaml:"apiKey"`
	// The API key parameter name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/agentic_applications_analyst_agent_persona#api_key_name AgenticApplicationsAnalystAgentPersona#api_key_name}
	ApiKeyName *string `field:"optional" json:"apiKeyName" yaml:"apiKeyName"`
	// The client ID for authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/agentic_applications_analyst_agent_persona#client_id AgenticApplicationsAnalystAgentPersona#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Input only. The client secret for authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/agentic_applications_analyst_agent_persona#client_secret AgenticApplicationsAnalystAgentPersona#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// The URL to use for retrieving the OAuth token.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/agentic_applications_analyst_agent_persona#oauth_token_url AgenticApplicationsAnalystAgentPersona#oauth_token_url}
	OauthTokenUrl *string `field:"optional" json:"oauthTokenUrl" yaml:"oauthTokenUrl"`
	// The custom prompt for the MCP agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/agentic_applications_analyst_agent_persona#prompt AgenticApplicationsAnalystAgentPersona#prompt}
	Prompt *string `field:"optional" json:"prompt" yaml:"prompt"`
}

