// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package colabschedule


type ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpec struct {
	// machine_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/colab_schedule#machine_spec ColabSchedule#machine_spec}
	MachineSpec *ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecMachineSpec `field:"optional" json:"machineSpec" yaml:"machineSpec"`
	// network_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/colab_schedule#network_spec ColabSchedule#network_spec}
	NetworkSpec *ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecNetworkSpec `field:"optional" json:"networkSpec" yaml:"networkSpec"`
	// persistent_disk_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/colab_schedule#persistent_disk_spec ColabSchedule#persistent_disk_spec}
	PersistentDiskSpec *ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpecPersistentDiskSpec `field:"optional" json:"persistentDiskSpec" yaml:"persistentDiskSpec"`
}

