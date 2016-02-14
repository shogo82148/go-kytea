package kytea

// #cgo pkg-config: kytea
// #include <stdlib.h>
// #include "ckytea.h"
import "C"

type KyTeaWord struct {
	word *C.kytea_word_t
}

func (w KyTeaWord) Surface(util StringUtil) string {
	s := C.kytea_word_surface(w.word, util.util)
	defer C.kytea_std_string_destroy(s)
	return C.GoString(C.kytea_std_string_cstring(s))
}

func (w KyTeaWord) Tag(i, j int, util StringUtil) (string, float64) {
	tag := C.kytea_word_tag(w.word, C.int(i), C.int(j), util.util)
	defer C.kytea_std_string_destroy(tag.feature)
	return C.GoString(C.kytea_std_string_cstring(tag.feature)), float64(tag.score)
}

func (w KyTeaWord) TagsLen(i int) int {
	return int(C.kytea_word_tags_len(w.word, C.int(i)))
}
