package command

import (
	"strings"

	"github.com/DaviAraujoCC/dns-manager-gke/gcp"
	"github.com/DaviAraujoCC/dns-manager-gke/utils"
	log "github.com/sirupsen/logrus"
)

func CheckCreateDnsEntries(validServices, validDNSEntries map[string]string, dnsGCPService *gcp.DnsService) {
	for s := range validServices {

		serviceName := s
		dnsName := utils.ReturnDNSName(serviceName)

		if validDNSEntries[dnsName] == "" {

			log.Println("Creating DNS entry for service: ", serviceName)
			err := dnsGCPService.CreateRecordSet(serviceName, validServices[s])
			if err != nil {
				log.Warn(err)
			}

		}

	}

}

func CheckDeleteDnsEntries(validServices, validDNSEntries map[string]string, dnsGCPService *gcp.DnsService) {

	for dns := range validDNSEntries {

		serviceName := strings.SplitN(dns, ".", -1)[0]

		if validServices[serviceName] == "" {

			log.Println("Removing DNS entry for service: ", dns)
			err := dnsGCPService.DeleteRecordSet(serviceName)
			if err != nil {
				log.Warn(err)
			}

		}
	}

}

func CheckUpdateDnsEntries(validServices, validDNSEntries map[string]string, dnsGCPService *gcp.DnsService) {
	for dns := range validDNSEntries {

		serviceName := strings.SplitN(dns, ".", -1)[0]

		if validServices[serviceName] != validDNSEntries[dns] && validServices[serviceName] != "" {

			log.Println("Updating DNS entry for service: ", serviceName)
			err := dnsGCPService.UpdateRecordSet(serviceName, validServices[serviceName])
			if err != nil {
				log.Warn(err)
			}

		}

	}

}
