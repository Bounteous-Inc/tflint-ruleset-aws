// This file generated by `generator/main.go`. DO NOT EDIT

package api

import (
	"fmt"
	"log"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
    "github.com/terraform-linters/tflint-ruleset-aws/aws"
)

// AwsRouteInvalidNatGatewayRule checks whether attribute value actually exists
type AwsRouteInvalidNatGatewayRule struct {
	resourceType  string
	attributeName string
	data          map[string]bool
	dataPrepared  bool
}

// NewAwsRouteInvalidNatGatewayRule returns new rule with default attributes
func NewAwsRouteInvalidNatGatewayRule() *AwsRouteInvalidNatGatewayRule {
	return &AwsRouteInvalidNatGatewayRule{
		resourceType:  "aws_route",
		attributeName: "nat_gateway_id",
		data:          map[string]bool{},
		dataPrepared:  false,
	}
}

// Name returns the rule name
func (r *AwsRouteInvalidNatGatewayRule) Name() string {
	return "aws_route_invalid_nat_gateway"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsRouteInvalidNatGatewayRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsRouteInvalidNatGatewayRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsRouteInvalidNatGatewayRule) Link() string {
	return ""
}

// Check checks whether the attributes are included in the list retrieved by DescribeNatGateways
func (r *AwsRouteInvalidNatGatewayRule) Check(rr tflint.Runner) error {
    runner := rr.(*aws.Runner)

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		if !r.dataPrepared {
			log.Print("[DEBUG] invoking DescribeNatGateways")
			var err error
			r.data, err = runner.AwsClient.DescribeNatGateways()
			if err != nil {
				err := &tflint.Error{
					Code:    tflint.ExternalAPIError,
					Level:   tflint.ErrorLevel,
					Message: "An error occurred while invoking DescribeNatGateways",
					Cause:   err,
				}
				log.Printf("[ERROR] %s", err)
				return err
			}
			r.dataPrepared = true
		}

		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			if !r.data[val] {
				runner.EmitIssueOnExpr(
					r,
					fmt.Sprintf(`"%s" is invalid NAT gateway ID.`, val),
					attribute.Expr,
				)
			}
			return nil
		})
	})
}
