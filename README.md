# Welcome AWS CDK Workshop project!

This is a simple project for Go development with CDK.

> The bootstrapping stack created through cdk bootstrap still exists. If you plan on using the CDK in the future (we hope you do!) do not delete this stack.

## Useful links

 * [AWS CDK Workshop](https://cdkworkshop.com/) & [the workshop source code](https://github.com/aws-samples/aws-cdk-intro-workshop)
 * [Security-related changes not in this list](https://github.com/aws/aws-cdk/issues/1299)
 * Logs tools like [SAM](https://github.com/aws/aws-sam-cli) & [awslog](https://github.com/jorgebastida/awslogs)
 * [Docs λ](https://docs.aws.amazon.com/lambda/latest/dg/welcome.html) & [Code signing](https://docs.aws.amazon.com/lambda/latest/dg/configuration-codesigning.html)
 * [Not for PROD!! Go dynamoviewer lib](https://pkg.go.dev/github.com/cdklabs/cdk-dynamo-table-viewer-go/dynamotableviewer#section-readme)

 > Construct libraries are a very powerful concept. They allow you to add complex capabilities to your apps with minimum effort. However, you must understand that with great power comes great responsibility. Constructs can add IAM permissions, expose data to the public or cause your application not to function. We are working on providing you tools for protecting your app, and identifying potential security issues with your stacks, but it is your responsibility to understand how certain constructs that you use impact your application, and to make sure you only use construct libraries from vendors you trust 

 ## Useful commands

 * `cdk deploy`      deploy this stack to your default AWS account/region
 * `cdk diff`        compare deployed stack with current state
 * `cdk synth`       emits the synthesized CloudFormation template
 * `cdk destroy`     destroy this stack
 * `cdk bootstrap`   deploy CloudFormation/CDKToolkit stack and S3 bucket
 * `go test`         run unit tests

The `cdk.json` file tells the CDK Toolkit how to execute your app.