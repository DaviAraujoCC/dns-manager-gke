package gcp

import (
	"github.com/DaviAraujoCC/dns-manager-gke/utils"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/dns/v1"
)

func (s *DnsService) ListRecordSetEntries() (*dns.ResourceRecordSetsListResponse, error) {

	entries, err := s.ResourceRecordSets.List(s.Project, s.ManagedZone).Do()
	if err != nil {
		return nil, err
	}

	return entries, nil
}

func (s *DnsService) CreateRecordSet(name, ip string) error {

	dnsName := utils.ReturnDNSName(name)

	recordSet := &dns.ResourceRecordSet{
		Name:    dnsName,
		Type:    "A",
		Ttl:     60,
		Rrdatas: []string{ip},
	}

	_, err := s.ResourceRecordSets.Create(s.Project, s.ManagedZone, recordSet).Do()
	if err != nil {
		return err
	}

	log.Info("Created DNS entry for service: ", name)

	return nil
}

func (s *DnsService) UpdateRecordSet(name, ip string) error {

	dnsName := utils.ReturnDNSName(name)

	recordSet, err := s.ResourceRecordSets.Get(s.Project, s.ManagedZone, dnsName, "A").Do()
	if err != nil {
		return err
	}

	recordSet.Rrdatas = []string{ip}

	_, err = s.ResourceRecordSets.Patch(s.Project, s.ManagedZone, dnsName, "A", recordSet).Do()
	if err != nil {
		return err
	}

	log.Info("Updated DNS entry for service: ", name)

	return nil
}

func (s *DnsService) DeleteRecordSet(name string) error {

	dnsName := utils.ReturnDNSName(name)

	_, err := s.ResourceRecordSets.Delete(s.Project, s.ManagedZone, dnsName, "A").Do()
	if err != nil {
		return err
	}

	log.Info("Deleted DNS entry for service: ", name)
	return nil
}
