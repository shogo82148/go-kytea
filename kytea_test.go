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

	t.Log(sentence.WordsLen())
}
