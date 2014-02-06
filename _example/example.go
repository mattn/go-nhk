package main

import (
	"fmt"
	"github.com/mattn/go-nhk"
	"log"
	"os"
)

func main() {
	client := nhk.NewClient(os.Getenv("NHK_PROGRAM_APIKEY"))

	// 東京のNHK総合1
	// http://api-portal.nhk.or.jp/doc-request#explain_area
	// http://api-portal.nhk.or.jp/doc-request#explain_service
	pl, err := client.ProgramList("130", "g1", nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("東京のNHK総合1 今日の番組一覧")
	for _, p := range pl {
		fmt.Println(p)
	}
	fmt.Println()

	// 東京のNHK総合1 アニメ特撮
	// http://www.arib.or.jp/english/html/overview/doc/2-STD-B10v5_1.pdf
	gl, err := client.ProgramGenre("130", "e1", 0x7, 0x0, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("東京のNHK Eテレ 今日のアニメ特撮一覧")
	for _, p := range gl {
		fmt.Println(p)
	}
	fmt.Println()

	// 東京のNHK Eテレ1 リトルチャロ2
	p, err := client.ProgramInfo("130", "e1", "2014020600268")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("東京のNHK Eテレ1 リトルチャロ")
	fmt.Println(p)
	fmt.Println()

	// 東京のNHK Eテレ1 現在放送中
	n, err := client.NowOnAir("130", "e1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("東京のNHK Eテレ1 現在放送中")
	fmt.Println(n)
	fmt.Println()
}
