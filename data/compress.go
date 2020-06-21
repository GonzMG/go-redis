package data

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"

	"github.com/golang/snappy"
	"github.com/rasky/go-lzo"
)

/****************************
*			GZIP			*
*****************************/

func EncodeGzip(stringToCompress string) ([]byte, error) {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	_, err := w.Write([]byte(stringToCompress))
	w.Close()
	if err != nil {
		return []byte{}, err
	}
	return b.Bytes(), nil
}

func DecodeGzip(bytesToDecode []byte) (string, error) {
	b := bytes.NewBuffer(bytesToDecode)
	r, err := gzip.NewReader(b)
	if err != nil {
		return "", err
	}
	result, err := ioutil.ReadAll(r)
	if err != nil {
		return "", nil
	}
	r.Close()
	return string(result), nil
}

/************************
*			LZO			*
*************************/

func EncodeLZO(stringToCompress string) []byte {
	return lzo.Compress1X([]byte(stringToCompress))
}

func DecodeLZO(bytesToDecode []byte) string {
	b := bytes.NewBuffer(bytesToDecode)
	out, err := lzo.Decompress1X(b, 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	return string(out)
}

/****************************
*			Snappy			*
*****************************/

func EncodeSnappy(stringToCompress string) ([]byte, error) {
	var b bytes.Buffer
	w := snappy.NewWriter(&b)
	_, err := w.Write([]byte(stringToCompress))
	w.Close()
	if err != nil {
		return []byte{}, err
	}
	return b.Bytes(), nil
}

func DecodeSnappy(bytesToDecode []byte) (string, error) {
	b := bytes.NewBuffer(bytesToDecode)
	r := snappy.NewReader(b)
	result, err := ioutil.ReadAll(r)
	if err != nil {
		return "", nil
	}
	return string(result), nil
}
