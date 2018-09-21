package zshhist

func isMeta(c byte) bool {
	return c == null || c == meta || c == marker || pound <= c && c <= bang || snull <= c && c <= nularg
}

// Metafy a string.
func Metafy(str string) string {
	bs := []byte(str)
	var metaCount int
	for _, b := range bs {
		if isMeta(b) {
			metaCount++
		}
	}
	if metaCount == 0 {
		return str
	}
	cs := make([]byte, len(bs)+metaCount)
	var j int
	for _, b := range bs {
		if isMeta(b) {
			cs[j] = meta
			j++
			cs[j] = b ^ 32
		} else {
			cs[j] = b
		}
		j++
	}
	return string(cs)
}
