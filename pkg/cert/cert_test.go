package cert_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/hsmtkk/http2/pkg/cert"
	"github.com/stretchr/testify/assert"
)

func TestMakeCert(t *testing.T) {
	dir, err := ioutil.TempDir("", "temp")
	assert.Nil(t, err, "should be nil")
	defer os.RemoveAll(dir)

	certPath := filepath.Join(dir, "cert.pem")
	keyPath := filepath.Join(dir, "key.pem")
	err = cert.New().MakeCert(certPath, keyPath)
	assert.Nil(t, err, "should be nil")
	assert.True(t, fileExists(certPath), "should be true")
	assert.True(t, fileExists(keyPath), "should be true")
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
