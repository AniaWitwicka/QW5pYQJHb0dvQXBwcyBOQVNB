package reader_implementations

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"time"
)

type NasaReader struct {
	urlBase               string
	apiKey                string
	maxConcurrentRequests int
	fs                    *flag.FlagSet
}

func NewNasaReader() NasaReader {
	nr := NasaReader{
		urlBase: "https://api.nasa.gov/planetary/apod",
		fs:      flag.NewFlagSet("NASA", flag.PanicOnError),
	}
	nr.fs.StringVar(&nr.apiKey, "API_KEY", "DEMO_KEY", "Api key used for NASA API requests")
	nr.fs.IntVar(&nr.maxConcurrentRequests, "CONCURRENT_REQUESTS", 5, "Limit of concurrent requests to NASA API")
	return nr
}

type nasaResponse struct {
	Copyright      string `json:"copyright"`
	Date           string `json:"date"`
	Explanation    string `json:"explanation"`
	Hdurl          string `json:"hdurl"`
	MediaType      string `json:"media_type"`
	ServiceVersion string `json:"service_version"`
	Title          string `json:"title"`
	Url            string `json:"url"`
}

//TODO: add error reporting
func (n NasaReader) GetUrl(date time.Time) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", n.urlBase, nil)
	if err != nil {
		fmt.Println("Errored while creating request: ", err)
	}
	q := req.URL.Query()
	q.Add("api_key", n.apiKey)
	q.Add("date", date.Format("2006-01-02"))
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Errored while sending request to the server: ", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println("Error while sending request to the server: ", resp.StatusCode, resp.Status)
	}
	responseJson := nasaResponse{}
	json.NewDecoder(resp.Body).Decode(&responseJson)
	return responseJson.Url
}

func (n NasaReader) MaxConcurrentRequests() int {
	return n.maxConcurrentRequests
}
