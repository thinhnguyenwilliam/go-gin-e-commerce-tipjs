package global

import (
	"github.com/thinhcompany/ecommerce-ver-2/pkg/logger"
	"github.com/thinhcompany/ecommerce-ver-2/pkg/setting"
	"gorm.io/gorm"
)

var (
	ConfigGlobal *setting.Config
	AppLogger    *logger.LoggerZap
	Mdb          *gorm.DB
)
