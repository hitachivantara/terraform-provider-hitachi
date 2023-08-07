package vssbstorage

import (
	spmanager "terraform-provider-hitachi/hitachi/storage/vssb/provisioner"
	vssbmodel "terraform-provider-hitachi/hitachi/storage/vssb/provisioner/model"
)

// vssbStorageManager contain information for storage setting
type vssbStorageManager struct {
	storageSetting vssbmodel.StorageDeviceSettings
}

// A private function to construct an newVssbStorageManagerEx
func newVssbStorageManagerEx(storageSetting vssbmodel.StorageDeviceSettings) (*vssbStorageManager, error) {
	psm := &vssbStorageManager{
		storageSetting: vssbmodel.StorageDeviceSettings{
			Username:       storageSetting.Username,
			Password:       storageSetting.Password,
			ClusterAddress: storageSetting.ClusterAddress,
		},
	}
	return psm, nil
}

// NewEx returns a new storage Provisioner
func NewEx(storageSetting vssbmodel.StorageDeviceSettings) (spmanager.VssbStorageManager, error) {
	return newVssbStorageManagerEx(storageSetting)
}

func New(userName, password, clusterAddress string) (spmanager.VssbStorageManager, error) {
	psm := &vssbStorageManager{
		storageSetting: vssbmodel.StorageDeviceSettings{
			Username:       userName,
			Password:       password,
			ClusterAddress: clusterAddress,
		},
	}
	return psm, nil
}
