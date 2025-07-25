package global

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"github.com/thinhcompany/ecommerce-ver-2/pkg/logger"
	"github.com/thinhcompany/ecommerce-ver-2/pkg/setting"
	"gorm.io/gorm"
)

var (
	ConfigGlobal  *setting.Config
	AppLogger     *logger.LoggerZap
	Mdb           *gorm.DB
	Rdb           *redis.Client
	Mdbc          *sql.DB
	KafkaProducer *kafka.Writer
)
