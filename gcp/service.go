package gcp

import (
	"context"

	"google.golang.org/api/dns/v1"
)

type DnsService struct {
	*dns.Service
	Project     string
	ManagedZone string
}

func NewDnsService(project, manZone string) (*DnsService, error) {
	ctx := context.Background()
	dnsService, err := dns.NewService(ctx)
	if err != nil {
		return nil, err
	}
	return &DnsService{dnsService, project, manZone}, nil
}
