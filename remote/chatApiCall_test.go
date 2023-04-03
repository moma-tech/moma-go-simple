package remote_test

import (
	"testing"

	"top.moma.go.simple/remote"
)

func Test_remoteCall(t *testing.T) {
	var response = remote.ChatApiCall("abc", "content")
	if response != "null" {
		t.Failed()
	}
}
