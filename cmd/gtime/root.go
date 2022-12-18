package gtime

import (
	"fmt"
	"os"
	"time"

	"github.com/rchaganti/gtime/pkg/gtime"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "gtime",
	Short: "gtime - a command line global time converter tool",
	Long: `gtime is a command line timezone converter tool.

	You can use gtime to either convert a specified time value or local time across one or more timezones specified.`,
	Run: func(c *cobra.Command, args []string) {
		t := viper.GetString("targetTime")
		tz := viper.GetStringSlice("timezone")
		result := gtime.ConvertTime(t, tz)
		gtime.PrettyPrint(result)
	},
}

func init() {
	cobra.OnInitialize(initializeConfig)
	// Add flags
	rootCmd.Flags().String("targetTime", time.Now().Format("15:04"), "Specify the local time [24h format] to be converted. For example: 19:00")
	rootCmd.Flags().StringSlice("timezone", []string{}, "Specify the timezone(s) to which targetTime needs to be converted. For example: Asia/Shanghai")

	viper.BindPFlag("targetTime", rootCmd.Flags().Lookup("targetTime"))
	viper.BindPFlag("timezone", rootCmd.Flags().Lookup("timezone"))
}

func initializeConfig() {
	viper.AddConfigPath(gtime.ConfigPath)
	viper.SetConfigName(gtime.ConfigName)
	viper.SetConfigType(gtime.ConfigType)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		//fmt.Println(err)
	}

	/*
		cmd.Flags().VisitAll(func(f *pflag.Flag) {
			configName := f.Name

			if !f.Changed && v.IsSet(configName) {
				val := v.Get(configName)
				cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
			}
		})
	*/
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Ooops. There was an error while executing command '%s'", err)
		os.Exit(1)
	}
}
