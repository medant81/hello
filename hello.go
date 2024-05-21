package hello

import "fmt"
import "github.com/medant81/hello/version1"
import "github.com/medant81/hello/version2"

func Hello() {
	fmt.Println("Version v1.1.0")
	version1.Version1()
	version2.Version2()
}
