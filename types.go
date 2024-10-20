package main

type GTResponse struct {
	GT string `json:"guest_token"`
}

type JSONResponse struct {
	Data struct {
		ThreadedConversationWithInjectionsV2 struct {
			Instructions []struct {
				Type    string `json:"type"`
				Entries []struct {
					Content   struct {
						Items []struct {
							EntryID string `json:"entryId"`
							Item    struct {
								ItemContent struct {
									ItemType     string `json:"itemType"`
									Typename     string `json:"__typename"`
									TweetResults struct {
										Result struct {
											Typename          string `json:"__typename"`
											RestID            string `json:"rest_id"`
											HasBirdwatchNotes bool   `json:"has_birdwatch_notes"`
											Core              struct {
												UserResults struct {
													Result struct {
														Typename                   string `json:"__typename"`
														ID                         string `json:"id"`
														RestID                     string `json:"rest_id"`
														AffiliatesHighlightedLabel struct {
														} `json:"affiliates_highlighted_label"`
														HasGraduatedAccess bool   `json:"has_graduated_access"`
														IsBlueVerified     bool   `json:"is_blue_verified"`
														ProfileImageShape  string `json:"profile_image_shape"`
														Legacy             struct {
															CreatedAt           string `json:"created_at"`
															DefaultProfile      bool   `json:"default_profile"`
															DefaultProfileImage bool   `json:"default_profile_image"`
															Description         string `json:"description"`
															Entities            struct {
																Description struct {
																	Urls []any `json:"urls"`
																} `json:"description"`
															} `json:"entities"`
															HasCustomTimelines      bool     `json:"has_custom_timelines"`
															IsTranslator            bool     `json:"is_translator"`
															Location                string   `json:"location"`
															Name                    string   `json:"name"`
															PinnedTweetIdsStr       []string `json:"pinned_tweet_ids_str"`
															ProfileBannerURL        string   `json:"profile_banner_url"`
															ProfileImageURLHTTPS    string   `json:"profile_image_url_https"`
															ProfileInterstitialType string   `json:"profile_interstitial_type"`
															ScreenName              string   `json:"screen_name"`
															TranslatorType          string   `json:"translator_type"`
														} `json:"legacy"`
													} `json:"result"`
												} `json:"user_results"`
											} `json:"core"`
											Legacy struct {
												Entities          struct {
													Hashtags     []any `json:"hashtags"`
													Symbols      []any `json:"symbols"`
													Timestamps   []any `json:"timestamps"`
													Urls         []any `json:"urls"`
													UserMentions []struct {
														IDStr      string `json:"id_str"`
														Name       string `json:"name"`
														ScreenName string `json:"screen_name"`
														Indices    []int  `json:"indices"`
													} `json:"user_mentions"`
												} `json:"entities"`
												FullText             string `json:"full_text"`
											} `json:"legacy"`
										} `json:"result"`
									} `json:"tweet_results"`
									TweetDisplayType string `json:"tweetDisplayType"`
								} `json:"itemContent"`
							} `json:"item"`
						} `json:"items"`
					} `json:"content,omitempty"`
				} `json:"entries,omitempty"`
			} `json:"instructions"`
		} `json:"threaded_conversation_with_injections_v2"`
	} `json:"data"`
}
