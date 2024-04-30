package zshhist

const (
	null   = 0x00
	meta   = 0x83
	marker = 0xa2
	pound  = 0x84
	nularg = 0xa1
	bang   = 0x9c
	snull  = 0x9d
)

func isMeta(c byte) bool {
	return c == null || c == meta || c == marker ||
		pound <= c && c <= bang || snull <= c && c <= nularg
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

// Unmetafy a metafied line
func Unmetafy(str string) string {
	var j int
	bs := []byte(str)
	for i := 0; i < len(bs); i++ {
		if bs[i] == meta {
			i++
			bs[j] = bs[i] ^ 32
		} else {
			bs[j] = bs[i]
		}
		j++
	}
	return string(bs[:j])
}
