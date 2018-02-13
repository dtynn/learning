package CountAndSay

func countAndSay(n int) string {
	if n <= 1 {
		return "1"
	}

	lastbytes := []byte(countAndSay(n - 1))
	bytes := make([]byte, 0, len(lastbytes)*2)

	lastbyte := lastbytes[0]
	lastcount := 1
	for i := 1; i < len(lastbytes); i++ {
		b := lastbytes[i]
		if b == lastbyte {
			lastcount++
			continue
		}

		bytes = append(bytes, '0'+byte(lastcount), lastbyte)
		lastbyte = b
		lastcount = 1
	}

	bytes = append(bytes, '0'+byte(lastcount), lastbyte)

	return string(bytes)
}
