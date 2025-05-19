package mocks

import (
	"gocleaner/config"
	"os"
	"time"
)

// MockFileInfo é uma implementação mock de os.FileInfo para testes
type MockFileInfo struct {
	FileName    string
	FileSize    int64
	FileMode    os.FileMode
	FileModTime time.Time
	FileIsDir   bool
	FileSys     interface{}
}

func (m MockFileInfo) Name() string       { return m.FileName }
func (m MockFileInfo) Size() int64        { return m.FileSize }
func (m MockFileInfo) Mode() os.FileMode  { return m.FileMode }
func (m MockFileInfo) ModTime() time.Time { return m.FileModTime }
func (m MockFileInfo) IsDir() bool        { return m.FileIsDir }
func (m MockFileInfo) Sys() interface{}   { return m.FileSys }

// MockConfigLoader é um function type que simula config.LoadConfig
type MockConfigLoader func(string) (*config.Config, error)

// MockFileRemover é um function type que simula os.RemoveAll
type MockFileRemover func(path string) error

// MockEmailSender é um function type que simula o envio de e-mail
type MockEmailSender func(cfg *config.Config, items []string) error