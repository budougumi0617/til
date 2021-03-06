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

resource "aws_subnet" "public_0" {
  vpc_id                  = aws_vpc.example.id
  cidr_block              = "10.0.1.0/24"
  availability_zone       = "ap-northeast-1a"
  map_public_ip_on_launch = true
}

resource "aws_subnet" "public_1" {
  vpc_id     = aws_vpc.example.id
  cidr_block = "10.0.2.0/24"
  # public_1 とは違うAZ
  availability_zone       = "ap-northeast-1c"
  map_public_ip_on_launch = true
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
resource "aws_route_table_association" "public_0" {
  subnet_id      = aws_subnet.public_0.id
  route_table_id = aws_route_table.public.id
}

resource "aws_route_table_association" "public_1" {
  subnet_id      = aws_subnet.public_1.id
  route_table_id = aws_route_table.public.id
}


# --------------------------------------
# private network setting
# --------------------------------------

# プライベートサブネット。パブリックサブネットと異なるCIDRブロックを指定する。
resource "aws_subnet" "private_0" {
  vpc_id            = aws_vpc.example.id
  cidr_block        = "10.0.65.0/24"
  availability_zone = "ap-northeast-1a"
  # Specify true to indicate that instances launched into the subnet should be assigned a public IP address.
  map_public_ip_on_launch = false
}

resource "aws_subnet" "private_1" {
  vpc_id                  = aws_vpc.example.id
  cidr_block              = "10.0.66.0/24"
  availability_zone       = "ap-northeast-1c"
  map_public_ip_on_launch = false
}

# ルートテーブルもAZごとに作成する。
resource "aws_route_table" "private_0" {
  vpc_id = aws_vpc.example.id
}

resource "aws_route_table" "private_1" {
  vpc_id = aws_vpc.example.id
}

resource "aws_route_table_association" "private_0" {
  subnet_id      = aws_subnet.private_0.id
  route_table_id = aws_route_table.private_0.id
}

resource "aws_route_table_association" "private_1" {
  subnet_id      = aws_subnet.private_1.id
  route_table_id = aws_route_table.private_1.id
}

# NATゲートウェイ。プライベートネットワークからインターネットにアクセスする用。

# NATゲートウェイを利用するために必要なEIP（Elastic IP Address）
# これを使うと、通常起動するたびに動的に変わるIPを固定できる。
resource "aws_eip" "nat_gateway_0" {
  vpc        = true
  depends_on = [aws_internet_gateway.example]
}

resource "aws_eip" "nat_gateway_1" {
  vpc        = true
  depends_on = [aws_internet_gateway.example]
}

# NATゲートウェイも冗長化しておく
resource "aws_nat_gateway" "nat_gateway_0" {
  allocation_id = aws_eip.nat_gateway_0.id
  subnet_id     = aws_subnet.public_0.id
  depends_on    = [aws_internet_gateway.example]
}

resource "aws_nat_gateway" "nat_gateway_1" {
  allocation_id = aws_eip.nat_gateway_1.id
  # NATゲートウェイはパブリックサブネットに配置する。
  subnet_id     = aws_subnet.public_1.id
  # depends_onを使うと明示的に依存関係を定義できる。
  # インターネットゲートウェイ作成後にEIPやNATゲートウェイを作成することを保証できる。
  depends_on    = [aws_internet_gateway.example]
}

resource "aws_route" "private_0" {
  route_table_id         = aws_route_table.private_0.id
  # ネットワークゲートウェイのIDを設定する。aws_route.publicとは異なる。
  nat_gateway_id         = aws_nat_gateway.nat_gateway_0.id
  destination_cidr_block = "0.0.0.0/0"
}

resource "aws_route" "private_1" {
  route_table_id         = aws_route_table.private_1.id
  nat_gateway_id         = aws_nat_gateway.nat_gateway_1.id
  destination_cidr_block = "0.0.0.0/0"
}

# --------------------------------------
# Route53 settings
# --------------------------------------
// 外部で作成されたホストゾーンを参照する場合はdataを使う
data "aws_route53_zone" "budougumi0617" {
  name = "budougumi0617.net"
}

// 新規作成する場合は次のように作る
resource "aws_route53_zone" "test_budougumi0617" {
  name = "test.budougumi0617.net"
}


resource "aws_route53_record" "example" {
  zone_id = data.aws_route53_zone.budougumi0617.zone_id
  name    = data.aws_route53_zone.budougumi0617.name
  type    = "A" // AWS独自拡張のALIASレコード。ドメイン名→IPアドレスという流れで名前解決されるので速い。

  alias {
    name                   = aws_lb.example.dns_name
    zone_id                = aws_lb.example.zone_id
    evaluate_target_health = true
  }
}

# --------------------------------------
# SSL settings
# --------------------------------------
resource "aws_acm_certificate" "example" {
  domain_name               = data.aws_route53_zone.budougumi0617.name
  subject_alternative_names = []
  validation_method         = "DNS" // 自動更新できる。

  lifecycle {
    create_before_destroy = true
  }
}


resource "aws_route53_record" "example_certificate" {
  name    = aws_acm_certificate.example.domain_validation_options[0].resource_record_name
  type    = aws_acm_certificate.example.domain_validation_options[0].resource_record_type
  records = [aws_acm_certificate.example.domain_validation_options[0].resource_record_value]
  zone_id = data.aws_route53_zone.budougumi0617.id
  ttl     = 60
}

// apply時にSSL証明書の検証が完了するまで待機する設定。
resource "aws_acm_certificate_validation" "example" {
  certificate_arn         = aws_acm_certificate.example.arn
  validation_record_fqdns = [aws_route53_record.example_certificate.fqdn]
}

resource "aws_lb_listener" "https" {
  load_balancer_arn = aws_lb.example.arn
  port              = "443"
  protocol          = "HTTPS"
  certificate_arn   = aws_acm_certificate.example.arn
  ssl_policy        = "ELBSecurityPolicy-2016-08"

  default_action {
    type = "fixed-response"

    fixed_response {
      content_type = "text/plain"
      message_body = "これは『HTTPS』です"
      status_code  = "200"
    }
  }
}

// HTTPからHTTPSにするリダイレクトルール
resource "aws_lb_listener" "redirect_http_to_https" {
  load_balancer_arn = aws_lb.example.arn
  port              = "8080"
  protocol          = "HTTP"

  default_action {
    type = "redirect"

    redirect {
      port        = "443"
      protocol    = "HTTPS"
      status_code = "HTTP_301"
    }
  }
}

resource "aws_lb_target_group" "example" {
  name                 = "example"
  vpc_id               = aws_vpc.example.id
  target_type          = "ip" // ECS FargateではIPアドレスによるルーティングが必要。
  port                 = 80
  protocol             = "HTTP"
  deregistration_delay = 300

  health_check {
    path                = "/"
    healthy_threshold   = 5
    unhealthy_threshold = 2
    timeout             = 5
    interval            = 30
    matcher             = 200
    port                = "traffic-port"
    protocol            = "HTTP"
  }

  depends_on = [aws_lb.example] // ECSサービスが完成してからという依存関係
}

resource "aws_lb_listener_rule" "example" {
  listener_arn = aws_lb_listener.https.arn
  priority     = 100

  action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.example.arn
  }

  condition {
    field  = "path-pattern"
    values = ["/*"]
  }
}
