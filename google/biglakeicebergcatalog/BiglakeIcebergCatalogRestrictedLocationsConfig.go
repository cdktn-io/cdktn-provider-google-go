// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package biglakeicebergcatalog


type BiglakeIcebergCatalogRestrictedLocationsConfig struct {
	// A list of GCS locations (e.g., 'gs://my-other-bucket/...') that are permitted for use by resources within this catalog. Each entry can be either a GCS bucket or a path within it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/biglake_iceberg_catalog#restricted_locations BiglakeIcebergCatalog#restricted_locations}
	RestrictedLocations *[]*string `field:"optional" json:"restrictedLocations" yaml:"restrictedLocations"`
}

