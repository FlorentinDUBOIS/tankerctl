package cmd

import (
	"bytes"
	"encoding/xml"
	"net/http"
	"strconv"
	"strings"
	"time"
	"errors"

	"github.com/FlorentinDUBOIS/tankerctl/core"
	"github.com/paulrosania/go-charset/charset"
	_ "github.com/paulrosania/go-charset/data" // needed by go-charset/charset
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var scrapeCmd = &cobra.Command{
	Use:   "scrape",
	Short: "Scrape data from OpenData",
	Run:   scrape,
}

func init() {
	scrapeCmd.Flags().StringP("output-directory", "o", "", "Set the path to save metrics")

	viper.BindPFlags(scrapeCmd.Flags())

	RootCmd.AddCommand(scrapeCmd)
}

func scrape(cmd *cobra.Command, args []string) {

	url := viper.GetString("opendata.instant")

	log.Infof("Scrape from %s", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)

		return
	}

	defer resp.Body.Close()

	log.Info("Read response and unzip it")
	contents, err := core.ReadAndUnzip(resp)
	if err != nil {
		log.Fatal(err)

		return
	}
	sensisions := make([]string, 0)
	for _, content := range contents {
		container := &core.XMLContainer{}
		decoder := xml.NewDecoder(bytes.NewReader(content))
		decoder.CharsetReader = charset.NewReader

		log.Info("Decode xml response")
		if err := decoder.Decode(container); err != nil {
			log.Error(err)

			continue
		}

		log.Info("Parse into sensision metrics")
		metrics, err := container.ToSensision()
		if err != nil {
			log.Error(err)

			continue
		}

		sensisions = append(sensisions, metrics...)
	}

	path := viper.GetString("output-directory")
	name := strconv.FormatInt(time.Now().Unix(), 10)
	if len(path) <= 0 {
		log.Fatal(errors.New("No output directory provided"))

		return
	}

	log.Infof("Save metrics into: %s/%s", path, name)
	if err := core.Save(path, name, strings.Join(sensisions, "\n")); err != nil {
		log.Fatal(err)

		return
	}
}
