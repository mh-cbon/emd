package gh

import (
	"testing"
)

func TestIs(t *testing.T) {
	data := []string{
		"http://github.com/",
	}
	p := New()
	want := true
	for i, url := range data {
		got := p.Is(url)
		if got != want {
			t.Errorf("At %v want: %v got %v in %q", i, want, got, url)
		}
	}
}

func TestIsNot(t *testing.T) {
	data := []string{
		"http://gitlab.com/",
	}
	p := New()
	want := false
	for i, url := range data {
		got := p.Is(url)
		if got != want {
			t.Errorf("At %v want: %v got %v in %q", i, want, got, url)
		}
	}
}

func TestGetUserName(t *testing.T) {
	data := map[string]string{
		"http://github.com/":                "",
		"http://github.com/mh-cbon":         "mh-cbon",
		"http://github.com/mh-cbon/emd":     "mh-cbon",
		"http://github.com/mh-cbon/emd/cmd": "mh-cbon",
		"github.com/mh-cbon":                "mh-cbon",
		"github.com/mh-cbon/emd":            "mh-cbon",
		"github.com/mh-cbon/emd/cmd":        "mh-cbon",
		"/github.com/mh-cbon":               "mh-cbon",
		"/github.com/mh-cbon/emd":           "mh-cbon",
		"/github.com/mh-cbon/emd/cmd":       "mh-cbon",
	}
	p := New()
	for url, want := range data {
		p.SetURL(url)
		got := p.GetUserName()
		if got != want {
			t.Errorf("Want: %q got %q in %q", want, got, url)
		}
	}
}

func TestGetProjectName(t *testing.T) {
	data := map[string]string{
		"http://github.com/":                "",
		"http://github.com/mh-cbon":         "",
		"http://github.com/mh-cbon/emd":     "emd",
		"http://github.com/mh-cbon/emd/cmd": "emd",
		"github.com/mh-cbon":                "",
		"github.com/mh-cbon/emd":            "emd",
		"github.com/mh-cbon/emd/cmd":        "emd",
		"/github.com/mh-cbon":               "",
		"/github.com/mh-cbon/emd":           "emd",
		"/github.com/mh-cbon/emd/cmd":       "emd",
	}
	p := New()
	for url, want := range data {
		p.SetURL(url)
		got := p.GetProjectName()
		if got != want {
			t.Errorf("Want: %q got %q in %q", want, got, url)
		}
	}
}

func TestGetProjectPath(t *testing.T) {
	data := map[string]string{
		"http://github.com/":                "",
		"http://github.com/mh-cbon":         "",
		"http://github.com/mh-cbon/emd":     "",
		"http://github.com/mh-cbon/emd/cmd": "/cmd",
		"github.com/mh-cbon":                "",
		"github.com/mh-cbon/emd":            "",
		"github.com/mh-cbon/emd/cmd":        "/cmd",
		"/github.com/mh-cbon":               "",
		"/github.com/mh-cbon/emd":           "",
		"/github.com/mh-cbon/emd/cmd":       "/cmd",
	}
	p := New()
	for url, want := range data {
		p.SetURL(url)
		got := p.GetProjectPath()
		if got != want {
			t.Errorf("Want: %q got %q in %q", want, got, url)
		}
	}
}

func TestGetProjectURL(t *testing.T) {
	data := map[string]string{
		"http://github.com/":                "",
		"http://github.com/mh-cbon":         "",
		"http://github.com/mh-cbon/emd":     "github.com/mh-cbon/emd",
		"http://github.com/mh-cbon/emd/cmd": "github.com/mh-cbon/emd",
		"github.com/mh-cbon":                "",
		"github.com/mh-cbon/emd":            "github.com/mh-cbon/emd",
		"github.com/mh-cbon/emd/cmd":        "github.com/mh-cbon/emd",
		"/github.com/mh-cbon":               "",
		"/github.com/mh-cbon/emd":           "github.com/mh-cbon/emd",
		"/github.com/mh-cbon/emd/cmd":       "github.com/mh-cbon/emd",
	}
	p := New()
	for url, want := range data {
		p.SetURL(url)
		got := p.GetProjectURL()
		if got != want {
			t.Errorf("Want: %q got %q in %q", want, got, url)
		}
	}
}

func TestGetURL(t *testing.T) {
	data := map[string]string{
		"http://github.com/":                "",
		"http://github.com/mh-cbon":         "",
		"http://github.com/mh-cbon/emd":     "github.com/mh-cbon/emd",
		"http://github.com/mh-cbon/emd/cmd": "github.com/mh-cbon/emd/cmd",
		"github.com/mh-cbon":                "",
		"github.com/mh-cbon/emd":            "github.com/mh-cbon/emd",
		"github.com/mh-cbon/emd/cmd":        "github.com/mh-cbon/emd/cmd",
		"/github.com/mh-cbon":               "",
		"/github.com/mh-cbon/emd":           "github.com/mh-cbon/emd",
		"/github.com/mh-cbon/emd/cmd":       "github.com/mh-cbon/emd/cmd",
	}
	p := New()
	for url, want := range data {
		p.SetURL(url)
		got := p.GetURL()
		if got != want {
			t.Errorf("Want: %q got %q in %q", want, got, url)
		}
	}
}
