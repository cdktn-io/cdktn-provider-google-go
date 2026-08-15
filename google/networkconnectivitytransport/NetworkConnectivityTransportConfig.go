// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package networkconnectivitytransport

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type NetworkConnectivityTransportConfig struct {
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
	// Name of the resource, see google.aip.dev/122 for resource naming.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/network_connectivity_transport#name NetworkConnectivityTransport#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The region of this resource.
	//
	// This is required to construct the resource name, but is not sent to the API since the region is already contained in the parent field.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/network_connectivity_transport#region NetworkConnectivityTransport#region}
	Region *string `field:"required" json:"region" yaml:"region"`
	// Resource URL of the remoteTransportProfile that this Transport is connecting to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/network_connectivity_transport#remote_profile NetworkConnectivityTransport#remote_profile}
	RemoteProfile *string `field:"required" json:"remoteProfile" yaml:"remoteProfile"`
	// List of IP Prefixes that will be advertised to the remote provider. Both IPv4 and IPv6 addresses are supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/network_connectivity_transport#advertised_routes NetworkConnectivityTransport#advertised_routes}
	AdvertisedRoutes *[]*string `field:"optional" json:"advertisedRoutes" yaml:"advertisedRoutes"`
	// Bandwidth of the Transport. This must be one of the supported bandwidths for the remote profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/network_connectivity_transport#bandwidth NetworkConnectivityTransport#bandwidth}
	Bandwidth *string `field:"optional" json:"bandwidth" yaml:"bandwidth"`
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/network_connectivity_transport#deletion_policy NetworkConnectivityTransport#deletion_policy}
	DeletionPolicy *string `field:"optional" json:"deletionPolicy" yaml:"deletionPolicy"`
	// An optional description of this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/network_connectivity_transport#description NetworkConnectivityTransport#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/network_connectivity_transport#id NetworkConnectivityTransport#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Optional labels in key:value format. For more information about labels, see [Requirements for labels](https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements).
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field 'effective_labels' for all of the labels present on the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/network_connectivity_transport#labels NetworkConnectivityTransport#labels}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// [Output only] The maximum transmission unit (MTU) of a packet that can be sent over this transport.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/network_connectivity_transport#mtu_limit NetworkConnectivityTransport#mtu_limit}
	MtuLimit *float64 `field:"optional" json:"mtuLimit" yaml:"mtuLimit"`
	// Resource URL of the Network that will be peered with this Transport.
	//
	// This field must be provided during resource creation and cannot be changed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/network_connectivity_transport#network NetworkConnectivityTransport#network}
	Network *string `field:"optional" json:"network" yaml:"network"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/network_connectivity_transport#project NetworkConnectivityTransport#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// Key used for establishing a connection with the remote transport.
	//
	// This key can only be provided if the profile supports an INPUT key flow and the resource is in the PENDING_KEY state.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/network_connectivity_transport#provided_activation_key NetworkConnectivityTransport#provided_activation_key}
	ProvidedActivationKey *string `field:"optional" json:"providedActivationKey" yaml:"providedActivationKey"`
	// The user supplied account id for the CSP associated with the remote profile.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/network_connectivity_transport#remote_account_id NetworkConnectivityTransport#remote_account_id}
	RemoteAccountId *string `field:"optional" json:"remoteAccountId" yaml:"remoteAccountId"`
	// IP version stack for the established connectivity. Possible values: ["IPV4_IPV6", "IPV4_ONLY"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/network_connectivity_transport#stack_type NetworkConnectivityTransport#stack_type}
	StackType *string `field:"optional" json:"stackType" yaml:"stackType"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.44.0/docs/resources/network_connectivity_transport#timeouts NetworkConnectivityTransport#timeouts}
	Timeouts *NetworkConnectivityTransportTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

