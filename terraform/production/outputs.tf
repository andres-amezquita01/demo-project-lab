output "production_lb" {
  value = aws_alb.lb.dns_name
}