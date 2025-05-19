package core

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"gocleaner/config"
	"gocleaner/core/email"
	"gocleaner/internal"
)

func RunCleaner(cfg *config.Config) {
	internal.SetupLogger()
	defer internal.LogFile.Close()

	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		panic(err)
	}

	cutoff := time.Now().AddDate(0, 0, -cfg.DaysThreshold)
	deletedChan := make(chan string, 1000)
	var wg sync.WaitGroup

	// Remove arquivos
	wg.Add(1)
	go func() {
		defer wg.Done()
		filepath.Walk(cfg.Directory, func(path string, info os.FileInfo, err error) error {
			if err == nil && !info.IsDir() && info.ModTime().Before(cutoff) {
				if removeWithLog(path) {
					deletedChan <- path
				}
			}
			return nil
		})
	}()

	// Remove pastas seguras
	wg.Add(1)
	go func() {
		defer wg.Done()
		var dirs []string
		filepath.Walk(cfg.Directory, func(path string, info os.FileInfo, err error) error {
			if err == nil && info.IsDir() {
				dirs = append(dirs, path)
			}
			return nil
		})

		for i := len(dirs) - 1; i >= 0; i-- {
			path := dirs[i]
			if path == cfg.Directory {
				continue
			}
			recent := false
			filepath.Walk(path, func(p string, i os.FileInfo, e error) error {
				if e == nil && !i.IsDir() && i.ModTime().After(cutoff) {
					recent = true
					return filepath.SkipDir
				}
				return nil
			})
			if !recent {
				if removeWithLog(path) {
					deletedChan <- path
				}
			}
		}
	}()

	go func() {
		wg.Wait()
		close(deletedChan)
	}()

	var deletedItems []string
	for item := range deletedChan {
		deletedItems = append(deletedItems, item)
	}

	if err := email.SendReport(cfg, deletedItems); err != nil {
		internal.Logger.Printf("Erro ao enviar e-mail: %v", err)
	}

	fmt.Println("Processo concluído.")
}

func removeWithLog(path string) bool {
	if err := os.RemoveAll(path); err != nil {
		internal.Logger.Printf("Erro ao remover %s: %v", path, err)
		return false
	}
	internal.Logger.Printf("Removido: %s", path)
	return true
}
