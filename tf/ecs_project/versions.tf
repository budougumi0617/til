terraform {
  required_version = ">= 0.12"
}

provider "aws" {
  version = "~> 2.41"
  region  = "ap-northeast-1"
}