package kytea

// #cgo pkg-config: kytea
// #include <stdlib.h>
// #include "ckytea.h"
import "C"

// Sentence provides the interface for KyTea's internal format of sentences.
type Sentence struct {
	sentence *C.kytea_sentence_t
}

func (s Sentence) Destroy() {
	C.kytea_sentence_destroy(s.sentence)
}

// WordsLen returns the number of words in the sentence.
func (s Sentence) WordsLen() int {
	return int(C.kytea_sentence_words_len(s.sentence))
}

// KyTeaWord gets the information of the i-th word as KyTea's internal format.
func (s Sentence) KyTeaWord(i int) KyTeaWord {
	return KyTeaWord{
		word: C.kytea_sentence_word_at(s.sentence, C.int(i)),
	}
}

// KyTeaWord gets the information of the i-th word as Go struct.
func (s Sentence) Word(i int, util StringUtil) Word {
	return s.KyTeaWord(i).Word(util)
}

// Words convers the sentence from Kytea's internal format to Go struct.
func (s Sentence) Words(util StringUtil) []Word {
	wordsLen := s.WordsLen()
	words := make([]Word, wordsLen)
	for i := 0; i < wordsLen; i++ {
		words[i] = s.KyTeaWord(i).Word(util)
	}
	return words
}
