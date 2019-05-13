package urlshort

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type RedirectRegistry struct {
	Redirects map[string]Redirect
}

// Create a new registry
func NewRedirectRegistry() *RedirectRegistry {
	var redirectReg RedirectRegistry
	redirectReg.Redirects = make(map[string]Redirect)
	return &redirectReg
}

// Get a redirect with the provided source path
func (redReg *RedirectRegistry) Get(src string) (Redirect, bool) {
	redirect, ok := redReg.Redirects[src]
	return redirect, ok
}

// Add a Redirect to the registry
func (redReg *RedirectRegistry) Add(redirect *Redirect) error {
	_, ok := redReg.Get(redirect.Src)
	if ok {
		return errors.New("Source is already mapped")
	}
	redReg.Redirects[redirect.Src] = *redirect
	return nil
}

// Add a slice of Redirects to registry
func (redReg *RedirectRegistry) AddFromSlice(redirects []Redirect) error {
	for _, redirect := range redirects {
		err := redReg.Add(&redirect)
		if err != nil {
			return err
		}
	}
	return nil
}

// Add to registry by loading json array from file
func (redReg *RedirectRegistry) AddFromJSON(jsonPath string) error {
	if jsonPath == "" {
		return nil
	}

	// get json into slice of Redirects
	jsonRedirects, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		return err
	}
	redirects := make([]Redirect, 0)
	err = json.Unmarshal(jsonRedirects, &redirects)
	if err != nil {
		return err
	}

	// add redirects to the registry
	err = redReg.AddFromSlice(redirects)
	if err != nil {
		return err
	}
	return nil
}
