// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cesagent


type CesAgentToolsets struct {
	// The resource name of the toolset. Format: 'projects/{project}/locations/{location}/apps/{app}/toolsets/{toolset}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/ces_agent#toolset CesAgent#toolset}
	Toolset *string `field:"required" json:"toolset" yaml:"toolset"`
	// The tools IDs to filter the toolset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/ces_agent#tool_ids CesAgent#tool_ids}
	ToolIds *[]*string `field:"optional" json:"toolIds" yaml:"toolIds"`
}

