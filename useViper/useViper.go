package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type ConfigStructure struct {
	MacPass     string `mapstructure:"macos"`
	LinuxPass   string `mapstructure:"linux"`
	WindowsPass string `mapstructure:"windows"`
	PostHost    string `mapstructure:"postgres"`
	MySQLHost   string `mapstructure:"mysql"`
	MongoHost   string `mapstructure:"mongodb"`
}

var CONFIG = ".config.json"

func aliasNormalizeFunc(f *pflag.FlagSet, n string) pflag.NormalizedName {
	switch n {
	case "pass":
		n = "password"
	case "ps":
		n = "password"
	}
	return pflag.NormalizedName(n)
}

func main() {
	pflag.StringP("name", "n", "Matt", "Name Parameter")
	pflag.StringP("password", "p", "xxxxx", "Password")
	pflag.StringP("config", "c", ".config.json", "Configuration file")
	pflag.CommandLine.SetNormalizeFunc(aliasNormalizeFunc)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	name := viper.GetString("name")
	password := viper.GetString("password")
	config := viper.GetString("config")

	fmt.Println(name, password)

	viper.BindEnv("GOMAXPROCS")
	val := viper.Get("GOMAXPROCS")
	if val != nil {
		fmt.Println("GOMAXPROCS: ", val)
	}

	viper.SetConfigType("json")
	viper.SetConfigFile(config)
	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())
	viper.ReadInConfig()

	if viper.IsSet("active") {
		val := viper.GetBool("active")
		if val {
			postgres := viper.Get("postgres")
			fmt.Println(postgres)
		}
	}

	var t ConfigStructure
	_ = viper.Unmarshal(&t)
	fmt.Println(t)

}
