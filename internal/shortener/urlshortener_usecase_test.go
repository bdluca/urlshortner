package shortener

import (
	"fmt"
	"testing"

	"github.com/brunodeluca/gophercises/urlshort/internal/shortener/shorturl"
	"github.com/stretchr/testify/assert"
)

func Test_useCaseURLShortner_RegisterURL(t *testing.T) {
	type fields struct {
		shortURLRepo shorturl.Repository
	}

	tests := []struct {
		name    string
		fields  fields
		input   string
		wantErr bool
		err     error
	}{
		{
			name:    "an invalid url should not be registered",
			input:   "aaaa",
			wantErr: true,
			err:     fmt.Errorf("invalid url"),
			fields: fields{
				shortURLRepo: testShortURLRepo{
					mockSave: func(url string) (shorturl.ShortURL, error) {
						return shorturl.ShortURL{}, nil
					},
				},
			},
		},
		{
			name:    "an invalid url should not be registered",
			input:   "http:/aaa",
			wantErr: true,
			err:     fmt.Errorf("invalid url"),
			fields: fields{
				shortURLRepo: testShortURLRepo{
					mockSave: func(url string) (shorturl.ShortURL, error) {
						return shorturl.ShortURL{}, nil
					},
				},
			},
		},
		{
			name:    "valid url should not return an error",
			input:   "http://google.com",
			wantErr: false,
			err:     nil,
			fields: fields{
				shortURLRepo: testShortURLRepo{
					mockSave: func(url string) (shorturl.ShortURL, error) {
						return shorturl.ShortURL{}, nil
					},
				},
			},
		},
		{
			name:    "failure to save url should return an error",
			input:   "http://google.com",
			wantErr: true,
			err:     fmt.Errorf("error saving url"),
			fields: fields{
				shortURLRepo: testShortURLRepo{
					mockSave: func(url string) (shorturl.ShortURL, error) {
						return shorturl.ShortURL{}, fmt.Errorf("error saving url")
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UseCaseURLShortener{ShortURLRepo: tt.fields.shortURLRepo}
			_, err := u.RegisterURL(tt.input)
			assert.True(t, (err != nil) == tt.wantErr, "expected an error return")
			assert.Equal(t, tt.err, err)
		})
	}
}

type testShortURLRepo struct {
	mockSave func(url string) (shorturl.ShortURL, error)
	mockGet  func(id string) (shorturl.ShortURL, error)
}

func (t testShortURLRepo) Save(url string) (shorturl.ShortURL, error) {
	return t.mockSave(url)
}

func (t testShortURLRepo) Get(id string) (shorturl.ShortURL, error) {
	return t.mockGet(id)
}
