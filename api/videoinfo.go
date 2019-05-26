package api

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const VIDEO_INFO_ENDPOINT = "http://api.ce.nicovideo.jp/nicoapi/v1/video.info?v="

type NicovideoVideoResponse struct {
	XMLName xml.Name `xml:"nicovideo_video_response" json:"xml_name"`
	Text    string   `xml:",chardata" json:"text"`
	Status  string   `xml:"status,attr" json:"status"`
	Video   struct {
		Text                      string `xml:",chardata" json:"text"`
		ID                        string `xml:"id" json:"id"`
		UserID                    string `xml:"user_id" json:"user_id"`
		Deleted                   string `xml:"deleted" json:"deleted"`
		Title                     string `xml:"title" json:"title"`
		Description               string `xml:"description" json:"description"`
		LengthInSeconds           string `xml:"length_in_seconds" json:"length_in_seconds"`
		ThumbnailURL              string `xml:"thumbnail_url" json:"thumbnail_url"`
		UploadTime                string `xml:"upload_time" json:"upload_time"`
		FirstRetrieve             string `xml:"first_retrieve" json:"first_retrieve"`
		DefaultThread             string `xml:"default_thread" json:"default_thread"`
		ViewCounter               string `xml:"view_counter" json:"view_counter"`
		MylistCounter             string `xml:"mylist_counter" json:"mylist_counter"`
		OptionFlagCommunity       string `xml:"option_flag_community" json:"option_flag_community"`
		OptionFlagNicowari        string `xml:"option_flag_nicowari" json:"option_flag_nicowari"`
		OptionFlagMiddleThumbnail string `xml:"option_flag_middle_thumbnail" json:"option_flag_middle_thumbnail"`
		OptionFlagDmcPlay         string `xml:"option_flag_dmc_play" json:"option_flag_dmc_play"`
		CommunityID               string `xml:"community_id" json:"community_id"`
		VitaPlayable              string `xml:"vita_playable" json:"vita_playable"`
		PpvVideo                  string `xml:"ppv_video" json:"ppv_video"`
		Permission                string `xml:"permission" json:"permission"`
		ProviderType              string `xml:"provider_type" json:"provider_type"`
		Options                   struct {
			Text           string `xml:",chardata" json:"text"`
			Mobile         string `xml:"mobile,attr" json:"mobile"`
			Sun            string `xml:"sun,attr" json:"sum"`
			LargeThumbnail string `xml:"large_thumbnail,attr" json:"large_thumbnail"`
			Adult          string `xml:"adult,attr" json:"adult"`
		} `xml:"options" json:"options"`
	} `xml:"video" json:"video"`
	Thread struct {
		Text        string `xml:",chardata" json:"text"`
		ID          string `xml:"id" json:"id"`
		NumRes      string `xml:"num_res" json:"numres"`
		Summary     string `xml:"summary" json:"summry"`
		CommunityID string `xml:"community_id" json:"community_id"`
		GroupType   string `xml:"group_type" json:"grouptype"`
	} `xml:"thread" json:"thread"`
	Tags struct {
		Text    string `xml:",chardata" json:"text"`
		TagInfo []struct {
			Text string `xml:",chardata" json:"text"`
			Tag  string `xml:"tag" json:"tag"`
			Area string `xml:"area" json:"area"`
		} `xml:"tag_info" json:"tag_info"`
	} `xml:"tags" json:"tags"`
}

func GetVideoInfo(videoID string) (*NicovideoVideoResponse, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", VIDEO_INFO_ENDPOINT+videoID, nil)
	if err != nil {
		log.Printf("[ERROR] failed create get video info request, videoID: %s\n", videoID)
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		log.Printf("[ERROR] failed get video info, videoID: %s\n", videoID)
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("[ERROR] responsed HTTP status code is %d, videoID: %s", res.StatusCode, videoID)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("[ERROR] failed read response body, videoID: %s\n", videoID)
		return nil, err
	}

	videoInfo := new(NicovideoVideoResponse)
	err = xml.Unmarshal(body, videoInfo)
	if err != nil {
		log.Printf("[ERROR] failed unmarshal xml, videoID: %s\n", videoID)
		return nil, err
	}

	return videoInfo, nil
}
