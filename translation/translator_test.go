package translation_test

import (
	"github.com/jswizzy/hello-api/translation"
	"testing"
)

func TestTranslate(t *testing.T) {
	// Arrange
	tt := []struct {
		Word        string
		Language    string
		Translation string
	}{
		{ //<2>
			Word:        "hello",
			Language:    "english",
			Translation: "hello",
		},
		{
			Word:        "hello",
			Language:    "german",
			Translation: "hallo",
		},
		{
			Word:        "hello",
			Language:    "finnish",
			Translation: "hei",
		},
		{ // <1>
			Word:        "bye",
			Language:    "dutch",
			Translation: "",
		},
		{
			Word:        "hello",
			Language:    "dutch",
			Translation: "",
		},
		{ // <1>
			Word:        "bye",
			Language:    "german",
			Translation: "",
		},
		{
			Word:        "hello",
			Language:    "German",
			Translation: "hallo",
		},
		{
			Word:        "Hello",
			Language:    "german",
			Translation: "hallo",
		},
		{
			Word:        "hello ",
			Language:    "german",
			Translation: "hallo",
		},
		{
			Word:        "hello",
			Language:    "french",
			Translation: "bonjour",
		},
	}
	underTest := translation.NewStaticService()
	for _, test := range tt {
		// Act
		res := underTest.Translate(test.Word, test.Language)

		// Assert
		if res != test.Translation {
			t.Errorf(
				`expected "%s" to be "%s" from "%s" but received "%s"`,
				test.Word, test.Language, test.Translation, res)
		}
	}
}
