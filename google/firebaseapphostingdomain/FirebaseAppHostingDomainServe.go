// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package firebaseapphostingdomain


type FirebaseAppHostingDomainServe struct {
	// redirect block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/6.50.0/docs/resources/firebase_app_hosting_domain#redirect FirebaseAppHostingDomain#redirect}
	Redirect *FirebaseAppHostingDomainServeRedirect `field:"optional" json:"redirect" yaml:"redirect"`
}

