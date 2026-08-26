// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package colabschedule


type ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJob struct {
	// Required. The display name of the Notebook Execution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/colab_schedule#display_name ColabSchedule#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The Cloud Storage location to upload the result to. Format:'gs://bucket-name'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/colab_schedule#gcs_output_uri ColabSchedule#gcs_output_uri}
	GcsOutputUri *string `field:"required" json:"gcsOutputUri" yaml:"gcsOutputUri"`
	// custom_environment_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/colab_schedule#custom_environment_spec ColabSchedule#custom_environment_spec}
	CustomEnvironmentSpec *ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobCustomEnvironmentSpec `field:"optional" json:"customEnvironmentSpec" yaml:"customEnvironmentSpec"`
	// dataform_repository_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/colab_schedule#dataform_repository_source ColabSchedule#dataform_repository_source}
	DataformRepositorySource *ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobDataformRepositorySource `field:"optional" json:"dataformRepositorySource" yaml:"dataformRepositorySource"`
	// encryption_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/colab_schedule#encryption_spec ColabSchedule#encryption_spec}
	EncryptionSpec *ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobEncryptionSpec `field:"optional" json:"encryptionSpec" yaml:"encryptionSpec"`
	// Max running time of the execution job in seconds (default 86400s / 24 hrs).
	//
	// A duration in seconds with up to nine fractional digits, ending with "s". Example: "3.5s".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/colab_schedule#execution_timeout ColabSchedule#execution_timeout}
	ExecutionTimeout *string `field:"optional" json:"executionTimeout" yaml:"executionTimeout"`
	// The user email to run the execution as.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/colab_schedule#execution_user ColabSchedule#execution_user}
	ExecutionUser *string `field:"optional" json:"executionUser" yaml:"executionUser"`
	// gcs_notebook_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/colab_schedule#gcs_notebook_source ColabSchedule#gcs_notebook_source}
	GcsNotebookSource *ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobGcsNotebookSource `field:"optional" json:"gcsNotebookSource" yaml:"gcsNotebookSource"`
	// The name of the kernel to use during notebook execution. If unset, the default kernel is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/colab_schedule#kernel_name ColabSchedule#kernel_name}
	KernelName *string `field:"optional" json:"kernelName" yaml:"kernelName"`
	// The labels with user-defined metadata to organize NotebookExecutionJobs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/colab_schedule#labels ColabSchedule#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// The NotebookRuntimeTemplate to source compute configuration from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/colab_schedule#notebook_runtime_template_resource_name ColabSchedule#notebook_runtime_template_resource_name}
	NotebookRuntimeTemplateResourceName *string `field:"optional" json:"notebookRuntimeTemplateResourceName" yaml:"notebookRuntimeTemplateResourceName"`
	// The service account to run the execution as.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/colab_schedule#service_account ColabSchedule#service_account}
	ServiceAccount *string `field:"optional" json:"serviceAccount" yaml:"serviceAccount"`
	// workbench_runtime block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.0/docs/resources/colab_schedule#workbench_runtime ColabSchedule#workbench_runtime}
	WorkbenchRuntime *ColabScheduleCreateNotebookExecutionJobRequestNotebookExecutionJobWorkbenchRuntime `field:"optional" json:"workbenchRuntime" yaml:"workbenchRuntime"`
}

