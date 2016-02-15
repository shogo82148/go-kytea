package kytea

// #cgo pkg-config: kytea
// #include <stdlib.h>
// #include "ckytea.h"
import "C"
import "unsafe"

// StringUtil is the string utility class. This allows you to convert from
// the appropriate string encoding to Kytea's internal format
type StringUtil struct {
	util *C.kytea_string_util_t
}

// NewSentence maps a plain text string to a KyteaString, and creates a sentence object
func (u StringUtil) NewSentence(sentence string) Sentence {
	length := C.size_t(len(sentence))
	if sentence == "" {
		sentence = "dummy"
	}
	s := *(**C.char)(unsafe.Pointer(&sentence))
	return Sentence{
		sentence: C.kytea_string_util_new_sentence(u.util, s, length),
	}
}
