terraform {
  backend "s3" {
    bucket = "devops-project-lab"
    key = "tf-infra/terraform.tfstate"
    region = "us-east-1"
    dynamodb_table = "terraform-state-locking"
    encrypt = true
  }
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }
  }
}
provider "aws" {
  region = "us-east-1"
}
data "aws_ecr_repository" "registry" {
  name = "final-demo"
}
