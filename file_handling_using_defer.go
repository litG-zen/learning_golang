// _Defer_ is used to ensure that a function call is
// performed later in a program's execution, usually for
// purposes of cleanup. `defer` is often used where e.g.
// `ensure` and `finally` would be used in other languages.

package main

import (
	"fmt"
	"os"
)

// Suppose we wanted to create a file, write to it,
// and then close when we're done. Here's how we could
// do that with `defer`.
func main() {

	// Immediately after getting a file object with
	// `createFile`, we defer the closing of that file
	// with `closeFile`. This will be executed at the end
	// of the enclosing function (`main`), after
	// `writeFile` has finished.
	var file_path string
	var f *os.File

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("panic!", err)
		os.Exit(1)
	} else {
		file_path = cwd + "/test.txt"
	}

	if !fileExists(file_path) {
		f = createFile(file_path)
	} else {
		f = openFile(file_path)
	}
	defer closeFile(f)
	writeFile(f)
}

func fileExists(p string) bool {
	fmt.Println(fmt.Sprintf("checking %s presence", p))
	if _, err := os.Stat(p); err != nil {
		return false
	} else {
		return true
	}
}

func openFile(p string) *os.File {
	fmt.Println("opening file!")
	f, err := os.OpenFile(p, os.O_APPEND, 777)
	if err != nil {
		fmt.Println("panic !", err)
	}
	return f
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	// It's important to check for errors when closing a
	// file, even in a deferred function.
	if err != nil {
		panic(err)
	}
}
