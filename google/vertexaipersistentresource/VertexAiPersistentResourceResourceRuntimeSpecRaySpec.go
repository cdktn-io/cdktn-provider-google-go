// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaipersistentresource


type VertexAiPersistentResourceResourceRuntimeSpecRaySpec struct {
	// Default image for user to choose a preferred ML framework (for example, TensorFlow or Pytorch) by choosing from [Vertex prebuilt images](https://cloud.google.com/vertex-ai/docs/training/pre-built-containers). Either this or the resource_pool_images is required. Use this field if you need all the resource pools to have the same Ray image. Otherwise, use the 'resource_pool_images' field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.9.0/docs/resources/vertex_ai_persistent_resource#image_uri VertexAiPersistentResource#image_uri}
	ImageUri *string `field:"optional" json:"imageUri" yaml:"imageUri"`
}

