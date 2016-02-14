package kytea

// #cgo pkg-config: kytea
// #include <stdlib.h>
// #include "ckytea.h"
import "C"

type Sentence struct {
	sentence *C.kytea_sentence_t
}

func (s Sentence) Destory() {
	C.kytea_sentence_destroy(s.sentence)
}

func (s Sentence) WordsLen() int {
	return int(C.kytea_sentence_words_len(s.sentence))
}

func (s Sentence) Word(i int) Word {
	return Word{
		word: C.kytea_sentence_word_at(s.sentence, C.int(i)),
	}
}
