package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func getVideo(str string) {
	fetchURL, TweetId := GetTweetID(str)
	guest_t := CreateGuestToken()
	vidURL := MakeRequest(fetchURL, guest_t)
	SaveFile(TweetId, vidURL)
}

func GetTweetID(tweetURL string) (string, string) {
	id := strings.Split(tweetURL, "/")
	return "https://api.x.com/graphql/OoJd6A50cv8GsifjoOHGfg/TweetResultByRestId?variables=%7B%22tweetId%22%3A%22" + id[5] + "%22%2C%22withCommunity%22%3Afalse%2C%22includePromotedContent%22%3Afalse%2C%22withVoice%22%3Afalse%7D&features=%7B%22creator_subscriptions_tweet_preview_api_enabled%22%3Atrue%2C%22communities_web_enable_tweet_community_results_fetch%22%3Atrue%2C%22c9s_tweet_anatomy_moderator_badge_enabled%22%3Atrue%2C%22articles_preview_enabled%22%3Atrue%2C%22responsive_web_edit_tweet_api_enabled%22%3Atrue%2C%22graphql_is_translatable_rweb_tweet_is_translatable_enabled%22%3Atrue%2C%22view_counts_everywhere_api_enabled%22%3Atrue%2C%22longform_notetweets_consumption_enabled%22%3Atrue%2C%22responsive_web_twitter_article_tweet_consumption_enabled%22%3Atrue%2C%22tweet_awards_web_tipping_enabled%22%3Afalse%2C%22creator_subscriptions_quote_tweet_preview_enabled%22%3Afalse%2C%22freedom_of_speech_not_reach_fetch_enabled%22%3Atrue%2C%22standardized_nudges_misinfo%22%3Atrue%2C%22tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled%22%3Atrue%2C%22rweb_video_timestamps_enabled%22%3Atrue%2C%22longform_notetweets_rich_text_read_enabled%22%3Atrue%2C%22longform_notetweets_inline_media_enabled%22%3Atrue%2C%22rweb_tipjar_consumption_enabled%22%3Atrue%2C%22responsive_web_graphql_exclude_directive_enabled%22%3Atrue%2C%22verified_phone_label_enabled%22%3Afalse%2C%22responsive_web_graphql_skip_user_profile_image_extensions_enabled%22%3Afalse%2C%22responsive_web_graphql_timeline_navigation_enabled%22%3Atrue%2C%22responsive_web_enhance_cards_enabled%22%3Afalse%7D&fieldToggles=%7B%22withArticleRichContentState%22%3Atrue%2C%22withArticlePlainText%22%3Afalse%2C%22withGrokAnalyze%22%3Afalse%2C%22withDisallowedReplyControls%22%3Afalse%7D", id[5]
}

func MakeRequest(url string, GT string) string {

	client := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header = http.Header{
		"Accept-Language":           {"en"},
		"Authorization":             {"Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA"},
		"Content-Type":              {"application/json"},
		"Priority":                  {"u=1,i"},
		"Sec-CH-UA":                 {`"Google Chrome";v="129", "Not=A?Brand";v="8", "Chromium";v="129"`},
		"Sec-CH-UA-Mobile":          {"?0"},
		"Sec-CH-UA-Platform":        {`"Windows"`},
		"Sec-Fetch-Dest":            {"empty"},
		"Sec-Fetch-Mode":            {"cors"},
		"Sec-Fetch-Site":            {"same-site"},
		"X-Client-Transaction-Id":   {"GyM1wcqfO0K94c3SO+6vcirG1CLtf953dtzx8Af3EGuQG6dE41XWkWnmlNGkEgHb8VNf2RnlZpYhgpXB6WfE46J8qwhcGA"},
		"X-Guest-Token":             {GT},
		"X-Twitter-Active-User":     {"yes"},
		"X-Twitter-Client-Language": {"en"},
		"Cookie":                    {`guest_id=172921024748603680; night_mode=2; guest_id_marketing=v1%3A172921024748603680; guest_id_ads=v1%3A172921024748603680; gt=1847067769851830364; personalization_id="v1_FvcPkyh9XhC26zctUraZqA=="`},
		"Referer":                   {"https://x.com/"},
		"Referrer-Policy":           {"strict-origin-when-cross-origin"},
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("status code:", res.StatusCode)
	if res.StatusCode != 200 {
		log.Fatal("BAD REQUEST", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	vidURL := GetLinks(body)

	fmt.Println(vidURL)

	return vidURL

}

func CreateGuestToken() string {

	client := http.Client{}

	req, err := http.NewRequest("POST", "https://api.twitter.com/1.1/guest/activate.json", nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header = http.Header{
		"Authorization": {"Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA"},
	}

	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	var gt GTResponse
	json.NewDecoder(res.Body).Decode(&gt)

	return gt.GT
}

func GetLinks(jsonObj []byte) string {
	var jsonBlob TweetJSON

	json.Unmarshal([]byte(jsonObj), &jsonBlob)

	url := jsonBlob.Data.TweetResult.Result.Legacy.ExtendedEntities.Media[0].VideoInfo.Variants

	return url[len(url)-1].URL

}

func SaveFile(fileName, vidURL string) {
	client := http.Client{}

	req, err := http.NewRequest("GET", vidURL, nil)

	if err != nil {
		log.Fatal(err)
	}
	res, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("./videos/" + fileName + ".mp4")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	file.WriteString(string(body))
}
