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

kytea_vector_string_t *kytea_get_ws(kytea_t *kytea, const char *input, size_t length) {
    kytea::Kytea *k = reinterpret_cast<kytea::Kytea*>(kytea);
    kytea::StringUtil* util = k->getStringUtil();
    std::vector<std::string>* vec = new std::vector<std::string>;
    std::string str(input, length);
    kytea::KyteaString surface_string(util->mapString(str));
    kytea::KyteaSentence sentence(surface_string, util->normalize(surface_string));
    k->calculateWS(sentence);

    const kytea::KyteaSentence::Words & words =  sentence.words;
    for(size_t i = 0; i < words.size(); i++) {
        vec->push_back(util->showString(words[i].surface));
    }

    return reinterpret_cast<kytea_vector_string_t*>(vec);
}

size_t kytea_vector_string_size(kytea_vector_string_t *vector_string) {
    std::vector<std::string>* vec = reinterpret_cast<std::vector<std::string>*>(vector_string);
    return vec->size();
}

const char* kytea_vector_string_at(kytea_vector_string_t *vector_string, int index) {
    std::vector<std::string>* vec = reinterpret_cast<std::vector<std::string>*>(vector_string);
    return vec->at(index).c_str();
}

void kytea_vector_string_destory(kytea_vector_string_t *vector_string) {
    delete reinterpret_cast<std::vector<std::string>*>(vector_string);
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
