package main

import (
	"cdk-workshop/hitcounter"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"

	"github.com/cdklabs/cdk-dynamo-table-viewer-go/dynamotableviewer"
)

type CdkWorkshopStackProps struct {
	awscdk.StackProps
}

func NewCdkWorkshopStack(scope constructs.Construct, id string, props *CdkWorkshopStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	helloHandler := awslambda.NewFunction(stack, jsii.String("HelloHandler"), &awslambda.FunctionProps{
		Code:    awslambda.Code_FromAsset(jsii.String("lambda"), nil),
		Runtime: awslambda.Runtime_NODEJS_16_X(),
		Handler: jsii.String("hello.handler"),
	})

	hitcounter := hitcounter.NewHitCounter(stack, "HelloHitCounter", &hitcounter.HitCounterProps{
		Downstream: helloHandler,
	})

	awsapigateway.NewLambdaRestApi(stack, jsii.String("Endpoint"), &awsapigateway.LambdaRestApiProps{
		Handler: hitcounter.Handler(),
	})

	dynamotableviewer.NewTableViewer(stack, jsii.String("ViewHitCounter"), &dynamotableviewer.TableViewerProps{
		Table:  hitcounter.Table(),
		Title:  jsii.String("Hello Hits"),
		SortBy: Ptr("-hits"), //https://stackoverflow.com/a/30716481
	})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewCdkWorkshopStack(app, "CdkWorkshopStack", &CdkWorkshopStackProps{})
	app.Synth(nil)
}

func Ptr[T any](v T) *T {
	return &v
}
