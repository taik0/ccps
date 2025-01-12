package mus

// Tools:
// https://github.com/vgmrips/vgmtools#vgm-text-writer-vgm2txt
import (
	"ccps/boards"
	"ccps/code"
)

func Build(v bool, board *boards.Board) *code.Code {
	// Here we open a VGM file, parse it and generate
	// a source file to be compiled with the Z80 source
	// code.
	return code.NewCode()
}
