package main

import b3gen "github.com/Jacob953/behavior3gen"

func main() {
	g := b3gen.NewGenerator(b3gen.Config{})
	g.GenerateNodeFrom(Register)
	g.Execute()
}
