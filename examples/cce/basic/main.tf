resource "huaweicloud_vpc" "myvpc" {
  name = var.vpc_name
  cidr = var.vpc_cidr
}

resource "huaweicloud_vpc_subnet" "mysubnet" {
  name       = var.subnet_name
  cidr       = var.subnet_cidr
  gateway_ip = var.subnet_gateway

  //dns is required for cce node installing
  primary_dns   = var.primary_dns
  secondary_dns = var.secondary_dns
  vpc_id        = huaweicloud_vpc.myvpc.id
}

resource "huaweicloud_vpc_eip" "myeip" {
  publicip {
    type = "5_bgp"
  }
  bandwidth {
    name        = var.bandwidth_name
    size        = 8
    share_type  = "PER"
    charge_mode = "traffic"
  }
}

data "huaweicloud_availability_zones" "myaz" {}

resource "huaweicloud_compute_keypair" "mykeypair" {
  name = var.key_pair_name
}
resource "huaweicloud_cce_cluster" "mycce" {
  name                   = var.cce_cluster_name
  flavor_id              = var.cce_cluster_flavor
  vpc_id                 = huaweicloud_vpc.myvpc.id
  subnet_id              = huaweicloud_vpc_subnet.mysubnet.id
  container_network_type = "overlay_l2"
  eip                    = huaweicloud_vpc_eip.myeip.address
}

resource "huaweicloud_cce_node" "mynode" {
  cluster_id        = huaweicloud_cce_cluster.mycce.id
  name              = var.node_name
  flavor_id         = var.node_flavor
  availability_zone = data.huaweicloud_availability_zones.myaz.names[0]
  key_pair          = huaweicloud_compute_keypair.mykeypair.name

  root_volume {
    size       = var.root_volume_size
    volumetype = var.root_volume_type
  }
  data_volumes {
    size       = var.data_volume_size
    volumetype = var.data_volume_type
  }
}

data "huaweicloud_images_image" "myimage" {
  name        = var.image_name
  most_recent = true
}

resource "huaweicloud_compute_instance" "myecs" {
  name                        = var.ecs_name
  image_id                    = data.huaweicloud_images_image.myimage.id
  flavor_id                   = var.ecs_flavor
  availability_zone           = data.huaweicloud_availability_zones.myaz.names[0]
  key_pair                    = huaweicloud_compute_keypair.mykeypair.name
  delete_disks_on_termination = true

  system_disk_type = var.root_volume_type
  system_disk_size = var.root_volume_size

  data_disks {
    type = var.data_volume_type
    size = var.data_volume_size
  }

  network {
    uuid = huaweicloud_vpc_subnet.mysubnet.id
  }
}

resource "huaweicloud_cce_node_attach" "test" {
  cluster_id = huaweicloud_cce_cluster.mycce.id
  server_id  = huaweicloud_compute_instance.myecs.id
  key_pair   = huaweicloud_compute_keypair.mykeypair.name
  os         = var.os
}
