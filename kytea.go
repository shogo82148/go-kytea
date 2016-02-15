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

func (k KyTea) ReadModel(path string) error {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))
	err := C.kytea_read_model(k.kytea, cpath)
	if err != nil {
		defer C.kytea_std_string_destroy(err)
		return errors.New(C.GoString(C.kytea_std_string_cstring(err)))
	}
	return nil
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

func (k KyTea) CalculateWS(s Sentence) error {
	err := C.kytea_calculate_ws(k.kytea, s.sentence)
	if err != nil {
		defer C.kytea_std_string_destroy(err)
		return errors.New(C.GoString(C.kytea_std_string_cstring(err)))
	}
	return nil
}

func (k KyTea) CalculateTags(s Sentence, i int) error {
	err := C.kytea_calculate_tags(k.kytea, s.sentence, C.int(i))
	if err != nil {
		defer C.kytea_std_string_destroy(err)
		return errors.New(C.GoString(C.kytea_std_string_cstring(err)))
	}
	return nil
}

func (k KyTea) CalculateAllTags(s Sentence) error {
	err := C.kytea_calculate_all_tags(k.kytea, s.sentence)
	if err != nil {
		defer C.kytea_std_string_destroy(err)
		return errors.New(C.GoString(C.kytea_std_string_cstring(err)))
	}
	return nil
}

func (k KyTea) Parse(str string) ([]Word, error) {
	config := k.Config()
	util := k.StringUtil()
	sentence := util.NewSentence(str)
	defer sentence.Destroy()

	if config.DoWS() {
		err := k.CalculateWS(sentence)
		if err != nil {
			return nil, err
		}
	}
	if config.DoTags() {
		err := k.CalculateAllTags(sentence)
		if err != nil {
			return nil, err
		}
	}

	return sentence.Words(util), nil
}
