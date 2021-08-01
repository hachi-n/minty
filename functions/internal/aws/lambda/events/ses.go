package events

import "time"

type AmazonSesEvent struct {
	Records []struct {
		EventSource  string `json:"eventSource"`
		EventVersion string `json:"eventVersion"`
		Ses          struct {
			Mail struct {
				CommonHeaders struct {
					Date       string   `json:"date"`
					From       []string `json:"from"`
					MessageID  string   `json:"messageId"`
					ReturnPath string   `json:"returnPath"`
					Subject    string   `json:"subject"`
					To         []string `json:"to"`
				} `json:"commonHeaders"`
				Destination []string `json:"destination"`
				Headers     []struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"headers"`
				HeadersTruncated bool      `json:"headersTruncated"`
				MessageID        string    `json:"messageId"`
				Source           string    `json:"source"`
				Timestamp        time.Time `json:"timestamp"`
			} `json:"mail"`
			Receipt struct {
				Action struct {
					FunctionArn    string `json:"functionArn"`
					InvocationType string `json:"invocationType"`
					Type           string `json:"type"`
				} `json:"action"`
				DkimVerdict struct {
					Status string `json:"status"`
				} `json:"dkimVerdict"`
				ProcessingTimeMillis int      `json:"processingTimeMillis"`
				Recipients           []string `json:"recipients"`
				SpamVerdict          struct {
					Status string `json:"status"`
				} `json:"spamVerdict"`
				SpfVerdict struct {
					Status string `json:"status"`
				} `json:"spfVerdict"`
				Timestamp    time.Time `json:"timestamp"`
				VirusVerdict struct {
					Status string `json:"status"`
				} `json:"virusVerdict"`
			} `json:"receipt"`
		} `json:"ses"`
	} `json:"Records"`
}
