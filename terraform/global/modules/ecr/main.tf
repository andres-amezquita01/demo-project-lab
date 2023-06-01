data "aws_ecr_repository" "registry" {
  name = var.registry
}