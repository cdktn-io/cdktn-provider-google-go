// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package colabnotebookexecution


type ColabNotebookExecutionWorkbenchRuntimeVmImage struct {
	// Use this VM image family to find the image; the newest image in this family will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/colab_notebook_execution#family ColabNotebookExecution#family}
	Family *string `field:"optional" json:"family" yaml:"family"`
	// Use VM image name to find the image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/colab_notebook_execution#name ColabNotebookExecution#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The name of the Google Cloud project that this VM image belongs to. Format: {project_id}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/colab_notebook_execution#project ColabNotebookExecution#project}
	Project *string `field:"optional" json:"project" yaml:"project"`
}

