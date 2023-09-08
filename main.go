package main

import (
	. "github.com/littletrainee/PongJongServer/ServerHolder"
)

func main() {
	var SH ServerHolder
	SH.Start()
	SH.Update()
}
