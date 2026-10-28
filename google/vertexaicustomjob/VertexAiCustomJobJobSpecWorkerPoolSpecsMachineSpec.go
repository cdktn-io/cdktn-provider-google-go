// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaicustomjob


type VertexAiCustomJobJobSpecWorkerPoolSpecsMachineSpec struct {
	// The number of accelerators to attach to the machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_custom_job#accelerator_count VertexAiCustomJob#accelerator_count}
	AcceleratorCount *float64 `field:"optional" json:"acceleratorCount" yaml:"acceleratorCount"`
	// The type of accelerator(s) that may be attached to the machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_custom_job#accelerator_type VertexAiCustomJob#accelerator_type}
	AcceleratorType *string `field:"optional" json:"acceleratorType" yaml:"acceleratorType"`
	// The type of the machine.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_custom_job#machine_type VertexAiCustomJob#machine_type}
	MachineType *string `field:"optional" json:"machineType" yaml:"machineType"`
	// reservation_affinity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_custom_job#reservation_affinity VertexAiCustomJob#reservation_affinity}
	ReservationAffinity *VertexAiCustomJobJobSpecWorkerPoolSpecsMachineSpecReservationAffinity `field:"optional" json:"reservationAffinity" yaml:"reservationAffinity"`
	// The topology of the TPUs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_custom_job#tpu_topology VertexAiCustomJob#tpu_topology}
	TpuTopology *string `field:"optional" json:"tpuTopology" yaml:"tpuTopology"`
}

