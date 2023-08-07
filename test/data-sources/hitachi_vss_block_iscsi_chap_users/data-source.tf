
/*
data "hitachi_vss_block_iscsi_chap_users" "chap_user_by_id" {
   vss_block_address   = ""
   target_chap_user = "a79c1a1d-2719-4e07-b800-faf9de73d0ae"
}

output "id_output" {
  value = data.hitachi_vss_block_iscsi_chap_users.chap_user_by_id
}

*/

data "hitachi_vss_block_iscsi_chap_users" "chap_user_by_name" {
  vss_block_address = "10.76.47.55"
  #target_chap_user = "rahul70"
}

output "name_output" {
  value = data.hitachi_vss_block_iscsi_chap_users.chap_user_by_name
}

data "hitachi_vss_block_iscsi_chap_users" "my_chap_users" {
  vss_block_address = "10.76.47.55"
}

output "my_iscsi_chap_users_output" {
  value = data.hitachi_vss_block_iscsi_chap_users.my_chap_users
}