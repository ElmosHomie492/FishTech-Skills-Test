# This is required to get the AWS region via ${data.aws_region.current}.
data "aws_region" "current" {
}

provider "aws" {
  region     = "us-east-1"
  access_key = "AKIA5JGRGIMK7RCXTLOC"
  secret_key = "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"
}

# Define a Lambda function.
#
# The handler is the name of the executable for go1.x runtime.
resource "aws_lambda_function" "ip_info" {
  function_name    = "ip_info"
  filename         = "ip_info.zip"
  handler          = "ip_info"
  source_code_hash = sha256(filebase64("ip_info.zip"))
  role             = aws_iam_role.ip_info.arn
  runtime          = "go1.x"
  memory_size      = 128
  timeout          = 1
}

# A Lambda function may access to other AWS resources such as S3 bucket. So an
# IAM role needs to be defined. This hello world example does not access to
# any resource, so the role is empty.
#
# The date 2012-10-17 is just the version of the policy language used here [1].
#
# [1]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_version.html
resource "aws_iam_role" "ip_info" {
  name               = "ip_info"
  assume_role_policy = <<POLICY
{
  "Version": "2012-10-17",
  "Statement": {
    "Action": "sts:AssumeRole",
    "Principal": {
      "Service": "lambda.amazonaws.com"
    },
    "Effect": "Allow"
  }
}
POLICY
}

# Allow API gateway to invoke the hello Lambda function.
resource "aws_lambda_permission" "ip_info" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.ip_info.arn
  principal     = "apigateway.amazonaws.com"
}

# A Lambda function is not a usual public REST API. We need to use AWS API
# Gateway to map a Lambda function to an HTTP endpoint.
resource "aws_api_gateway_resource" "ip_info" {
  rest_api_id = aws_api_gateway_rest_api.ip_info.id
  parent_id   = aws_api_gateway_rest_api.ip_info.root_resource_id
  path_part   = "ip_info"
}

resource "aws_api_gateway_rest_api" "ip_info" {
  name = "ip_info"
}

#           GET
# Internet -----> API Gateway
resource "aws_api_gateway_method" "ip_info" {
  rest_api_id   = aws_api_gateway_rest_api.ip_info.id
  resource_id   = aws_api_gateway_resource.ip_info.id
  http_method   = "GET"
  authorization = "NONE"
}

#              POST
# API Gateway ------> Lambda
# For Lambda the method is always POST and the type is always AWS_PROXY.
#
# The date 2015-03-31 in the URI is just the version of AWS Lambda.
resource "aws_api_gateway_integration" "ip_info" {
  rest_api_id             = aws_api_gateway_rest_api.ip_info.id
  resource_id             = aws_api_gateway_resource.ip_info.id
  http_method             = aws_api_gateway_method.ip_info.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = "arn:aws:apigateway:${data.aws_region.current.name}:lambda:path/2015-03-31/functions/${aws_lambda_function.ip_info.arn}/invocations"
}

# This resource defines the URL of the API Gateway.
resource "aws_api_gateway_deployment" "ip_info_v1" {
  depends_on = [
    "aws_api_gateway_integration.ip_info"
  ]
  rest_api_id = aws_api_gateway_rest_api.ip_info.id
  stage_name  = "v1"
}

# Set the generated URL as an output. Run `terraform output url` to get this.
output "url" {
  value = "${aws_api_gateway_deployment.ip_info_v1.invoke_url}${aws_api_gateway_resource.ip_info.path}"
}
