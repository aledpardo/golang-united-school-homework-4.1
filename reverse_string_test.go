package reverse_string

import "testing"

func TestReverseString(t *testing.T) {
	data := []struct {
		In  string
		Out string
	}{
		{In: "Hello, World!", Out: "!dlroW ,olleH"},
		{In: "Hello, 🗺!", Out: "!🗺 ,olleH"},
		{In: "Hello\nWorld!", Out: "!dlroW\nolleH"},
		{In: "Hello\r\nWorld!", Out: "!dlroW\r\nolleH"},
	}
	for _, d := range data {
		got := ReverseString(d.In)
		if d.Out != got {
			t.Errorf("Got %s; want %s", got, d.Out)
		}
	}
}
