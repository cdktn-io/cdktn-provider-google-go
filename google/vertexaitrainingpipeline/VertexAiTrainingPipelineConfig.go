// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaitrainingpipeline

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VertexAiTrainingPipelineConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The user-defined name of this TrainingPipeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#display_name VertexAiTrainingPipeline#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The ID to use for the uploaded Model, which will become the final component of the model resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#model_id VertexAiTrainingPipeline#model_id}
	ModelId *string `field:"required" json:"modelId" yaml:"modelId"`
	// A Google Cloud Storage path to the YAML file that defines the training task which is responsible for producing the model artifact, and may also include additional auxiliary work.
	//
	// The definition files that can be used here are found in gs://google-cloud-aiplatform/schema/trainingjob/definition/.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#training_task_definition VertexAiTrainingPipeline#training_task_definition}
	TrainingTaskDefinition *string `field:"required" json:"trainingTaskDefinition" yaml:"trainingTaskDefinition"`
	// The training task's parameter(s), as specified in the training_task_definition's inputs. This field uses JSON data as a string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#training_task_inputs VertexAiTrainingPipeline#training_task_inputs}
	TrainingTaskInputs *string `field:"required" json:"trainingTaskInputs" yaml:"trainingTaskInputs"`
	// Whether Terraform will be prevented from destroying the instance.
	//
	// Defaults to "DELETE".
	// When a 'terraform destroy' or 'terraform apply' would delete the instance,
	// the command will fail if this field is set to "PREVENT" in Terraform state.
	// When set to "ABANDON", the command will remove the resource from Terraform
	// management without updating or deleting the resource in the API.
	// When set to "DELETE", deleting the resource is allowed.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#deletion_policy VertexAiTrainingPipeline#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// encryption_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#encryption_spec VertexAiTrainingPipeline#encryption_spec}
	EncryptionSpec *VertexAiTrainingPipelineEncryptionSpec `field:"optional" json:"encryptionSpec" yaml:"encryptionSpec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#id VertexAiTrainingPipeline#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// input_data_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#input_data_config VertexAiTrainingPipeline#input_data_config}
	InputDataConfig *VertexAiTrainingPipelineInputDataConfig `field:"optional" json:"inputDataConfig" yaml:"inputDataConfig"`
	// The labels with user-defined metadata to organize TrainingPipelines.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#labels VertexAiTrainingPipeline#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// model_to_upload block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#model_to_upload VertexAiTrainingPipeline#model_to_upload}
	ModelToUpload *VertexAiTrainingPipelineModelToUpload `field:"optional" json:"modelToUpload" yaml:"modelToUpload"`
	// When specified, the model_to_upload will become a new version of this parent_model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#parent_model VertexAiTrainingPipeline#parent_model}
	ParentModel *string `field:"optional" json:"parentModel" yaml:"parentModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#project VertexAiTrainingPipeline#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The region of the training pipeline. eg us-central1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#region VertexAiTrainingPipeline#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#timeouts VertexAiTrainingPipeline#timeouts}
	Timeouts *VertexAiTrainingPipelineTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

