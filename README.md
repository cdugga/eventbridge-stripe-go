
# eventbridge-stripe-go 

`eventbridge-stripe-go` Builds a serverless architecture to handle [Stripe Webhook](https://stripe.com/docs/api/webhook_endpoints) events 

The AWS CDK stack deploys the following serverless components:
* A single API Gateway endpoint used as the target for the Stripe Webhook [customer.created](https://stripe.com/docs/api/events/types#event_types-customer.created) webhook event
* GO Lambda handler functions for handling the initial API Gateway request and subsequent Eventbridge event notification that new Stripe customer creation has occured
* DynamoDB table where new customer details are written
* Eventbridge event bus which orchestrates AWS servcies based on various events


## Setup 

Please refer to [AWS CDK Python workshop](https://cdkworkshop.com/30-python/20-create-project/200-virtualenv.html) for a more detailed set of instructions for initializing and using the python language with the CDK.  

Configure Python virtual env
```
python -m venv .env/
source .env/bin/activate
pip install -r requirements.txt
```

### Building Lambda handlers

Executable versions for each function have been included for convenience 

See [AWS Lambda deployment package in Go](https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html) for further instructions on how to package a Go lambda function. 

```
cd lambda/stripe-create-customer/
GOOS=linux go build -o createCustomerHandler github.com/cdugga/eventbridge-stripe-go/createCustomer

cd lambda/stripe-create-customer/
GOOS=linux go build -o stripeWebhookHandler github.com/cdugga/eventbridge-stripe-go/stripeWehbookHandler
```
Running programs target the Linux operating system. Use the GOOS runtime value to modify if required; [Go Runtime](GOOS is the running program's operating system target)

Deploy CDK stack
```
cdk synth
cdk deploy
```

## Tools used
* [Go](https://golang.org/)
* [Python](https://www.python.org/)
* [AWS CDK](https://github.com/aws/aws-cdk)
* [AWS APIGateway](https://aws.amazon.com/api-gateway/)
* [AWS Lambda](https://aws.amazon.com/lambda/)
* [AWS Eventbridge](https://aws.amazon.com/eventbridge/)
* [AWS Dynamodb](https://aws.amazon.com/dynamodb/)
