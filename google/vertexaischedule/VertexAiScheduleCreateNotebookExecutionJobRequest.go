// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaischedule


type VertexAiScheduleCreateNotebookExecutionJobRequest struct {
	// notebook_execution_job block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/vertex_ai_schedule#notebook_execution_job VertexAiSchedule#notebook_execution_job}
	NotebookExecutionJob *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJob `field:"required" json:"notebookExecutionJob" yaml:"notebookExecutionJob"`
	// The resource name of the Location to create the NotebookExecutionJob. Format: 'projects/{project}/locations/{location}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/vertex_ai_schedule#parent VertexAiSchedule#parent}
	Parent *string `field:"required" json:"parent" yaml:"parent"`
	// User specified ID for the NotebookExecutionJob.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/vertex_ai_schedule#notebook_execution_job_id VertexAiSchedule#notebook_execution_job_id}
	NotebookExecutionJobId *string `field:"optional" json:"notebookExecutionJobId" yaml:"notebookExecutionJobId"`
}

