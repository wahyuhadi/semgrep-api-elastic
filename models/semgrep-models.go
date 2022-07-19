package models

type SemgrepJSON struct {
	Errors []interface{} `json:"errors"`
	Paths  struct {
		Scanned []string `json:"scanned"`
		Skipped []struct {
			Path   string `json:"path"`
			Reason string `json:"reason"`
		} `json:"skipped"`
	} `json:"paths"`
	Results []struct {
		RepoURI string `json:"repo_url"`
		CheckID string `json:"check_id"`
		End     struct {
			Col    int `json:"col"`
			Line   int `json:"line"`
			Offset int `json:"offset"`
		} `json:"end"`
		Extra struct {
			Fingerprint string `json:"fingerprint"`
			IsIgnored   bool   `json:"is_ignored"`
			Lines       string `json:"lines"`
			Message     string `json:"message"`
			Metadata    struct {
			} `json:"metadata"`
			Metavars struct {
				VAR struct {
					AbstractContent string `json:"abstract_content"`
					End             struct {
						Col    int `json:"col"`
						Line   int `json:"line"`
						Offset int `json:"offset"`
					} `json:"end"`
					Start struct {
						Col    int `json:"col"`
						Line   int `json:"line"`
						Offset int `json:"offset"`
					} `json:"start"`
					UniqueID struct {
						Md5Sum string `json:"md5sum"`
						Type   string `json:"type"`
					} `json:"unique_id"`
				} `json:"$VAR"`
			} `json:"metavars"`
			Severity string `json:"severity"`
		} `json:"extra"`
		Path  string `json:"path"`
		Start struct {
			Col    int `json:"col"`
			Line   int `json:"line"`
			Offset int `json:"offset"`
		} `json:"start"`
	} `json:"results"`
	Version string `json:"version"`
}
