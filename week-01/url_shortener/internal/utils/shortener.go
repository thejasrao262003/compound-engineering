package utils

const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// EncodeBase62 converts a number to a base62 string
func Encode(num int64) string {
	if num == 0 {
		return string(base62Chars[0]);
	}

	var encoded[] byte

	for num > 0 {
		remainder := num % 62
		encoded = append([]byte{base62Chars[remainder]}, encoded...)
		num = num/62
	}

	return string(encoded)
}

func Decode(encoded string) int64 {
	var num int64 = 0
	for i:=0; i<len(encoded) ; i++ {
		num = num * 62 + int64(indexOf(encoded[i]))
	}

	return num;
}

func indexOf(ele byte) int64 {
	for i:=0 ; i<len(base62Chars) ; i++ {
		if ele == base62Chars[i] {
			return int64(i);
		}
	}
	return -1;
}