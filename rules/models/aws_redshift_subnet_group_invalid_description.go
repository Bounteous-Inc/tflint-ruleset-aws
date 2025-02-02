// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsRedshiftSubnetGroupInvalidDescriptionRule checks the pattern is valid
type AwsRedshiftSubnetGroupInvalidDescriptionRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
}

// NewAwsRedshiftSubnetGroupInvalidDescriptionRule returns new rule with default attributes
func NewAwsRedshiftSubnetGroupInvalidDescriptionRule() *AwsRedshiftSubnetGroupInvalidDescriptionRule {
	return &AwsRedshiftSubnetGroupInvalidDescriptionRule{
		resourceType:  "aws_redshift_subnet_group",
		attributeName: "description",
		max:           2147483647,
	}
}

// Name returns the rule name
func (r *AwsRedshiftSubnetGroupInvalidDescriptionRule) Name() string {
	return "aws_redshift_subnet_group_invalid_description"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRedshiftSubnetGroupInvalidDescriptionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRedshiftSubnetGroupInvalidDescriptionRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRedshiftSubnetGroupInvalidDescriptionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRedshiftSubnetGroupInvalidDescriptionRule) Check(runner tflint.Runner) error {
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
			if len(val) > r.max {
				runner.EmitIssue(
					r,
					"description must be 2147483647 characters or less",
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
