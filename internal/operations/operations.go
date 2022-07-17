package operations

import (
	"time"
)

func ListUrls(from, to string, reader ApiReader) ([]string, error) {
	fromDate, toDate, err := ParseAndValidateDateRanges(from, to)
	if err != nil {
		return nil, err
	}
	//Guard makes sure that the number of concurrent request isn't exceeded by blocking space in channel with empty struct
	guard := make(chan struct{}, reader.MaxConcurrentRequests())
	urlsResponse := make([]string, GetUrlLen(fromDate, toDate))
	counter := 0
	for d := fromDate; d.After(toDate) == false; d = d.AddDate(0, 0, 1) {
		guard <- struct{}{}
		go func(d time.Time, counter int) {
			urlsResponse[counter] = reader.GetUrl(d)
			<-guard
		}(d, counter)
		counter++
	}

	return urlsResponse, nil
}

