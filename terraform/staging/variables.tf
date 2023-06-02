#Variables used in remote backend
variable "region" {
  default = "us-east-1"
}
variable "environment" {
  default = "staging"
}
variable "go_app_staging_bucket" {
  default = "go-app-staging-terraform-state"
}
variable "dynamodb_table" {
  default = "staging_terraform_state_locking"
}
variable "bucket_key" {
  default = "staging/terraform.tfstate"
}

#Variables used in ...