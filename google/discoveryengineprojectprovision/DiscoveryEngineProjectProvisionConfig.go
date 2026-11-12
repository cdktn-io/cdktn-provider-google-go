// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package discoveryengineprojectprovision

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DiscoveryEngineProjectProvisionConfig struct {
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
	// Set to 'true' to specify that the caller has read and consents to the [Terms for data use](https://cloud.google.com/retail/data-use-terms).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/discovery_engine_project_provision#accept_data_use_terms DiscoveryEngineProjectProvision#accept_data_use_terms}
	AcceptDataUseTerms interface{} `field:"optional" json:"acceptDataUseTerms" yaml:"acceptDataUseTerms"`
	// The version of the [Terms for data use](https://cloud.google.com/retail/data-use-terms) that the caller has read and consents to. An acceptable version is '2022-11-23', and this may change over time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/discovery_engine_project_provision#data_use_terms_version DiscoveryEngineProjectProvision#data_use_terms_version}
	DataUseTermsVersion *string `field:"optional" json:"dataUseTermsVersion" yaml:"dataUseTermsVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/discovery_engine_project_provision#id DiscoveryEngineProjectProvision#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The geographic location in which to provision the project, such as 'global', 'us', or 'eu'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/discovery_engine_project_provision#location DiscoveryEngineProjectProvision#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/discovery_engine_project_provision#project DiscoveryEngineProjectProvision#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// saas_params block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/discovery_engine_project_provision#saas_params DiscoveryEngineProjectProvision#saas_params}
	SaasParams *DiscoveryEngineProjectProvisionSaasParams `field:"optional" json:"saasParams" yaml:"saasParams"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/discovery_engine_project_provision#timeouts DiscoveryEngineProjectProvision#timeouts}
	Timeouts *DiscoveryEngineProjectProvisionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

