resource "aws_vpc" "example" {
  # あとから変更できない。
  cidr_block           = "10.0.0.0/16"
  enable_dns_support   = true
  enable_dns_hostnames = true

  tags = {
    Name = "example"
  }
}

# --------------------------------------
# public network settings
# --------------------------------------

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

# ルーティングテーブルを定義する。
# ルートテーブルではVPC内の通信を有効にするため、ローカルルートが自動生成される。Terraformからは制御できない。
resource "aws_route_table" "public" {
  vpc_id = aws_vpc.example.id
}

# インターネットゲートウェイ経由でインターネットへデータを流すためにデフォルトルートを指定する。
resource "aws_route" "public" {
  route_table_id         = aws_route_table.public.id
  gateway_id             = aws_internet_gateway.example.id
  destination_cidr_block = "0.0.0.0/0"
}

# ルートテーブルとサブネットの関連付け。
resource "aws_route_table_association" "public" {
  route_table_id = aws_route_table.public.id
  subnet_id      = aws_subnet.public.id
}

# --------------------------------------
# private network setting
# --------------------------------------

# プライベートサブネット。パブリックサブネットと異なるCIDRブロックを指定する。
resource "aws_subnet" "private" {
  vpc_id     = aws_vpc.example.id
  cidr_block = "10.0.64.0/24"
  # Specify true to indicate that instances launched into the subnet should be assigned a public IP address.
  map_public_ip_on_launch = false
  availability_zone       = "ap-northeast-1a"
}

resource "aws_route_table" "private" {
  vpc_id = aws_vpc.example.id
}

resource "aws_route_table_association" "private" {
  subnet_id      = aws_subnet.private.id
  route_table_id = aws_route_table.private.id
}

# NATゲートウェイ。プライベートネットワークからインターネットにアクセスする用。

# NATゲートウェイを利用するために必要なEIP（Elastic IP Address）
# これを使うと、通常起動するたびに動的に変わるIPを固定できる。
resource "aws_eip" "nat_gateway" {
  vpc = true
  depends_on = [
    aws_internet_gateway.example
  ]
}

resource "aws_nat_gateway" "example" {
  allocation_id = aws_eip.nat_gateway.id
  # NATゲートウェイはパブリックサブネットに配置する。
  subnet_id     = aws_subnet.public.id
  # depends_onを使うと明示的に依存関係を定義できる。
  # インターネットゲートウェイ作成後にEIPやNATゲートウェイを作成することを保証できる。
  depends_on    = [aws_internet_gateway.example]
}

resource "aws_route" "private" {
  route_table_id         = aws_route_table.private.id
  # ネットワークゲートウェイのIDを設定する。aws_route.publicとは異なる。
  nat_gateway_id         = aws_nat_gateway.example.id
  destination_cidr_block = "0.0.0.0/0"
}

