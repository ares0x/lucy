package raft

import (
	"github.com/google/wire"
	"github.com/hashicorp/raft"
)

var ProviderSet = wire.NewSet(NewRaft)

func NewRaft() *raft.Raft {
	return nil
}
