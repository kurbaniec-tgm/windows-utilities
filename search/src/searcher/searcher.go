package searcher

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	. "search/src/opts"
	"sync"
)

func FindFiles(opts Opts) {

	/**
	matches, err := filepath.Glob(opts.Base + opts.Search)
	if err != nil {
		fmt.Println(err)
	}
	if len(matches) != 0 {
		fmt.Println("Found : ", matches)
	}*/
	//err := filepath.Walk(opts.Base, visit)

	/**
	err := filepath.Walk(opts.Base, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			matches, err := filepath.Glob(path + string(os.PathSeparator) + opts.Search)
			if err != nil {
				fmt.Println(err)
			}
			if len(matches) != 0 {
				for _, match := range matches {
					fmt.Println(info.Name() + string(os.PathSeparator) + match)
				}
			}
		}
		//fmt.Println(hey)
		//fmt.Printf("Visited: %s\n", path)
		return nil
	})
	fmt.Printf("filepath.Walk() returned %v\n", err)*/

	// Needed to wait for goroutines
	/**
	var wg sync.WaitGroup
	find(opts.Base, opts.Search, wg)
	wg.Wait()

	finished := make(chan bool)
	find2(opts.Base, opts.Search, finished)
	<- finished


	*/
	wg := &sync.WaitGroup{}
	wg.Add(1)
	find(opts.Base, opts.Search, wg)
	wg.Wait()

}

// WaitGroup needs to be a pointer, so that Go does not make a unnecessary
// copy that calls done.
// Source: https://stackoverflow.com/a/25234899
func find(path string, pattern string, wg *sync.WaitGroup) {
	nextWg := &sync.WaitGroup{}
	end := true
	dir, err := ioutil.ReadDir(path)
	for _, entry := range dir {
		if entry.IsDir() {
			end = false
			nextWg.Add(1)
			go find(filepath.Join(path, entry.Name()), pattern, nextWg)
		} else {
			fmt.Println(entry.Name())
		}
	}
	if err != nil {
		fmt.Println(err)
	}
	if !end {
		nextWg.Wait()
	}
	wg.Done()
}
