// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsRoute53ResolverFirewallRuleInvalidFirewallDomainListIDRule checks the pattern is valid
type AwsRoute53ResolverFirewallRuleInvalidFirewallDomainListIDRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsRoute53ResolverFirewallRuleInvalidFirewallDomainListIDRule returns new rule with default attributes
func NewAwsRoute53ResolverFirewallRuleInvalidFirewallDomainListIDRule() *AwsRoute53ResolverFirewallRuleInvalidFirewallDomainListIDRule {
	return &AwsRoute53ResolverFirewallRuleInvalidFirewallDomainListIDRule{
		resourceType:  "aws_route53_resolver_firewall_rule",
		attributeName: "firewall_domain_list_id",
		max:           64,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsRoute53ResolverFirewallRuleInvalidFirewallDomainListIDRule) Name() string {
	return "aws_route53_resolver_firewall_rule_invalid_firewall_domain_list_id"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRoute53ResolverFirewallRuleInvalidFirewallDomainListIDRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRoute53ResolverFirewallRuleInvalidFirewallDomainListIDRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRoute53ResolverFirewallRuleInvalidFirewallDomainListIDRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRoute53ResolverFirewallRuleInvalidFirewallDomainListIDRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"firewall_domain_list_id must be 64 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"firewall_domain_list_id must be 1 characters or higher",
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
