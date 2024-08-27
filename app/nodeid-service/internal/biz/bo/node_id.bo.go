package bo

import (
	"time"
)

const (
	DefaultMinNodeID         = 1
	DefaultMaxNodeID         = 1023
	DefaultIdleDuration      = time.Second * 100
	DefaultHeartbeatDuration = time.Second * 30
)

type NodeIDConfig struct {
	MinNodeID         int64
	MaxNodeID         int64
	IdleDuration      time.Duration
	HeartbeatInterval time.Duration
}

func (s *NodeIDConfig) Clone() *NodeIDConfig {
	res := *s
	return &res
}

// Initialization 初始化
func (s *NodeIDConfig) Initialization() {
	if s.MinNodeID < DefaultMinNodeID {
		s.MaxNodeID = DefaultMinNodeID
	}
	if s.MaxNodeID <= 0 {
		s.MaxNodeID = DefaultMaxNodeID
	}
	if s.MaxNodeID < s.MinNodeID {
		s.MaxNodeID = s.MinNodeID
	}
	if s.IdleDuration <= 0 {
		s.IdleDuration = DefaultIdleDuration
	}
	if s.HeartbeatInterval <= 0 {
		s.HeartbeatInterval = DefaultHeartbeatDuration
	}
	if s.HeartbeatInterval < s.IdleDuration {
		s.HeartbeatInterval = s.IdleDuration
	}
}

type GetNodeIdParam struct {
	InstanceId   string            // 实例ID
	InstanceName string            // 实例名称
	Metadata     map[string]string // 实例元数据
}
