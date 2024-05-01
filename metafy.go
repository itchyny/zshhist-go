package zshhist

const (
	null   = 0x00
	meta   = 0x83
	marker = 0xa2
)

var isMeta [256]bool

func init() {
	isMeta[null] = true
	for b := meta; b <= marker; b++ {
		isMeta[b] = true
	}
}

// Metafy a string.
func Metafy(str string) string {
	bs, i := []byte(str)[:0], 0
	for j := 0; j < len(str); j++ {
		if b := str[j]; isMeta[b] {
			bs = append(bs, str[i:j]...)
			bs = append(bs, meta, b^0x20)
			i = j + 1
		}
	}
	if i == 0 {
		return str
	}
	bs = append(bs, str[i:]...)
	return string(bs)
}

// Unmetafy a metafied string.
func Unmetafy(str string) string {
	bs, i := []byte(str)[:0], 0
	for j := 0; j < len(str); j++ {
		if str[j] == meta {
			bs = append(bs, str[i:j]...)
			j++
			bs = append(bs, str[j]^0x20)
			i = j + 1
		}
	}
	if i == 0 {
		return str
	}
	bs = append(bs, str[i:]...)
	return string(bs)
}
