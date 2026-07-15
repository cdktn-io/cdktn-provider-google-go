// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package biglakeicebergcatalog


type BiglakeIcebergCatalogFederatedCatalogOptionsRefreshOptionsRefreshScope struct {
	// A list of namespace filters to limit which namespaces are synchronized from the remote catalog.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.40.0/docs/resources/biglake_iceberg_catalog#namespace_filters BiglakeIcebergCatalog#namespace_filters}
	NamespaceFilters *[]*string `field:"optional" json:"namespaceFilters" yaml:"namespaceFilters"`
}

