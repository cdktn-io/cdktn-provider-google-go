// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cesexample


type CesExampleMessagesChunksAgentTransfer struct {
	// The agent to which the conversation is being transferred. The agent will handle the conversation from this point forward. Format: 'projects/{project}/locations/{location}/apps/{app}/agents/{agent}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/ces_example#target_agent CesExample#target_agent}
	TargetAgent *string `field:"required" json:"targetAgent" yaml:"targetAgent"`
}

