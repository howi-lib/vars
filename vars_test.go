package vars

import (
	"fmt"
	"testing"
)

type stringTest struct {
	key  string
	in   string
	want string
}

var stringTests = []stringTest{
	{"GOARCH", "amd64", "amd64"},
	{"GOHOSTARCH", "amd", "amd"},
	{"GOHOSTOS", "linux", "linux"},
	{"GOOS", "linux", "linux"},
	{"GOPATH", "/go-workspace", "/go-workspace"},
	{"GOROOT", "/usr/lib/golang", "/usr/lib/golang"},
	{"GOTOOLDIR", "/usr/lib/golang/pkg/tool/linux_amd64", "/usr/lib/golang/pkg/tool/linux_amd64"},
	{"GCCGO", "gccgo", "gccgo"},
	{"CC", "gcc", "gcc"},
	{"GOGCCFLAGS", "-fPIC -m64 -pthread -fmessage-length=0", "-fPIC -m64 -pthread -fmessage-length=0"},
	{"CXX", "g++", "g++"},
	{"PKG_CONFIG", "pkg-config", "pkg-config"},
	{"CGO_ENABLED", "1", "1"},
	{"CGO_CFLAGS", "-g -O2", "-g -O2"},
	{"CGO_CPPFLAGS", "", ""},
	{"CGO_CXXFLAGS", "-g -O2", "-g -O2"},
	{"CGO_FFLAGS", "-g -O2", "-g -O2"},
	{"CGO_LDFLAGS", "-g -O2", "-g -O2"},
}

func genStringTestBytes() []byte {
	var out []byte
	for _, data := range stringTests {
		line := fmt.Sprintf(`%s="%s"`+"\n", data.key, data.in)
		out = append(out, []byte(line)...)
	}
	return out
}

func TestNewValue(t *testing.T) {
	var tests = []struct {
		val  interface{}
		want string
	}{
		{nil, ""},
		{"", ""},
	}

	for _, tt := range tests {
		got := NewValue(tt.val).String()
		if got != tt.want {
			t.Errorf("want: %s got %s", tt.want, got)
		}
	}
}

func TestValueFromString(t *testing.T) {
	tests := []struct {
		name string
		val  string
		want string
	}{
		{"STRING", "some-string", "some-string"},
		{"STRING", "some-string with space ", "some-string with space"},
		{"STRING", " some-string with space", "some-string with space"},
		{"STRING", "1234567", "1234567"},
	}
	for _, tt := range tests {
		if got := NewValue(tt.val); got.String() != tt.want {
			t.Errorf("ValueFromString() = %q, want %q", got.String(), tt.want)
		}
		if rv := NewValue(tt.val); string(rv.Rune()) != tt.want {
			t.Errorf("Value.Rune() = %q, want %q", string(rv.Rune()), tt.want)
		}
	}
}
