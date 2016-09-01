package main

import(

	"fmt"
	"strings"
)


func pyramid(base string) string {

	length := len(base)

	if length < 3 {
	return "Error: Need at least 3 args"
	}

	return base
}





func cut_space(level string) string {

	cutted := strings.TrimSpace(level)
	return cutted
}





func switchandremove(level string) string{

	level_byte := []byte(level)
	level_slice := []byte(level[1:len(level_byte)-1])


	for i := 0 ; i < len(level_slice); i++{

		if level_slice[i] == byte('\\'){

		level_slice[i] = byte('/')
		} else {

			if level_slice[i] == byte('/'){

			level_slice[i] = byte('\\')
			}
		}
	}

	return string(level_slice)
}





func main() {
	base := "/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\"
	fmt.Println(switchandremove(base))
}
