// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaitrainingpipeline


type VertexAiTrainingPipelineModelToUpload struct {
	// The user-defined name of the Model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#display_name VertexAiTrainingPipeline#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The path to the directory containing the Model artifact and any of its supporting files.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#artifact_uri VertexAiTrainingPipeline#artifact_uri}
	ArtifactUri *string `field:"optional" json:"artifactUri" yaml:"artifactUri"`
	// container_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#container_spec VertexAiTrainingPipeline#container_spec}
	ContainerSpec *VertexAiTrainingPipelineModelToUploadContainerSpec `field:"optional" json:"containerSpec" yaml:"containerSpec"`
	// The description of the Model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#description VertexAiTrainingPipeline#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// encryption_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#encryption_spec VertexAiTrainingPipeline#encryption_spec}
	EncryptionSpec *VertexAiTrainingPipelineModelToUploadEncryptionSpec `field:"optional" json:"encryptionSpec" yaml:"encryptionSpec"`
	// explanation_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#explanation_spec VertexAiTrainingPipeline#explanation_spec}
	ExplanationSpec *VertexAiTrainingPipelineModelToUploadExplanationSpec `field:"optional" json:"explanationSpec" yaml:"explanationSpec"`
	// The labels with user-defined metadata to organize your Models.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#labels VertexAiTrainingPipeline#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// User provided version aliases so that a model version can be referenced using alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#version_aliases VertexAiTrainingPipeline#version_aliases}
	VersionAliases *[]*string `field:"optional" json:"versionAliases" yaml:"versionAliases"`
	// The description of this version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#version_description VertexAiTrainingPipeline#version_description}
	VersionDescription *string `field:"optional" json:"versionDescription" yaml:"versionDescription"`
}

