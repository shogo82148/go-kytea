package kytea

// #cgo pkg-config: kytea
// #include <stdlib.h>
// #include "ckytea.h"
import "C"
import "unsafe"

type Config struct {
	config *C.kytea_config_t
}

func (c Config) ParseRunCommandLine(args []string) {
	opts := make([]*C.char, 0, len(args))
	for _, arg := range args {
		opt := C.CString(arg)
		defer C.free(unsafe.Pointer(opt))
		opts = append(opts, opt)
	}
	C.kytea_config_parse_run_command_line(c.config, C.int(len(opts)), (**C.char)(&opts[0]))
}

// ModelFile returns the model file to use when analyzing text.
// Default values is "$INSTALL_DIRECTORY/share/kytea/model.bin".
func (c Config) ModelFile() string {
	path := C.kytea_config_get_model_file(c.config)
	defer C.kytea_std_string_destroy(path)
	return C.GoString(C.kytea_std_string_cstring(path))
}

func (c Config) SetModelFile(path string) {
	s := C.CString(path)
	defer C.free(unsafe.Pointer(s))
	C.kytea_config_set_model_file(c.config, s)
}

// DoWS is a switch for doing the word segmentation.
// KyTea.Parse skips the word segmentation if DoWS is false.
func (c Config) DoWS() bool {
	return C.kytea_config_get_do_ws(c.config) != 0
}

func (c Config) SetDoWS(flag bool) {
	if flag {
		C.kytea_config_set_do_ws(c.config, 1)
	} else {
		C.kytea_config_set_do_ws(c.config, 0)
	}
}

// WsConstraint is a character type that do not segment.
func (c Config) WsConstraint() string {
	s := C.kytea_config_get_ws_constraint(c.config)
	defer C.kytea_std_string_destroy(s)
	return C.GoString(C.kytea_std_string_cstring(s))
}

func (c Config) SetWsConstraint(str string) {
	length := C.size_t(len(str))
	if str == "" {
		str = "dummy"
	}
	s := *(**C.char)(unsafe.Pointer(&str))
	C.kytea_config_set_ws_constraint(c.config, s, length)
}

// DoTags is a switch for tagging.
// KyTea.Parse skips tagging if DoTags is false.
func (c Config) DoTags() bool {
	return C.kytea_config_get_do_tags(c.config) != 0
}

func (c Config) SetDoTags(flag bool) {
	if flag {
		C.kytea_config_set_do_tags(c.config, 1)
	} else {
		C.kytea_config_set_do_tags(c.config, 0)
	}
}

// DoTag is a switch for tagging a particular tag.
// For example, SetDoTag(1, false) will skip the first tag.
func (c Config) DoTag(i int) bool {
	return C.kytea_config_get_do_tag(c.config, C.int(i)) != 0
}

func (c Config) SetDoTag(i int, flag bool) {
	if flag {
		C.kytea_config_set_do_tag(c.config, C.int(i), 1)
	} else {
		C.kytea_config_set_do_tag(c.config, C.int(i), 0)
	}
}

// DoUnk is a switch for estimating the pronunciation of unkown words.
func (c Config) DoUnk() bool {
	return C.kytea_config_get_do_tags(c.config) != 0
}

func (c Config) SetDoUnk(flag bool) {
	if flag {
		C.kytea_config_set_do_unk(c.config, 1)
	} else {
		C.kytea_config_set_do_unk(c.config, 0)
	}
}

// NumTags the number of tags.
func (c Config) NumTags() int {
	return int(C.kytea_config_get_num_tags(c.config))
}

func (c Config) SetNumTags(num int) {
	C.kytea_config_set_num_tags(c.config, C.int(num))
}

// TagMax is the maximum number of tags to print for one word.
// default is 3. 0 implies no limit.
func (c Config) TagMax() int {
	return int(C.kytea_config_get_tag_max(c.config))
}

func (c Config) SetTagMax(num int) {
	C.kytea_config_set_tag_max(c.config, C.int(num))
}

// UnkTag is a tag to append to indicate words not in the dictionary.
func (c Config) UnkTag() string {
	s := C.kytea_config_get_unk_tag(c.config)
	defer C.kytea_std_string_destroy(s)
	return C.GoString(C.kytea_std_string_cstring(s))
}

func (c Config) SetUnkTag(str string) {
	length := C.size_t(len(str))
	if str == "" {
		str = "dummy"
	}
	s := *(**C.char)(unsafe.Pointer(&str))
	C.kytea_config_set_unk_tag(c.config, s, length)
}

// DefaultTag is tag for words that cannot be given any tag.
func (c Config) DefaultTag() string {
	s := C.kytea_config_get_default_tag(c.config)
	defer C.kytea_std_string_destroy(s)
	return C.GoString(C.kytea_std_string_cstring(s))
}

func (c Config) SetDefaultTag(str string) {
	length := C.size_t(len(str))
	if str == "" {
		str = "dummy"
	}
	s := *(**C.char)(unsafe.Pointer(&str))
	C.kytea_config_set_default_tag(c.config, s, length)
}

// UnkBeam is the width of the beam to use in beam search for unknown words.
func (c Config) UnkBeam() int {
	return int(C.kytea_config_get_unk_beam(c.config))
}

func (c Config) SetUnkBeam(num int) {
	C.kytea_config_set_unk_beam(c.config, C.int(num))
}
