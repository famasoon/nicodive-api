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
	XMLName xml.Name `xml:"nicovideo_video_response"`
	Text    string   `xml:",chardata"`
	Status  string   `xml:"status,attr"`
	Video   struct {
		Text                      string `xml:",chardata"`
		ID                        string `xml:"id"`
		UserID                    string `xml:"user_id"`
		Deleted                   string `xml:"deleted"`
		Title                     string `xml:"title"`
		Description               string `xml:"description"`
		LengthInSeconds           string `xml:"length_in_seconds"`
		ThumbnailURL              string `xml:"thumbnail_url"`
		UploadTime                string `xml:"upload_time"`
		FirstRetrieve             string `xml:"first_retrieve"`
		DefaultThread             string `xml:"default_thread"`
		ViewCounter               string `xml:"view_counter"`
		MylistCounter             string `xml:"mylist_counter"`
		OptionFlagCommunity       string `xml:"option_flag_community"`
		OptionFlagNicowari        string `xml:"option_flag_nicowari"`
		OptionFlagMiddleThumbnail string `xml:"option_flag_middle_thumbnail"`
		OptionFlagDmcPlay         string `xml:"option_flag_dmc_play"`
		CommunityID               string `xml:"community_id"`
		VitaPlayable              string `xml:"vita_playable"`
		PpvVideo                  string `xml:"ppv_video"`
		Permission                string `xml:"permission"`
		ProviderType              string `xml:"provider_type"`
		Options                   struct {
			Text           string `xml:",chardata"`
			Mobile         string `xml:"mobile,attr"`
			Sun            string `xml:"sun,attr"`
			LargeThumbnail string `xml:"large_thumbnail,attr"`
			Adult          string `xml:"adult,attr"`
		} `xml:"options"`
	} `xml:"video"`
	Thread struct {
		Text        string `xml:",chardata"`
		ID          string `xml:"id"`
		NumRes      string `xml:"num_res"`
		Summary     string `xml:"summary"`
		CommunityID string `xml:"community_id"`
		GroupType   string `xml:"group_type"`
	} `xml:"thread"`
	Tags struct {
		Text    string `xml:",chardata"`
		TagInfo []struct {
			Text string `xml:",chardata"`
			Tag  string `xml:"tag"`
			Area string `xml:"area"`
		} `xml:"tag_info"`
	} `xml:"tags"`
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
