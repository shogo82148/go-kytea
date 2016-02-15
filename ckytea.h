#ifdef __cplusplus
extern "C" {
#endif

typedef struct kytea_t kytea_t;
typedef struct kytea_sentence_t kytea_sentence_t;
typedef struct kytea_string_util_t kytea_string_util_t;
typedef struct kytea_config_t kytea_config_t;
typedef struct kytea_word_t kytea_word_t;
typedef struct kytea_std_string_t kytea_std_string_t;
typedef struct kytea_tag_t {
    kytea_std_string_t *feature;
    double score;
} kytea_tag_t;

// functions for handling Kytea
kytea_t *kytea_new(void);
void kytea_destroy(kytea_t *kytea);
kytea_std_string_t *kytea_read_model(kytea_t *kytea, const char *path);
kytea_string_util_t *kytea_get_string_util(kytea_t *kytea);
kytea_config_t *kytea_get_config(kytea_t *kytea);
kytea_std_string_t *kytea_calculate_ws(kytea_t *kytea, kytea_sentence_t *sentence);
kytea_std_string_t *kytea_calculate_tags(kytea_t *kytea, kytea_sentence_t *sentence, int i);
kytea_std_string_t *kytea_calculate_all_tags(kytea_t *kytea, kytea_sentence_t *sentence);

// functions for handling KyteaStringUtil
kytea_sentence_t *kytea_string_util_new_sentence(kytea_string_util_t *util, const char *input, size_t length);

// functions for handling KyteaConfig
void kytea_config_parse_run_command_line(kytea_config_t *config, int argc, const char **argv);
kytea_std_string_t *kytea_config_get_model_file(kytea_config_t *config);
void kytea_config_set_model_file(kytea_config_t *config, const char *path);
int kytea_config_get_do_ws(kytea_config_t *config);
void kytea_config_set_do_ws(kytea_config_t *config, int flag);
kytea_std_string_t *kytea_config_get_ws_constraint(kytea_config_t *config);
void kytea_config_set_ws_constraint(kytea_config_t *config, const char *str, size_t length);
int kytea_config_get_do_tags(kytea_config_t *config);
void kytea_config_set_do_tags(kytea_config_t *config, int flag);
int kytea_config_get_do_tag(kytea_config_t *config, int i);
void kytea_config_set_do_tag(kytea_config_t *config, int i, int flag);
int kytea_config_get_do_unk(kytea_config_t *config);
void kytea_config_set_do_unk(kytea_config_t *config, int flag);
int kytea_config_get_num_tags(kytea_config_t *config);
void kytea_config_set_num_tags(kytea_config_t *config, int num);
int kytea_config_get_tag_max(kytea_config_t *config);
void kytea_config_set_tag_max(kytea_config_t *config, int num);
kytea_std_string_t *kytea_config_get_unk_tag(kytea_config_t *config);
void kytea_config_set_unk_tag(kytea_config_t *config, const char *str, size_t length);
kytea_std_string_t *kytea_config_get_default_tag(kytea_config_t *config);
void kytea_config_set_default_tag(kytea_config_t *config, const char *str, size_t length);
int kytea_config_get_unk_beam(kytea_config_t *config);
void kytea_config_set_unk_beam(kytea_config_t *config, int num);


// functions for handling KyteaSentence
void kytea_sentence_destroy(kytea_sentence_t *sentence);
size_t kytea_sentence_words_len(kytea_sentence_t *sentence);
kytea_word_t *kytea_sentence_word_at(kytea_sentence_t *sentence, int i);

// functions for handling KyteaWord
kytea_std_string_t *kytea_word_surface(kytea_word_t *word, kytea_string_util_t *util);
kytea_tag_t kytea_word_tag(kytea_word_t *word, int i, int j, kytea_string_util_t *util);
size_t kytea_word_candidate_tags_len(kytea_word_t *word, int i);
size_t kytea_word_tags_len(kytea_word_t *word);

// functions for handling std::string
void kytea_std_string_destroy(kytea_std_string_t *str);
const char *kytea_std_string_cstring(kytea_std_string_t *str);

#ifdef __cplusplus
}
#endif
