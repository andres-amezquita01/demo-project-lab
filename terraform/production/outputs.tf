output "production_lb" {
  value = aws_alb.lb.dns_name
}
output "production_cluster_name" {
  value = aws_ecs_cluster.main_cluster.name
}
output "production_task_arn" {
  value = aws_ecs_task_definition.task.arn
}