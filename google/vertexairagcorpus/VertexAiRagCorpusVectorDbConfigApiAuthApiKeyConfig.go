// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexairagcorpus


type VertexAiRagCorpusVectorDbConfigApiAuthApiKeyConfig struct {
	// The SecretManager secret version resource name storing API key. e.g. projects/{project}/secrets/{secret}/versions/{version}.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_rag_corpus#api_key_secret_version VertexAiRagCorpus#api_key_secret_version}
	ApiKeySecretVersion *string `field:"optional" json:"apiKeySecretVersion" yaml:"apiKeySecretVersion"`
	// The API key string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.6.0/docs/resources/vertex_ai_rag_corpus#api_key_string VertexAiRagCorpus#api_key_string}
	ApiKeyString *string `field:"optional" json:"apiKeyString" yaml:"apiKeyString"`
}

