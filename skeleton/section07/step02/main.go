// STEP02: HTTPサーバを作ってガチャの結果をブラウザで表示しよう

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gohandson/gacha-ja/gacha"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	p := gacha.NewPlayer(10, 100)
	// ※本当はハンドラ間で競合が起きるのでNG
	play := gacha.NewPlay(p)

	http.HandleFunc("/draw", func(w http.ResponseWriter, r *http.Request) {
		if play.Draw() {
			// TODO: レスポンスとして最後の結果を出力する
			// 最後の結果はplay.Resultメソッドから取得できる
			fmt.Fprintln(w, play.Result())
		}

		if err := play.Err(); err != nil {
			// TODO: InternalServerErrorでエラーレスポンスを返す
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintln(w, "残り:", p.DrawableNum())
	})

	return http.ListenAndServe(":8080", nil)
}
