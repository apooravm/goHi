package main

import (
	"fmt"
	"os"
)

func main() {

	cow := `	\   ^__^
	 \  (oo)\_______
	    (__)\       )\/\
	        ||----w |
	        ||     ||
    `

	text := "	"
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
    			text += arg + " "
    	}
	} else {
		text += "Hi :D"
	}
	text_top := "	"
	text_bottom := "	"
	for i := 0; i < len(text); i++ {
		text_top += "_"
		text_bottom += "-"
	} 
	
    fmt.Println(text_top)
    fmt.Println(text)
    fmt.Println(text_bottom)
    fmt.Println(cow)
}