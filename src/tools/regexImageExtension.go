package tools

import "regexp"

//GetExtension deuelve la exntesión de una imagen
func GetExtension(str string) (res string) {
	var re = regexp.MustCompile(`(?m)(gif|png|jpe?g)`)

	for _, match := range re.FindAllString(str, -1) {
		res = match
	}

	return
}
