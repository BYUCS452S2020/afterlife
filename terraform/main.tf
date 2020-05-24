provider "digitalocean" {}

terraform {
  backend "s3" {
    bucket                      = "terraform-state-storage"
    endpoint                    = "sfo2.digitaloceanspaces.com"
    region                      = "us-west-1"
    key                         = "afterlife.tfstate"
    skip_requesting_account_id  = true
    skip_credentials_validation = true
    skip_get_ec2_platforms      = true
    skip_metadata_api_check     = true
  }
}

data "digitalocean_kubernetes_cluster" "default_prd" {
  name = "k8s-default-cluster-prd"
}

data "digitalocean_domain" "default" {
  name = "danielrandall.dev"
}

//data "digitalocean_loadbalancer" "default" {
//  name = "lb-k8s-cluster-default-prd"
//}

provider "kubernetes" {
  load_config_file = false
  host             = data.digitalocean_kubernetes_cluster.default_prd.endpoint
  token            = data.digitalocean_kubernetes_cluster.default_prd.kube_config[0].token
  cluster_ca_certificate = base64decode(
    data.digitalocean_kubernetes_cluster.default_prd.kube_config[0].cluster_ca_certificate
  )
}

//resource "digitalocean_certificate" "afterlife" {
//  name    = "cert-afterlife"
//  type    = "lets_encrypt"
//  domains = ["hello.danielrandall.dev"]
//}

//resource "digitalocean_record" "db" {
//  domain = data.digitalocean_domain.default.name
//  type   = "A"
//  name   = "db.afterlife"
//  value  = data.digitalocean_loadbalancer.default.ip
//}
//
//module "db" {
//  source = "github.com/dannyrandall/cloud//kubernetes-statefulset"
//
//  // required
//  name                 = "afterlife-db"
//  image                = "postgres"
//  image_version        = "12"
//  container_port       = 5432
//  repo_url             = "https://github.com/dannyrandall/afterlife"
//  storage_mount_path   = "/var/lib/postgresql/data"
//  storage_request_size = "2Gi"
//  storage_provisioner  = "dobs.csi.digitalocean.com"
//
//  // optional
//  public_urls = ["db.afterlife.danielrandall.dev"]
//  container_env = {
//    "POSTGRES_USER"     = "test"
//    "POSTGRES_PASSWORD" = "testPassword"
//    "POSTGRES_DB"       = "default"
//    "PGDATA"            = "/var/lib/postgresql/data/pgdata"
//  }
//  container_args = []
//}

//resource "digitalocean_record" "hello_world" {
//  domain = data.digitalocean_domain.default.name
//  type   = "A"
//  name   = "hello"
//  value  = data.digitalocean_loadbalancer.default.ip
//}
//
//module "hello_world" {
//  source = "github.com/dannyrandall/cloud//kubernetes-deployment"
//
//  // required
//  name           = "hellow-world"
//  image          = "nginxdemos/hello"
//  image_version  = "latest"
//  container_port = 80
//  repo_url       = "https://github.com/dannyrandall/afterlife"
//
//  // optional
//  public_urls    = ["hello.danielrandall.dev"]
//  health_check   = false
//  container_env  = {}
//  container_args = []
//}
