terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

# Configure the AWS Provider
provider "aws" {
  region = "us-east-1"
  secret_key = secrets.aws_access_key_id 
  access_key = secrets.aws_secret_access_key
}

resource "aws_instance" "web" {
  ami           = "ami-08a52ddb321b32a8c"
  instance_type = "t2.micro"

  tags = {
    Name = "go-app-instance"
  }
}