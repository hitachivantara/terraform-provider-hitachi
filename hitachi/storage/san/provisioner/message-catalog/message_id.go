package messagecatalog

// MessageID .
type MessageID uint64

const sidx = 2000

const (
	// Default1 .
	Default1 MessageID = iota + sidx
	// STORAGE SYSTEM
	INFO_GET_STORAGE_SYSTEM_BEGIN
	INFO_GET_STORAGE_SYSTEM_END
	ERR_GET_STORAGE_SYSTEM_FAILED

	// VOLUME
	INFO_GET_LUN_BEGIN
	INFO_GET_LUN_END
	ERR_GET_LUN_FAILED
	ERR_GET_ALL_LUN_FAILED
	ERR_CREATE_LUN_FAILED
	INFO_CREATE_LUN_BEGIN
	INFO_CREATE_LUN_END
	ERR_EXPAND_LUN_FAILED
	INFO_EXPAND_LUN_BEGIN
	INFO_EXPAND_LUN_END
	ERR_DELETE_LUN_FAILED
	INFO_DELETE_LUN_BEGIN
	INFO_DELETE_LUN_END
	INFO_GET_LUN_RANGE_BEGIN
	INFO_GET_LUN_RANGE_END
	ERR_UPDATE_LUN_FAILED
	INFO_UPDATE_LUN_BEGIN
	INFO_UPDATE_LUN_END

	// HOSTGROUP
	INFO_GET_HOSTGROUP_BEGIN
	INFO_GET_HOSTGROUP_END
	ERR_GET_HOSTGROUP_FAILED
	INFO_GET_ALL_HOSTGROUP_BEGIN
	INFO_GET_ALL_HOSTGROUP_END
	ERR_GET_ALL_HOSTGROUP_FAILED
	INFO_CREATE_HOSTGROUP_BEGIN
	INFO_CREATE_HOSTGROUP_END
	ERR_CREATE_HOSTGROUP_FAILED
	ERR_ADD_LDEV_TO_HOSTGROUP_FAILED
	ERR_ADD_WWN_TO_HOSTGROUP_FAILED
	ERR_SET_NICKNAME_TO_WWN_FAILED
	INFO_DELETE_HOSTGROUP_BEGIN
	INFO_DELETE_HOSTGROUP_END
	ERR_DELETE_HOSTGROUP_FAILED
	INFO_SET_MODE_OPTION_HOSTGROUP_BEGIN
	INFO_SET_MODE_OPTION_HOSTGROUP_END
	ERR_SET_MODE_OPTION_HOSTGROUP_FAILED
	INFO_DELETE_WWN_BEGIN
	INFO_DELETE_WWN_END
	ERR_DELETE_WWN_FAILED
	INFO_SET_NICKNAME_HOSTGROUP_BEGIN
	INFO_SET_NICKNAME_HOSTGROUP_END
	INFO_ADD_WWN_TO_HOSTGROUP_BEGIN
	INFO_ADD_WWN_TO_HOSTGROUP_END
	INFO_REMOVE_LUN_HOSTGROUP_BEGIN
	INFO_REMOVE_LUN_HOSTGROUP_END
	ERR_REMOVE_LUN_HOSTGROUP_FAILED
	INFO_ADD_LDEV_TO_HOSTGROUP_BEGIN
	INFO_ADD_LDEV_TO_HOSTGROUP_END

	// ISCSI TARGET
	INFO_GET_ISCSITARGET_BEGIN
	INFO_GET_ISCSITARGET_END
	ERR_GET_ISCSITARGET_FAILED
	INFO_GET_ALL_ISCSITARGET_BEGIN
	INFO_GET_ALL_ISCSITARGET_END
	ERR_GET_ALL_ISCSITARGET_FAILED
	INFO_SET_MODE_OPTION_ISCSITARGET_BEGIN
	INFO_SET_MODE_OPTION_ISCSITARGET_END
	ERR_SET_MODE_OPTION_ISCSITARGET_FAILED
	INFO_CREATE_ISCSITARGET_BEGIN
	INFO_CREATE_ISCSITARGET_END
	ERR_CREATE_ISCSITARGET_FAILED
	INFO_ADD_LDEV_TO_ISCSITARGET_BEGIN
	INFO_ADD_LDEV_TO_ISCSITARGET_END
	ERR_ADD_LDEV_TO_ISCSITARGET_FAILED
	INFO_ADD_IQN_NAME_TO_ISCSITARGET_BEGIN
	INFO_ADD_IQN_NAME_TO_ISCSITARGET_END
	ERR_ADD_IQN_NAME_TO_ISCSITARGET_FAILED
	INFO_ADD_NICKNAME_TO_ISCSITARGET_BEGIN
	INFO_ADD_NICKNAME_TO_ISCSITARGET_END
	ERR_ADD_NICKNAME_TO_ISCSITARGET_FAILED
	INFO_DELETE_IQN_NAME_BEGIN
	INFO_DELETE_IQN_NAME_END
	ERR_DELETE_IQN_NAME_FAILED
	INFO_DELETE_ISCSITARGET_BEGIN
	INFO_DELETE_ISCSITARGET_END
	ERR_DELETE_ISCSITARGET_FAILED

	// ISCSI TARGET CHAP USER
	INFO_GET_ISCSITARGET_CHAPUSER_BEGIN
	INFO_GET_ISCSITARGET_CHAPUSER_END
	ERR_GET_ISCSITARGET_CHAPUSER_FAILED
	INFO_GET_ISCSITARGET_CHAPUSERS_BEGIN
	INFO_GET_ISCSITARGET_CHAPUSERS_END
	ERR_GET_ISCSITARGET_CHAPUSERS_FAILED
	INFO_CREATE_ISCSITARGET_CHAPUSER_BEGIN
	INFO_CREATE_ISCSITARGET_CHAPUSER_END
	ERR_CREATE_ISCSITARGET_CHAPUSER_FAILED
	INFO_SET_ISCSITARGET_CHAPUSERNAME_BEGIN
	INFO_SET_ISCSITARGET_CHAPUSERNAME_END
	ERR_SET_ISCSITARGET_CHAPUSERNAME_FAILED
	INFO_SET_ISCSITARGET_CHAPUSERSECRET_BEGIN
	INFO_SET_ISCSITARGET_CHAPUSERSECRET_END
	ERR_SET_ISCSITARGET_CHAPUSERSECRET_FAILED
	INFO_DELETE_ISCSITARGET_CHAPUSER_BEGIN
	INFO_DELETE_ISCSITARGET_CHAPUSER_END
	ERR_DELETE_ISCSITARGET_CHAPUSER_FAILED
	INFO_CHANGE_ISCSITARGET_CHAPUSER_BEGIN
	INFO_CHANGE_ISCSITARGET_CHAPUSER_END
	ERR_CHANGE_ISCSITARGET_CHAPUSER_FAILED

	// STORAGE PORTS
	INFO_GET_STORAGE_PORTS_BEGIN
	INFO_GET_STORAGE_PORTS_END
	ERR_GET_STORAGE_PORTS_FAILED
	INFO_GET_STORAGE_PORTS_PORTID_BEGIN
	INFO_GET_STORAGE_PORTS_PORTID_END
	ERR_GET_STORAGE_PORTS_PORTID_FAILED

	// DYNAMIC POOL
	INFO_GET_DYNAMIC_POOLS_BEGIN
	INFO_GET_DYNAMIC_POOLS_END
	ERR_GET_DYNAMIC_POOLS_FAILED
	INFO_GET_DYNAMIC_POOL_ID_BEGIN
	INFO_GET_DYNAMIC_POOL_ID_END
	ERR_GET_DYNAMIC_POOL_ID_FAILED

	// PARITY GROUP
	INFO_GET_PARITY_GROUP_BEGIN
	INFO_GET_PARITY_GROUP_END
	ERR_GET_PARITY_GROUP_FAILED
)

