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
