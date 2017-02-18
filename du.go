package main

import (
	"code.cloudfoundry.org/bytefmt"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

var semaphore = make(chan struct{}, 20)

func main() {
	flag.Parse()
	roots := flag.Args()
	var wg sync.WaitGroup

	if len(roots) == 0 {
		roots = []string{"."}
	}
	filesizes := make(chan int64)

	for _, dir := range roots {
		wg.Add(1)
		go walkDir(dir, &wg, filesizes)
	}

	go func() {
		wg.Wait()
		close(filesizes)
	}()

	var nbytes, nfiles int64
	for size := range filesizes {
		nfiles++
		nbytes += size
	}
	printDiskUsage(nfiles, nbytes)

}

func printDiskUsage(files, sizes int64) {
	fmt.Fprintf(os.Stdout, "%d files, %s\n", files, bytefmt.ByteSize(uint64(sizes)))
}

func walkDir(dir string, n *sync.WaitGroup, filesizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, filesizes)
		} else {
			filesizes <- entry.Size()
		}

	}
}

func dirents(dir string) []os.FileInfo {

	semaphore <- struct{}{}
	defer func() { <-semaphore }()

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dirents: %s", err)
		return nil
	}
	return entries

}
