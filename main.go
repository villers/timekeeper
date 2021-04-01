package main

import (
	"encoding/json"
	"fmt"
	"github.com/andybrewer/mack"
	"github.com/go-vgo/robotgo"
	"github.com/villers/timekeeper/models"
	"time"
)

var entry []models.Entry

func getUrl(title string) string {
	var url string
	var err error

	switch title {
	case "Google Chrome", "Google Chrome Canary", "Brave Browser":
		{
			url, err = mack.Tell(title, " URL of active tab of front window")
		}

	case "webkit", "Safari":
		{
			url, err = mack.Tell(title, " URL of front document")
		}

	default:
		{
			fmt.Println("current process", title)
		}
	}

	if err != nil {
		fmt.Println("error: ", err)
	}

	return url
}

func getCurrentProcess() models.ProcessInfo {
	pid := robotgo.GetPID()
	processName, _ := robotgo.FindName(pid)

	return models.ProcessInfo{
		Pid:         pid,
		WindowName:  robotgo.GetTitle(),
		AppName:     robotgo.GetTitle(pid),
		ProcessName: processName,
		Url:         getUrl(processName),
	}
}

func main() {
	fmt.Println("Hello")

	for {
		p := getCurrentProcess()

		metadata := models.Metadata{
			WindowName: p.WindowName,
			Url:        p.Url,
			Pid:        p.Pid,
			Date:       time.Time{},
			Duration:   1,
		}

		if len(entry) == 0 || entry[len(entry)-1].Name != p.AppName {
			entry = append(entry, models.Entry{
				Name:      p.AppName,
				Date:      time.Now(),
				Metadatas: []models.Metadata{metadata},
			})
		} else {
			metadatas := entry[len(entry)-1].Metadatas
			if metadatas[len(metadatas)-1].WindowName == metadata.WindowName {
				metadatas[len(metadatas)-1].Duration += 1
			} else {
				fmt.Println("test")
				metadatas = append(metadatas, metadata)
				entry[len(entry)-1].Metadatas = metadatas
			}
		}

		entryBytes, err := json.Marshal(entry)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(entryBytes))
		time.Sleep(1 * time.Second)
		fmt.Println("--------------")
	}
}
