// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaischedule


type VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDirectNotebookSource struct {
	// The base64-encoded contents of the input notebook file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/vertex_ai_schedule#content VertexAiSchedule#content}
	Content *string `field:"optional" json:"content" yaml:"content"`
}

