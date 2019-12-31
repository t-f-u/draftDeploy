package convert

import (
	. "draftDeploy/app/model"
	"testing"
)

func TestConvert(t *testing.T) {
	data := ReadData{}
	data.BuildConfig = BuildConfig{
		Title:    "テスト",
		Chapters: nil,
	}
	data.Drafts = []Draft{
		{Data: []string{
			"あああ",
			"「いいいい」",
			"ううううう",
		}},
	}
	actual, _ := Convert(data, nil)
	if actual.Data[0] != "　あああ\n" {
		t.Log(actual.Data[0])
		t.Fail()
	}
	if actual.Data[1] != "「いいいい」\n" {
		t.Log(actual.Data[1])
		t.Fail()
	}
}

func TestAddNewLineMark(t *testing.T) {
	actual := addNewLineMark("あああ")
	if actual != "あああ\n" {
		t.Fail()
	}
}

func TestAddIndent(t *testing.T) {
	actual1 := addIndent("あああ")
	actual2 := addIndent("「あああ」")
	if actual1 != "　あああ" {
		t.Fail()
		t.Log(actual1)
	}
	if actual2 != "「あああ」" {
		t.Fail()
		t.Log(actual2)
	}
}

func TestStartsWithBracket(t *testing.T) {
	actual := startsWithBracket("「あああ」")
	if !actual {
		t.Fail()
	}
}
