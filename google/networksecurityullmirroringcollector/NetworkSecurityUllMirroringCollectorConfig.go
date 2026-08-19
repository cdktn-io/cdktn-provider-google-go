// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networksecurityullmirroringcollector

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NetworkSecurityUllMirroringCollectorConfig struct {
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
	// The engine resource to which this collector points to, for example: 'projects/123456789/locations/us-south1-d/ullMirroringEngines/my-engine'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/network_security_ull_mirroring_collector#engine NetworkSecurityUllMirroringCollector#engine}
	Engine *string `field:"required" json:"engine" yaml:"engine"`
	// The regional load balancer which the mirrored traffic should be forwarded to, for example: 'projects/123456789/regions/us-south1/forwardingRules/my-fr'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/network_security_ull_mirroring_collector#forwarding_rule NetworkSecurityUllMirroringCollector#forwarding_rule}
	ForwardingRule *string `field:"required" json:"forwardingRule" yaml:"forwardingRule"`
	// The cloud location of the collector, e.g. 'us-south1-d' or 'us-south1-e'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/network_security_ull_mirroring_collector#location NetworkSecurityUllMirroringCollector#location}
	Location *string `field:"required" json:"location" yaml:"location"`
	// The ID to use for the new collector, which will become the final component of the collector's resource name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/network_security_ull_mirroring_collector#ull_mirroring_collector_id NetworkSecurityUllMirroringCollector#ull_mirroring_collector_id}
	UllMirroringCollectorId *string `field:"required" json:"ullMirroringCollectorId" yaml:"ullMirroringCollectorId"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/network_security_ull_mirroring_collector#deletion_policy NetworkSecurityUllMirroringCollector#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/network_security_ull_mirroring_collector#id NetworkSecurityUllMirroringCollector#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Labels are key/value pairs that help to organize and filter resources.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/network_security_ull_mirroring_collector#labels NetworkSecurityUllMirroringCollector#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/network_security_ull_mirroring_collector#project NetworkSecurityUllMirroringCollector#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.45.0/docs/resources/network_security_ull_mirroring_collector#timeouts NetworkSecurityUllMirroringCollector#timeouts}
	Timeouts *NetworkSecurityUllMirroringCollectorTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

