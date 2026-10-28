// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package bigtablememorylayer


type BigtableMemoryLayerTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/bigtable_memory_layer#create BigtableMemoryLayer#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/bigtable_memory_layer#delete BigtableMemoryLayer#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

