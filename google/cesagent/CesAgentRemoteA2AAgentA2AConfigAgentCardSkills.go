// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cesagent


type CesAgentRemoteA2AAgentA2AConfigAgentCardSkills struct {
	// A detailed description of the skill.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_agent#description CesAgent#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// A unique identifier for the agent's skill.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_agent#id CesAgent#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// A human-readable name for the skill.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_agent#name CesAgent#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A set of keywords describing the skill's capabilities.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_agent#tags CesAgent#tags}
	Tags *[]*string `field:"required" json:"tags" yaml:"tags"`
	// Example prompts or scenarios that this skill can handle.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_agent#examples CesAgent#examples}
	Examples *[]*string `field:"optional" json:"examples" yaml:"examples"`
	// The set of supported input media types for this skill, overriding the agent's defaults.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_agent#input_modes CesAgent#input_modes}
	InputModes *[]*string `field:"optional" json:"inputModes" yaml:"inputModes"`
	// The set of supported output media types for this skill, overriding the agent's defaults.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_agent#output_modes CesAgent#output_modes}
	OutputModes *[]*string `field:"optional" json:"outputModes" yaml:"outputModes"`
}

