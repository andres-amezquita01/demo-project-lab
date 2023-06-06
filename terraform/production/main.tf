terraform {
  backend "s3" {
    bucket = "go-app-production-terraform-state"
    key = "production/terraform.tfstate"
    region = "us-east-1"
    dynamodb_table = "production_terraform_state_locking"
    encrypt = true
  }
  
  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "~> 4.0"
    }
  }
}
provider "aws" {
  region = var.region
}

# resource "aws_s3_bucket" "terraform_state" {
#   bucket = var.go_app_staging_bucket
#   lifecycle {
#     prevent_destroy = true
#   }
# }
# resource "aws_s3_bucket_versioning" "terraform_state_versioning" {
#   bucket = aws_s3_bucket.terraform_state.id
#   versioning_configuration {
#     status = "Enabled"
#   }
# }
# resource "aws_s3_bucket_server_side_encryption_configuration" "encryption" {
#   bucket = aws_s3_bucket.terraform_state.id
#   rule {
#     apply_server_side_encryption_by_default {
#      sse_algorithm = "AES256" 
#     }
#   }
# }
# resource "aws_dynamodb_table" "terraform_locks" {
#   name = var.dynamodb_table
#   billing_mode = "PAY_PER_REQUEST"
#   hash_key = "LockID"
#   attribute {
#     name = "LockID"
#     type = "S"
#   }
# }