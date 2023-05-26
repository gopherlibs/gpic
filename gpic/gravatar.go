package gpic

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"net/mail"
	"net/url"
	"strconv"
	"strings"
)

const GR_HOSTNAME = "https://www.gravatar.com"
const GR_MAX_SIZE = 2048

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

type gravatar struct {
	avatarEmbed
	email          string
	emailHash      string
	defaultImage   string
	size           int16
	rating         rating
	disableDefault bool
}

func (this *gravatar) IsDefault() bool {

	notDefault, _ := checkURL(this)
	if notDefault {
		return false
	}

	return true
}

func (this *gravatar) SetDefault(urlS string) error {

	theURL, err := url.Parse(urlS)
	if err != nil {
		return err
	}

	this.defaultImage = theURL.String()

	return nil
}

func (this *gravatar) setDisableDefault(disable bool) {
	this.disableDefault = disable
}

func (this *gravatar) SetSize(size int16) error {

	if size == -1 || size > GR_MAX_SIZE {
		size = GR_MAX_SIZE
	}

	if size < -1 {
		return errors.New("Size cannot be a negative number under -1.")
	}

	this.size = size

	return nil
}

func (this *gravatar) URL() (*url.URL, error) {

	path := "/avatar/" + this.emailHash + ".jpg" + "?"

	v := url.Values{}
	if this.size != 0 {
		v.Add("size", strconv.Itoa(int(this.size)))
	}
	v.Add("rating", this.rating.String())

	if this.disableDefault {
		v.Add("d", "404")
	}

	if this.defaultImage != "" {
		v.Add("default", this.defaultImage)
	}

	return url.Parse(GR_HOSTNAME + path + v.Encode())
}

/*
 * Retrieve a Gravatar image by email address.
 */
func NewGravatar(email string) (*gravatar, error) {

	avatar := new(gravatar)

	avatar.disableDefault = false

	if !isValidEmail(email) {
		return nil, errors.New("Not a valid URL.")
	}

	emailHash, err := hashEmail(email)
	if err != nil {
		return nil, err
	}
	avatar.emailHash = emailHash
	avatar.email = email

	err = avatar.SetSize(80)
	if err != nil {
		return nil, err
	}

	avatar.rating = RatingG

	return avatar, nil

}

/* Original entrypoint into this module. Only supports Gravatar. This function
 * is deprecated and will go away before the 1.0 release. Instead, the new
 * NewAvatar function should be used instead.
 */
func NewImage(inputs ...string) (*gravatar, error) {

	i := new(gravatar)

	for idx, input := range inputs {

		i.disableDefault = false

		if !isValidEmail(input) {
			continue
		}

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

		if len(inputs) != idx+1 {
			i.disableDefault = true
			validURL, err := checkURL(i)
			if err != nil {
				return nil, err
			}

			if validURL {
				break
			}
		}
	}

	return i, nil

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

func isValidEmail(address string) bool {
	_, err := mail.ParseAddress(address)
	return err == nil
}
