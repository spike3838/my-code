package main

import "fmt"


func pyramid(pyramid string) string {

	white_space := " "
	state := pyramid
	for i := 0; i < len(pyramid); i++{

		buffer := ""
		new_level := switchandremove(state)
		state = new_level

		for j := 0; j <= i; j++{

			buffer += white_space
		}

		if len(new_level) < 2 {
			return pyramid
		} else {
			new_level += "\n"
			buffer += new_level
			new_level = buffer
			new_level += pyramid
			pyramid = new_level
		}
	}

	return pyramid
}

func switchandremove(level string) string{

	level_slice := []byte(level[1:len(level)-1])

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
	base := "/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\/\\"
	fmt.Println(pyramid(base))

}
