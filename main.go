package main

import (
    hellomod "git.shannonai.com/taoshibo/hellomod"
    hellomodV2 "git.shannonai.com/taoshibo/hellomod/v2"
    "git.shannonai.com/taoshibo/go-module-demo/util"
    "golang.org/x/text/message"
)

func main() {
    p := message.NewPrinter(message.MatchLanguage("en"))
    p.Println(123456.78)
    hellomod.SayHello()
    hellomodV2.SayHello()
    util.Logger.Info("info")
    util.Logger.Fatal("fatal")
}
