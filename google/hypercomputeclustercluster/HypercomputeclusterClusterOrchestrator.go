// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package hypercomputeclustercluster


type HypercomputeclusterClusterOrchestrator struct {
	// slurm block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.39.0/docs/resources/hypercomputecluster_cluster#slurm HypercomputeclusterCluster#slurm}
	Slurm *HypercomputeclusterClusterOrchestratorSlurm `field:"optional" json:"slurm" yaml:"slurm"`
}

