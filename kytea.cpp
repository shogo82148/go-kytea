// implementation for KyTea C Interface

#include <string>
#include <vector>

#include <kytea/kytea.h>
#include <kytea/kytea-struct.h>
#include <kytea/string-util.h>

#include "ckytea.h"

using namespace std;
using namespace kytea;


// functions for handling Kytea

kytea_t *kytea_new(void) {
    return reinterpret_cast<kytea_t*>(new kytea::Kytea());
}

void kytea_destroy(kytea_t *kytea) {
    delete reinterpret_cast<kytea::Kytea*>(kytea);
}

void kytea_read_model(kytea_t *kytea, const char *path) {
    kytea::Kytea *k = reinterpret_cast<kytea::Kytea*>(kytea);
    k->readModel(path);
}

kytea_string_util_t *kytea_get_string_util(kytea_t *kytea) {
    kytea::Kytea *k = reinterpret_cast<kytea::Kytea*>(kytea);
    return reinterpret_cast<kytea_string_util_t*>(k->getStringUtil());
}

kytea_config_t *kytea_get_config(kytea_t *kytea) {
    kytea::Kytea *k = reinterpret_cast<kytea::Kytea*>(kytea);
    return reinterpret_cast<kytea_config_t*>(k->getConfig());
}

void kytea_calculate_ws(kytea_t *kytea, kytea_sentence_t *sentence) {
    Kytea *k = reinterpret_cast<Kytea*>(kytea);
    KyteaSentence *s = reinterpret_cast<KyteaSentence*>(sentence);
    k->calculateWS(*s);
}

void kytea_calculate_tags(kytea_t *kytea, kytea_sentence_t *sentence, int i) {
    Kytea *k = reinterpret_cast<Kytea*>(kytea);
    KyteaSentence *s = reinterpret_cast<KyteaSentence*>(sentence);
    k->calculateTags(*s, i);
}

void kytea_calculate_all_tags(kytea_t *kytea, kytea_sentence_t *sentence) {
    Kytea *k = reinterpret_cast<Kytea*>(kytea);
    KyteaSentence *s = reinterpret_cast<KyteaSentence*>(sentence);
    KyteaConfig* config = k->getConfig();
    for(int i = 0; i < config->getNumTags(); i++)
        k->calculateTags(*s, i);
}


// functions for handling KyteaStringUtil

kytea_sentence_t *kytea_string_util_new_sentence(kytea_string_util_t *util_t, const char *input, size_t length) {
    string str(input, length);
    StringUtil* util = reinterpret_cast<StringUtil*>(util_t);
    KyteaString surface_string(util->mapString(str));
    KyteaSentence *sentence = new KyteaSentence(surface_string, util->normalize(surface_string));
    return reinterpret_cast<kytea_sentence_t*>(sentence);
}


// functions for handling KytaeConfig

void kytea_config_parse_run_command_line(kytea_config_t *config, int argc, const char **argv) {
    KyteaConfig *c = reinterpret_cast<KyteaConfig*>(config);
    c->parseRunCommandLine(argc, argv);
}

kytea_std_string_t *kytea_config_get_model_file(kytea_config_t *config) {
    KyteaConfig *c = reinterpret_cast<KyteaConfig*>(config);
    return reinterpret_cast<kytea_std_string_t*>(new string(c->getModelFile()));
}

void kytea_config_set_model_file(kytea_config_t *config, const char *path) {
    KyteaConfig *c = reinterpret_cast<KyteaConfig*>(config);
    c->setModelFile(path);
}

int kytea_config_get_do_ws(kytea_config_t *config) {
    KyteaConfig *c = reinterpret_cast<KyteaConfig*>(config);
    return c->getDoWS();
}

void kytea_config_set_do_ws(kytea_config_t *config, int flag) {
    KyteaConfig *c = reinterpret_cast<KyteaConfig*>(config);
    c->setDoWS(flag);
}

kytea_std_string_t *kytea_config_get_ws_constraint(kytea_config_t *config) {
    KyteaConfig *c = reinterpret_cast<KyteaConfig*>(config);
    return reinterpret_cast<kytea_std_string_t*>(new string(c->getWsConstraint()));
}

void kytea_config_set_ws_constraint(kytea_config_t *config, const char *str, size_t length) {
    KyteaConfig *c = reinterpret_cast<KyteaConfig*>(config);
    c->setWsConstraint(string(str, length));
}

int kytea_config_get_do_tags(kytea_config_t *config) {
    KyteaConfig *c = reinterpret_cast<KyteaConfig*>(config);
    return c->getDoTags() ? 1 : 0;
}

void kytea_config_set_do_tags(kytea_config_t *config, int flag) {
    KyteaConfig *c = reinterpret_cast<KyteaConfig*>(config);
    c->setDoTags(flag);
}

int kytea_config_get_do_tag(kytea_config_t *config, int i) {
    KyteaConfig *c = reinterpret_cast<KyteaConfig*>(config);
    return c->getDoTag(i) ? 1 : 0;
}

