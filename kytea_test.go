package kytea

import "testing"

func TestKyTea(t *testing.T) {
	kytea, err := New()
	if err != nil {
		t.Error("unexpected error", err)
	}
	defer kytea.Destory()

	kytea.ReadModel("/usr/local/share/kytea/model.bin")

	t.Log(kytea.GetWS("こんにちは世界"))
}
