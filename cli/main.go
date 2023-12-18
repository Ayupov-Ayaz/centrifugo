package main

import "ayupov-ayaz/centrifugo/cli/cmd/run"

func main() {
	if err := run.Run(); err != nil {
		panic(err)
	}
}
