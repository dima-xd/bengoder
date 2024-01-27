package decoder

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
)

type Decoder struct {
	data []byte
	pos  int
}

var d Decoder

func Decode(reader *bufio.Reader) (interface{}, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	d = Decoder{data, 0}

	return decode(), nil
}

func decode() interface{} {
	switch d.data[d.pos] {
	case 'd':
		d.pos++

		dictionary := decodeDictionary()

		return dictionary
	case 'l':
		d.pos++

		list := decodeList()

		return list
	case 'i':
		d.pos++

		number := decodeInt()

		return number
	default:
		return decodeString()
	}
}

func decodeDictionary() map[string]interface{} {
	dictionary := map[string]interface{}{}

	for {
		if d.data[d.pos] == 'e' {
			d.pos++
			break
		}

		length := decodeStringLength()
		key := decodeKey(length)
		value := decode()

		dictionary[key] = value
	}

	return dictionary
}

func decodeList() []interface{} {
	list := []interface{}{}

	for {
		if d.data[d.pos] == 'e' {
			d.pos++
			break
		}

		value := decode()
		list = append(list, value)
	}

	return list
}

func decodeInt() int64 {
	endIndex := bytes.IndexByte(d.data[d.pos:], 'e')
	index := endIndex + d.pos
	number, _ := strconv.ParseInt(string(d.data[d.pos:index]), 10, 64)
	d.pos += endIndex + 1

	return number
}

func decodeStringLength() int {
	delimiterIndex := bytes.IndexByte(d.data[d.pos:], ':')
	index := delimiterIndex + d.pos
	length, _ := strconv.Atoi(string(d.data[d.pos:index]))
	d.pos += delimiterIndex + 1

	return length
}

func decodeKey(length int) string {
	key := string(d.data[d.pos : d.pos+length])
	d.pos += length

	return key
}

func decodeString() string {
	length := decodeStringLength()
	return decodeKey(length)
}
