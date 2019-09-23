package snowflake

import (
	"github.com/surh-go/snowflake/models"
	"github.com/surh-go/snowflake/service"

	//
	_ "github.com/surh-go/logrus_conf"
)

// NextID 下一个ID
func NextID() *models.ID {
	return service.NextID()
}

// ParseID 解析ID
func ParseID(id *int64) *models.Info {
	return service.ParseID(id)
}
