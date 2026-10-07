// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cestoolset


type CesToolsetMcpToolsetToolOverrides struct {
	// The name of the tool to be overridden.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_toolset#tool CesToolset#tool}
	Tool *string `field:"required" json:"tool" yaml:"tool"`
	// The description override for the tool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_toolset#description_override CesToolset#description_override}
	DescriptionOverride *string `field:"optional" json:"descriptionOverride" yaml:"descriptionOverride"`
	// The name override for the tool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_toolset#name_override CesToolset#name_override}
	NameOverride *string `field:"optional" json:"nameOverride" yaml:"nameOverride"`
}

