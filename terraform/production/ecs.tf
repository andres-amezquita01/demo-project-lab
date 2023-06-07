data "aws_iam_role" "task_execution_role" {
  name = var.ecs_task_execution_role
}
data "aws_ecr_repository" "registry" {
  name = var.registry
}
resource "aws_security_group" "sg" {
  name        = "${var.environment}-fargate-sg"
  description = "Allow HTTP inbound traffic"
  vpc_id      = aws_vpc.main_vpc.id

  ingress {
    description = "Allow HTTP for 8080 from lb"
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    security_groups = [ aws_security_group.lb-sg.id ]
  }
  ingress {
    description = "Allow HTTP for 8080 from scraping monitor"
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
    # security_groups = [ aws_security_group.monitor-sg.id ]
  }
  ingress {
    description = "Allow HTTP for 9100"
    from_port   = 9100
    to_port     = 9100
    protocol    = "tcp"
    security_groups = [ aws_security_group.monitor-sg.id ]
  }
  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }
}
resource "aws_ecs_cluster" "main_cluster" {
  name = "${var.environment}-cluster"
  setting {
    name  = "containerInsights"
    value = "enabled"
  }
  tags = {
    "environment" = "${var.environment}"
  }
}
resource "aws_ecs_cluster_capacity_providers" "main_cp" {
  cluster_name = aws_ecs_cluster.main_cluster.name
  capacity_providers = [ "FARGATE" ]
}
resource "aws_cloudwatch_log_group" "lg-service" {
  name = "/ecs/${var.environment}-td"
}
resource "aws_ecs_task_definition" "task" {
  family = "${var.environment}-td"
  task_role_arn = data.aws_iam_role.task_execution_role.arn
  execution_role_arn = data.aws_iam_role.task_execution_role.arn
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
        "name": "${var.environment}-image",
        "image": "${data.aws_ecr_repository.registry.repository_url}:${var.image_tag}",
        "cpu": 512,
        "portMappings": [
            {
                "name": "${var.environment}-image-8080-tcp",
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
                "awslogs-group": "/ecs/${var.environment}-td",
                "awslogs-region": "us-east-1",
                "awslogs-stream-prefix": "ecs"
            }        
        }
        # "healthCheck": {
        #       "retries": 3,
        #       "command": [
        #           "CMD-SHELL",
        #           "curl -f http://localhost:8080/ || exit 1"
        #       ],
        #       "timeout": 5,
        #       "interval": 30,
        #       "startPeriod": null
        # }
    }
  ])
  depends_on = [ aws_cloudwatch_log_group.lg-service ]
}
resource "aws_ecs_service" "service" {
  name = "${var.environment}-service"
  cluster = aws_ecs_cluster.main_cluster.id
  task_definition = aws_ecs_task_definition.task.arn
  desired_count = 1
  launch_type = "FARGATE"
  network_configuration {
    subnets = [aws_subnet.public_subnets[0].id,aws_subnet.public_subnets[1].id]
    security_groups = [aws_security_group.sg.id]
    assign_public_ip = true
  }
  load_balancer {
    target_group_arn = aws_lb_target_group.app-tg.arn
    container_name = "${var.environment}-image"
    container_port = 8080
  }
}
