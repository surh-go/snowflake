package service

import (
	"github.com/surh-go/snowflake/models"
)

// ParseID 解析ID
func ParseID(id *int64) *models.Info {
	timeLeftShift := MaxSequenceBit
	if machineID > 0 {
		timeLeftShift += MaxMachineBit
	}
	timeDiff := *id >> timeLeftShift
	idTime := timeDiff + Epoch

	var info models.Info
	info.Timestamp = idTime
	info.MachineID = machineID
	return &info
}
