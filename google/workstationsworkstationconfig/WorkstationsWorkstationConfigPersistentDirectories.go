// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workstationsworkstationconfig


type WorkstationsWorkstationConfigPersistentDirectories struct {
	// gce_hd block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.36.0/docs/resources/workstations_workstation_config#gce_hd WorkstationsWorkstationConfigA#gce_hd}
	GceHd *WorkstationsWorkstationConfigPersistentDirectoriesGceHd `field:"optional" json:"gceHd" yaml:"gceHd"`
	// gce_pd block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.36.0/docs/resources/workstations_workstation_config#gce_pd WorkstationsWorkstationConfigA#gce_pd}
	GcePd *WorkstationsWorkstationConfigPersistentDirectoriesGcePd `field:"optional" json:"gcePd" yaml:"gcePd"`
	// Location of this directory in the running workstation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.36.0/docs/resources/workstations_workstation_config#mount_path WorkstationsWorkstationConfigA#mount_path}
	MountPath *string `field:"optional" json:"mountPath" yaml:"mountPath"`
}

