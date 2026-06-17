// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package pubsubtopic


type PubsubTopicSchemaSettings struct {
	// The name of the schema that messages published should be validated against.
	//
	// Format is projects/{project}/schemas/{schema}.
	// The value of this field will be _deleted-schema_
	// if the schema has been deleted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.37.0/docs/resources/pubsub_topic#schema PubsubTopic#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// The encoding of messages validated against schema. Default value: "ENCODING_UNSPECIFIED" Possible values: ["ENCODING_UNSPECIFIED", "JSON", "BINARY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.37.0/docs/resources/pubsub_topic#encoding PubsubTopic#encoding}
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// The minimum (inclusive) revision allowed for validating messages.
	//
	// If empty or not present, allow any revision to be validated against last_revision or any revision created before.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.37.0/docs/resources/pubsub_topic#first_revision_id PubsubTopic#first_revision_id}
	FirstRevisionId *string `field:"optional" json:"firstRevisionId" yaml:"firstRevisionId"`
	// The maximum (inclusive) revision allowed for validating messages.
	//
	// If empty or not present, allow any revision to be validated against first_revision or any revision created after.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.37.0/docs/resources/pubsub_topic#last_revision_id PubsubTopic#last_revision_id}
	LastRevisionId *string `field:"optional" json:"lastRevisionId" yaml:"lastRevisionId"`
}

