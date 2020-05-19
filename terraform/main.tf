// variable "do_token" {}

provider "digitalocean" {
  // token = var.do_token
}

terraform {
  backend "s3" {
    endpoint                    = "sfo2.digitaloceanspaces.com"
    region                      = "us-west-1"
    bucket                      = "terraform-state-storage"
    key                         = "terraform.tfstate"
    skip_requesting_account_id  = true
    skip_credentials_validation = true
    skip_get_ec2_platforms      = true
    skip_metadata_api_check     = true
  }
}

data "digitalocean_kubernetes_cluster" "default_prd" {
  name = "k8s-default-cluster-prd"
}

provider "kubernetes" {
  load_config_file = false
  host             = data.digitalocean_kubernetes_cluster.default_prd.endpoint
  token            = data.digitalocean_kubernetes_cluster.default_prd.kube_config[0].token
  cluster_ca_certificate = base64decode(
    data.digitalocean_kubernetes_cluster.default_prd.kube_config[0].cluster_ca_certificate
  )
}
