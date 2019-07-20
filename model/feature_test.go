package model

import (
	"reflect"
	"testing"

	"github.com/buckhx/gobert/tokenize"
	"github.com/buckhx/gobert/vocab"
)

func TestSequenceFeature(t *testing.T) {
	voc := vocab.New([]string{"[CLS]", "[SEP]", "the", "dog", "is", "hairy", "."})
	tkz := tokenize.NewTokenizer(voc)
	for _, test := range []struct {
		text    string
		feature Feature
	}{
		// TODO more tests, but this one covers some good edge cases
		{"the dog is hairy. ||| the ||| a dog is hairy", Feature{
			ID:       0,
			Tokens:   []string{"[CLS]", "the", "dog", "[SEP]", "the", "[SEP]", "[UNK]", "[SEP]"},
			TokenIDs: []int32{0, 2, 3, 1, 2, 1, -1, 1},
			Mask:     []int{1, 1, 1, 1, 1, 1, 1, 1},
			TypeIDs:  []int{0, 0, 0, 0, 1, 1, 2, 2},
		}},
	} {
		f := SequenceFeature(tkz, 8, test.text)
		if !reflect.DeepEqual(f, test.feature) {
			t.Errorf("Invalid Sequence Feature - Want: %+v, Got: %+v", test.feature, f)
		}
	}
}

func Test_sequenceTruncate(t *testing.T) {
	for _, test := range []struct {
		seqs   [][]string
		len    int
		tokens [][]string
	}{
		{nil, 1, nil},
		{[][]string{}, 1, [][]string{}},
		{[][]string{{"a1"}, {"b1"}, {"c1", "c2"}}, -1, [][]string{{}, {}, {}}},
		{[][]string{{"a1"}, {"b1"}, {"c1", "c2"}}, 0, [][]string{{}, {}, {}}},
		{[][]string{{"a1"}, {"b1"}, {"c1", "c2"}}, 1, [][]string{{"a1"}, {}, {}}},
		{[][]string{{"a1"}, {"b1"}, {"c1", "c2"}}, 3, [][]string{{"a1"}, {"b1"}, {"c1"}}},
		{[][]string{{"a1"}, {"b1"}, {"c1", "c2"}}, 10, [][]string{{"a1"}, {"b1"}, {"c1", "c2"}}},
	} {
		toks := truncate(test.seqs, test.len)
		if !reflect.DeepEqual(toks, test.tokens) {
			t.Errorf("Invalid Sequence Truncate - Want: %+v, Got: %+v", test.tokens, toks)
		}
	}
}
