package main

import(

	"fmt"
	"strings"
)


func pyramid(pyramid string) string {

	length := len(pyramid)

	if length < 3 {
	return "Error: Need at least 3 args"
	}


	for i := 0; i < 10; i++ {

		new_level := switchandremove(pyramid)

		if len(new_level) < 2 {
		return pyramid
		} else {

		new_level += "\n"
		new_level += pyramid
		pyramid = new_level
		}

	}

	return pyramid
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
	fmt.Println(pyramid(base))
}
