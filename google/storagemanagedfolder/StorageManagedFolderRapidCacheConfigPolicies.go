// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package storagemanagedfolder


type StorageManagedFolderRapidCacheConfigPolicies struct {
	// The unique identifier of the rapid cache.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/storage_managed_folder#rapid_cache_id StorageManagedFolder#rapid_cache_id}
	RapidCacheId *string `field:"required" json:"rapidCacheId" yaml:"rapidCacheId"`
	// The ingest-on-write policy for objects in the managed folder.
	//
	// When
	// set to 'enabled', objects are automatically ingested into the cache
	// when they are written to the managed folder.
	// Possible values: ["enabled", "unspecified"]
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/storage_managed_folder#ingest_on_write StorageManagedFolder#ingest_on_write}
	IngestOnWrite *string `field:"optional" json:"ingestOnWrite" yaml:"ingestOnWrite"`
}

