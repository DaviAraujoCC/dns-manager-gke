package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/DaviAraujoCC/dns-manager-gke/command"
	"github.com/DaviAraujoCC/dns-manager-gke/config"
	"github.com/DaviAraujoCC/dns-manager-gke/gcp"
	"github.com/DaviAraujoCC/dns-manager-gke/k8s/controller"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/api/dns/v1"
	corev1 "k8s.io/api/core/v1"
)

func main() {

	log.Info("Starting DNS Manager")

	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Gathering info about cluster...")

	obj, err := controller.NewObjectsController(cfg.Namespace)
	if err != nil {
		log.Fatal(err)
	}

	k8sServices, err := obj.ListServices()
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Gathering info about DNS...")

	dnsGCPService, err := gcp.NewDnsService(cfg.ProjectId, cfg.ManagedZone)
	if err != nil {
		log.Fatal(err)
	}

	entries, err := dnsGCPService.ListRecordSetEntries()
	if err != nil {
		log.Fatal(err)
	}

	validDNSEntries := getValidDnsEntries(entries)
	validServices := getValidServices(k8sServices)

	//update dns for services if necessary (ip doesn't match)
	command.CheckUpdateDnsEntries(validServices, validDNSEntries, dnsGCPService)

	//create dns for services if necessary (dns doesn't present)
	command.CheckCreateDnsEntries(validServices, validDNSEntries, dnsGCPService)

	if !viper.GetBool("IGNORE_DELETE_RECORD") {
		//delete dns for services if necessary (service doesn't exist)
		command.CheckDeleteDnsEntries(validServices, validDNSEntries, dnsGCPService)
	}

	log.Info("OK")

}

func getValidServices(services *corev1.ServiceList) map[string]string {

	var validServices = make(map[string]string)
	for _, s := range services.Items {
		if s.Spec.Type == "LoadBalancer" && s.Status.LoadBalancer.Ingress != nil {
			validServices[s.Name] = s.Status.LoadBalancer.Ingress[0].IP
		}
	}

	return validServices
}

func getValidDnsEntries(entries *dns.ResourceRecordSetsListResponse) map[string]string {

	var validDnsEntries = make(map[string]string)
	for _, e := range entries.Rrsets {
		dnsSuffix := strings.Split(e.Name, ".")[1:]
		dnsSuffixN := strings.Join(dnsSuffix[:len(dnsSuffix)-1], ".")
		if e.Type == "A" && len(e.Rrdatas) > 0 && match(viper.GetString("DNS_SUFFIX"), dnsSuffixN) {
			validDnsEntries[e.Name] = e.Rrdatas[0]
		}
	}

	return validDnsEntries
}

func match(v, s string) bool {
	regexString := fmt.Sprintf("^%s\\.$", s)
	regex, _ := regexp.MatchString(regexString, v)
	return regex
}
