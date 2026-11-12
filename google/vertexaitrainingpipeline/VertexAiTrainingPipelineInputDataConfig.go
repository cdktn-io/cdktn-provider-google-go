// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaitrainingpipeline


type VertexAiTrainingPipelineInputDataConfig struct {
	// The ID of the Dataset in the same Project and Location which data will be used to train the Model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#dataset_id VertexAiTrainingPipeline#dataset_id}
	DatasetId *string `field:"required" json:"datasetId" yaml:"datasetId"`
	// Applicable only to custom training with Datasets that have DataItems and Annotations.
	//
	// Cloud Storage URI that points to a YAML file describing the annotation schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#annotation_schema_uri VertexAiTrainingPipeline#annotation_schema_uri}
	AnnotationSchemaUri *string `field:"optional" json:"annotationSchemaUri" yaml:"annotationSchemaUri"`
	// Applicable only to Datasets that have DataItems and Annotations. A filter on Annotations of the Dataset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#annotations_filter VertexAiTrainingPipeline#annotations_filter}
	AnnotationsFilter *string `field:"optional" json:"annotationsFilter" yaml:"annotationsFilter"`
	// bigquery_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#bigquery_destination VertexAiTrainingPipeline#bigquery_destination}
	BigqueryDestination *VertexAiTrainingPipelineInputDataConfigBigqueryDestination `field:"optional" json:"bigqueryDestination" yaml:"bigqueryDestination"`
	// filter_split block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#filter_split VertexAiTrainingPipeline#filter_split}
	FilterSplit *VertexAiTrainingPipelineInputDataConfigFilterSplit `field:"optional" json:"filterSplit" yaml:"filterSplit"`
	// fraction_split block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#fraction_split VertexAiTrainingPipeline#fraction_split}
	FractionSplit *VertexAiTrainingPipelineInputDataConfigFractionSplit `field:"optional" json:"fractionSplit" yaml:"fractionSplit"`
	// gcs_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#gcs_destination VertexAiTrainingPipeline#gcs_destination}
	GcsDestination *VertexAiTrainingPipelineInputDataConfigGcsDestination `field:"optional" json:"gcsDestination" yaml:"gcsDestination"`
	// Whether to persist the ML use assignment to data item system labels.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#persist_ml_use_assignment VertexAiTrainingPipeline#persist_ml_use_assignment}
	PersistMlUseAssignment interface{} `field:"optional" json:"persistMlUseAssignment" yaml:"persistMlUseAssignment"`
	// predefined_split block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#predefined_split VertexAiTrainingPipeline#predefined_split}
	PredefinedSplit *VertexAiTrainingPipelineInputDataConfigPredefinedSplit `field:"optional" json:"predefinedSplit" yaml:"predefinedSplit"`
	// Only applicable to Datasets that have SavedQueries.
	//
	// The ID of a SavedQuery under the Dataset specified by dataset_id used for filtering Annotations for training.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#saved_query_id VertexAiTrainingPipeline#saved_query_id}
	SavedQueryId *string `field:"optional" json:"savedQueryId" yaml:"savedQueryId"`
	// stratified_split block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#stratified_split VertexAiTrainingPipeline#stratified_split}
	StratifiedSplit *VertexAiTrainingPipelineInputDataConfigStratifiedSplit `field:"optional" json:"stratifiedSplit" yaml:"stratifiedSplit"`
	// timestamp_split block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#timestamp_split VertexAiTrainingPipeline#timestamp_split}
	TimestampSplit *VertexAiTrainingPipelineInputDataConfigTimestampSplit `field:"optional" json:"timestampSplit" yaml:"timestampSplit"`
}

