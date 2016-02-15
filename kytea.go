package kytea

// #cgo pkg-config: kytea
// #include <stdlib.h>
// #include "ckytea.h"
import "C"
import (
	"errors"
	"unsafe"
)

type KyTea struct {
	kytea *C.kytea_t
}

func New() (KyTea, error) {
	kytea := C.kytea_new()
	if kytea == nil {
		return KyTea{}, errors.New("kytae is not created")
	}
	return KyTea{kytea: kytea}, nil
}

func (k KyTea) Destroy() {
	C.kytea_destroy(k.kytea)
}

func (k KyTea) ReadModel(path string) {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))
	C.kytea_read_model(k.kytea, cpath)
}

func (k KyTea) StringUtil() StringUtil {
	return StringUtil{
		util: C.kytea_get_string_util(k.kytea),
	}
}

func (k KyTea) Config() Config {
	return Config{
		config: C.kytea_get_config(k.kytea),
	}
}

func (k KyTea) CalculateWS(s Sentence) {
	C.kytea_calculate_ws(k.kytea, s.sentence)
}

func (k KyTea) CalculateTags(s Sentence, i int) {
	C.kytea_calculate_tags(k.kytea, s.sentence, C.int(i))
}

func (k KyTea) CalculateAllTags(s Sentence) {
	C.kytea_calculate_all_tags(k.kytea, s.sentence)
}

func (k KyTea) Parse(str string) ([]Word, error) {
	config := k.Config()
	util := k.StringUtil()
	sentence := util.NewSentence(str)
	defer sentence.Destroy()

	if config.DoWS() {
		k.CalculateWS(sentence)
	}
	if config.DoTags() {
		k.CalculateAllTags(sentence)
	}

	return sentence.Words(util), nil
}
