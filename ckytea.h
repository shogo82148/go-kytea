#ifdef __cplusplus
extern "C" {
#endif

typedef struct kytea_t kytea_t;
typedef struct kytea_vector_string_t kytea_vector_string_t;
typedef struct kytea_sentence_t kytea_sentence_t;
typedef struct kytea_string_util_t kytea_string_util_t;

kytea_t *kytea_new(void);
void kytea_destroy(kytea_t *kytea);
void kytea_read_model(kytea_t *kytea, const char *path);
kytea_string_util_t *kytea_get_string_util(kytea_t *kytea);
kytea_vector_string_t *kytea_get_ws(kytea_t *kytea, const char *input, size_t length);
size_t kytea_vector_string_size(kytea_vector_string_t *vector_string);
const char* kytea_vector_string_at(kytea_vector_string_t *vector_string, int index);
void kytea_vector_string_destory(kytea_vector_string_t *vector_string);
    kytea_sentence_t *kytea_string_util_new_sentence(kytea_string_util_t *util, const char *input, size_t length);
void kytea_sentence_destroy(kytea_sentence_t *sentence);

#ifdef __cplusplus
}
#endif
