data "aws_ecs_cluster" "final_demo_cluster" {
  cluster_name = var.cluster_name
}

data "aws_ecs_service" "final_demo_service" {
  service_name = var.service_name
  cluster_arn = data.aws_ecs_cluster.final_demo_cluster.arn
}
data "aws_ecs_task_definition" "go-app-td" {
  task_definition = var.task_family
}

resource "aws_ecs_task_definition" "my_task" {
  family = "${var.test_family_td}"
  task_role_arn = "arn:aws:iam::282335569253:role/ecsTaskExecutionRole"
  execution_role_arn = "arn:aws:iam::282335569253:role/ecsTaskExecutionRole"
  network_mode = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu = "512"
  memory = "1024"
  runtime_platform {
    cpu_architecture = "X86_64" 
    operating_system_family = "LINUX"
  }
  container_definitions = jsonencode([
    {
        "name": "test-image",
        "image": "282335569253.dkr.ecr.us-east-1.amazonaws.com/final-demo:${var.image_tag}",
        "cpu": 0,
        "portMappings": [
            {
                "name": "test-image-8080-tcp",
                "containerPort": 8080,
                "hostPort": 8080,
                "protocol": "tcp",
                "appProtocol": "http"
            }
        ],
        "essential": true,
        "environment": [],
        "environmentFiles": [],
        "mountPoints": [],
        "volumesFrom": [],
        "ulimits": [],
        "logConfiguration": {
            "logDriver": "awslogs",
            "options": {
                "awslogs-create-group": "true",
                "awslogs-group": "/ecs/test-td",
                "awslogs-region": "us-east-1",
                "awslogs-stream-prefix": "ecs"
            }
        }
    }
  ])
}