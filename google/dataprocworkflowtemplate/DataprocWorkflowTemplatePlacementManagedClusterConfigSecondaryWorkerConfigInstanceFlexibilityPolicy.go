// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package dataprocworkflowtemplate


type DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigInstanceFlexibilityPolicy struct {
	// instance_selection_list block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/dataproc_workflow_template#instance_selection_list DataprocWorkflowTemplate#instance_selection_list}
	InstanceSelectionList interface{} `field:"optional" json:"instanceSelectionList" yaml:"instanceSelectionList"`
	// provisioning_model_mix block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/dataproc_workflow_template#provisioning_model_mix DataprocWorkflowTemplate#provisioning_model_mix}
	ProvisioningModelMix *DataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigInstanceFlexibilityPolicyProvisioningModelMix `field:"optional" json:"provisioningModelMix" yaml:"provisioningModelMix"`
}

