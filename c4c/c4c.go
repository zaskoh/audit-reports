package c4c

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/zaskoh/c4c-reports/config"
	"github.com/zaskoh/c4c-reports/logger"
	"github.com/zaskoh/c4c-reports/types"
	"github.com/zaskoh/discordbooter"
	"go.uber.org/zap"
)

var c4cReports map[string]types.C4cReports

func init() {
	c4cReports = make(map[string]types.C4cReports)
}

func Start() error {

	// first load current known data
	err := loadBackupFile()
	if err != nil {
		return err
	}

	go func() {
		for {
			// get current reports
			currentData, err := getCurrentReports()
			if err != nil {
				logger.Error("failed to load data from c4c: " + err.Error())
			}

			// check for new entries
			newReports := 0
			for _, data := range currentData {
				if _, exists := c4cReports[data.Slug]; !exists {
					// add new entry
					c4cReports[data.Slug] = data
					newReports++

					logger.Info("New Report available",
						zap.Any("Data", data),
					)

					if config.DiscordConfig.Active {
						err := discordbooter.SendMessage(
							config.DiscordConfig.Channel,
							"New Report "+data.Slug+" available: "+data.Site,
						)
						if err != nil {
							logger.Error("failed to send to discord", zap.Error(err))
						}
					}
				}
			}

			// if we had new entries, update
			if newReports > 0 {
				updateC4cReports()
			}

			logger.Info("checked for new reports and got: " + strconv.Itoa(newReports))

			time.Sleep(time.Duration(config.Base.ReportCheckInterval) * time.Second)
		}
	}()

	return nil
}

func updateC4cReports() {
	// save tokenList
	c4cReportsFile, _ := json.MarshalIndent(c4cReports, "", "  ")
	err := os.WriteFile(config.Base.ReportFile, c4cReportsFile, 0666)
	if err != nil {
		log.Fatalf("error in writing pairList - %s", err)
	}
}

func getCurrentReports() (map[string]types.C4cReports, error) {
	result := make(map[string]types.C4cReports, 0)
	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("GET", "https://code4rena.com/page-data/reports/page-data.json", nil)
	if err != nil {
		return result, err
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return result, err
	}

	jsonDataFromHttp, err := io.ReadAll(res.Body)
	if err != nil {
		return result, err
	}

	var jsonData types.C4cPageResponse
	err = json.Unmarshal([]byte(jsonDataFromHttp), &jsonData)
	if err != nil {
		return result, err
	}

	for _, val := range jsonData.Result.Data.Reports.Edges {
		result[val.Node.Frontmatter.Slug] = types.C4cReports{
			Slug:     val.Node.Frontmatter.Slug,
			Findings: val.Node.Frontmatter.Findings,
			Site:     "https://code4rena.com/reports/" + val.Node.Frontmatter.Slug,
		}
	}

	return result, nil
}

func loadBackupFile() error {
	jsonC4cReports, err := os.OpenFile(config.Base.ReportFile, os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	jsonByteValue, _ := io.ReadAll(jsonC4cReports)
	json.Unmarshal([]byte(jsonByteValue), &c4cReports)
	return nil
}
