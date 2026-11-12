// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaitrainingpipeline


type VertexAiTrainingPipelineModelToUploadEncryptionSpec struct {
	// The Cloud KMS resource identifier of the customer-managed encryption key used to protect the model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_training_pipeline#kms_key_name VertexAiTrainingPipeline#kms_key_name}
	KmsKeyName *string `field:"required" json:"kmsKeyName" yaml:"kmsKeyName"`
}

