// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package storagemanagedfolder


type StorageManagedFolderRapidCacheConfig struct {
	// policies block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/storage_managed_folder#policies StorageManagedFolder#policies}
	Policies interface{} `field:"required" json:"policies" yaml:"policies"`
}

