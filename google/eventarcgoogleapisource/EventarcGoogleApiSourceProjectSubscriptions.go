// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package eventarcgoogleapisource


type EventarcGoogleApiSourceProjectSubscriptions struct {
	// A list of projects to receive events from.
	//
	// All the projects must be in the same org. The listed projects should have
	// the format project/{identifier} where identifier can be either the
	// project id for project number. A single list may contain both formats. At
	// most 100 projects can be listed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/8.11.0/docs/resources/eventarc_google_api_source#list EventarcGoogleApiSource#list}
	List *[]*string `field:"required" json:"list" yaml:"list"`
}

