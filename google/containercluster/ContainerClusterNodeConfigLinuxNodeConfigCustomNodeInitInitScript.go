// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterNodeConfigLinuxNodeConfigCustomNodeInitInitScript struct {
	// The Secret Manager secret URI of the init script.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/container_cluster#gcp_secret_manager_secret_uri ContainerCluster#gcp_secret_manager_secret_uri}
	GcpSecretManagerSecretUri *string `field:"optional" json:"gcpSecretManagerSecretUri" yaml:"gcpSecretManagerSecretUri"`
	// The GCS generation of the init script.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/container_cluster#gcs_generation ContainerCluster#gcs_generation}
	GcsGeneration *float64 `field:"optional" json:"gcsGeneration" yaml:"gcsGeneration"`
	// The GCS URI of the init script.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/container_cluster#gcs_uri ContainerCluster#gcs_uri}
	GcsUri *string `field:"optional" json:"gcsUri" yaml:"gcsUri"`
}

