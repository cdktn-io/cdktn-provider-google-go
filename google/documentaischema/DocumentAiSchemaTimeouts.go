// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package documentaischema


type DocumentAiSchemaTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/document_ai_schema#create DocumentAiSchema#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/document_ai_schema#delete DocumentAiSchema#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/document_ai_schema#update DocumentAiSchema#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

