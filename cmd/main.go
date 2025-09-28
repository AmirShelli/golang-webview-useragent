package main

import "github.com/webview/webview_go"

func main() {
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("Basic Example")
	w.SetSize(480, 320, webview.HintNone)
	w.SetUserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/139.0.0.0 Safari/537.36")
	w.Navigate("https://www.whatismybrowser.com/detect/what-is-my-user-agent/")
	w.Run()
}
