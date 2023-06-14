package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

func ReturnDNSName(name string) string {
	return fmt.Sprintf("%s.%s", name, viper.GetString("DNS_SUFFIX"))
}
