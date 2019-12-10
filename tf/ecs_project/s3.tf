# 非公開S3バケット
resource "aws_s3_bucket" "private" {
  bucket = "budougumi0617-pragmatic-terraform-on-aws"

  versioning {
    enabled = true
  }

  server_side_encryption_configuration {
    rule {
      apply_server_side_encryption_by_default {
        sse_algorithm = "AES256"
      }
    }
  }

  # バケットにデータが残っていても強制的に削除する。
  force_destroy = true
}

# パブリックアクセスを禁止する
resource "aws_s3_bucket_public_access_block" "private" {
  bucket                  = aws_s3_bucket.private.id
  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

# 公開用S3バケット
resource "aws_s3_bucket" "public" {
  bucket = "public-pragmatic-terraform-on-aws"
  acl    = "public-read" # デフォルトはprivateなので明示的に公開する

  cors_rule {
    allowed_origins = ["https://example.com"]
    allowed_methods = ["GET"]
    allowed_headers = ["*"]
    max_age_seconds = 3000
  }

  # バケットにデータが残っていても強制的に削除する。
  force_destroy = true
}

# ログローテション用バケット。ALB用のアクセスロブ用バケット。
resource "aws_s3_bucket" "alb_log" {
  bucket = "alb-log-pragmatic-terraform-on-aws"

  # ライフサイクルを設定することで、無尽蔵に増え続けるのを防ぐ。
  lifecycle_rule {
    enabled = true

    expiration {
      days = "180"
    }
  }

  # バケットにデータが残っていても強制的に削除する。
  force_destroy = true
}


resource "aws_s3_bucket_policy" "alb_log" {
  bucket = aws_s3_bucket.alb_log.id
  policy = data.aws_iam_policy_document.alb_log.json
}

data "aws_iam_policy_document" "alb_log" {
  statement {
    effect    = "Allow"
    actions   = ["s3:PutObject"]
    resources = ["arn:aws:s3:::${aws_s3_bucket.alb_log.id}/*"]

    principals {
      type = "AWS"
      # リージョンごとに異なるアカウントID
      # https://docs.aws.amazon.com/ja_jp/elasticloadbalancing/latest/classic/enable-access-logs.html#attach-bucket-policy
      # ap-northeast-1	アジアパシフィック (東京)	582318560864
      identifiers = ["582318560864"]
    }
  }
}
