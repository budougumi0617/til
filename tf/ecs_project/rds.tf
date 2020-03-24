# 13.1.1 DBパラメータグループ
resource "aws_db_parameter_group" "example" {
  name   = "example"
  family = "mysql5.7"

  # 設定のパラメータ名と値のセット
  parameter {
    name  = "character_set_database"
    value = "utf8mb4"
  }

  parameter {
    name  = "character_set_server"
    value = "utf8mb4"
  }
}

# 13.1.2 DBオプショングループ
# データベースエンジンのオプション機能を追加できる。
resource "aws_db_option_group" "example" {
  name                 = "example"
  engine_name          = "mysql"
  major_engine_version = "5.7"

  option {
    # 監査プラグインを有効にしている。
    option_name = "MARIADB_AUDIT_PLUGIN"
  }
}

# 13.1.3 DBサブネットグループ
resource "aws_db_subnet_group" "example" {
  name       = "example"
  subnet_ids = [aws_subnet.private_0.id, aws_subnet.private_1.id]
}

# セキュリティーグループの定義
module "mysql_sg" {
  source      = "./security_group"
  name        = "mysql-sg"
  vpc_id      = aws_vpc.example.id
  port        = 3306
  cidr_blocks = [aws_vpc.example.cidr_block]
}

# 13.1.4 DBインスタンス
resource "aws_db_instance" "example" {
  identifier        = "example"
  engine            = "mysql"
  engine_version    = "5.7.23"
  instance_class    = "db.t3.small"
  allocated_storage = 20
  storage_type      = "gp2"
  storage_encrypted = true
  kms_key_id        = aws_kms_key.example.arn
  username          = "admin"
  # 本来は後で設定を変更しておくこと
  # aws rds modify-db-instance --db-instance-identifier 'example' --master-user-password　'NewMasterPassword'
  password = "VeryStrongPassword!"
  # aws_db_subnet_groupで異なるAZのサブネットを指定できているならば有効にできる。
  multi_az = true
  # VPC外からのアクセスを遮断する設定。
  publicly_accessible        = false
  backup_window              = "09:10-09:40"
  backup_retention_period    = 30
  maintenance_window         = "mon:10:10-mon:10:40"
  auto_minor_version_upgrade = false
  # !!!!!!! 本当はtrueにしておくべき設定。練習用の節約のため、削除を許可しておく。
  deletion_protection = false
  # !!!!!!! 本当はfalseにしておくべき設定。スナップショットの作成をスキップする
  skip_final_snapshot    = true
  port                   = 3306
  apply_immediately      = false
  vpc_security_group_ids = [module.mysql_sg.security_group_id]
  parameter_group_name   = aws_db_parameter_group.example.name
  option_group_name      = aws_db_option_group.example.name
  db_subnet_group_name   = aws_db_subnet_group.example.name

  lifecycle {
    ignore_changes = [password]
  }
}