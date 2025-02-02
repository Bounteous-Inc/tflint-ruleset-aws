// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsCodedeployDeploymentGroupInvalidAppNameRule checks the pattern is valid
type AwsCodedeployDeploymentGroupInvalidAppNameRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
}

// NewAwsCodedeployDeploymentGroupInvalidAppNameRule returns new rule with default attributes
func NewAwsCodedeployDeploymentGroupInvalidAppNameRule() *AwsCodedeployDeploymentGroupInvalidAppNameRule {
	return &AwsCodedeployDeploymentGroupInvalidAppNameRule{
		resourceType:  "aws_codedeploy_deployment_group",
		attributeName: "app_name",
		max:           100,
		min:           1,
	}
}

// Name returns the rule name
func (r *AwsCodedeployDeploymentGroupInvalidAppNameRule) Name() string {
	return "aws_codedeploy_deployment_group_invalid_app_name"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCodedeployDeploymentGroupInvalidAppNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsCodedeployDeploymentGroupInvalidAppNameRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsCodedeployDeploymentGroupInvalidAppNameRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCodedeployDeploymentGroupInvalidAppNameRule) Check(runner tflint.Runner) error {
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
					"app_name must be 100 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"app_name must be 1 characters or higher",
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
