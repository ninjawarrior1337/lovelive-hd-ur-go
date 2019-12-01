// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    cardResponse, err := UnmarshalCardResponse(bytes)
//    bytes, err = cardResponse.Marshal()

package CardResponse

import "encoding/json"

func UnmarshalCardResponse(data []byte) (CardResponse, error) {
	var r CardResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *CardResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type CardResponse struct {
	Count    *int64      `json:"count,omitempty"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []Result    `json:"results"`
}

type Result struct {
	ID                                *int64        `json:"id,omitempty"`
	GameID                            *int64        `json:"game_id,omitempty"`
	Idol                              *Idol         `json:"idol,omitempty"`
	JapaneseCollection                *string       `json:"japanese_collection,omitempty"`
	TranslatedCollection              *string       `json:"translated_collection,omitempty"`
	Rarity                            *string       `json:"rarity,omitempty"`
	Attribute                         *string       `json:"attribute,omitempty"`
	JapaneseAttribute                 *string       `json:"japanese_attribute,omitempty"`
	IsPromo                           *bool         `json:"is_promo,omitempty"`
	PromoItem                         interface{}   `json:"promo_item"`
	PromoLink                         interface{}   `json:"promo_link"`
	ReleaseDate                       *string       `json:"release_date,omitempty"`
	JapanOnly                         *bool         `json:"japan_only,omitempty"`
	Event                             interface{}   `json:"event"`
	OtherEvent                        interface{}   `json:"other_event"`
	IsSpecial                         *bool         `json:"is_special,omitempty"`
	HP                                *int64        `json:"hp,omitempty"`
	MinimumStatisticsSmile            *int64        `json:"minimum_statistics_smile,omitempty"`
	MinimumStatisticsPure             *int64        `json:"minimum_statistics_pure,omitempty"`
	MinimumStatisticsCool             *int64        `json:"minimum_statistics_cool,omitempty"`
	NonIdolizedMaximumStatisticsSmile *int64        `json:"non_idolized_maximum_statistics_smile,omitempty"`
	NonIdolizedMaximumStatisticsPure  *int64        `json:"non_idolized_maximum_statistics_pure,omitempty"`
	NonIdolizedMaximumStatisticsCool  *int64        `json:"non_idolized_maximum_statistics_cool,omitempty"`
	IdolizedMaximumStatisticsSmile    *int64        `json:"idolized_maximum_statistics_smile,omitempty"`
	IdolizedMaximumStatisticsPure     *int64        `json:"idolized_maximum_statistics_pure,omitempty"`
	IdolizedMaximumStatisticsCool     *int64        `json:"idolized_maximum_statistics_cool,omitempty"`
	Skill                             *string       `json:"skill,omitempty"`
	JapaneseSkill                     *string       `json:"japanese_skill,omitempty"`
	SkillDetails                      *string       `json:"skill_details,omitempty"`
	JapaneseSkillDetails              interface{}   `json:"japanese_skill_details"`
	CenterSkill                       *string       `json:"center_skill,omitempty"`
	CenterSkillExtraType              *string       `json:"center_skill_extra_type,omitempty"`
	CenterSkillDetails                *string       `json:"center_skill_details,omitempty"`
	JapaneseCenterSkill               *string       `json:"japanese_center_skill,omitempty"`
	JapaneseCenterSkillDetails        *string       `json:"japanese_center_skill_details,omitempty"`
	CardImage                         *string       `json:"card_image"`
	CardIdolizedImage                 *string       `json:"card_idolized_image,omitempty"`
	EnglishCardImage                  interface{}   `json:"english_card_image"`
	EnglishCardIdolizedImage          interface{}   `json:"english_card_idolized_image"`
	RoundCardImage                    *string       `json:"round_card_image"`
	RoundCardIdolizedImage            *string       `json:"round_card_idolized_image,omitempty"`
	EnglishRoundCardImage             interface{}   `json:"english_round_card_image"`
	EnglishRoundCardIdolizedImage     interface{}   `json:"english_round_card_idolized_image"`
	VideoStory                        interface{}   `json:"video_story"`
	JapaneseVideoStory                interface{}   `json:"japanese_video_story"`
	WebsiteURL                        *string       `json:"website_url,omitempty"`
	NonIdolizedMaxLevel               *int64        `json:"non_idolized_max_level,omitempty"`
	IdolizedMaxLevel                  *int64        `json:"idolized_max_level,omitempty"`
	TransparentImage                  *string       `json:"transparent_image"`
	TransparentIdolizedImage          *string       `json:"transparent_idolized_image,omitempty"`
	CleanUr                           *string       `json:"clean_ur"`
	CleanUrIdolized                   *string       `json:"clean_ur_idolized,omitempty"`
	SkillUpCards                      []interface{} `json:"skill_up_cards"`
	UrPair                            *UrPair       `json:"ur_pair"`
	TotalOwners                       *int64        `json:"total_owners,omitempty"`
	TotalWishlist                     *int64        `json:"total_wishlist,omitempty"`
	RankingAttribute                  *int64        `json:"ranking_attribute,omitempty"`
	RankingRarity                     *int64        `json:"ranking_rarity,omitempty"`
	RankingSpecial                    *int64        `json:"ranking_special"`
}

type Idol struct {
	Note         *string `json:"note,omitempty"`
	School       *string `json:"school,omitempty"`
	Name         *string `json:"name,omitempty"`
	Year         *string `json:"year,omitempty"`
	Chibi        *string `json:"chibi,omitempty"`
	MainUnit     *string `json:"main_unit,omitempty"`
	JapaneseName *string `json:"japanese_name,omitempty"`
	ChibiSmall   *string `json:"chibi_small,omitempty"`
	SubUnit      *string `json:"sub_unit,omitempty"`
}

type UrPair struct {
	ReverseDisplayIdolized *bool `json:"reverse_display_idolized,omitempty"`
	ReverseDisplay         *bool `json:"reverse_display,omitempty"`
	Card                   *Card `json:"card,omitempty"`
}

type Card struct {
	Note           *string `json:"note,omitempty"`
	Attribute      *string `json:"attribute,omitempty"`
	RoundCardImage *string `json:"round_card_image,omitempty"`
	ID             *int64  `json:"id,omitempty"`
	Name           *string `json:"name,omitempty"`
}
