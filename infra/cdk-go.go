package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func NewCdkGoStack(scope constructs.Construct, id string, props *awscdk.StackProps) awscdk.Stack {
	stack := awscdk.NewStack(scope, &id, props)

	subnets := []*awsec2.SubnetConfiguration{
		{
			CidrMask:   jsii.Number(24),
			Name:       jsii.String("subnet-1"),
			SubnetType: awsec2.SubnetType_PUBLIC,
		},
	}
	vpc := awsec2.NewVpc(stack, jsii.String("testVPC"), &awsec2.VpcProps{
		VpcName:             jsii.String("FirstTestVPC"),
		EnableDnsHostnames:  jsii.Bool(true),
		EnableDnsSupport:    jsii.Bool(true),
		NatGateways:         jsii.Number(0),
		SubnetConfiguration: &subnets,
		MaxAzs:              jsii.Number(1),
	})

	// ECR
	awsecr.NewRepository(stack, jsii.String("ecr-repo"), &awsecr.RepositoryProps{
		RepositoryName: jsii.String("ecr-test"),
	})
	// EC2
	awsec2.NewInstance(stack, jsii.String("ec2-machine"), &awsec2.InstanceProps{
		InstanceType: awsec2.InstanceType_Of(awsec2.InstanceClass_T3, awsec2.InstanceSize_MICRO),
		MachineImage: awsec2.NewAmazonLinuxImage(&awsec2.AmazonLinuxImageProps{
			Generation: awsec2.AmazonLinuxGeneration_AMAZON_LINUX_2,
		}),
		Vpc: vpc,
		VpcSubnets: &awsec2.SubnetSelection{
			SubnetType: awsec2.SubnetType_PUBLIC,
		},
	})
	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewCdkGoStack(app, "CdkGoTest", &awscdk.StackProps{Env: env()})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	return &awscdk.Environment{
		Account: jsii.String("550610942901"),
		Region:  jsii.String("us-east-1"),
	}
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
