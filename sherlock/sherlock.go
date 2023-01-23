package sherlock

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

var sherlockReports map[string]types.SherlockReports

func init() {
	sherlockReports = make(map[string]types.SherlockReports)
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
			reportList, err := getReportList()
			if err != nil {
				logger.Error("failed to load data from sherlock: " + err.Error())
			}

			// check for new entries
			newReports := 0
			for _, data := range reportList {
				if _, exists := sherlockReports[data.Slug]; !exists {
					// add new entry
					sherlockReports[data.Slug] = data
					newReports++

					logger.Info("New Report available",
						zap.Any("Data", data),
					)

					if config.DiscordConfig.Active {
						err := discordbooter.SendMessage(
							config.DiscordConfig.Channel,
							"New Sherlock-Report "+data.Slug+" available: "+data.Site,
						)
						if err != nil {
							logger.Error("failed to send to discord", zap.Error(err))
						}
					}
				}
			}

			// if we had new entries, update
			if newReports > 0 {
				updateSherlockReports()
			}

			logger.Info("checked for new reports and got: " + strconv.Itoa(newReports))

			time.Sleep(time.Duration(config.Base.ReportCheckInterval) * time.Second)
		}
	}()

	return nil
}

func updateSherlockReports() {
	// save tokenList
	sherlockReportsFile, _ := json.MarshalIndent(sherlockReports, "", "  ")
	err := os.WriteFile(config.Base.ReportFileSherlock, sherlockReportsFile, 0666)
	if err != nil {
		log.Fatalf("error in writing pairList - %s", err)
	}
}

func getReportList() (map[string]types.SherlockReports, error) {
	result := make(map[string]types.SherlockReports, 0)
	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("GET", "https://mainnet-contest.sherlock.xyz/contests", nil)
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

	var jsonData []types.SherlockListResult
	err = json.Unmarshal([]byte(jsonDataFromHttp), &jsonData)
	if err != nil {
		return result, err
	}

	for _, val := range jsonData {
		if val.Status == "FINISHED" && !val.Private {
			result[val.TemplateRepoName] = types.SherlockReports{
				Slug:     val.TemplateRepoName,
				Findings: "https://app.sherlock.xyz/audits/contests/" + strconv.Itoa(val.ID),
				Site:     "https://app.sherlock.xyz/audits/contests/" + strconv.Itoa(val.ID),
			}
		}
	}

	return result, nil
}

func loadBackupFile() error {
	jsonSherlockReports, err := os.OpenFile(config.Base.ReportFileSherlock, os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	jsonByteValue, _ := io.ReadAll(jsonSherlockReports)
	json.Unmarshal([]byte(jsonByteValue), &sherlockReports)
	return nil
}
