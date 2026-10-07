// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataprocworkflowtemplate


type DataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfigAttachedDiskConfig struct {
	// Optional. Size of the attached disk, specified in GB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/dataproc_workflow_template#disk_size_gb DataprocWorkflowTemplate#disk_size_gb}
	DiskSizeGb *float64 `field:"optional" json:"diskSizeGb" yaml:"diskSizeGb"`
	// Optional. The disk type of the attached disk. Currently only supports Hyperdisks: `hyperdisk-balanced`, `hyperdisk-extreme`, `hyperdisk-ml`, `hyperdisk-throughput`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/dataproc_workflow_template#disk_type DataprocWorkflowTemplate#disk_type}
	DiskType *string `field:"optional" json:"diskType" yaml:"diskType"`
	// Optional.
	//
	// Indicates how many IOPS to provision for the disk. This sets the number of I/O operations per second that the disk can handle.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/dataproc_workflow_template#provisioned_iops DataprocWorkflowTemplate#provisioned_iops}
	ProvisionedIops *float64 `field:"optional" json:"provisionedIops" yaml:"provisionedIops"`
	// Optional.
	//
	// Indicates how much throughput to provision for the disk. This sets the number of throughput mb per second that the disk can handle.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/dataproc_workflow_template#provisioned_throughput DataprocWorkflowTemplate#provisioned_throughput}
	ProvisionedThroughput *float64 `field:"optional" json:"provisionedThroughput" yaml:"provisionedThroughput"`
}

