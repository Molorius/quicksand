package main

import "www.github.com/Molorius/quicksand/pkg/bedrock"

func main() {
	b := bedrock.Bedrock{ServerDir: "/opt/mcpe"}
	b.Start()
	b.Stop()
}
