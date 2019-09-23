package service

import (
	"sync"
)

var (
	// Epoch 开始时间戳，固定小于当前时间的毫秒，用于缩减时间位数
	Epoch int64 = 1546272000000
	// MaxMachineBit 最大机器标识位数，8->255 10->1023
	MaxMachineBit uint8 = 8
	// MaxSequenceBit 最大序号数位数，12->4095
	MaxSequenceBit uint8 = 12
	// maxMachineID 最大机器标识号
	maxMachineID uint16 = (1 << MaxMachineBit) - 1
	// timestamp 毫秒时间戳
	timestamp int64
	// sequence 每毫秒内的队列
	sequence int64
	// mu 同步锁
	mu sync.Mutex
	// machineID 机器标识号
	machineID uint16 = 1
)
