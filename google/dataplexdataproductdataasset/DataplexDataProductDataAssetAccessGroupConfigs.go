// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataplexdataproductdataasset


type DataplexDataProductDataAssetAccessGroupConfigs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/dataplex_data_product_data_asset#access_group DataplexDataProductDataAsset#access_group}.
	AccessGroup *string `field:"required" json:"accessGroup" yaml:"accessGroup"`
	// IAM roles granted on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/dataplex_data_product_data_asset#iam_roles DataplexDataProductDataAsset#iam_roles}
	IamRoles *[]*string `field:"optional" json:"iamRoles" yaml:"iamRoles"`
}

