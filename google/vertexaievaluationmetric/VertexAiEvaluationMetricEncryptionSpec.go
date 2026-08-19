// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaievaluationmetric


type VertexAiEvaluationMetricEncryptionSpec struct {
	// Required.
	//
	// The Cloud KMS resource identifier of the customer managed encryption key
	// used to protect a resource. Has the form:
	// 'projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key'.
	// The key needs to be in the same region as where the resource is created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/vertex_ai_evaluation_metric#kms_key_name VertexAiEvaluationMetric#kms_key_name}
	KmsKeyName *string `field:"optional" json:"kmsKeyName" yaml:"kmsKeyName"`
}

