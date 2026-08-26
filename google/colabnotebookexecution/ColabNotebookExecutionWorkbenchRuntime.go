// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package colabnotebookexecution


type ColabNotebookExecutionWorkbenchRuntime struct {
	// vm_image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/colab_notebook_execution#vm_image ColabNotebookExecution#vm_image}
	VmImage *ColabNotebookExecutionWorkbenchRuntimeVmImage `field:"required" json:"vmImage" yaml:"vmImage"`
}

