// This file generated by `generator/`. DO NOT EDIT

package models

import (
	"fmt"

	"github.com/terraform-linters/tflint-plugin-sdk/hclext"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsRoute53ZoneAssociationInvalidVpcRegionRule checks the pattern is valid
type AwsRoute53ZoneAssociationInvalidVpcRegionRule struct {
	tflint.DefaultRule

	resourceType  string
	attributeName string
	max           int
	min           int
	enum          []string
}

// NewAwsRoute53ZoneAssociationInvalidVpcRegionRule returns new rule with default attributes
func NewAwsRoute53ZoneAssociationInvalidVpcRegionRule() *AwsRoute53ZoneAssociationInvalidVpcRegionRule {
	return &AwsRoute53ZoneAssociationInvalidVpcRegionRule{
		resourceType:  "aws_route53_zone_association",
		attributeName: "vpc_region",
		max:           64,
		min:           1,
		enum: []string{
			"us-east-1",
			"us-east-2",
			"us-west-1",
			"us-west-2",
			"eu-west-1",
			"eu-west-2",
			"eu-west-3",
			"eu-central-1",
			"eu-central-2",
			"ap-east-1",
			"me-south-1",
			"us-gov-west-1",
			"us-gov-east-1",
			"us-iso-east-1",
			"us-iso-west-1",
			"us-isob-east-1",
			"me-central-1",
			"ap-southeast-1",
			"ap-southeast-2",
			"ap-southeast-3",
			"ap-south-1",
			"ap-south-2",
			"ap-northeast-1",
			"ap-northeast-2",
			"ap-northeast-3",
			"eu-north-1",
			"sa-east-1",
			"ca-central-1",
			"cn-north-1",
			"af-south-1",
			"eu-south-1",
			"eu-south-2",
			"ap-southeast-4",
		},
	}
}

// Name returns the rule name
func (r *AwsRoute53ZoneAssociationInvalidVpcRegionRule) Name() string {
	return "aws_route53_zone_association_invalid_vpc_region"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRoute53ZoneAssociationInvalidVpcRegionRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRoute53ZoneAssociationInvalidVpcRegionRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRoute53ZoneAssociationInvalidVpcRegionRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsRoute53ZoneAssociationInvalidVpcRegionRule) Check(runner tflint.Runner) error {
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
					"vpc_region must be 64 characters or less",
					attribute.Expr.Range(),
				)
			}
			if len(val) < r.min {
				runner.EmitIssue(
					r,
					"vpc_region must be 1 characters or higher",
					attribute.Expr.Range(),
				)
			}
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					fmt.Sprintf(`"%s" is an invalid value as vpc_region`, truncateLongMessage(val)),
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
