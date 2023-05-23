package gpic

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/mail"
	"net/url"
	"strconv"
	"strings"
)

const hostname = "https://www.gravatar.com"

type rating int

const (
	RatingG rating = iota
	RatingPG
	RatingR
	RatingX
)

func (r rating) String() string {
	return [...]string{"g", "pg", "r", "x"}[r]
}

type Image struct {
	email        	string
	emailHash    	string
	defaultImage 	string
	size         	uint16
	rating       	rating
	disableDefault	bool
}

type GithubUser struct {
	AvatarURL	string	`json:"avatar_url"`
}

func (this *Image) SetDefault(urlS string) error {

	theURL, err := url.Parse(urlS)
	if err != nil {
		return err
	}

	this.defaultImage = theURL.String()

	return nil
}

func (i *Image) SetSize(size uint16) error {

	if size == 0 {
		return errors.New("gravatar: image size cannot be 0px")
	} else if size > 2048 {
		return errors.New("gravatar: image size cannot be larger than 2048px")
	}

	i.size = size

	return nil
}

func (i *Image) URL() (*url.URL, error) {

	path := "/avatar/" + i.emailHash + ".jpg" + "?"

	v := url.Values{}
	v.Add("size", strconv.Itoa(int(i.size)))
	v.Add("rating", i.rating.String())

	if i.disableDefault {
		v.Add("d", "404")
	}

	if i.defaultImage != "" {
		v.Add("default", i.defaultImage)
	}

	return url.Parse(hostname + path + v.Encode())
}

// checkURL returns true when we want to use the image and false when we do not
func (i *Image) checkURL() (bool, error) {
	
	url, err := i.URL()
	if err != nil {
		return false, err
	}
	resp, err := http.Get(url.String())
	if err != nil {
		return false, err
	}

	if resp.StatusCode == 404 {
		return false, nil 
	} else if resp.StatusCode == 200 {
		return true, nil
	} else {
		return false, errors.New("gravatar: unexpected status code: " + resp.Status)
	}
}

func NewImage(inputs ...string) (*Image, error) {

	var i Image

	for idx, input := range inputs {

		i.disableDefault = false

		if isValidEmail(input) {
			emailHash, err := hashEmail(input)
			if err != nil {
				return nil, err
			}
			i.emailHash = emailHash
			i.email = input

			err = i.SetSize(80)
			if err != nil {
				return nil, err
			}

			i.rating = RatingG

			if len(inputs) != idx + 1 {
				i.disableDefault = true
				validURL, err := i.checkURL() 
				if err != nil {
					return nil, err	
				}

				if validURL {
					break 
				}
			} 
		}

		if strings.HasPrefix(input, "gh:") {
			
			ghUser := strings.TrimPrefix(input, "gh:")
			url := "https://api.github.com/users/" + ghUser

			resp, err := http.Get(url)
			if err != nil {
				return nil, err
			}

			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			
			var githubUser GithubUser 
			err = json.Unmarshal(body, &githubUser)
			if err != nil {
				return nil, err
			}

			fmt.Println(githubUser.AvatarURL) //Need to figure out what to do with this

		}
		
	}

	return &i, nil

}

func hashEmail(email string) (string, error) {

	if email == "" {
		return "", errors.New("gravatar: email address cannot be an empty string")
	}

	email = strings.TrimSpace(email)
	email = strings.ToLower(email)
	hash := md5.Sum([]byte(email))

	return hex.EncodeToString(hash[:]), nil
}

func isValidEmail(address string) (bool) {
	_, err := mail.ParseAddress(address)
	return err == nil 
}
