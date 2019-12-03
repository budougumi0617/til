provider "aws" {
  region = "ap-northeast-1"
}

// 変数宣言。実行時に-var 'example_instance_type=t3.nano'というように変更できる。
variable "example_instance_type" {
  default = "t3.micro"
}

// コマンド実行時に上書きできない。
locals {
  example_instance_type = "t3.micro"
}

// 外部データを参照する。filterで検索条件を指定して最新のAMIを取得しているだけ。
data "aws_ami" "recent_amazon_linux_2" {
  most_recent = true
  owners      = ["amazon"]

  filter {
    name   = "name"
    values = ["amzn2-ami-hvm-2.0.????????-x86_64-gp2"]
  }

  filter {
    name   = "state"
    values = ["available"]
  }
}

// セキュリティグループの定義
resource "aws_security_group" "example_ec2" {
  name = "example-ec2"

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
  
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_instance" "example" {
  ami           = data.aws_ami.recent_amazon_linux_2.image_id
  instance_type = var.example_instance_type
  // 条件分岐も利用できる
  // instance_type = var.env == "prod" ? "m5.large" : "t3.micro"
  // instance_type = locals.example_instance_type
  vpc_security_group_ids = [aws_security_group.example_ec2.id]


  tags = {
    Name = "example"
  }
  user_data = <<EOF
    #!/bin/bash
    yum install -y httpd
    systemctl start httpd.service
EOF
}

/**
 * apply時に実行結果の最後に結果が表示される。
 *
 * $terraform apply
 * aws_instance.example: Refreshing state... [id=i-057617a019734946a]

 * Apply complete! Resources: 0 added, 0 changed, 0 destroyed.

 * Outputs:

 * example_instance_id = i-057617a019734946a
*/
output "example_instance_id" {
  value = aws_instance.example.id
}

output "exmple_public_dns" {
  value = aws_instance.example.public_dns
}
