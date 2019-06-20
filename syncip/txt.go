package syncip

import (
	// "fmt"
	"github.com/cloudflare/cloudflare-go"
)

// CreateTXTRecord 创建TXT记录
func (s *SyncIP) CreateTXTRecord(fullName, value string) (*cloudflare.DNSRecordResponse, error) {
	err := s.check()
	if err != nil {
		return nil, err
	}

	resp, err := s.api.CreateDNSRecord(s.zone.ID, cloudflare.DNSRecord{
		Type:    "TXT",
		Name:    fullName,
		Content: value,
		TTL:     120,
	})
	return resp, err
}

// SetTXTRecord 设置TXT记录
func (s *SyncIP) SetTXTRecord(fullName, value string) error {
	err := s.check()
	if err != nil {
		return err
	}

	records, err := s.api.DNSRecords(s.zone.ID, cloudflare.DNSRecord{
		Type: "TXT",
		Name: fullName,
	})
	if err != nil {
		return err
	}

	var r cloudflare.DNSRecord
	if len(records) == 0 {
		resp, err := s.CreateTXTRecord(fullName, value)
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

	err = s.api.UpdateDNSRecord(s.zone.ID, r.ID, r)
	return err
}
