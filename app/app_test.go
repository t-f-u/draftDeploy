package app

import "testing"

func TestStart(t *testing.T) {
	err := Start("/Users/t_nagao/go/src/draftDeploy/test_draft", "/Users/t_nagao/go/src/draftDeploy/test_out")
	if err != nil {
		t.Fatal(err)
	}
}
