package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"semgrep-to-elastic/models"
	"strings"

	"github.com/projectdiscovery/gologger"
)

var (
	Blocker = false
	repo    = flag.String("r", "https://gitlab-ci-token:[MASKED]@unknow", "repo name")
	URI     = "http://127.0.0.1:8080" // API
)

func main() {
	flag.Parse()
	sc := bufio.NewScanner(os.Stdin)
	var semgrepJson models.SemgrepJSON
	isBlocker := false
	for sc.Scan() {
		err := json.Unmarshal([]byte(sc.Text()), &semgrepJson)
		if err != nil {
			gologger.Info().Str("Error", "parsing to models").Msg("Failed parsing to models")
			continue
		}

		if len(semgrepJson.Results) == 0 {
			gologger.Info().Str("Info", "No Issue Found").Msg("Yeayyyy, No issue found in this repo")
			continue
		}

		for _, data := range semgrepJson.Results {
			repo_name := strings.Split(*repo, "@")
			line := fmt.Sprintf("%s -> %s:%v:%v ", repo_name[1], data.Path, data.Start.Col, data.Start.Line)
			msg := fmt.Sprintf("[%s]  [%s] -> %s ", data.CheckID, data.Extra.Severity, line)
			gologger.Info().Str("Warning", "Issue Found").Msg(msg)
			if data.Extra.Severity == "WARNING" {
				isBlocker = true
			}
		}

		for _, data := range semgrepJson.Results {
			repo_name := strings.Split(*repo, "@")
			data.RepoURI = repo_name[1]
			postBody, _ := json.Marshal(data)
			responseBody := bytes.NewBuffer(postBody)
			//Leverage Go's HTTP Post function to make request
			resp, err := http.Post(URI, "application/json", responseBody)
			//Handle Error
			if err != nil {
				gologger.Info().Str("Warning", "Cant connect to api").Msg(err.Error())
				continue
			}
			if resp.StatusCode != http.StatusOK {
				gologger.Info().Str("Warning", "Cant connect to api").Msg(fmt.Sprintf("HTTP CODE = %v", resp.StatusCode))
			}

		}
	}

	if isBlocker {
		if Blocker {
			gologger.Error().Str("Warning", "Issue must fix").Msg("[!] You must fixing this issue !!!")
			os.Exit(1)
		}
	}

}
