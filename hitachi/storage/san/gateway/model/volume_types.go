package sanstorage

// LogicalUnit .
type LogicalUnit struct {
	// these are returned from gateway
	LdevID             int    `json:"ldevId"`
	ClprID             int    `json:"clprId"`
	EmulationType      string `json:"emulationType"`
	ByteFormatCapacity string `json:"byteFormatCapacity"`
	BlockCapacity      uint64 `json:"blockCapacity"`
	NumOfPorts         int    `json:"numOfPorts"`
	Ports              []struct {
		PortID          string `json:"portId"`
		HostGroupNumber int    `json:"hostGroupNumber"`
		HostGroupName   string `json:"hostGroupName"`
		Lun             int    `json:"lun"`
	} `json:"ports"`
	Attributes                []string `json:"attributes"`
	Label                     string   `json:"label"`
	Status                    string   `json:"status"`
	MpBladeID                 int      `json:"mpBladeId"`
	Ssid                      string   `json:"ssid"`
	PoolID                    int      `json:"poolId"`
	NumOfUsedBlock            uint64   `json:"numOfUsedBlock"`
	IsFullAllocationEnabled   bool     `json:"isFullAllocationEnabled"`
	ResourceGroupID           int      `json:"resourceGroupId"`
	DataReductionStatus       string   `json:"dataReductionStatus"`
	DataReductionMode         string   `json:"dataReductionMode"`
	DataReductionProcessMode  string   `json:"dataReductionProcessMode"`
	DataReductionProgressRate int      `json:"dataReductionProgressRate"`
	IsAluaEnabled             bool     `json:"isAluaEnabled"`
	NaaID                     string   `json:"naaId"`
	ParityGroupId             []string `json:"parityGroupIds"`

	// below will be populated by provisioner
	TotalCapacityInMB uint64 `json:"totalCapacityInMB"`
	FreeCapacityInMB  uint64 `json:"freeCapacityInMB"`
	UsedCapacityInMB  uint64 `json:"usedCapacityInMB"`
}

// LogicalUnits .
type LogicalUnits struct {
	ListOfLun []LogicalUnit `json:"data,omitempty"`
}

// CreateLunRequest .
type CreateLunRequest struct {
	LdevID            *int    `json:"ldevId,omitempty"`
	Name              *string `json:"name,omitempty"`
	PoolID            *int    `json:"poolId"`
	ParityGroupID     *string `json:"parityGroupId,omitempty"`
	CapacityInGB      uint64  `json:"capacityInGB"`
	DataReductionMode *string `json:"dataReductionMode,omitempty"`
}

// CreateLunRequestGwy .
type CreateLunRequestGwy struct {
	LdevID             *int    `json:"ldevId,omitempty"`
	PoolID             *int    `json:"poolId"`
	ParityGroupID      *string `json:"parityGroupId,omitempty"`
	ByteFormatCapacity string  `json:"byteFormatCapacity"`
	DataReductionMode  *string `json:"dataReductionMode,omitempty"`
}

// UpdateLunRequestGwy .
type UpdateLunRequestGwy struct {
	Label             *string `json:"label,omitempty"`
	DataReductionMode *string `json:"dataReductionMode,omitempty"`
}

// UpdateLunRequest .
type UpdateLunRequest struct {
	LdevID            *int    `json:"ldevId,omitempty"`
	Name              *string `json:"name,omitempty"`
	CapacityInGB      *uint64 `json:"capacityInGB"`
	DataReductionMode *string `json:"dataReductionMode,omitempty"`
}

// ExpandLunRequestGwy .
type ExpandLunRequestGwy struct {
	Parameters ExpandLunParameters `json:"parameters,omitempty"`
}

// ExpandLunParameters .
type ExpandLunParameters struct {
	AdditionalByteFormatCapacity string `json:"additionalByteFormatCapacity,omitempty"`
	// AdditionalBlockCapacity      int    `json:"additionalBlockCapacity,omitempty"`
	// EnhancedExpansion            bool   `json:"enhancedExpansion,omitempty"`
}
