// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dialogflowcompanionagent


type DialogflowCompanionAgentCesToolSpecsParameterOverrides struct {
	// The key of the context reference in ingested context references.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dialogflow_companion_agent#context_reference_key DialogflowCompanionAgent#context_reference_key}
	ContextReferenceKey *string `field:"required" json:"contextReferenceKey" yaml:"contextReferenceKey"`
	// The name of the tool input parameter to override, as defined in the tool's input schema. For example, 'user_email'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dialogflow_companion_agent#parameter_name DialogflowCompanionAgent#parameter_name}
	ParameterName *string `field:"required" json:"parameterName" yaml:"parameterName"`
	// The path into the JSON object of the latest content ingested under 'context_reference_key' (e.g. 'email' or 'user.email'). If omitted, the entire content under 'context_reference_key' is used. If the path cannot be resolved when the tool is called, the value predicted by the model is used (or omitted if the model did not predict it).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dialogflow_companion_agent#field_path DialogflowCompanionAgent#field_path}
	FieldPath *string `field:"optional" json:"fieldPath" yaml:"fieldPath"`
}

