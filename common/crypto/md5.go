package crypto

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
)

//对字符串进行MD5哈希
func MD5(data string) string {
	t := md5.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

//对字符串进行SHA1哈希
func SHA1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}
