// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package vertexaicacheconfig

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type VertexAiCacheConfigConfig struct {
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
	// If set to true, disables GenAI caching. Otherwise caching is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_cache_config#disable_cache VertexAiCacheConfig#disable_cache}
	DisableCache interface{} `field:"required" json:"disableCache" yaml:"disableCache"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_cache_config#id VertexAiCacheConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_cache_config#project VertexAiCacheConfig#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// The region of the Vertex AI API endpoint used to manage the cache config, e.g. 'us-central1'. The cache config is a project-level setting, so this only selects which regional endpoint requests are sent to. If not specified, the provider-level region is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_cache_config#region VertexAiCacheConfig#region}
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/vertex_ai_cache_config#timeouts VertexAiCacheConfig#timeouts}
	Timeouts *VertexAiCacheConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

