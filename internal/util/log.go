package util

import (
    "os"
    "strings"
    "encoding/json"
    "path/filepath"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"

    . "git.shannonai.com/taoshibo/go-module-demo/internal/common"
)

var Logger *zap.SugaredLogger

func init() {
    for _, path := range Config.Log.OutputPaths {
        path, _ := filepath.Abs(path)
        os.MkdirAll(filepath.Dir(path), os.ModePerm)
    }
    for _, path := range Config.Log.ErrorOutputPaths {
        path, _ := filepath.Abs(path)
        os.MkdirAll(filepath.Dir(path), os.ModePerm)
    }

    var cfg zap.Config
    rawJSON, _ := json.Marshal(Config.Log)
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
    // if in debug mode, print log to stdout
    if strings.EqualFold(Config.Log.Level, "DEBUG") {
        cfg.OutputPaths = []string{"stdout"}
        cfg.ErrorOutputPaths = []string{"stderr"}
    }

    cfg.EncoderConfig = zap.NewProductionEncoderConfig()
    cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

    //var err error
    if prod_logger, err := cfg.Build(); err != nil {
        panic(err)
    } else {
        Logger = prod_logger.Sugar()
    }
}