var enumStrings = map[interface{}]string{
	Default1: "Default1",
	// STORAGE SYSTEM
	INFO_GET_STORAGE_SYSTEM_BEGIN: "INFO_GET_STORAGE_SYSTEM_BEGIN",
	INFO_GET_STORAGE_SYSTEM_END:   "INFO_GET_STORAGE_SYSTEM_END",
	ERR_GET_STORAGE_SYSTEM_FAILED: "ERR_GET_STORAGE_SYSTEM_FAILED",

	// VOLUME
	INFO_GET_LUN_BEGIN:       "INFO_GET_LUN_BEGIN",
	INFO_GET_LUN_END:         "INFO_GET_LUN_END",
	ERR_GET_LUN_FAILED:       "ERR_GET_LUN_FAILED",
	ERR_GET_ALL_LUN_FAILED:   "ERR_GET_ALL_LUN_FAILED",
	ERR_CREATE_LUN_FAILED:    "ERR_CREATE_LUN_FAILED",
	INFO_CREATE_LUN_BEGIN:    "INFO_CREATE_LUN_BEGIN",
	INFO_CREATE_LUN_END:      "INFO_CREATE_LUN_END",
	ERR_EXPAND_LUN_FAILED:    "ERR_EXPAND_LUN_FAILED",
	INFO_EXPAND_LUN_BEGIN:    "INFO_EXPAND_LUN_BEGIN",
	INFO_EXPAND_LUN_END:      "INFO_EXPAND_LUN_END",
	ERR_DELETE_LUN_FAILED:    "ERR_DELETE_LUN_FAILED",
	INFO_DELETE_LUN_BEGIN:    "INFO_DELETE_LUN_BEGIN",
	INFO_DELETE_LUN_END:      "INFO_DELETE_LUN_END",
	INFO_GET_LUN_RANGE_BEGIN: "INFO_GET_LUN_RANGE_BEGIN",
	INFO_GET_LUN_RANGE_END:   "INFO_GET_LUN_RANGE_END",
	ERR_UPDATE_LUN_FAILED:    "ERR_UPDATE_LUN_FAILED",
	INFO_UPDATE_LUN_BEGIN:    "INFO_UPDATE_LUN_BEGIN",
	INFO_UPDATE_LUN_END:      "INFO_UPDATE_LUN_END",

	// HOSTGROUP
	INFO_GET_HOSTGROUP_BEGIN:             "INFO_GET_HOSTGROUP_BEGIN",
	INFO_GET_HOSTGROUP_END:               "INFO_GET_HOSTGROUP_END",
	ERR_GET_HOSTGROUP_FAILED:             "ERR_GET_HOSTGROUP_FAILED",
	INFO_GET_ALL_HOSTGROUP_BEGIN:         "INFO_GET_ALL_HOSTGROUP_BEGIN",
	INFO_GET_ALL_HOSTGROUP_END:           "INFO_GET_ALL_HOSTGROUP_END",
	ERR_GET_ALL_HOSTGROUP_FAILED:         "ERR_GET_ALL_HOSTGROUP_FAILED",
	INFO_CREATE_HOSTGROUP_BEGIN:          "INFO_CREATE_HOSTGROUP_BEGIN",
	INFO_CREATE_HOSTGROUP_END:            "INFO_CREATE_HOSTGROUP_END",
	ERR_CREATE_HOSTGROUP_FAILED:          "ERR_CREATE_HOSTGROUP_FAILED",
	ERR_ADD_LDEV_TO_HOSTGROUP_FAILED:     "ERR_ADD_LDEV_TO_HOSTGROUP_FAILED",
	ERR_ADD_WWN_TO_HOSTGROUP_FAILED:      "ERR_ADD_WWN_TO_HOSTGROUP_FAILED",
	ERR_SET_NICKNAME_TO_WWN_FAILED:       "ERR_SET_NICKNAME_TO_WWN_FAILED",
	INFO_DELETE_HOSTGROUP_BEGIN:          "INFO_DELETE_HOSTGROUP_BEGIN",
	INFO_DELETE_HOSTGROUP_END:            "INFO_DELETE_HOSTGROUP_END",
	ERR_DELETE_HOSTGROUP_FAILED:          "ERR_DELETE_HOSTGROUP_FAILED",
	INFO_SET_MODE_OPTION_HOSTGROUP_BEGIN: "INFO_SET_MODE_OPTION_HOSTGROUP_BEGIN",
	INFO_SET_MODE_OPTION_HOSTGROUP_END:   "INFO_SET_MODE_OPTION_HOSTGROUP_END",
	ERR_SET_MODE_OPTION_HOSTGROUP_FAILED: "ERR_SET_MODE_OPTION_HOSTGROUP_FAILED",
	INFO_DELETE_WWN_BEGIN:                "INFO_DELETE_WWN_BEGIN",
	INFO_DELETE_WWN_END:                  "INFO_DELETE_WWN_END",
	ERR_DELETE_WWN_FAILED:                "ERR_DELETE_WWN_FAILED",
	INFO_SET_NICKNAME_HOSTGROUP_BEGIN:    "INFO_SET_NICKNAME_HOSTGROUP_BEGIN",
	INFO_SET_NICKNAME_HOSTGROUP_END:      "INFO_SET_NICKNAME_HOSTGROUP_END",
	INFO_ADD_WWN_TO_HOSTGROUP_BEGIN:      "INFO_ADD_WWN_TO_HOSTGROUP_BEGIN",
	INFO_ADD_WWN_TO_HOSTGROUP_END:        "INFO_ADD_WWN_TO_HOSTGROUP_END",
	INFO_REMOVE_LUN_HOSTGROUP_BEGIN:      "INFO_REMOVE_LUN_HOSTGROUP_BEGIN",
	INFO_REMOVE_LUN_HOSTGROUP_END:        "INFO_REMOVE_LUN_HOSTGROUP_END",
	ERR_REMOVE_LUN_HOSTGROUP_FAILED:      "ERR_REMOVE_LUN_HOSTGROUP_FAILED",
	INFO_ADD_LDEV_TO_HOSTGROUP_BEGIN:     "INFO_ADD_LDEV_TO_HOSTGROUP_BEGIN",
	INFO_ADD_LDEV_TO_HOSTGROUP_END:       "INFO_ADD_LDEV_TO_HOSTGROUP_END",

	// ISCSI TARGET
	INFO_GET_ISCSITARGET_BEGIN:             "INFO_GET_ISCSITARGET_BEGIN",
	INFO_GET_ISCSITARGET_END:               "INFO_GET_ISCSITARGET_END",
	ERR_GET_ISCSITARGET_FAILED:             "ERR_GET_ISCSITARGET_FAILED",
	INFO_GET_ALL_ISCSITARGET_BEGIN:         "INFO_GET_ALL_ISCSITARGET_BEGIN",
	INFO_GET_ALL_ISCSITARGET_END:           "INFO_GET_ALL_ISCSITARGET_END",
	ERR_GET_ALL_ISCSITARGET_FAILED:         "ERR_GET_ALL_ISCSITARGET_FAILED",
	INFO_SET_MODE_OPTION_ISCSITARGET_BEGIN: "INFO_SET_MODE_OPTION_ISCSITARGET_BEGIN",
	INFO_SET_MODE_OPTION_ISCSITARGET_END:   "INFO_SET_MODE_OPTION_ISCSITARGET_END",
	ERR_SET_MODE_OPTION_ISCSITARGET_FAILED: "ERR_SET_MODE_OPTION_ISCSITARGET_FAILED",
	INFO_CREATE_ISCSITARGET_BEGIN:          "INFO_CREATE_ISCSITARGET_BEGIN",
	INFO_CREATE_ISCSITARGET_END:            "INFO_CREATE_ISCSITARGET_END",
	ERR_CREATE_ISCSITARGET_FAILED:          "ERR_CREATE_ISCSITARGET_FAILED",
	INFO_ADD_LDEV_TO_ISCSITARGET_BEGIN:     "INFO_ADD_LDEV_TO_ISCSITARGET_BEGIN",
	INFO_ADD_LDEV_TO_ISCSITARGET_END:       "INFO_ADD_LDEV_TO_ISCSITARGET_END",
	ERR_ADD_LDEV_TO_ISCSITARGET_FAILED:     "ERR_ADD_LDEV_TO_ISCSITARGET_FAILED",
	INFO_ADD_IQN_NAME_TO_ISCSITARGET_BEGIN: "INFO_ADD_IQN_NAME_TO_ISCSITARGET_BEGIN",
	INFO_ADD_IQN_NAME_TO_ISCSITARGET_END:   "INFO_ADD_IQN_NAME_TO_ISCSITARGET_END",
	ERR_ADD_IQN_NAME_TO_ISCSITARGET_FAILED: "ERR_ADD_IQN_NAME_TO_ISCSITARGET_FAILED",
	INFO_ADD_NICKNAME_TO_ISCSITARGET_BEGIN: "INFO_ADD_NICKNAME_TO_ISCSITARGET_BEGIN",
	INFO_ADD_NICKNAME_TO_ISCSITARGET_END:   "INFO_ADD_NICKNAME_TO_ISCSITARGET_END",
	ERR_ADD_NICKNAME_TO_ISCSITARGET_FAILED: "ERR_ADD_NICKNAME_TO_ISCSITARGET_FAILED",
	INFO_DELETE_IQN_NAME_BEGIN:             "INFO_DELETE_IQN_NAME_BEGIN",
	INFO_DELETE_IQN_NAME_END:               "INFO_DELETE_IQN_NAME_END",
	ERR_DELETE_IQN_NAME_FAILED:             "ERR_DELETE_IQN_NAME_FAILED",
	INFO_DELETE_ISCSITARGET_BEGIN:          "INFO_DELETE_ISCSITARGET_BEGIN",
	INFO_DELETE_ISCSITARGET_END:            "INFO_DELETE_ISCSITARGET_END",
	ERR_DELETE_ISCSITARGET_FAILED:          "ERR_DELETE_ISCSITARGET_FAILED",

	// ISCSI TARGET CHAP USERS
	INFO_GET_ISCSITARGET_CHAPUSER_BEGIN:       "INFO_GET_ISCSITARGET_CHAPUSER_BEGIN",
	INFO_GET_ISCSITARGET_CHAPUSER_END:         "INFO_GET_ISCSITARGET_CHAPUSER_END",
	ERR_GET_ISCSITARGET_CHAPUSER_FAILED:       "ERR_GET_ISCSITARGET_CHAPUSER_FAILED",
	INFO_GET_ISCSITARGET_CHAPUSERS_BEGIN:      "INFO_GET_ISCSITARGET_CHAPUSERS_BEGIN",
	INFO_GET_ISCSITARGET_CHAPUSERS_END:        "INFO_GET_ISCSITARGET_CHAPUSERS_END",
	ERR_GET_ISCSITARGET_CHAPUSERS_FAILED:      "ERR_GET_ISCSITARGET_CHAPUSERS_FAILED",
	INFO_CREATE_ISCSITARGET_CHAPUSER_BEGIN:    "INFO_CREATE_ISCSITARGET_CHAPUSER_BEGIN",
	INFO_CREATE_ISCSITARGET_CHAPUSER_END:      "INFO_CREATE_ISCSITARGET_CHAPUSER_END",
	ERR_CREATE_ISCSITARGET_CHAPUSER_FAILED:    "ERR_CREATE_ISCSITARGET_CHAPUSER_FAILED",
	INFO_SET_ISCSITARGET_CHAPUSERNAME_BEGIN:   "INFO_SET_ISCSITARGET_CHAPUSERNAME_BEGIN",
	INFO_SET_ISCSITARGET_CHAPUSERNAME_END:     "INFO_SET_ISCSITARGET_CHAPUSERNAME_END",
	ERR_SET_ISCSITARGET_CHAPUSERNAME_FAILED:   "ERR_SET_ISCSITARGET_CHAPUSERNAME_FAILED",
	INFO_SET_ISCSITARGET_CHAPUSERSECRET_BEGIN: "INFO_SET_ISCSITARGET_CHAPUSERSECRET_BEGIN",
	INFO_SET_ISCSITARGET_CHAPUSERSECRET_END:   "INFO_SET_ISCSITARGET_CHAPUSERSECRET_END",
	ERR_SET_ISCSITARGET_CHAPUSERSECRET_FAILED: "ERR_SET_ISCSITARGET_CHAPUSERSECRET_FAILED",
	INFO_DELETE_ISCSITARGET_CHAPUSER_BEGIN:    "INFO_DELETE_ISCSITARGET_CHAPUSER_BEGIN",
	INFO_DELETE_ISCSITARGET_CHAPUSER_END:      "INFO_DELETE_ISCSITARGET_CHAPUSER_END",
	ERR_DELETE_ISCSITARGET_CHAPUSER_FAILED:    "ERR_DELETE_ISCSITARGET_CHAPUSER_FAILED",
	INFO_CHANGE_ISCSITARGET_CHAPUSER_BEGIN:    "INFO_CHANGE_ISCSITARGET_CHAPUSER_BEGIN",
	INFO_CHANGE_ISCSITARGET_CHAPUSER_END:      "INFO_CHANGE_ISCSITARGET_CHAPUSER_END",
	ERR_CHANGE_ISCSITARGET_CHAPUSER_FAILED:    "ERR_CHANGE_ISCSITARGET_CHAPUSER_FAILED",

	// STORAGE PORTS
	INFO_GET_STORAGE_PORTS_BEGIN:        "INFO_GET_STORAGE_PORTS_BEGIN",
	INFO_GET_STORAGE_PORTS_END:          "INFO_GET_STORAGE_PORTS_END",
	ERR_GET_STORAGE_PORTS_FAILED:        "ERR_GET_STORAGE_PORTS_FAILED",
	INFO_GET_STORAGE_PORTS_PORTID_BEGIN: "INFO_GET_STORAGE_PORTS_PORTID_BEGIN",
	INFO_GET_STORAGE_PORTS_PORTID_END:   "INFO_GET_STORAGE_PORTS_PORTID_END",
	ERR_GET_STORAGE_PORTS_PORTID_FAILED: "ERR_GET_STORAGE_PORTS_PORTID_FAILED",

	// DYNAMIC POOL
	INFO_GET_DYNAMIC_POOLS_BEGIN:   "INFO_GET_DYNAMIC_POOLS_BEGIN",
	INFO_GET_DYNAMIC_POOLS_END:     "INFO_GET_DYNAMIC_POOLS_END",
	ERR_GET_DYNAMIC_POOLS_FAILED:   "ERR_GET_DYNAMIC_POOLS_FAILED",
	INFO_GET_DYNAMIC_POOL_ID_BEGIN: "INFO_GET_DYNAMIC_POOL_ID_BEGIN",
	INFO_GET_DYNAMIC_POOL_ID_END:   "INFO_GET_DYNAMIC_POOL_ID_END",
	ERR_GET_DYNAMIC_POOL_ID_FAILED: "ERR_GET_DYNAMIC_POOL_ID_FAILED",

	// PARITY GROUP
	INFO_GET_PARITY_GROUP_BEGIN: "INFO_GET_PARITY_GROUP_BEGIN",
	INFO_GET_PARITY_GROUP_END:   "INFO_GET_PARITY_GROUP_END",
	ERR_GET_PARITY_GROUP_FAILED: "ERR_GET_PARITY_GROUP_FAILED",
}

func (s MessageID) String() string { return enumStrings[s] }

// GetEnumString .
func GetEnumString(m interface{}) string {
	if m, ok := m.(MessageID); ok {
		return m.String()
	}

	return "UNKNOWN"
}