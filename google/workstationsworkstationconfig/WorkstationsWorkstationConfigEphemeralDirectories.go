// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package workstationsworkstationconfig


type WorkstationsWorkstationConfigEphemeralDirectories struct {
	// gce_pd block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/workstations_workstation_config#gce_pd WorkstationsWorkstationConfigA#gce_pd}
	GcePd *WorkstationsWorkstationConfigEphemeralDirectoriesGcePd `field:"optional" json:"gcePd" yaml:"gcePd"`
	// Location of this directory in the running workstation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/workstations_workstation_config#mount_path WorkstationsWorkstationConfigA#mount_path}
	MountPath *string `field:"optional" json:"mountPath" yaml:"mountPath"`
}

