package bo

import (
	"encoding/json"
	enumv1 "github.com/go-micro-saas/nodeid-service/api/nodeid-service/v1/enums"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
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
	NodeEpoch         time.Time
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
	if s.HeartbeatInterval >= s.IdleDuration {
		s.HeartbeatInterval = s.IdleDuration / 3
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
	InstanceId  string `json:"instance_id"`
	NodeID      int64  `json:"node_id"`
	AccessToken string `json:"access_token"`
}

func (s *RenewalNodeIdParam) MarshalToJSON() ([]byte, error) {
	buf, err := json.Marshal(s)
	if err != nil {
		return nil, errorpkg.WithStack(errorpkg.ErrorInternalServer(err.Error()))
	}
	return buf, nil
}

func (s *RenewalNodeIdParam) UnmarshalFromJSON(buf []byte) error {
	err := json.Unmarshal(buf, s)
	if err != nil {
		return errorpkg.WithStack(errorpkg.ErrorInternalServer(err.Error()))
	}
	return nil
}

type RenewalNodeIDReply struct {
	Status    enumv1.NodeIDStatusEnum_Status
	ExpiredAt time.Time
}

type ReleaseNodeIdParam struct {
	InstanceId  string
	NodeID      int64
	AccessToken string
}
