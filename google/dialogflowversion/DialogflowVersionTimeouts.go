// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dialogflowversion


type DialogflowVersionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/dialogflow_version#create DialogflowVersion#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/dialogflow_version#delete DialogflowVersion#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/dialogflow_version#update DialogflowVersion#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

