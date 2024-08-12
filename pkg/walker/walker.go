package walker

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type Opt struct {
	Symlink bool
	Dir     bool
	File    bool
	Ext     string
}

func innerFind(root string, opt *Opt, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()

	entries, err := os.ReadDir(root)
	if err != nil {
		return
	}

	for _, entry := range entries {
		path := filepath.Join(root, entry.Name())
		info, err := os.Lstat(path)
		if err != nil {
			continue
		}

		if info.IsDir() {
			if opt.Dir {
				results <- path
			}
			wg.Add(1)
			go innerFind(path, opt, wg, results)
		} else if info.Mode()&os.ModeSymlink != 0 && opt.Symlink {
			target, err := os.Readlink(path)
			if err != nil {
				return
			}
			if _, err := os.Stat(target); err == nil {
				path += " -> " + target
			} else if os.IsNotExist(err) {
				path += " -> " + "[broken]"
			} else {
				continue
			}
			results <- path
		} else {
			if opt.File && (opt.Ext == "" || opt.Ext != "" && ("."+opt.Ext) == filepath.Ext(path)) {
				results <- path
				continue
			}
		}
	}
}

func Find(root string, opt *Opt) {
	var wg sync.WaitGroup
	results := make(chan string)

	go func() {
		for result := range results {
			fmt.Println("Found:", result)
		}
	}()
	wg.Add(1)
	go innerFind(root, opt, &wg, results)

	wg.Wait()

	close(results)
}
