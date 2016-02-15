package kytea_test

import (
	"fmt"

	"github.com/shogo82148/go-kytea"
)

func ExampleKyTea_Parse() {
	tagger, err := kytea.New()
	if err != nil {
		panic(err)
	}
	defer tagger.Destroy()

	err = tagger.ReadModel(tagger.Config().ModelFile())
	if err != nil {
		panic(err)
	}

	result, err := tagger.Parse("こんにちは世界")
	if err != nil {
		panic(err)
	}
	for _, word := range result {
		fmt.Println(word)
	}
	// Output:
	// こんにちは/感動詞/こんにちは
	// 世界/名詞/せかい
}
