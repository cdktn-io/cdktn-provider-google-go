// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package memorystoreaclpolicy


type MemorystoreAclPolicyRules struct {
	// The rule to be applied to the username.
	//
	// Ex: "on >password123 ~* +@all"
	// The format of the rule is defined by Valkey OSS:
	// https://valkey.io/topics/acl/
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/memorystore_acl_policy#rule MemorystoreAclPolicy#rule}
	Rule *string `field:"required" json:"rule" yaml:"rule"`
	// Specifies the IAM user or service account to be added to the ACL policy.
	//
	// This username will be directly set on the Valkey OSS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.46.1/docs/resources/memorystore_acl_policy#username MemorystoreAclPolicy#username}
	Username *string `field:"required" json:"username" yaml:"username"`
}

