// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package accesscontextmanagerserviceperimeters


type AccessContextManagerServicePerimetersServicePerimetersSpecEgressPoliciesEgressFromSourcesPscEndpoint struct {
	// The full resource name of the global forwarding rule that identifies a Private Service Connect endpoint. Forwarding rule format: '//compute.googleapis.com/projects/{PROJECT_ID}/global/forwardingRules/{FORWARDING_RULE_ID}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/access_context_manager_service_perimeters#forwarding_rule AccessContextManagerServicePerimeters#forwarding_rule}
	ForwardingRule *string `field:"optional" json:"forwardingRule" yaml:"forwardingRule"`
}

