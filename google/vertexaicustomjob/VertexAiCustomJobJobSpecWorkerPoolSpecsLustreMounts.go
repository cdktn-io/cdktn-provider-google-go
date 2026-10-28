// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaicustomjob


type VertexAiCustomJobJobSpecWorkerPoolSpecsLustreMounts struct {
	// The name of the Lustre filesystem.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_custom_job#filesystem VertexAiCustomJob#filesystem}
	Filesystem *string `field:"required" json:"filesystem" yaml:"filesystem"`
	// IP address of the Lustre instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_custom_job#instance_ip VertexAiCustomJob#instance_ip}
	InstanceIp *string `field:"required" json:"instanceIp" yaml:"instanceIp"`
	// Destination mount path. The Lustre file system will be mounted for the user under /mnt/lustre/<mount_point>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_custom_job#mount_point VertexAiCustomJob#mount_point}
	MountPoint *string `field:"required" json:"mountPoint" yaml:"mountPoint"`
	// The unique identifier of the Lustre volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_custom_job#volume_handle VertexAiCustomJob#volume_handle}
	VolumeHandle *string `field:"required" json:"volumeHandle" yaml:"volumeHandle"`
}

