package main

import (
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/aws/jsii-runtime-go"
)

func TestCdkGoStack(t *testing.T) {
	// GIVEN
	app := awscdk.NewApp(nil)

	// WHEN
	stack := NewCdkGoStack(app, "MyStack", nil)

	// THEN
	template := assertions.Template_FromStack(stack)

	template.HasResourceProperties(jsii.String("AWS::SQS::Queue"), map[string]interface{}{
		"VisibilityTimeout": 300,
	})
	template.ResourceCountIs(jsii.String("AWS::SNS::Topic"), jsii.Number(1))
}
