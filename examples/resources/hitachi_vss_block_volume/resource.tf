resource "hitachi_vss_block_volume" "volumecreate" {
  vss_block_address = "172.25.58.151"
  name              = "test-volume-newCol"
  capacity_gb       = 1.9
  storage_pool      = "SP01"
  compute_nodes     = []
  nick_name         = "Vss_volume_changesnk"


}

output "volumecreateData" {
  value = resource.hitachi_vss_block_volume.volumecreate
}
