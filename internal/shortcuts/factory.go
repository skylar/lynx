package shortcuts

import (
	"crypto/sha1"
	"encoding/base64"
	"net/url"
)

// Generator the type to generate keys(short urls)
type Generator func(uri string) string

const shortcodeLength = 9;

// HashingGenerator is the defautl url generator
var HashingGenerator = func(uriString string) string {
	hasher := sha1.New()
	hasher.Write([]byte(uriString))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	return sha[0:shortcodeLength]
}

// Factory is responsible to generate keys(short urls)
type Factory struct {
	generator Generator
}

// NewFactory receives a generator and a store and returns a new url Factory.
func NewFactory(generator Generator) *Factory {
	return &Factory{
		generator: generator,
	}
}

func (f *Factory) isValid(uriString string) bool {
	_, err := url.ParseRequestURI(uriString)
	return err == nil
}

// Gen generates the key.
func (f *Factory) gen(uriString string) (string, error) {
	uri, err := url.ParseRequestURI(uriString)
	if err != nil {
		return "", err
	}

	// use the normalized, parsed string to generate a key
	key := f.generator(uri.String())
	return key, nil
}
