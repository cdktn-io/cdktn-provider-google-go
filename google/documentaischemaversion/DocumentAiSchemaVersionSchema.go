// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package documentaischemaversion


type DocumentAiSchemaVersionSchema struct {
	// Description of the schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/document_ai_schema_version#description DocumentAiSchemaVersion#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Display name to show users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/document_ai_schema_version#display_name DocumentAiSchemaVersion#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Document-level prompt provided by the user.
	//
	// This custom text is injected into the AI model's prompt to provide extra, document-wide guidance for processing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/document_ai_schema_version#document_prompt DocumentAiSchemaVersion#document_prompt}
	DocumentPrompt *string `field:"optional" json:"documentPrompt" yaml:"documentPrompt"`
	// entity_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/document_ai_schema_version#entity_types DocumentAiSchemaVersion#entity_types}
	EntityTypes interface{} `field:"optional" json:"entityTypes" yaml:"entityTypes"`
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/document_ai_schema_version#metadata DocumentAiSchemaVersion#metadata}
	Metadata *DocumentAiSchemaVersionSchemaMetadata `field:"optional" json:"metadata" yaml:"metadata"`
}

