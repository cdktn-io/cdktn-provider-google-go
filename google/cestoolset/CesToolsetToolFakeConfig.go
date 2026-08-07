// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cestoolset


type CesToolsetToolFakeConfig struct {
	// code_block block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/ces_toolset#code_block CesToolset#code_block}
	CodeBlock *CesToolsetToolFakeConfigCodeBlock `field:"optional" json:"codeBlock" yaml:"codeBlock"`
	// Whether the tool is using fake mode.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/ces_toolset#enable_fake_mode CesToolset#enable_fake_mode}
	EnableFakeMode interface{} `field:"optional" json:"enableFakeMode" yaml:"enableFakeMode"`
}

