// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaischedule


type VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJob struct {
	// custom_environment_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/vertex_ai_schedule#custom_environment_spec VertexAiSchedule#custom_environment_spec}
	CustomEnvironmentSpec *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpec `field:"optional" json:"customEnvironmentSpec" yaml:"customEnvironmentSpec"`
	// dataform_repository_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/vertex_ai_schedule#dataform_repository_source VertexAiSchedule#dataform_repository_source}
	DataformRepositorySource *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySource `field:"optional" json:"dataformRepositorySource" yaml:"dataformRepositorySource"`
	// direct_notebook_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/vertex_ai_schedule#direct_notebook_source VertexAiSchedule#direct_notebook_source}
	DirectNotebookSource *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDirectNotebookSource `field:"optional" json:"directNotebookSource" yaml:"directNotebookSource"`
	// The display name of the NotebookExecutionJob.
	//
	// The name can be up to 128 characters long and can consist of any UTF-8 characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/vertex_ai_schedule#display_name VertexAiSchedule#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// encryption_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/vertex_ai_schedule#encryption_spec VertexAiSchedule#encryption_spec}
	EncryptionSpec *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobEncryptionSpec `field:"optional" json:"encryptionSpec" yaml:"encryptionSpec"`
	// Max running time of the execution job in seconds (default 86400s / 24 hrs).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/vertex_ai_schedule#execution_timeout VertexAiSchedule#execution_timeout}
	ExecutionTimeout *string `field:"optional" json:"executionTimeout" yaml:"executionTimeout"`
	// The user email to run the execution as. Only supported by Colab runtimes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/vertex_ai_schedule#execution_user VertexAiSchedule#execution_user}
	ExecutionUser *string `field:"optional" json:"executionUser" yaml:"executionUser"`
	// gcs_notebook_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/vertex_ai_schedule#gcs_notebook_source VertexAiSchedule#gcs_notebook_source}
	GcsNotebookSource *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobGcsNotebookSource `field:"optional" json:"gcsNotebookSource" yaml:"gcsNotebookSource"`
	// The Cloud Storage location to upload the result to. Format: 'gs://bucket-name'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/vertex_ai_schedule#gcs_output_uri VertexAiSchedule#gcs_output_uri}
	GcsOutputUri *string `field:"optional" json:"gcsOutputUri" yaml:"gcsOutputUri"`
	// The name of the kernel to use during notebook execution. If unset, the default kernel is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/vertex_ai_schedule#kernel_name VertexAiSchedule#kernel_name}
	KernelName *string `field:"optional" json:"kernelName" yaml:"kernelName"`
	// The labels with user-defined metadata to organize NotebookExecutionJobs.
	//
	// Label keys and values can be no longer than 64 characters (Unicode codepoints), can only contain lowercase letters, numeric characters, underscores and dashes. International characters are allowed. See https://goo.gl/xmQnxf for more information and examples of labels. System reserved label keys are prefixed with "aiplatform.googleapis.com/" and are immutable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/vertex_ai_schedule#labels VertexAiSchedule#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The NotebookRuntimeTemplate to source compute configuration from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/vertex_ai_schedule#notebook_runtime_template_resource_name VertexAiSchedule#notebook_runtime_template_resource_name}
	NotebookRuntimeTemplateResourceName *string `field:"optional" json:"notebookRuntimeTemplateResourceName" yaml:"notebookRuntimeTemplateResourceName"`
	// The user-defined parameters to use during notebook execution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/vertex_ai_schedule#parameters VertexAiSchedule#parameters}
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// The service account to run the execution as.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/vertex_ai_schedule#service_account VertexAiSchedule#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// workbench_runtime block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/vertex_ai_schedule#workbench_runtime VertexAiSchedule#workbench_runtime}
	WorkbenchRuntime *VertexAiScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobWorkbenchRuntime `field:"optional" json:"workbenchRuntime" yaml:"workbenchRuntime"`
}

