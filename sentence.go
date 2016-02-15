package kytea

// #cgo pkg-config: kytea
// #include <stdlib.h>
// #include "ckytea.h"
import "C"

type Sentence struct {
	sentence *C.kytea_sentence_t
}

func (s Sentence) Destroy() {
	C.kytea_sentence_destroy(s.sentence)
}

func (s Sentence) WordsLen() int {
	return int(C.kytea_sentence_words_len(s.sentence))
}

func (s Sentence) KyTeaWord(i int) KyTeaWord {
	return KyTeaWord{
		word: C.kytea_sentence_word_at(s.sentence, C.int(i)),
	}
}

func (s Sentence) Word(i int, util StringUtil) Word {
	return s.KyTeaWord(i).Word(util)
}

func (s Sentence) Words(util StringUtil) []Word {
	wordsLen := s.WordsLen()
	words := make([]Word, wordsLen)
	for i := 0; i < wordsLen; i++ {
		words[i] = s.KyTeaWord(i).Word(util)
	}
	return words
}
