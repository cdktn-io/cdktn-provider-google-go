// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package resourcemanagerboundary


type ResourceManagerBoundaryResourceFilter struct {
	// tag_filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/resource_manager_boundary#tag_filter ResourceManagerBoundary#tag_filter}
	TagFilter *ResourceManagerBoundaryResourceFilterTagFilter `field:"required" json:"tagFilter" yaml:"tagFilter"`
}