void kytea_config_set_do_tag(kytea_config_t *config, int i, int flag) {
    KyteaConfig *c = reinterpret_cast<KyteaConfig*>(config);
    c->setDoTag(i, flag);
}

int kytea_config_get_do_unk(kytea_config_t *config) {
    KyteaConfig *c = reinterpret_cast<KyteaConfig*>(config);
    return c->getDoUnk();
}

void kytea_config_set_do_unk(kytea_config_t *config, int flag) {
    KyteaConfig *c = reinterpret_cast<KyteaConfig*>(config);
    c->setDoUnk(flag);
}

int kytea_config_get_num_tags(kytea_config_t *config) {
    KyteaConfig *c = reinterpret_cast<KyteaConfig*>(config);
    return c->getNumTags();
}

void kytea_config_set_num_tags(kytea_config_t *config, int num) {
    KyteaConfig *c = reinterpret_cast<KyteaConfig*>(config);
    c->setNumTags(num);
}

int kytea_config_get_tag_max(kytea_config_t *config) {
    KyteaConfig *c = reinterpret_cast<KyteaConfig*>(config);
    return c->getTagMax();
}

void kytea_config_set_tag_max(kytea_config_t *config, int num) {
    KyteaConfig *c = reinterpret_cast<KyteaConfig*>(config);
    c->setTagMax(num);
}

kytea_std_string_t *kytea_config_get_unk_tag(kytea_config_t *config) {
    KyteaConfig *c = reinterpret_cast<KyteaConfig*>(config);
    return reinterpret_cast<kytea_std_string_t*>(new string(c->getUnkTag()));
}

void kytea_config_set_unk_tag(kytea_config_t *config, const char *str, size_t length) {
    KyteaConfig *c = reinterpret_cast<KyteaConfig*>(config);
    c->setUnkTag(string(str, length));
}

kytea_std_string_t *kytea_config_get_default_tag(kytea_config_t *config) {
    KyteaConfig *c = reinterpret_cast<KyteaConfig*>(config);
    return reinterpret_cast<kytea_std_string_t*>(new string(c->getDefaultTag()));
}

void kytea_config_set_default_tag(kytea_config_t *config, const char *str, size_t length) {
    KyteaConfig *c = reinterpret_cast<KyteaConfig*>(config);
    c->setDefaultTag(string(str, length));
}

int kytea_config_get_unk_beam(kytea_config_t *config) {
    KyteaConfig *c = reinterpret_cast<KyteaConfig*>(config);
    return c->getUnkBeam();
}

void kytea_config_set_unk_beam(kytea_config_t *config, int num) {
    KyteaConfig *c = reinterpret_cast<KyteaConfig*>(config);
    c->setUnkBeam(num);
}

// functions for handling KyteaSentence

void kytea_sentence_destroy(kytea_sentence_t *sentence) {
    delete reinterpret_cast<KyteaSentence*>(sentence);
}

size_t kytea_sentence_words_len(kytea_sentence_t *sentence) {
    return reinterpret_cast<KyteaSentence*>(sentence)->words.size();
}

kytea_word_t *kytea_sentence_word_at(kytea_sentence_t *sentence, int i) {
    return reinterpret_cast<kytea_word_t*>(&(reinterpret_cast<KyteaSentence*>(sentence)->words[i]));
}


// functions for handling KyteaWord

kytea_std_string_t *kytea_word_surface(kytea_word_t *word, kytea_string_util_t *util) {
    StringUtil *u = reinterpret_cast<StringUtil*>(util);
    KyteaWord *w = reinterpret_cast<KyteaWord*>(word);
    string *str = new string(u->showString(w->surface));
    return reinterpret_cast<kytea_std_string_t*>(str);
}

kytea_tag_t kytea_word_tag(kytea_word_t *word, int i, int j, kytea_string_util_t *util) {
    StringUtil *u = reinterpret_cast<StringUtil*>(util);
    KyteaWord *w = reinterpret_cast<KyteaWord*>(word);
    KyteaTag &tag = w->tags[i][j];
    string *str = new string(u->showString(tag.first));
    kytea_tag_t ret = {reinterpret_cast<kytea_std_string_t*>(str), tag.second};
    return ret;
}

size_t kytea_word_candidate_tags_len(kytea_word_t *word, int i) {
    KyteaWord *w = reinterpret_cast<KyteaWord*>(word);
    return w->tags[i].size();
}

size_t kytea_word_tags_len(kytea_word_t *word) {
    KyteaWord *w = reinterpret_cast<KyteaWord*>(word);
    return w->tags.size();
}


// functions for handling std::string

void kytea_std_string_destroy(kytea_std_string_t *str) {
    delete reinterpret_cast<string*>(str);
}

const char *kytea_std_string_cstring(kytea_std_string_t *str) {
    return reinterpret_cast<string*>(str)->c_str();
}
