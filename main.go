package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {

	tweetFlag := flag.String("t", "empty", "tweet link")
	flag.Parse()

	fetchURL := GetTweetID(*tweetFlag)
	guest_t := CreateGuestToken()
	MakeRequest(fetchURL, guest_t)
}

func GetTweetID(tweetURL string) string {
	id := strings.Split(tweetURL, "/")
	return "https://api.x.com/graphql/OoJd6A50cv8GsifjoOHGfg/TweetResultByRestId?variables=%7BtweetId%3A%20%22" + id[5] + "%22%2C%20withCommunity%3A%20false%2C%20includePromotedContent%3A%20false%2C%20withVoice%3A%20false%2C%20%7D&features=%7B%20rweb_tipjar_consumption_enabled%3A%20false%2C%20responsive_web_graphql_exclude_directive_enabled%3A%20false%2C%20verified_phone_label_enabled%3A%20false%2C%20creator_subscriptions_tweet_preview_api_enabled%3A%20false%2C%20responsive_web_graphql_timeline_navigation_enabled%3A%20false%2C%20responsive_web_graphql_skip_user_profile_image_extensions_enabled%3A%20false%2C%20communities_web_enable_tweet_community_results_fetch%3A%20false%2C%20c9s_tweet_anatomy_moderator_badge_enabled%3A%20false%2C%20articles_preview_enabled%3A%20false%2C%20responsive_web_edit_tweet_api_enabled%3A%20false%2C%20graphql_is_translatable_rweb_tweet_is_translatable_enabled%3A%20false%2C%20view_counts_everywhere_api_enabled%3A%20false%2C%20longform_notetweets_consumption_enabled%3A%20false%2C%20responsive_web_twitter_article_tweet_consumption_enabled%3A%20false%2C%20tweet_awards_web_tipping_enabled%3A%20false%2C%20creator_subscriptions_quote_tweet_preview_enabled%3A%20false%2C%20freedom_of_speech_not_reach_fetch_enabled%3A%20false%2C%20standardized_nudges_misinfo%3A%20false%2C%20tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled%3A%20false%2C%20rweb_video_timestamps_enabled%3A%20false%2C%20longform_notetweets_rich_text_read_enabled%3A%20false%2C%20longform_notetweets_inline_media_enabled%3A%20false%2C%20responsive_web_enhance_cards_enabled%3A%20false%2C%20%7D&fieldToggles=%7B%20withArticleRichContentState%3A%20false%2C%20withArticlePlainText%3A%20false%2C%20withGrokAnalyze%3A%20false%2C%20withDisallowedReplyControls%3A%20false%2C%20%7D"
}

func MakeRequest(url string, GT string) {

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

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Gövde okuma hatası:", err)
		return
	}

	// Yanıtı yazdırma
	fmt.Println("Yanıt gövdesi:")
	fmt.Println(string(body))

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
