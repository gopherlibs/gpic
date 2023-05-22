package gpic

import "testing"

func TestNewImage(t *testing.T) {

	email := "Ricardo@Feliciano.Tech"
	emailHash := "f6d625c59c19ea57fe2c3d7968a56f29"

	i, err := NewImage(email)
	if err != nil {
		t.Errorf("Failed creating new image.")
	}

	if i.emailHash != emailHash {
		t.Errorf("got hash: %q; expected: %q", i.emailHash, emailHash)
	}
}

func TestImageURL(t *testing.T) {

	samples := []struct {
		email        string
		emailHash    string
		defaultImage string
		rating       rating
		size         uint16
		expectedURL  string
	}{
		{"Ricardo@Feliciano.Tech", "f6d625c59c19ea57fe2c3d7968a56f29", "", 0, 0, "https://www.gravatar.com/avatar/f6d625c59c19ea57fe2c3d7968a56f29.jpg?rating=g&size=80"},
		{"Ricardo@Feliciano.Tech", "f6d625c59c19ea57fe2c3d7968a56f29", "", 0, 50, "https://www.gravatar.com/avatar/f6d625c59c19ea57fe2c3d7968a56f29.jpg?rating=g&size=50"},
		{"Ricardo@Feliciano.Tech", "f6d625c59c19ea57fe2c3d7968a56f29", "", RatingR, 50, "https://www.gravatar.com/avatar/f6d625c59c19ea57fe2c3d7968a56f29.jpg?rating=r&size=50"},
		{"Ricardo@Feliciano.Tech", "f6d625c59c19ea57fe2c3d7968a56f29", "", RatingR, 3000, "https://www.gravatar.com/avatar/f6d625c59c19ea57fe2c3d7968a56f29.jpg?rating=r&size=80"},
		{"Ricardo@Feliciano.Tech", "f6d625c59c19ea57fe2c3d7968a56f29", "identicon", 0, 0, "https://www.gravatar.com/avatar/f6d625c59c19ea57fe2c3d7968a56f29.jpg?default=identicon&rating=g&size=80"},
		{"Ricardo@Feliciano.Tech", "f6d625c59c19ea57fe2c3d7968a56f29", "retro", RatingPG, 2000, "https://www.gravatar.com/avatar/f6d625c59c19ea57fe2c3d7968a56f29.jpg?default=retro&rating=pg&size=2000"},
	}

	for _, sample := range samples {

		i, err := NewImage(sample.email)
		if err != nil {
			t.Errorf("Failed creating new image.")
		}

		i.defaultImage = sample.defaultImage
		i.SetSize(sample.size)
		i.rating = sample.rating

		if url, _ := i.URL(); url.String() != sample.expectedURL {
			t.Errorf(" got url: %q; expected: %q", url, sample.expectedURL)
		}
	}
}
