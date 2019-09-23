package snowflake

import (
	"testing"

	"github.com/sirupsen/logrus"

	//
	_ "github.com/surh-go/logrus_conf"
)

// TestNextID 测试下一个ID
func TestNextID(t *testing.T) {
	logrus.Info(NextID().String())
}

// TestParseID 测试解析ID
func TestParseID(t *testing.T) {
	var id int64 = 24074082846773248
	info := ParseID(&id)
	logrus.Info("Timestamp: ", info.Timestamp)
	logrus.Info("MachineID: ", info.MachineID)
}
