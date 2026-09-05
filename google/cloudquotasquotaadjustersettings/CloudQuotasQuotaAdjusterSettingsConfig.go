// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package cloudquotasquotaadjustersettings

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type CloudQuotasQuotaAdjusterSettingsConfig struct {
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
	// Required. The configured value of the enablement at the given resource. Possible values: ["ENABLED", "DISABLED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/cloud_quotas_quota_adjuster_settings#enablement CloudQuotasQuotaAdjusterSettings#enablement}
	Enablement *string `field:"required" json:"enablement" yaml:"enablement"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/cloud_quotas_quota_adjuster_settings#id CloudQuotasQuotaAdjusterSettings#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The parent of the quota preference. Allowed parent format is "projects/[project-id / number]".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/cloud_quotas_quota_adjuster_settings#parent CloudQuotasQuotaAdjusterSettings#parent}
	Parent *string `field:"optional" json:"parent" yaml:"parent"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/cloud_quotas_quota_adjuster_settings#timeouts CloudQuotasQuotaAdjusterSettings#timeouts}
	Timeouts *CloudQuotasQuotaAdjusterSettingsTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

