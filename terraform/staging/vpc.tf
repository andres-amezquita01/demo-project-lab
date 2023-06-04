resource "aws_vpc" "main_vpc" {
  cidr_block = var.vpc_block
  tags = {
    Name = "${var.environment}-vpc"
  }
}
resource "aws_subnet" "public_subnets" {
 count      = length(var.public_subnet_cidrs)
 vpc_id     = aws_vpc.main_vpc.id
 cidr_block = element(var.public_subnet_cidrs, count.index)
 availability_zone = element(var.availability_zones,count.index)
 tags = {
   Name = "${var.environment} public subnet ${count.index + 1}"
 }
} 
resource "aws_subnet" "private_subnets" {
 count      = length(var.private_subnet_cidrs)
 vpc_id     = aws_vpc.main_vpc.id
 cidr_block = element(var.private_subnet_cidrs, count.index)
 availability_zone = element(var.availability_zones,count.index)
 tags = {
   Name = "${var.environment} private subnet ${count.index + 1}"
 }
}
resource "aws_internet_gateway" "main_gw" {
  vpc_id = aws_vpc.main_vpc.id
  tags = {
    Name = "${var.environment}-gw"
  } 
}
resource "aws_route_table" "public_rt" {
  vpc_id = aws_vpc.main_vpc.id

  route {
    cidr_block = var.internet_destination_block
    gateway_id = aws_internet_gateway.main_gw.id
  }
  tags = {
    Name = "${var.environment}-public-rt"
  }
}
resource "aws_route_table_association" "staging_public_rt_association" {
  count = length(var.public_subnet_cidrs)
  subnet_id = element(aws_subnet.public_subnets[*].id,count.index)
  route_table_id = aws_route_table.public_rt.id
}
