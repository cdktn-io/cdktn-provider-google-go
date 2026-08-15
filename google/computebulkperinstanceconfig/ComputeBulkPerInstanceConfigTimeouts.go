// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package computebulkperinstanceconfig


type ComputeBulkPerInstanceConfigTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/compute_bulk_per_instance_config#create ComputeBulkPerInstanceConfig#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/compute_bulk_per_instance_config#delete ComputeBulkPerInstanceConfig#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

