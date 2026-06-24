// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package containercluster


type ContainerClusterNodePoolNodeConfigKubeletConfigCrashLoopBackOff struct {
	// The maximum duration the backoff delay can accrue to for container restarts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.38.0/docs/resources/container_cluster#max_container_restart_period ContainerCluster#max_container_restart_period}
	MaxContainerRestartPeriod *string `field:"optional" json:"maxContainerRestartPeriod" yaml:"maxContainerRestartPeriod"`
}

