package xgo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type XDowlander struct {
	url string
}

func (d *XDowlander) XGO(url string) *XDowlander {

	return &XDowlander{
		url: url,
	}
}

func (d *XDowlander) GetTweetID(tweetURL string) (string, error) {
	id := strings.Split(tweetURL, "/")

	if id[5] != "" {
		return id[5], nil
	}
	return "", errors.New("something wrong")
}

func (d *XDowlander) MakeRequest(r *http.Response) {

	client := http.Client{}

	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		log.Fatal(err)
	}
	d.SetHeader(req)

	defer req.Body.Close()

	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)

}
func (d *XDowlander) SetHeader(r *http.Request) {
	r.Header.Add("accept", "*/*")
	r.Header.Add("accept-language", "en")
	r.Header.Add("autharization", "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA")
	r.Header.Add("content-type", "application/json")
	r.Header.Add("priority", "u=1,i")
	r.Header.Add("sec-ch-ua", `"Google Chrome";v="129", "Not=A?Brand";v="8", "Chromium";v="129"`)
	r.Header.Add("sec-ch-ua-mobile", "?0")
	r.Header.Add("sec-ch-ua-platform", `"Windows"`)
	r.Header.Add("sec-fetch-dest", "empty")
	r.Header.Add("sec-fetch-mode", "cors")
	r.Header.Add("sec-fetch-site", "same-site")
	r.Header.Add("x-client-transaction-id", "GyM1wcqfO0K94c3SO+6vcirG1CLtf953dtzx8Af3EGuQG6dE41XWkWnmlNGkEgHb8VNf2RnlZpYhgpXB6WfE46J8qwhcGA")
	r.Header.Add("x-guest-token", "1847067769851830364")
	r.Header.Add("x-twitter-active-user", "yes")
	r.Header.Add("x-twitter-client-language", "en")
	r.Header.Add("cookie", `guest_id=172921024748603680; night_mode=2; guest_id_marketing=v1%3A172921024748603680; guest_id_ads=v1%3A172921024748603680; gt=1847067769851830364; personalization_id="v1_FvcPkyh9XhC26zctUraZqA=="`)
	r.Header.Add("referer", "https://x.com/")
	r.Header.Add("referrer-policy", "strict-origin-when-cross-origin")

}

