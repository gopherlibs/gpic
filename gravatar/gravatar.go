package gravatar

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
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
	email        string
	emailHash    string
	defaultImage string
	size         uint16
	rating       rating
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

	if i.defaultImage != "" {
		v.Add("default", url.QueryEscape(i.defaultImage))
	}

	return url.Parse(hostname + path + v.Encode())
}

func NewImage(email string) (*Image, error) {

	var i Image

	emailHash, err := hashEmail(email)
	if err != nil {
		return nil, err
	}
	i.emailHash = emailHash
	i.email = email

	err = i.SetSize(80)
	if err != nil {
		return nil, err
	}

	i.rating = RatingG

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
