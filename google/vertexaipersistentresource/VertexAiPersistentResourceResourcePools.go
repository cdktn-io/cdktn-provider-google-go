// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaipersistentresource


type VertexAiPersistentResourceResourcePools struct {
	// machine_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/vertex_ai_persistent_resource#machine_spec VertexAiPersistentResource#machine_spec}
	MachineSpec *VertexAiPersistentResourceResourcePoolsMachineSpec `field:"required" json:"machineSpec" yaml:"machineSpec"`
	// autoscaling_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/vertex_ai_persistent_resource#autoscaling_spec VertexAiPersistentResource#autoscaling_spec}
	AutoscalingSpec *VertexAiPersistentResourceResourcePoolsAutoscalingSpec `field:"optional" json:"autoscalingSpec" yaml:"autoscalingSpec"`
	// disk_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/vertex_ai_persistent_resource#disk_spec VertexAiPersistentResource#disk_spec}
	DiskSpec *VertexAiPersistentResourceResourcePoolsDiskSpec `field:"optional" json:"diskSpec" yaml:"diskSpec"`
	// The unique ID in a PersistentResource for referring to this resource pool.
	//
	// User can specify it if necessary. Otherwise, it's generated
	// automatically.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/vertex_ai_persistent_resource#id VertexAiPersistentResource#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The total number of machines to use for this resource pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.43.0/docs/resources/vertex_ai_persistent_resource#replica_count VertexAiPersistentResource#replica_count}
	ReplicaCount *string `field:"optional" json:"replicaCount" yaml:"replicaCount"`
}

