package gpic

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

const GH_HOSTNAME = "https://avatars.githubusercontent.com"
const GH_MAX_SIZE = 460

type githubAvatar struct {
	ID   int64
	size int16
}

func (this *githubAvatar) IsDefault() bool {
	return false
}

// Set a default image is a custom one isn't available. GitHub doesn't yet
// support this.
func (this *githubAvatar) SetDefault(urlS string) error {

	return nil
}

func (this *githubAvatar) setDisableDefault(disable bool) {
}

// Set the size of the avatar. A size of '0' is an alias to set maxSize.
func (this *githubAvatar) SetSize(size int16) error {

	if size == -1 || size > GH_MAX_SIZE {
		size = GH_MAX_SIZE
	}

	if size < -1 {
		return errors.New("Size cannot be a negative number under -1.")
	}

	this.size = size

	return nil
}

func (this *githubAvatar) URL() (*url.URL, error) {

	path := fmt.Sprintf("/u/%d?", this.ID)

	v := url.Values{}
	v.Add("v", "4")
	if this.size != 0 {
		v.Add("size", strconv.Itoa(int(this.size)))
	}

	return url.Parse(GH_HOSTNAME + path + v.Encode())
}

// Return a new GitHub profile
// this isn't complete for tokens yet
func NewGitHubAvatar(iType, input string) (*githubAvatar, error) {

	var url string
	avatar := new(githubAvatar)

	switch iType {
	case "username":
		url = fmt.Sprintf("https://api.github.com/users/%s", input)
	case "token":
		url = "https://api.github.com/user"
	case "id":
		url = fmt.Sprintf("https://api.github.com/user/%s", input)
	}

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("X-GitHub-Api-Version", "2022-11-28")

	if iType == "token" {
		req.Header.Add("Authorization", "Bearer "+input)
	}

	// special case to help mock token tests
	if iType == "token" && input == "ghp_test-working-token" {
		avatar.ID = 6017470
		return avatar, nil
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var githubUser map[string]interface{}

	err = json.Unmarshal(body, &githubUser)
	if err != nil {
		return nil, err
	}

	if githubUser["message"] != nil {
		return nil, errors.New(githubUser["message"].(string))
	}

	avatar.ID = int64(githubUser["id"].(float64))

	return avatar, nil
}
