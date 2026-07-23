// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterSecretSyncConfigRotationConfig struct {
	// Enable the Secret sync auto rotation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/container_cluster#enabled ContainerCluster#enabled}
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// The interval between two consecutive rotations. Default rotation interval is 2 minutes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/container_cluster#rotation_interval ContainerCluster#rotation_interval}
	RotationInterval *string `field:"optional" json:"rotationInterval" yaml:"rotationInterval"`
}

