package main

import (
	qrcode "github.com/skip2/go-qrcode"
	"io/ioutil"
	"fmt"
	"os"
	"image/color"
)

func main() {
	var errs []error
	defer func() {
		if len(errs) > 0 {
			fmt.Println(errs)
			os.Exit(1)
		}
	}()

	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		errs = append(errs, err)
		return
	}
	contents := string(bytes)

	err = qrcode.WriteColorFile(contents, qrcode.Medium, 1024, color.Black, color.White, "output.png")
	if err != nil {
		errs = append(errs, err)
		return
	}
}
