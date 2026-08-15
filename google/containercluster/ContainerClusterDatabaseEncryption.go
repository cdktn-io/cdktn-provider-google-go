// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterDatabaseEncryption struct {
	// ENCRYPTED, ALL_OBJECTS_ENCRYPTION_ENABLED or DECRYPTED.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/container_cluster#state ContainerCluster#state}
	State *string `field:"required" json:"state" yaml:"state"`
	// The key to use to encrypt/decrypt secrets.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/container_cluster#key_name ContainerCluster#key_name}
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
}

