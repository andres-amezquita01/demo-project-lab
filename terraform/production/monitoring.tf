resource "aws_security_group" "monitor-sg" {
  name        = "${var.environment}-monitor-sg"
  description = "Allow ssh, prometheus & grafana inbound traffic"
  vpc_id      = aws_vpc.main_vpc.id
  ingress {
    description = "Allow SSH"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
    # security_groups = [ aws_security_group.lb-sg.id ]

  }
  ingress {
    description = "Allow prometheus for 9090"
    from_port   = 9090
    to_port     = 9090
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
    #security_groups = [ aws_security_group.lb-sg.id ]
  }
ingress {
    description = "Allow grafana for 3000"
    from_port   = 3000
    to_port     = 3000
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
    #security_groups = [ aws_security_group.lb-sg.id ]
  }
  
  egress {
    from_port        = 0
    to_port          = 0
    protocol         = "-1"
    cidr_blocks      = ["0.0.0.0/0"]
    ipv6_cidr_blocks = ["::/0"]
  }
}

resource "aws_instance" "monitor" {
  instance_type = "${var.ec2_type}"
  ami = "${var.ami}"
  subnet_id = aws_subnet.public_subnets[0].id
  key_name = "${var.key_name}"
  associate_public_ip_address = "true"
  security_groups = [ aws_security_group.monitor-sg.id ]
  user_data = "${file("../scripts/install_docker.sh")}"
  tags = {
    "Name" =  "${var.environment}-monitor"
  }
  lifecycle {
    # ignore_changes = [
    #     block_device_mappings
    # ]
    ignore_changes = all
  }
}
