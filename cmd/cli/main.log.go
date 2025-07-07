package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	encoder := getEncoderLog()
	writer := getWriterSync()
	core := zapcore.NewCore(encoder, writer, zapcore.InfoLevel)

	logger := zap.New(core, zap.AddCaller())
	defer logger.Sync()

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 2))

	// Create a basic Zap logger (development mode)
	// logger, _ := zap.NewDevelopment() // prettier output
	// defer logger.Sync()               // flushes buffer

	// sugar := logger.Sugar()

	// name := "William"
	// age := 40

	// sugar.Infof("Hello name: %s, age: %d", name, age)

	// Create a production logger (JSON format, includes timestamp, caller, etc.)
	// logger, err := zap.NewProduction()
	// if err != nil {
	// 	panic(err) // Handle logger creation error
	// }
	// defer logger.Sync() // flush logs

	// sugar := logger.Sugar()

	// name := "William"
	// age := 40

	// sugar.Infof("Hello name: %s, age: %d", name, age)
}

// getEncoderLog returns a JSON encoder (used in production logging)
func getEncoderLog() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder   // readable time format
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // INFO, ERROR, etc.
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder // file:line

	return zapcore.NewJSONEncoder(encoderConfig)
}

func getWriterSync() zapcore.WriteSyncer {
	os.MkdirAll("log", os.ModePerm)
	file, err := os.OpenFile("log/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err) // In production, consider logging this error elsewhere
	}

	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)

	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
