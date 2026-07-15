// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package biglakeicebergcatalog


type BiglakeIcebergCatalogFederatedCatalogOptionsRefreshOptions struct {
	// refresh_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/biglake_iceberg_catalog#refresh_schedule BiglakeIcebergCatalog#refresh_schedule}
	RefreshSchedule *BiglakeIcebergCatalogFederatedCatalogOptionsRefreshOptionsRefreshSchedule `field:"optional" json:"refreshSchedule" yaml:"refreshSchedule"`
	// refresh_scope block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/biglake_iceberg_catalog#refresh_scope BiglakeIcebergCatalog#refresh_scope}
	RefreshScope *BiglakeIcebergCatalogFederatedCatalogOptionsRefreshOptionsRefreshScope `field:"optional" json:"refreshScope" yaml:"refreshScope"`
}

