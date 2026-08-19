// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cestool


type CesToolAgentTool struct {
	// Required. The name of the agent tool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/ces_tool#name CesTool#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Optional. The resource name of the agent that is the entry point of the tool. Format: projects/{project}/locations/{location}/agents/{agent}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/ces_tool#agent CesTool#agent}
	Agent *string `field:"optional" json:"agent" yaml:"agent"`
	// Optional. Description of the tool's purpose.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/ces_tool#description CesTool#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

