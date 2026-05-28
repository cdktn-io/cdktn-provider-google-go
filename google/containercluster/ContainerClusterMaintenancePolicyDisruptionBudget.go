// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterMaintenancePolicyDisruptionBudget struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.34.0/docs/resources/container_cluster#minor_version_disruption_interval ContainerCluster#minor_version_disruption_interval}.
	MinorVersionDisruptionInterval *string `field:"optional" json:"minorVersionDisruptionInterval" yaml:"minorVersionDisruptionInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.34.0/docs/resources/container_cluster#patch_version_disruption_interval ContainerCluster#patch_version_disruption_interval}.
	PatchVersionDisruptionInterval *string `field:"optional" json:"patchVersionDisruptionInterval" yaml:"patchVersionDisruptionInterval"`
}

