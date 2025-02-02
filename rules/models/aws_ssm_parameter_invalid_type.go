// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsSsmParameterInvalidTypeRule checks the pattern is valid
type AwsSsmParameterInvalidTypeRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsSsmParameterInvalidTypeRule returns new rule with default attributes
func NewAwsSsmParameterInvalidTypeRule() *AwsSsmParameterInvalidTypeRule {
	return &AwsSsmParameterInvalidTypeRule{
		resourceType:  "aws_ssm_parameter",
		attributeName: "type",
		enum: []string{
			"String",
			"StringList",
			"SecureString",
		},
	}
}

// Name returns the rule name
func (r *AwsSsmParameterInvalidTypeRule) Name() string {
	return "aws_ssm_parameter_invalid_type"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsSsmParameterInvalidTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsSsmParameterInvalidTypeRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsSsmParameterInvalidTypeRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsSsmParameterInvalidTypeRule) Check(runner tflint.Runner) error {
	logger.Trace("Check `%s` rule", r.Name())

	resources, err := runner.GetResourceContent(r.resourceType, &hclext.BodySchema{
		Attributes: []hclext.AttributeSchema{
			{Name: r.attributeName},
		},
	}, nil)
	if err != nil {
		return err
	}

	for _, resource := range resources.Blocks {
		attribute, exists := resource.Body.Attributes[r.attributeName]
		if !exists {
			continue
		}

		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val, nil)

		err = runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as type`, truncateLongMessage(val)),
					attribute.Expr.Range(),
				)
			}
			return nil
		})
		if err != nil {
			return err
		}
	}

	return nil
}
