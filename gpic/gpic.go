package gpic

import (
	"errors"
	"net/http"
	"net/url"
	"strings"
)

/*
 * An avatar is an image that represents a user. This is an interface that is
 * fulfilled by back provider. Current providers are Gravatar and GitHub.
 *
 * a size of '0' means default size, '-1' means max size.
 */
type Avatar interface {
	IsDefault() bool
	SetDefault(string) error
	setDisableDefault(bool)
	SetSize(int16) error
	URL() (*url.URL, error)
}

type avatarEmbed struct {
}

// checkURL returns true when we want to use the image and false when we do not
func checkURL(avatar Avatar) (bool, error) {

	url, err := avatar.URL()
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
		return false, errors.New("unexpected status code: " + resp.Status)
	}
}

/*
 * Returns a new avatar to be display in an app via an image URL. This
 * function accepts one or more 'inputs' that are to be processed by a
 * provider. The provider is determined by the format of the input. These are
 * the possibilities:
 *
 * `example@email.com` - Gravatar - email address
 * `ghp_someSortOfToken` - GitHub - API token
 * `gh:someUsername` - GitHub - username
 * `ghi:12345678` - GitHub - user ID
 * `glpat-someSortOfToken` - GitLab - API token
 * `gl:someUsername` - GitLab - username
 */
func NewAvatar(inputs ...string) (Avatar, error) {

	var avatar Avatar
	var err error
	var last bool

	// loop through each possible input
	for idx, input := range inputs {

		err = nil //reset

		if (idx + 1) < len(inputs) {
			last = false
		} else {
			last = true
		}

		// check for email address (for Gravatar)
		if isValidEmail(input) {

			avatar, err = NewGravatar(input)
			if err != nil {
				continue
			} else if !last {
				avatar.setDisableDefault(true)
			}

			if avatar.IsDefault() && !last {

				continue
			}

			break
		}

		// check for GitHub token
		if strings.HasPrefix(input, "ghp_") {

			avatar, err = NewGitHubAvatar("token", input)
			if err != nil {
				continue
			}

			break
		}

		// check for GitHub username
		result, found := strings.CutPrefix(input, "gh:")
		if found {

			avatar, err = NewGitHubAvatar("username", result)
			if err != nil {
				continue
			}

			break
		}

		// check for GitHub ID
		result, found = strings.CutPrefix(input, "ghi:")
		if found {

			avatar, err = NewGitHubAvatar("id", result)

			if err != nil {
				continue
			}

			break
		} 
		
		// check for GitLab token
		if strings.HasPrefix(input, "glpat-") {

			avatar, err = NewGitLabAvatar("token", input)
			if err != nil {
				continue
			}

			break
		}

		// check for GitLab username
		result, found = strings.CutPrefix(input, "gl:")
		if found {

			avatar, err = NewGitLabAvatar("username", result)
			if err != nil {
				continue
			}

			break

		}
		
		// check for GitLab id
		result, found = strings.CutPrefix(input, "gli:")
		if found {

			avatar, err = NewGitLabAvatar("id", result)
			if err != nil {
				continue
			}

			break

		} else {
			if (idx + 1) == len(inputs) {
				return nil, errors.New("None of the inputs were valid.")
			}
		}

	}

	return avatar, err

}
