package configs

import (
	"strings"
	"github.com/spf13/viper"
)

var (
	passphrase string
	fang       *viper.Viper
)

func init() {
	fang = viper.New()
	fang.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	fang.SetConfigType("yaml")

}
