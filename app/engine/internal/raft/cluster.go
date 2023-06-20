package raft

import (
	"fmt"
	zaplog "github.com/go-kratos/kratos/contrib/log/zap/v2"
	"github.com/hashicorp/raft"
)

type cluster struct {
	Stores     *raft.InmemStore
	log        Logger
	transports raft.Transport
}

type Logger interface {
	Log(v ...interface{})
	Logf(s string, v ...interface{})
}

type LoggerAdapter struct {
	log zaplog.Logger
}

// Log a message to the contained debug log
func (a *LoggerAdapter) Log(v ...interface{}) {

}

// Logf will record a formatted message to the contained debug log
func (a *LoggerAdapter) Logf(s string, v ...interface{}) {
	a.log.Info(fmt.Sprintf(s, v...))
}

func NewCluster() *cluster {
	zap.NewLogger()
}
