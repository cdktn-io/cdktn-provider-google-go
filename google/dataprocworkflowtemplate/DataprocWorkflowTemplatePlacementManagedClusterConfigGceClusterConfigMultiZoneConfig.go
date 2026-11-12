// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataprocworkflowtemplate


type DataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigMultiZoneConfig struct {
	// Optional. Specifies the target shape of the cluster. Possible values: TARGET_SHAPE_UNSPECIFIED, ANY.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/dataproc_workflow_template#target_shape DataprocWorkflowTemplate#target_shape}
	TargetShape *string `field:"optional" json:"targetShape" yaml:"targetShape"`
}

