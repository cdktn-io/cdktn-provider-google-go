// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataplexdataproduct


type DataplexDataProductAccessGroupsDefaultIamRoleConfig struct {
	// Default IAM role associated with the access group, such as 'roles/bigquery.dataViewer'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dataplex_data_product#role DataplexDataProduct#role}
	Role *string `field:"required" json:"role" yaml:"role"`
}

