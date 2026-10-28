// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaicustomjob


type VertexAiCustomJobJobSpecWorkerPoolSpecs struct {
	// container_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_custom_job#container_spec VertexAiCustomJob#container_spec}
	ContainerSpec *VertexAiCustomJobJobSpecWorkerPoolSpecsContainerSpec `field:"optional" json:"containerSpec" yaml:"containerSpec"`
	// disk_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_custom_job#disk_spec VertexAiCustomJob#disk_spec}
	DiskSpec *VertexAiCustomJobJobSpecWorkerPoolSpecsDiskSpec `field:"optional" json:"diskSpec" yaml:"diskSpec"`
	// lustre_mounts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_custom_job#lustre_mounts VertexAiCustomJob#lustre_mounts}
	LustreMounts interface{} `field:"optional" json:"lustreMounts" yaml:"lustreMounts"`
	// machine_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_custom_job#machine_spec VertexAiCustomJob#machine_spec}
	MachineSpec *VertexAiCustomJobJobSpecWorkerPoolSpecsMachineSpec `field:"optional" json:"machineSpec" yaml:"machineSpec"`
	// nfs_mounts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_custom_job#nfs_mounts VertexAiCustomJob#nfs_mounts}
	NfsMounts interface{} `field:"optional" json:"nfsMounts" yaml:"nfsMounts"`
	// python_package_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_custom_job#python_package_spec VertexAiCustomJob#python_package_spec}
	PythonPackageSpec *VertexAiCustomJobJobSpecWorkerPoolSpecsPythonPackageSpec `field:"optional" json:"pythonPackageSpec" yaml:"pythonPackageSpec"`
	// The number of worker replicas to use for this worker pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_custom_job#replica_count VertexAiCustomJob#replica_count}
	ReplicaCount *float64 `field:"optional" json:"replicaCount" yaml:"replicaCount"`
}

