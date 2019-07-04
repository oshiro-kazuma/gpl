package main

func JoinString(strings []string) string {
	s, sep := "", ""
	for _, str := range strings {
		s += sep + str
		sep = " "
	}
	return s
}
