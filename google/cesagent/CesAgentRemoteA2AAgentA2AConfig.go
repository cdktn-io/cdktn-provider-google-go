// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cesagent


type CesAgentRemoteA2AAgentA2AConfig struct {
	// agent_card block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_agent#agent_card CesAgent#agent_card}
	AgentCard *CesAgentRemoteA2AAgentA2AConfigAgentCard `field:"optional" json:"agentCard" yaml:"agentCard"`
	// Reference to the agent in the Agent Registry. Format: 'projects/{project}/locations/{location}/agents/{agent}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_agent#agent_registry CesAgent#agent_registry}
	AgentRegistry *string `field:"optional" json:"agentRegistry" yaml:"agentRegistry"`
	// api_authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_agent#api_authentication CesAgent#api_authentication}
	ApiAuthentication *CesAgentRemoteA2AAgentA2AConfigApiAuthentication `field:"optional" json:"apiAuthentication" yaml:"apiAuthentication"`
	// If not empty, interactions with the remote A2A agent will use this context ID.
	//
	// This context_id field can refer to a session variable like
	// '$context.variables.order_agent_session_id'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_agent#context_id CesAgent#context_id}
	ContextId *string `field:"optional" json:"contextId" yaml:"contextId"`
	// Mapping of input variable names of remote agent to GECX variable names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_agent#input_variable_mapping CesAgent#input_variable_mapping}
	InputVariableMapping *map[string]*string `field:"optional" json:"inputVariableMapping" yaml:"inputVariableMapping"`
	// Mapping of output variable names of remote agent to GECX variable names.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_agent#output_variable_mapping CesAgent#output_variable_mapping}
	OutputVariableMapping *map[string]*string `field:"optional" json:"outputVariableMapping" yaml:"outputVariableMapping"`
	// Whether streaming is enabled for the remote agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_agent#streaming_enabled CesAgent#streaming_enabled}
	StreamingEnabled interface{} `field:"optional" json:"streamingEnabled" yaml:"streamingEnabled"`
}

