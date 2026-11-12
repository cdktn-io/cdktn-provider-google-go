// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaievaluationitem

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VertexAiEvaluationItemConfig struct {
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
	// The user-friendly display name for the EvaluationItem.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_evaluation_item#display_name VertexAiEvaluationItem#display_name}
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// The type of the EvaluationItem. Possible values: ["REQUEST", "RESULT"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_evaluation_item#evaluation_item_type VertexAiEvaluationItem#evaluation_item_type}
	EvaluationItemType *string `field:"required" json:"evaluationItemType" yaml:"evaluationItemType"`
	// The region of the EvaluationItem. eg us-central1.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_evaluation_item#region VertexAiEvaluationItem#region}
	Region *string `field:"required" json:"region" yaml:"region"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_evaluation_item#deletion_policy VertexAiEvaluationItem#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// evaluation_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_evaluation_item#evaluation_request VertexAiEvaluationItem#evaluation_request}
	EvaluationRequest *VertexAiEvaluationItemEvaluationRequest `field:"optional" json:"evaluationRequest" yaml:"evaluationRequest"`
	// The Cloud Storage object where the request or response is stored. Exactly one of 'evaluation_request' or 'gcs_uri' can be set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_evaluation_item#gcs_uri VertexAiEvaluationItem#gcs_uri}
	GcsUri *string `field:"optional" json:"gcsUri" yaml:"gcsUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_evaluation_item#id VertexAiEvaluationItem#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels for the EvaluationItem.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_evaluation_item#labels VertexAiEvaluationItem#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// User-defined metadata for this evaluation item as a JSON string.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_evaluation_item#metadata VertexAiEvaluationItem#metadata}
	Metadata *string `field:"optional" json:"metadata" yaml:"metadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_evaluation_item#project VertexAiEvaluationItem#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_evaluation_item#timeouts VertexAiEvaluationItem#timeouts}
	Timeouts *VertexAiEvaluationItemTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

