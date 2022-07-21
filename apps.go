package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"semgrep-to-elastic/models"

	"github.com/projectdiscovery/gologger"
)

var (
	Blocker = flag.Bool("b", false, "block")
	repo    = flag.String("r", "https://gitlab-ci-token:[MASKED]@unknow", "repo name")
	URI     = flag.String("h", "https://gitlab-ci-token:[MASKED]@unknow", "API")
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
			line := fmt.Sprintf("%s -> %s:%v:%v ", *repo, data.Path, data.Start.Col, data.Start.Line)
			position := fmt.Sprintf("[%s]  [%s] -> %s ", data.CheckID, data.Extra.Severity, line)
			log.Println(position)
			log.Println(data.Extra.Lines)
			fmt.Println("=======================================================")
			if data.Extra.Severity == "WARNING" {
				isBlocker = true
			}
		}
		for _, data := range semgrepJson.Results {
			data.RepoURI = *repo
			postBody, _ := json.Marshal(data)
			responseBody := bytes.NewBuffer(postBody)
			//Leverage Go's HTTP Post function to make request
			resp, err := http.Post(*URI, "application/json", responseBody)
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
		if *Blocker {
			gologger.Error().Str("Warning", "Issue must fix").Msg("[!] You must fixing this issue !!!")
			os.Exit(1)
		}
	}

}
