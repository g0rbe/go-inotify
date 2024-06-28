package main

import (
	"fmt"
	"os"

	"github.com/g0rbe/go-inotify"
)

func main() {

	w, err := inotify.Init(0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to Init(): %s\n", err)
		os.Exit(1)
	}
	defer w.Close()

	if len(os.Args) == 1 {
		fmt.Printf("Usage: %s paths...\n", os.Args[0])
		os.Exit(0)
	}

	paths := os.Args[1:]

	for i := range paths {

		fmt.Printf("Start watching %s\n", paths[i])

		_, err = w.AddWatch(paths[i], inotify.IN_ALL_EVENTS)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to watch %s: %s\n", paths[i], err)
			os.Exit(1)
		}
	}

	for {

		e, err := w.Read()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to Read(): %s\n", err)
			os.Exit(1)
		}

		for j := range e {
			fmt.Printf("%s\n", e[j])
		}
	}
}
