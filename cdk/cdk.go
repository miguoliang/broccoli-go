package main

import (
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscognito"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3deployment"
	"github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type CdkStackProps struct {
	awscdk.StackProps
}

func NewCdkStack(scope constructs.Construct, id string, props *CdkStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	userPool := awscognito.NewUserPool(stack, jsii.String("broccoli-go-user-pool"), &awscognito.UserPoolProps{
		UserPoolName: jsii.String("broccoli-go-user-pool"),
		SignInAliases: &awscognito.SignInAliases{
			Username: jsii.Bool(true),
			Email:    jsii.Bool(true),
		},
		AutoVerify: &awscognito.AutoVerifiedAttrs{
			Email: jsii.Bool(true),
		},
		StandardAttributes: &awscognito.StandardAttributes{
			Email: &awscognito.StandardAttribute{
				Required: jsii.Bool(true),
			},
		},
		CustomAttributes: &map[string]awscognito.ICustomAttribute{
			"stripe_customer_id": awscognito.NewStringAttribute(&awscognito.StringAttributeProps{}),
		},
	})

	awscognito.NewUserPoolClient(stack, jsii.String("broccoli-go-user-pool-client"), &awscognito.UserPoolClientProps{
		UserPool:           userPool,
		UserPoolClientName: jsii.String("broccoli-go-user-pool-client"),
		AuthFlows: &awscognito.AuthFlow{
			UserPassword: jsii.Bool(true),
			UserSrp:      jsii.Bool(true),
		},
		OAuth: &awscognito.OAuthSettings{
			Flows: &awscognito.OAuthFlows{
				AuthorizationCodeGrant: jsii.Bool(true),
			},
			CallbackUrls: &[]*string{
				jsii.String("http://localhost:5173/"),
			},
			LogoutUrls: &[]*string{
				jsii.String("http://localhost:5173/"),
			},
		},
	})

	awscognito.NewCfnUserPoolDomain(stack, jsii.String("broccoli-go-user-pool-domain"), &awscognito.CfnUserPoolDomainProps{
		Domain:     jsii.String("broccoli-go-user-pool-domain"),
		UserPoolId: userPool.UserPoolId(),
	})

	authorizer := awsapigateway.NewCognitoUserPoolsAuthorizer(stack, jsii.String("broccoli-go-authorizer"), &awsapigateway.CognitoUserPoolsAuthorizerProps{
		CognitoUserPools: &[]awscognito.IUserPool{userPool},
		IdentitySource:   jsii.String("method.request.header.Authorization"),
	})

	function := awscdklambdagoalpha.NewGoFunction(stack, jsii.String("broccoli-go"),
		&awscdklambdagoalpha.GoFunctionProps{
			Runtime: awslambda.Runtime_PROVIDED_AL2023(),
			Environment: &map[string]*string{
				"GIN_MODE":              jsii.String(os.Getenv("GIN_MODE")),
				"STRIPE_SECRET_KEY":     jsii.String(os.Getenv("STRIPE_SECRET_KEY")),
				"STRIPE_WEBHOOK_SECRET": jsii.String(os.Getenv("STRIPE_WEBHOOK_SECRET")),
			},
			Architecture: awslambda.Architecture_ARM_64(),
			Entry:        jsii.String("../cmd"),
		})

	app := awsapigateway.NewRestApi(stack, jsii.String("broccoli-go-api"), nil)

	webhook := app.Root().AddResource(jsii.String("webhook"), nil)
	webhook.AddProxy(&awsapigateway.ProxyResourceOptions{
		AnyMethod:          jsii.Bool(true),
		DefaultIntegration: awsapigateway.NewLambdaIntegration(function, nil),
	})

	api := app.Root().AddResource(jsii.String("api"), nil)
	api.AddProxy(&awsapigateway.ProxyResourceOptions{
		AnyMethod:          jsii.Bool(true),
		DefaultIntegration: awsapigateway.NewLambdaIntegration(function, nil),
		DefaultMethodOptions: &awsapigateway.MethodOptions{
			AuthorizationType: awsapigateway.AuthorizationType_COGNITO,
			Authorizer:        authorizer,
		},
	})

	// host static website on S3
	bucket := awss3.NewBucket(stack, jsii.String("broccoli-go-bucket"), &awss3.BucketProps{
		WebsiteIndexDocument: jsii.String("index.html"),
		Encryption:           awss3.BucketEncryption_S3_MANAGED,
		PublicReadAccess:     jsii.Bool(false),
	})

	awss3deployment.NewBucketDeployment(stack, jsii.String("broccoli-go-bucket-deployment"), &awss3deployment.BucketDeploymentProps{
		Sources: &[]awss3deployment.ISource{
			awss3deployment.Source_Asset(jsii.String("../web/dist"), nil),
		},
		DestinationBucket: bucket,
	})

	// host static docs on S3
	docsBucket := awss3.NewBucket(stack, jsii.String("broccoli-go-docs-bucket"), &awss3.BucketProps{
		WebsiteIndexDocument: jsii.String("index.html"),
		Encryption:           awss3.BucketEncryption_S3_MANAGED,
		PublicReadAccess:     jsii.Bool(false),
	})

	awss3deployment.NewBucketDeployment(stack, jsii.String("broccoli-go-docs-bucket-deployment"), &awss3deployment.BucketDeploymentProps{
		Sources: &[]awss3deployment.ISource{
			awss3deployment.Source_Asset(jsii.String("../docs/build"), nil),
		},
		DestinationBucket: docsBucket,
	})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewCdkStack(app, "CdkStack", &CdkStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return nil

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
