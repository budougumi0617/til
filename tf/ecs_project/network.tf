resource "aws_vpc" "example" {
  # あとから変更できない。
  cidr_block           = "10.0.0.0/16"
  enable_dns_support   = true
  enable_dns_hostnames = true

  tags = {
    Name = "example"
  }
}

resource "aws_subnet" "public" {
  vpc_id     = aws_vpc.example.id
  cidr_block = "10.0.0.0/24"
  # サブネットで起動したインスタンスにパブリックIPアドレスを自動的に割り当ててくれる。
  map_public_ip_on_launch = true
  availability_zone       = "ap-northeast-1a"
}

# VPC-インターネット間で通信をできるようにする。
resource "aws_internet_gateway" "example" {
  vpc_id = aws_vpc.example.id
}