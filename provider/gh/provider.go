package gh

import (
	"strings"
)

// Provider detects GH url
type Provider struct {
	URL string
}

// New provider to work with url
func New() *Provider {
	return &Provider{}
}

// Is the current url about gh?
func (p *Provider) Is(url string) bool {
	return strings.Index(url, "github.com/") > -1
}

// SetURL of the provider
func (p *Provider) SetURL(url string) {
	p.URL = url
}

func (p *Provider) cleanURL() string {
	url := p.URL
	if url[:4] == "http" {
		url = url[4:]
	} else if url[:5] == "https" {
		url = url[5:]
	}
	if url[:3] == "://" {
		url = url[3:]
	}
	if url[:1] == "/" {
		url = url[1:]
	}
	return url
}

// GetUserName of the the current url.
func (p *Provider) GetUserName() string {
	url := p.cleanURL()
	ss := strings.Split(url, "/")
	if len(ss) > 0 {
		return ss[1]
	}
	return ""
}

// GetProjectName of the the current url.
func (p *Provider) GetProjectName() string {
	url := p.cleanURL()
	ss := strings.Split(url, "/")
	if len(ss) > 2 {
		return ss[2]
	}
	return ""
}

// GetProjectPath of the the current url.
func (p *Provider) GetProjectPath() string {
	url := p.cleanURL()
	ss := strings.Split(url, "/")
	if len(ss) > 3 {
		return "/" + strings.Join(ss[3:], "/")
	}
	return ""
}

// GetProjectURL of the the current url.
func (p *Provider) GetProjectURL() string {
	ret := []string{}
	if x := p.GetProviderURL(); x != "" {
		if y := p.GetUserName(); y != "" {
			if z := p.GetProjectName(); z != "" {
				ret = append(ret, x)
				ret = append(ret, y)
				ret = append(ret, z)
			}
		}
	}
	return strings.Join(ret, "/")
}

// GetURL of the the current url.
func (p *Provider) GetURL() string {
	ret := []string{}
	if x := p.GetProjectURL(); x != "" {
		ret = append(ret, x)
		if y := p.GetProjectPath(); y != "" {
			ret = append(ret, y[1:]) //rm front /
		}
	}
	return strings.Join(ret, "/")
}

// GetProviderURL of the the current url.
func (p *Provider) GetProviderURL() string {
	return "github.com"
}

// GetProviderID of the the current url.
func (p *Provider) GetProviderID() string {
	return "github"
}
