package service

import (
	"time"

	"github.com/surh-go/snowflake/models"
)

// NextID 下一个ID
func NextID() *models.ID {
	mu.Lock()

	now := time.Now().UnixNano() / 1000000
	if now == timestamp {
		sequence++
	} else {
		timestamp = now
		sequence = 0
	}

	var id models.ID
	timeDiff := now - Epoch
	if machineID == 0 {
		timeLeftShift := MaxSequenceBit
		id = models.ID((timeDiff << timeLeftShift) | sequence)
	} else {
		timeLeftShift := MaxSequenceBit + MaxMachineBit
		id = models.ID((timeDiff << timeLeftShift) | (int64(machineID) << MaxSequenceBit) | sequence)
	}

	mu.Unlock()
	return &id
}
