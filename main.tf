data "aws_region" "current" {}

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

resource "aws_lambda_permission" "ip_info" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.ip_info.arn
  principal     = "apigateway.amazonaws.com"
}

resource "aws_api_gateway_resource" "ip_info" {
  rest_api_id = aws_api_gateway_rest_api.ip_info.id
  parent_id   = aws_api_gateway_rest_api.ip_info.root_resource_id
  path_part   = "ip_info"
}

resource "aws_api_gateway_rest_api" "ip_info" {
  name = "ip_info"
}

resource "aws_api_gateway_method" "ip_info" {
  rest_api_id   = aws_api_gateway_rest_api.ip_info.id
  resource_id   = aws_api_gateway_resource.ip_info.id
  http_method   = "GET"
  authorization = "NONE"
}

resource "aws_api_gateway_integration" "ip_info" {
  rest_api_id             = aws_api_gateway_rest_api.ip_info.id
  resource_id             = aws_api_gateway_resource.ip_info.id
  http_method             = aws_api_gateway_method.ip_info.http_method
  integration_http_method = "POST"
  type                    = "AWS_PROXY"
  uri                     = "arn:aws:apigateway:${data.aws_region.current.name}:lambda:path/2015-03-31/functions/${aws_lambda_function.ip_info.arn}/invocations"
}

resource "aws_api_gateway_deployment" "ip_info_v1" {
  depends_on = [
    "aws_api_gateway_integration.ip_info"
  ]
  rest_api_id = aws_api_gateway_rest_api.ip_info.id
  stage_name  = "v1"
}

output "url" {
  value = "${aws_api_gateway_deployment.ip_info_v1.invoke_url}${aws_api_gateway_resource.ip_info.path}?ip=142.251.46.142"
}
