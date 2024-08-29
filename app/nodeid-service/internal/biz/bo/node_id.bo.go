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

// Initialization 初始化
func (s *NodeIDConfig) Initialization() {
	if s.MinNodeID < DefaultMinNodeID {
		s.MinNodeID = DefaultMinNodeID
	}
	if s.MaxNodeID <= 0 {
		s.MaxNodeID = DefaultMaxNodeID
	}
	if s.MaxNodeID < s.MinNodeID {
		s.MaxNodeID = s.MinNodeID
	}
	if s.IdleDuration <= time.Second {
		s.IdleDuration = DefaultIdleDuration
	}
	if s.HeartbeatInterval <= time.Second {
		s.HeartbeatInterval = DefaultHeartbeatDuration
	}
	if s.HeartbeatInterval < s.IdleDuration {
		s.HeartbeatInterval = s.IdleDuration - time.Second/2
	}
}

func (s *NodeIDConfig) Clone() *NodeIDConfig {
	res := *s
	return &res
}

func (s *NodeIDConfig) IsValidNodeID(nodeID int64) bool {
	return nodeID >= s.MinNodeID && nodeID <= s.MaxNodeID
}

func (s *NodeIDConfig) NextExpireTime(t time.Time) time.Time {
	return t.Add(s.IdleDuration)
}

func (s *NodeIDConfig) PreviousExpiredTime(t time.Time) time.Time {
	return t.Add(-s.IdleDuration)
}

type GetNodeIdParam struct {
	InstanceId   string            // 实例ID
	InstanceName string            // 实例名称
	Metadata     map[string]string // 实例元数据
}

type RenewalNodeIdParam struct {
	ID         uint64
	InstanceId string
	NodeID     int64
}
