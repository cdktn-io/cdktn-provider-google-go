// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package colabschedule


type ColabScheduleCreatePipelineJobRequest struct {
	// pipeline_job block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/colab_schedule#pipeline_job ColabSchedule#pipeline_job}
	PipelineJob *ColabScheduleCreatePipelineJobRequestPipelineJob `field:"required" json:"pipelineJob" yaml:"pipelineJob"`
	// The resource name of the Location to create the PipelineJob in. Format: 'projects/{project}/locations/{location}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/colab_schedule#parent ColabSchedule#parent}
	Parent *string `field:"optional" json:"parent" yaml:"parent"`
}

