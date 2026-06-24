// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dialogflowcxtoolversion


type DialogflowCxToolVersionToolDataStoreSpec struct {
	// data_store_connections block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/dialogflow_cx_tool_version#data_store_connections DialogflowCxToolVersion#data_store_connections}
	DataStoreConnections interface{} `field:"required" json:"dataStoreConnections" yaml:"dataStoreConnections"`
	// fallback_prompt block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/dialogflow_cx_tool_version#fallback_prompt DialogflowCxToolVersion#fallback_prompt}
	FallbackPrompt *DialogflowCxToolVersionToolDataStoreSpecFallbackPrompt `field:"required" json:"fallbackPrompt" yaml:"fallbackPrompt"`
}

