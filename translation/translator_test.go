package translation_test

import (
	"github.com/jswizzy/hello-api/translation"
	"testing"
)

func TestTranslate(t *testing.T) {
	tt := map[string]struct {
		word     string
		language string
		want     string
	}{
		"English": {
			word:     "hello",
			language: "english",
			want:     "hello",
		},
		"German": {
			word:     "hello",
			language: "german",
			want:     "hallo",
		},
		"Finnish": {
			word:     "hello",
			language: "finnish",
			want:     "hei",
		},
		"Dutch": {
			word:     "hello",
			language: "dutch",
			want:     "",
		},
		"French": {
			word:     "hello",
			language: "french",
			want:     "bonjour",
		},
		"Missing": {
			word:     "bye",
			language: "dutch",
			want:     "",
		},
		"Wrong": {
			word:     "bye",
			language: "german",
			want:     "",
		},
		"Capital Word": {
			word:     "Hello",
			language: "german",
			want:     "hallo",
		},
		"Capital Language": {
			word:     "hello",
			language: "German",
			want:     "hallo",
		},
		"Space": {
			word:     "hello ",
			language: "german",
			want:     "hallo",
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			res := translation.Translate(tc.word, tc.language)
			if res != tc.want {
				t.Errorf(`expected "%s" but received "%s"`, tc.want, res)
			}
		})
	}

}
