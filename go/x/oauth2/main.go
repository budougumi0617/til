package main

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	conf := &oauth2.Config{
		// echo MDAwMDAwMDA6YWJjZGVm | base64 -D
		// 00000000:abcdef%
		ClientID:     "00000000",
		ClientSecret: "abcdef",
		Scopes:       []string{"SCOPE1", "SCOPE2"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://sample.apigw.opencanvas.ne.jp/bizsol/v1/banks/0001/oauth/token",
			TokenURL: "https://sample.apigw.opencanvas.ne.jp/bizsol/v1/banks/0001/oauth/token",
		},
	}
	code := "cb46420e53c24580a4c4e0fe8f390575"
	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", tok)
	cid := tok.Extra("contractor_id").(string)
	fmt.Printf("controctor_id = %q\n", cid)
	fmt.Printf("expiry = %s\n", tok.Expiry)
}
