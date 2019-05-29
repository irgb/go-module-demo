package main

import (
    "golang.org/x/text/message"
    hellomod "git.shannonai.com/taoshibo/hellomod"
    hellomodV2 "git.shannonai.com/taoshibo/hellomod/v2"
    "github.com/irgb/go-demo-lib/pkg/dump"
    "github.com/irgb/go-module-demo/internal/util"
    "github.com/irgb/go-module-demo/internal/common"
)

func main() {
    p := message.NewPrinter(message.MatchLanguage("en"))
    p.Println(123456.78)
    hellomod.SayHello()
    hellomodV2.SayHello()
    util.Logger.Info("info")
    util.Logger.Fatal("fatal")
    dump.Dump("ahaha")
}
