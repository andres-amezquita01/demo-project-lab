#Global
variable "environment" {
  default = "staging"
}

#Variables used in main.tf for remote backend
variable "region" {
  default = "us-east-1"
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

#Variables used in ECS

variable "ecs_task_execution_role" {
  default = "ecsTaskExecutionRole"
}
variable "registry" {
  default = "final-demo"
}
variable "image_tag" {
  default = "latest"
}
#Variables used in VPC
variable "vpc_block" {
  default = "10.0.0.0/16"
}
variable "internet_destination_block" {
  default = "0.0.0.0/0"
}
variable "public_subnet_cidrs" {
 type        = list(string)
 description = "Public Subnet CIDR values"
 default     = ["10.0.1.0/24", "10.0.2.0/24"]
}
 
variable "private_subnet_cidrs" {
 type        = list(string)
 description = "Private Subnet CIDR values"
 default     = ["10.0.3.0/24", "10.0.4.0/24"]
}
variable "availability_zones" {
 type        = list(string)
 description = "Availability Zones list"
 default     = ["us-east-1a", "us-east-1b"]
}