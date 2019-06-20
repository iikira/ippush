package syncip

import (
	// "fmt"
	"github.com/cloudflare/cloudflare-go"
)

// SetARecord 设置a记录
func (s *SyncIP) SetARecord(name, value string) error {
	err := s.check()
	if err != nil {
		return err
	}

	records, err := s.api.DNSRecords(s.zone.ID, cloudflare.DNSRecord{
		Type: "A",
		Name: name + "." + s.zone.Name,
	})
	if err != nil {
		return err
	}
	if len(records) == 0 {
		return ErrNoSuchRecord
	}

	r := records[0]

	// 如果没有变化则不提交
	if r.Content == value {
		return nil
	}

	r.Content = value

	err = s.api.UpdateDNSRecord(s.zone.ID, r.ID, r)
	return err
}
