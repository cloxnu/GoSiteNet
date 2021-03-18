package main

import (
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding/htmlindex"
	"io"
)

// https://github.com/PuerkitoBio/goquery/wiki/Tips-and-tricks#handle-non-utf8-html-pages
func detectContentCharset(body io.Reader) string {
	r := bufio.NewReader(body)
	if data, err := r.Peek(1024); err == nil {
		_, name, _ := charset.DetermineEncoding(data, "")
		return name
		//if _, name, ok := charset.DetermineEncoding(data, ""); ok {
		//	return name
		//}
	}
	return "utf-8"
}

// DecodeHTMLBody returns an decoding reader of the html Body for the specified `charset`
// If `charset` is empty, DecodeHTMLBody tries to guess the encoding from the content
func DecodeHTMLBody(body io.Reader, charset string) (io.Reader, error) {
	if charset == "" {
		charset = detectContentCharset(body)
	}
	e, err := htmlindex.Get(charset)
	if err != nil {
		return nil, err
	}
	if name, _ := htmlindex.Name(e); name != "utf-8" {
		body = e.NewDecoder().Reader(body)
	}
	return body, nil
}
