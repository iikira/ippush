package syncip

import (
	"context"
	"github.com/cloudflare/cloudflare-go"
)

// SetTXTRecord 设置TXT记录
func (s *SyncIP) SetTXTRecord(ctx context.Context, fullName, value string) error {
	err := s.check()
	if err != nil {
		return err
	}

	records, err := s.api.DNSRecords(ctx, s.zone.ID, cloudflare.DNSRecord{
		Type: "TXT",
		Name: fullName,
	})
	if err != nil {
		return err
	}

	var r cloudflare.DNSRecord
	if len(records) == 0 {
		// 创建记录
		resp, err := s.api.CreateDNSRecord(ctx, s.zone.ID, cloudflare.DNSRecord{
			Type:    "TXT",
			Name:    fullName,
			Content: value,
			TTL:     120,
		})
		if err != nil {
			return err
		}
		r = resp.Result
	} else {
		r = records[0]
	}

	// 如果没有变化则不提交
	if r.Content == value {
		return nil
	}

	r.Content = value
	r.TTL = 120

	err = s.api.UpdateDNSRecord(ctx, s.zone.ID, r.ID, r)
	return err
}
