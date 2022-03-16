package syncip

import (
	"context"
	"github.com/cloudflare/cloudflare-go"
)

type (
	// SyncIP 同步更新ip结构体
	SyncIP struct {
		api  *cloudflare.API
		zone *cloudflare.Zone
	}
)

// NewSyncIP 初始化
func NewSyncIP(apiKey, apiEmail string) (*SyncIP, error) {
	api, err := cloudflare.New(apiKey, apiEmail)
	if err != nil {
		return nil, err
	}

	return &SyncIP{
		api: api,
	}, nil
}

// SetZone 设置zone, 通过域名
func (s *SyncIP) SetZone(ctx context.Context, name string) error {
	err := s.checkAPI()
	if err != nil {
		return err
	}

	if s.zone != nil && s.zone.Name == name {
		// Name 一致, 直接返回
		return nil
	}

	id, err := s.api.ZoneIDByName(name)
	if err != nil {
		return err
	}

	zone, err := s.api.ZoneDetails(ctx, id)
	if err != nil {
		return err
	}

	s.zone = &zone
	return nil
}

// IsSetZone 是否设置了Zone
func (s *SyncIP) IsSetZone() bool {
	return s.zone != nil
}

// lazyInit 做空检查
func (s *SyncIP) lazyInit() {

}

func (s *SyncIP) check() error {
	err := s.checkAPI()
	if err != nil {
		return err
	}
	err = s.checkZone()
	return err
}

func (s *SyncIP) checkAPI() error {
	if s.api == nil {
		return ErrNotSetAPI
	}
	return nil
}

func (s *SyncIP) checkZone() error {
	if s.zone == nil {
		return ErrNotSetZone
	}
	return nil
}
