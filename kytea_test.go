package kytea

import "testing"

func TestKyTea(t *testing.T) {
	kytea, err := New()
	if err != nil {
		t.Error("unexpected error", err)
	}
	defer kytea.Destroy()

	kytea.ReadModel(kytea.Config().ModelFile())

	util := kytea.StringUtil()
	sentence := util.NewSentence("こんにちは世界")
	defer sentence.Destroy()

	kytea.CalculateWS(sentence)
	kytea.CalculateAllTags(sentence)
	var got, expected string

	if got, expected := sentence.WordsLen(), 2; got != expected {
		t.Errorf("want %d, got %d", expected, got)
	}
	if got, expected := sentence.KyTeaWord(0).Surface(util), "こんにちは"; got != expected {
		t.Errorf("want %s, got %s", expected, got)
	}
	if got, expected := sentence.KyTeaWord(1).Surface(util), "世界"; got != expected {
		t.Errorf("want %s, got %s", expected, got)
	}

	got, _ = sentence.KyTeaWord(0).Tag(0, 0, util)
	expected = "感動詞"
	if got != expected {
		t.Errorf("want %s, got %s", expected, got)
	}
	got, _ = sentence.KyTeaWord(1).Tag(0, 0, util)
	expected = "名詞"
	if got != expected {
		t.Errorf("want %s, got %s", expected, got)
	}

	got, _ = sentence.KyTeaWord(0).Tag(1, 0, util)
	expected = "こんにちは"
	if got != expected {
		t.Errorf("want %s, got %s", expected, got)
	}
	got, _ = sentence.KyTeaWord(1).Tag(1, 0, util)
	expected = "せかい"
	if got != expected {
		t.Errorf("want %s, got %s", expected, got)
	}

	if sentence.KyTeaWord(0).CandidateTagsLen(0) < 1 {
		t.Errorf("expected CandidateTagsLen is larger than 1, got %d", sentence.KyTeaWord(0).CandidateTagsLen(0))
	}

	if sentence.KyTeaWord(0).TagsLen() != 2 {
		t.Errorf("expected 2, got %d", sentence.KyTeaWord(0).TagsLen())
	}
}
