package kytea

import "testing"

func TestKyTea(t *testing.T) {
	kytea, err := New()
	if err != nil {
		t.Error("unexpected error", err)
	}
	defer kytea.Destory()

	kytea.ReadModel("/usr/local/share/kytea/model.bin")

	util := kytea.StringUtil()
	sentence := util.NewSentence("こんにちは世界")
	defer sentence.Destory()

	kytea.CalculateWS(sentence)
	kytea.CalculateAllTags(sentence)

	if got, expected := sentence.Word(0).Surface(util), "こんにちは"; got != expected {
		t.Errorf("want %s, got %s", expected, got)
	}
	if got, expected := sentence.Word(1).Surface(util), "世界"; got != expected {
		t.Errorf("want %s, got %s", expected, got)
	}

	if got, expected := sentence.Word(0).Tag(0, 0, util), "感動詞"; got != expected {
		t.Errorf("want %s, got %s", expected, got)
	}
	if got, expected := sentence.Word(1).Tag(0, 0, util), "名詞"; got != expected {
		t.Errorf("want %s, got %s", expected, got)
	}

	if got, expected := sentence.Word(0).Tag(1, 0, util), "こんにちは"; got != expected {
		t.Errorf("want %s, got %s", expected, got)
	}
	if got, expected := sentence.Word(1).Tag(1, 0, util), "せかい"; got != expected {
		t.Errorf("want %s, got %s", expected, got)
	}

	if sentence.Word(0).TagsLen(0) < 1 {
		t.Errorf("expected TagsLen is larger than 1, got %d", sentence.Word(0).TagsLen(0))
	}
}
