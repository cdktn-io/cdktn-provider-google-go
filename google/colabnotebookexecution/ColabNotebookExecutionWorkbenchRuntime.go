// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package colabnotebookexecution


type ColabNotebookExecutionWorkbenchRuntime struct {
	// custom_container_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/colab_notebook_execution#custom_container_image ColabNotebookExecution#custom_container_image}
	CustomContainerImage *ColabNotebookExecutionWorkbenchRuntimeCustomContainerImage `field:"optional" json:"customContainerImage" yaml:"customContainerImage"`
	// vm_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/colab_notebook_execution#vm_image ColabNotebookExecution#vm_image}
	VmImage *ColabNotebookExecutionWorkbenchRuntimeVmImage `field:"optional" json:"vmImage" yaml:"vmImage"`
}

