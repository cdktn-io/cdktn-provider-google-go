// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataplexdataproduct


type DataplexDataProductAccessGroupsPrincipal struct {
	// Email of the Google Group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/dataplex_data_product#google_group DataplexDataProduct#google_group}
	GoogleGroup *string `field:"optional" json:"googleGroup" yaml:"googleGroup"`
	// Specifies the email of the producer service account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/dataplex_data_product#service_account DataplexDataProduct#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
}

