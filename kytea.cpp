// implementation for KyTea C Interface

#include <string>
#include <vector>

#include <kytea/kytea.h>
#include <kytea/kytea-struct.h>
#include <kytea/string-util.h>

#include "ckytea.h"

using namespace std;
using namespace kytea;

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

kytea_sentence_t *kytea_string_util_new_sentence(kytea_string_util_t *util_t, const char *input, size_t length) {
    string str(input, length);
    StringUtil* util = reinterpret_cast<StringUtil*>(util_t);
    KyteaString surface_string(util->mapString(str));
    KyteaSentence *sentence = new KyteaSentence(surface_string, util->normalize(surface_string));
    return reinterpret_cast<kytea_sentence_t*>(sentence);
}

void kytea_sentence_destroy(kytea_sentence_t *sentence) {
    delete reinterpret_cast<KyteaSentence*>(sentence);
}

size_t kytea_sentence_words_len(kytea_sentence_t *sentence) {
    return reinterpret_cast<KyteaSentence*>(sentence)->words.size();
}

kytea_word_t *kytea_sentence_word_at(kytea_sentence_t *sentence, int i) {
    return reinterpret_cast<kytea_word_t*>(&(reinterpret_cast<KyteaSentence*>(sentence)->words[i]));
}

kytea_std_string_t *kytea_word_surface(kytea_word_t *word, kytea_string_util_t *util) {
    StringUtil *u = reinterpret_cast<StringUtil*>(util);
    KyteaWord *w = reinterpret_cast<KyteaWord*>(word);
    string *str = new string(u->showString(w->surface));
    return reinterpret_cast<kytea_std_string_t*>(str);
}

kytea_std_string_t *kytea_word_tag(kytea_word_t *word, int i, int j, kytea_string_util_t *util) {
    StringUtil *u = reinterpret_cast<StringUtil*>(util);
    KyteaWord *w = reinterpret_cast<KyteaWord*>(word);
    string *str = new string(u->showString(w->tags[i][j].first));
    return reinterpret_cast<kytea_std_string_t*>(str);
}

size_t kytea_word_tags_len(kytea_word_t *word, int i) {
    KyteaWord *w = reinterpret_cast<KyteaWord*>(word);
    return w->tags[i].size();
}

void kytea_std_string_destroy(kytea_std_string_t *str) {
    delete reinterpret_cast<string*>(str);
}

const char *kytea_std_string_cstring(kytea_std_string_t *str) {
    return reinterpret_cast<string*>(str)->c_str();
}
