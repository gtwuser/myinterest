terraform {
  backend "s3" {
    bucket = "testbucketeer"
    key = "terraform/tfstate.tfstate"
    region = "us-east-1"
  }
}