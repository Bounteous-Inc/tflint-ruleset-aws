// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"
	"log"
	"regexp"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsWafv2WebACLInvalidDescriptionRule checks the pattern is valid
type AwsWafv2WebACLInvalidDescriptionRule struct {
	resourceType  string
	attributeName string
	max           int
	min           int
	pattern       *regexp.Regexp
}

// NewAwsWafv2WebACLInvalidDescriptionRule returns new rule with default attributes
func NewAwsWafv2WebACLInvalidDescriptionRule() *AwsWafv2WebACLInvalidDescriptionRule {
	return &AwsWafv2WebACLInvalidDescriptionRule{
		resourceType:  "aws_wafv2_web_acl",
		attributeName: "description",
		max:           256,
		min:           1,
		pattern:       regexp.MustCompile(`^[\w+=:#@/\-,\.][\w+=:#@/\-,\.\s]+[\w+=:#@/\-,\.]$`),
	}
}

// Name returns the rule name
func (r *AwsWafv2WebACLInvalidDescriptionRule) Name() string {
	return "aws_wafv2_web_acl_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsWafv2WebACLInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsWafv2WebACLInvalidDescriptionRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsWafv2WebACLInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsWafv2WebACLInvalidDescriptionRule) Check(runner tflint.Runner) error {
	log.Printf("[TRACE] Check `%s` rule", r.Name())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		return runner.EnsureNoError(err, func() error {
			if len(val) > r.max {
				runner.EmitIssueOnExpr(
					r,
					"description must be 256 characters or less",
					attribute.Expr,
				)
			}
			if len(val) < r.min {
				runner.EmitIssueOnExpr(
					r,
					"description must be 1 characters or higher",
					attribute.Expr,
				)
			}
			if !r.pattern.MatchString(val) {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" does not match valid pattern %s`, truncateLongMessage(val), `^[\w+=:#@/\-,\.][\w+=:#@/\-,\.\s]+[\w+=:#@/\-,\.]$`),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
