// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datastreamstream


type DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierPostgresqlIdentifier struct {
	// The schema name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/datastream_stream#schema DatastreamStream#schema}
	Schema *string `field:"required" json:"schema" yaml:"schema"`
	// The table name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/datastream_stream#table DatastreamStream#table}
	Table *string `field:"required" json:"table" yaml:"table"`
}

