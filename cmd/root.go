/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tuxiedev/gotweet/pkg/structs"

	"github.com/spf13/viper"
)

var cfgFile string

var twitterAPIKey, twitterAPIKeySecret, twitterAccessToken, twitterAccessSecret string

var keywords []string

var rootCmd = &cobra.Command{
	Use:   "gotweet",
	Short: "Read tweets and publish them somewhere",
	Long:  `For more information, checkout github.com/tuxiedev`,
}

func getTwitterConfigs() structs.TwitterConfig {
	return structs.TwitterConfig{
		Credentials: structs.TwitterCredentials{
			APIKey:       twitterAPIKey,
			APISecret:    twitterAPIKeySecret,
			AccessToken:  twitterAccessToken,
			AccessSecret: twitterAccessSecret,
		},
		Keywords: keywords,
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

	pf := rootCmd.PersistentFlags()

	viper.AutomaticEnv()

	buildFlags(pf, []RequiredFlag{
		{
			&twitterAPIKey,
			"twitter-api-key",
			"TWITTER_API_KEY",
			"Twitter API key",
		},
		{
			&twitterAPIKeySecret,
			"twitter-api-secret",
			"TWITTER_API_SECRET",
			"Twitter API secret",
		},
		{
			&twitterAccessToken,
			"twitter-access-token",
			"TWITTER_ACCESS_TOKEN",
			"Twitter access token",
		},
		{
			&twitterAccessSecret,
			"twitter-access-secret",
			"TWITTER_ACCESS_SECRET",
			"Twitter access secret",
		},
	})

	pf.StringArrayVarP(&keywords, "keywords", "k", viper.GetStringSlice("TWITTER_KEYWORDS"), "keywords to filter the stream on")

	cobra.MarkFlagRequired(pf, "keywords")

}
