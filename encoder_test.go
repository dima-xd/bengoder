package bengoder_test

import (
	"testing"

	"github.com/dima-xd/bengoder"
)

type testInput struct {
	Announce     string
	AnnounceList []string `bengoder:"announce-list"`
}

func TestEncode(t *testing.T) {
	test := testInput{"udp://tracker.openbittorrent.com:80/announce",
		[]string{"udp://tracker.openbittorrent.com:80/announce", "udp://tracker.publicbt.com:80/announce"}}

	expected := "d8:announce44:udp://tracker.openbittorrent.com:80/announce13:announce-listll44:udp://tracker.openbittorrent.com:80/announceel38:udp://tracker.publicbt.com:80/announceeee"

	actual, err := bengoder.Encode(test)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if expected != actual {
		t.Fatalf("'%s' expected value is not equal to '%s' actual value", expected, actual)
	}
}
