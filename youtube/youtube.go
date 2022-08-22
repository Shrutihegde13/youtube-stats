package youtube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Response struct {
	Kind  string  `json:"kind"`
	Items []Items `json:"items"`
}

type Items struct {
	Kind  string `json:"kind"`
	Id    string `json:"id"`
	Stats Stats  `json:"statistics"`
}

type Stats struct {
	Views       string `json:"viewCount"`
	Subscribers string `json:"subscriberCount"`
	Videos      string `json:"videoCount"`
}

func GetSubscribers() (Items, error) {
	var response Response

	req, err := http.NewRequest("GET", "https://www.googleapis.com/youtube/v3/channels", nil)

	if err != nil {
		fmt.Println(err)
		return Items{}, err
	}

	q := req.URL.Query()

	q.Add("key", os.Getenv("YOUTUBE_KEY"))
	q.Add("id", os.Getenv("CHANNEL_ID"))
	q.Add("part", "statistics")

	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return Items{}, err
	}

	defer resp.Body.Close()

	fmt.Println("Response Status: ", resp.Status)

	body, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &response)

	if err != nil {
		return Items{}, err
	}

	return response.Items[0], nil

}
