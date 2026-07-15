// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cesexample


type CesExampleMessagesChunksToolCallToolsetTool struct {
	// The resource name of the Toolset from which this tool is derived. Format: 'projects/{project}/locations/{location}/apps/{app}/toolsets/{toolset}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/ces_example#toolset CesExample#toolset}
	Toolset *string `field:"required" json:"toolset" yaml:"toolset"`
	// The tool ID to filter the tools to retrieve the schema for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/ces_example#tool_id CesExample#tool_id}
	ToolId *string `field:"optional" json:"toolId" yaml:"toolId"`
}

