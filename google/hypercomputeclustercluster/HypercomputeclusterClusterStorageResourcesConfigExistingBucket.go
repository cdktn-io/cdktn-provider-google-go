// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package hypercomputeclustercluster


type HypercomputeclusterClusterStorageResourcesConfigExistingBucket struct {
	// Name of the Cloud Storage bucket to import.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.36.0/docs/resources/hypercomputecluster_cluster#bucket HypercomputeclusterCluster#bucket}
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
}

