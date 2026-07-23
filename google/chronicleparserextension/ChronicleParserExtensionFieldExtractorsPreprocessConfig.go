// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package chronicleparserextension


type ChronicleParserExtensionFieldExtractorsPreprocessConfig struct {
	// GROK Regex to extract the structured part of the log. syntax documentation: www.elastic.co/guide/en/logstash/current/plugins-filters-grok.html.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/chronicle_parser_extension#grok_regex ChronicleParserExtension#grok_regex}
	GrokRegex *string `field:"optional" json:"grokRegex" yaml:"grokRegex"`
	// Target field name for the structured part of the log. This should match a SEMANTIC identifier from the grok expression.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.41.0/docs/resources/chronicle_parser_extension#target ChronicleParserExtension#target}
	Target *string `field:"optional" json:"target" yaml:"target"`
}

