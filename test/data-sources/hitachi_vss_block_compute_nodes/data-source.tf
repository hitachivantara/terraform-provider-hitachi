data "hitachi_vss_block_compute_nodes" "computenodes" {
  vss_block_address = "10.76.47.55"
  #compute_node_name = "ComputeNode-RESTAPI123" // Optional
  
  #compute_node_name = "MongoNode1" // Optional   
}

output "nodeoutput" {
  value = data.hitachi_vss_block_compute_nodes.computenodes
}
