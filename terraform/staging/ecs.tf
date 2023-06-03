#min 54 load balancer
#create target group
data "aws_iam_role" "task_execution_role" {
  name = var.ecs_task_execution_role
}

# resource "aws_security_group" "name" {
  
# }

resource "aws_ecs_cluster" "main_cluster" {
  name = "${var.environment}-cluster"
  setting {
    name  = "containerInsights"
    value = "enabled"
  }
}

resource "aws_ecs_cluster_capacity_providers" "main_cp" {
  cluster_name = aws_ecs_cluster.main_cluster.name
  capacity_providers = [ "FARGATE" ]
}