package kytea

// #cgo pkg-config: kytea
// #include <stdlib.h>
// #include "ckytea.h"
import "C"

type Word struct {
	Surface string
	Tags    [][]Tag
}

func (w Word) String() string {
	s := w.Surface
	if w.Tags != nil {
		for _, tag := range w.Tags {
			if tag != nil && len(tag) > 0 {
				s += "/" + tag[0].Feature
			}
		}
	}
	return s
}

type Tag struct {
	Feature string
	Score   float64
}

func (t Tag) String() string {
	return t.Feature
}

// KyTeaWord provides the interface for KyTea's internal format of words.
type KyTeaWord struct {
	word *C.kytea_word_t
}

// Surface returns the surface of the word.
func (w KyTeaWord) Surface(util StringUtil) string {
	s := C.kytea_word_surface(w.word, util.util)
	defer C.kytea_std_string_destroy(s)
	return C.GoString(C.kytea_std_string_cstring(s))
}

// Tag returns a particular tag and its score.
func (w KyTeaWord) Tag(i, j int, util StringUtil) (string, float64) {
	tag := C.kytea_word_tag(w.word, C.int(i), C.int(j), util.util)
	defer C.kytea_std_string_destroy(tag.feature)
	return C.GoString(C.kytea_std_string_cstring(tag.feature)), float64(tag.score)
}

// CandidateTagsLen is the number of i-th tags.
func (w KyTeaWord) CandidateTagsLen(i int) int {
	return int(C.kytea_word_candidate_tags_len(w.word, C.int(i)))
}

// TagsLen is the number of tags.
func (w KyTeaWord) TagsLen() int {
	return int(C.kytea_word_tags_len(w.word))
}

// Word converts the word from Kytea's internal format to Go struct.
func (w KyTeaWord) Word(util StringUtil) Word {
	surface := w.Surface(util)
	tagsLen := w.TagsLen()
	tags := make([][]Tag, tagsLen)
	for i := 0; i < tagsLen; i++ {
		candidateTagsLen := w.CandidateTagsLen(i)
		tmp := make([]Tag, candidateTagsLen)
		for j := 0; j < candidateTagsLen; j++ {
			tmp[j].Feature, tmp[j].Score = w.Tag(i, j, util)
		}
		tags[i] = tmp
	}
	return Word{
		Surface: surface,
		Tags:    tags,
	}
}
