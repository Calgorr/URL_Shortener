package model

import (
	"crypto/md5"
	"encoding/base64"
	"strings"
)

type Link struct {
	Address   string
	Hash      string
	UsedTimes int
}

func NewLink(url string) *Link {
	link := new(Link)
	link.Address = url
	md5 := md5.Sum([]byte(url))
	link.Hash = strings.ReplaceAll(strings.ReplaceAll(base64.StdEncoding.EncodeToString(md5[:])[:6], "/", "_"), "+", "-")
	return link
}
