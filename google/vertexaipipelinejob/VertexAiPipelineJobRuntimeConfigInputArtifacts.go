// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaipipelinejob


type VertexAiPipelineJobRuntimeConfigInputArtifacts struct {
	// Artifact resource id from MLMD.
	//
	// Which is the last portion of an artifact resource name: 'projects/{project}/locations/{location}/metadataStores/default/artifacts/{artifact_id}'. The artifact must be within the same project, location and default metadatastore as the pipeline.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_pipeline_job#artifact_id VertexAiPipelineJob#artifact_id}
	ArtifactId *string `field:"required" json:"artifactId" yaml:"artifactId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_pipeline_job#artifact_key VertexAiPipelineJob#artifact_key}.
	ArtifactKey *string `field:"required" json:"artifactKey" yaml:"artifactKey"`
}

