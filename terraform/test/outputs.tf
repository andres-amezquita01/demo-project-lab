output "ecs_cluster" {
  value = data.aws_ecs_cluster.final_demo_cluster
}

output "ecs_service" {
  value = data.aws_ecs_service.final_demo_service
}
output "ecs_goapp_td" {
  value = data.aws_ecs_task_definition.go-app-td
}