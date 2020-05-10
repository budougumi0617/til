# 12.2.2 データベースのユーザー名の定義
resource "aws_ssm_parameter" "db_username" {
  name        = "/db/username"
  value       = "root"
  type        = "String"
  description = "データベースの接続ユーザ名"
}

# 暗号化
resource "aws_ssm_parameter" "db_password" {
  name        = "/db/password"
  value       = "uninitialized"
  type        = "SecureString"
  description = "データベースのパスワード"

  # 文字列自体は変更管理の対象にしない。
  lifecycle {
    ignore_changes = [
      value,
    ]
  }
}