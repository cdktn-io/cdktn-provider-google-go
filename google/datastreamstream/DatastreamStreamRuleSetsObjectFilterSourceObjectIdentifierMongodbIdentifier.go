// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package datastreamstream


type DatastreamStreamRuleSetsObjectFilterSourceObjectIdentifierMongodbIdentifier struct {
	// The MongoDB collection name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/datastream_stream#collection DatastreamStream#collection}
	Collection *string `field:"required" json:"collection" yaml:"collection"`
	// The MongoDB database name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/datastream_stream#database DatastreamStream#database}
	Database *string `field:"required" json:"database" yaml:"database"`
}

