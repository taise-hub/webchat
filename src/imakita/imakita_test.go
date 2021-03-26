package imakita

import (
	"net/http"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetKeyPhrase(t *testing.T) {
	type args struct {
		format 	 string
		sentence string
	}
	tests := map[string]struct {
		args args
	}{
		"キーワードをmap[string]uint型として抽出できる": {
			args: args {
				format: "json",
				sentence: "いづれの御時にか女御更衣あまた侍ひ給ひけるなかに、いとやむごとなき際にはあらぬが、すぐれてときめき給うありけり。",
			},
		},
	}
	for tName, test := range tests {
		t.Run(tName, func(t *testing.T) {
			keyWords, err := getKeyPhrase(test.args.format, test.args.sentence)
			if err != nil {
				t.Fatalf("%s", err)
			}
			assert.NotNil(t, keyWords)
		})
	}
}

func TestGetSentence(t *testing.T) {
	type args struct {
		keyWord string
	}
	tests := map[string]struct {
		args args
	} {
		"キーワードを与えて文章を取得することができる。": {
			args: args{
				keyWord: "ずんだ餅",
			},
		},
	}
	client := new(http.Client)
	for tName, test := range tests {
		t.Run(tName, func(t *testing.T) {
			sentence, err := getSentence(test.args.keyWord, client)
			if err != nil {
				t.Fatalf("%s", err)
			}
			assert.NotNil(t, sentence)
		})
	}
}

func TestImakita(t *testing.T) {
	type args struct {
		sentence string
	}
	tests := map[string]struct {
		args args
	}{
		"今北産業できる！！！！": {
			args: args {
				sentence: "いづれの御時にか女御更衣あまた侍ひ給ひけるなかに、いとやむごとなき際にはあらぬが、すぐれてときめき給うありけり。",
			},
		},
	}
	for tName, test := range tests {
		t.Run(tName, func(t *testing.T) {
			imakita, err := Imakita(test.args.sentence)
			if err != nil {
				t.Fatalf("%s",err)
			}
			assert.NotNil(t, imakita)
		})
	}
}