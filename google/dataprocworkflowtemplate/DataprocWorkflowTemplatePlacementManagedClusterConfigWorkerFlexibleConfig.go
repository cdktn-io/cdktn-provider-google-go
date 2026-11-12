// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataprocworkflowtemplate


type DataprocWorkflowTemplatePlacementManagedClusterConfigWorkerFlexibleConfig struct {
	// Optional. Size of the boot disk in GB for each instance in the instance group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dataproc_workflow_template#boot_disk_size_gb DataprocWorkflowTemplate#boot_disk_size_gb}
	BootDiskSizeGb *float64 `field:"optional" json:"bootDiskSizeGb" yaml:"bootDiskSizeGb"`
	// Optional. The core count for each instance in the instance group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dataproc_workflow_template#core_count DataprocWorkflowTemplate#core_count}
	CoreCount *float64 `field:"optional" json:"coreCount" yaml:"coreCount"`
	// Optional. The hardware class of instances in the instance group. Possible values: HARDWARE_CLASS_UNSPECIFIED, STANDARD, PERFORMANCE.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dataproc_workflow_template#hardware_class DataprocWorkflowTemplate#hardware_class}
	HardwareClass *string `field:"optional" json:"hardwareClass" yaml:"hardwareClass"`
	// Optional. Number of instances in the instance group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dataproc_workflow_template#num_instances DataprocWorkflowTemplate#num_instances}
	NumInstances *float64 `field:"optional" json:"numInstances" yaml:"numInstances"`
}

