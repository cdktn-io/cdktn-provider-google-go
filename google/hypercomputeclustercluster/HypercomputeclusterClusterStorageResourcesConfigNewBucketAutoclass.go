// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package hypercomputeclustercluster


type HypercomputeclusterClusterStorageResourcesConfigNewBucketAutoclass struct {
	// Enables Auto-class feature.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/hypercomputecluster_cluster#enabled HypercomputeclusterCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Terminal storage class of the autoclass bucket Possible values: NEARLINE ARCHIVE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.35.0/docs/resources/hypercomputecluster_cluster#terminal_storage_class HypercomputeclusterCluster#terminal_storage_class}
	TerminalStorageClass *string `field:"optional" json:"terminalStorageClass" yaml:"terminalStorageClass"`
}

