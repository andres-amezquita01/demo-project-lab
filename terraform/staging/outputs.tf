output "staging_lb" {
  value = aws_alb.lb.dns_name
}
output "staging_cluster_name" {
  value = aws_ecs_cluster.main_cluster.name
}
output "task_arn" {
  value = aws_ecs_task_definition.task.arn
}