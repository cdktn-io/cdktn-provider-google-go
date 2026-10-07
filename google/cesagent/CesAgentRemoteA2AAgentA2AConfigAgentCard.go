// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cesagent


type CesAgentRemoteA2AAgentA2AConfigAgentCard struct {
	// A description of the agent's domain of action/solution space.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_agent#description CesAgent#description}
	Description *string `field:"required" json:"description" yaml:"description"`
	// A human-readable name for the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_agent#name CesAgent#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// skills block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_agent#skills CesAgent#skills}
	Skills interface{} `field:"required" json:"skills" yaml:"skills"`
	// supported_interfaces block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_agent#supported_interfaces CesAgent#supported_interfaces}
	SupportedInterfaces interface{} `field:"required" json:"supportedInterfaces" yaml:"supportedInterfaces"`
	// The version of the agent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/ces_agent#version CesAgent#version}
	Version *string `field:"required" json:"version" yaml:"version"`
}

