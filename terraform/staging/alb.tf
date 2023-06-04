resource "aws_security_group" "lb-sg" {
  name = "${var.environment}-lb-sg"
  description = "Allow HTTP traffic and go-app port traffic"
  vpc_id = aws_vpc.main_vpc.id
  ingress {
    description = "Allow traffic for 8080"
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]

  }
  ingress {
    description = "Allow HTTP traffic"
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]

  }
  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }
}
resource "aws_alb" "lb" {
  name = "${var.environment}-lb"
  internal = false
  load_balancer_type = "application"
  security_groups = [ aws_security_group.lb-sg.id ]
  subnets = [ aws_subnet.public_subnets[0].id,aws_subnet.public_subnets[1].id ]
  ip_address_type = "ipv4"

}
resource "aws_lb_target_group" "app-tg" {
  name = "${var.environment}-lb-tg"
  port = 80
  protocol = "HTTP"
  vpc_id = aws_vpc.main_vpc.id
  target_type = "ip"
  health_check {
    enabled = true
    port = 8080
    protocol = "HTTP"
    path = "/health"
    healthy_threshold = 5
    unhealthy_threshold = 2
    timeout = 5
    interval = 6
    matcher = "200-299"
  }
}

resource "aws_lb_listener" "listener" {
  load_balancer_arn = aws_alb.lb.arn
  port = 80
  protocol = "HTTP"
  default_action {
    type = "forward"
    target_group_arn = aws_lb_target_group.app-tg.arn
  }
}