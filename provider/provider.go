package provider

import (
	"github.com/mh-cbon/emd/provider/gh"
)

//Provider identify an url.
type Provider interface {
	Is(string) bool
	SetURL(string)
	GetUserName() string
	GetProjectName() string
	GetProjectPath() string
	GetProviderURL() string
	GetProviderID() string
	GetProjectURL() string
}

// Providers if a facade of many Provider.
type Providers struct {
	URL       string
	Providers []Provider
}

// New providers for url.
func New(url string) Providers {
	return Providers{URL: url}
}

// Default makes a new Providers facade with pre loaded gh provider.
func Default(url string) Providers {
	ret := New(url)
	ret.Add(gh.New())
	return ret
}

// Add a concrete provider.
func (p Providers) Add(provider ...Provider) Providers {
	p.Providers = append(p.Providers, provider...)
	return p
}

// GetUserName of the the current url.
func (p Providers) selectProvider() Provider {
	ret := &NotFoundProvider{}
	for _, pp := range p.Providers {
		if pp.Is(p.URL) {
			pp.SetURL(p.URL)
			return pp
		}
	}
	return ret
}

// GetUserName of the the current url.
func (p Providers) GetUserName() string {
	return p.selectProvider().GetUserName()
}

// GetProjectName of the the current url.
func (p Providers) GetProjectName() string {
	return p.selectProvider().GetProjectName()
}

// GetProviderURL of the the current url.
func (p Providers) GetProviderURL() string {
	return p.selectProvider().GetProviderURL()
}

// GetProviderID of the the current url.
func (p Providers) GetProviderID() string {
	return p.selectProvider().GetProviderID()
}

// GetProjectPath of the the current url.
func (p Providers) GetProjectPath() string {
	return p.selectProvider().GetProjectPath()
}

// NotFoundProvider for an url not identified
type NotFoundProvider struct{}

var notFound = "not found"

// Is always return false.
func (p *NotFoundProvider) Is(u string) bool {
	return false
}

// SetURL si no op.
func (p *NotFoundProvider) SetURL(u string) {
}

// GetUserName of the the current url.
func (p *NotFoundProvider) GetUserName() string {
	return notFound
}

// GetProjectName of the the current url.
func (p *NotFoundProvider) GetProjectName() string {
	return notFound
}

// GetProviderURL of the the current url.
func (p *NotFoundProvider) GetProviderURL() string {
	return notFound
}

// GetProviderID of the the current url.
func (p *NotFoundProvider) GetProviderID() string {
	return notFound
}

// GetProjectPath of the the current url.
func (p *NotFoundProvider) GetProjectPath() string {
	return notFound
}

// GetProjectURL of the the current url.
func (p *NotFoundProvider) GetProjectURL() string {
	return notFound
}

// GetURL of the the current url.
func (p *NotFoundProvider) GetURL() string {
	return notFound
}
