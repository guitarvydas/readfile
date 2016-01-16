package readfile

import (
	"bufio"
	"github.com/guitarvydas/ip"
	"io"
	"os"
)

func Read(name string, path <-chan string, out chan ip.IP) {
	inFile, err := os.Open(<-path)
	defer sendEof(name, out)
	if err != nil {
		os.Exit(1)
	}
	defer inFile.Close()
	r := bufio.NewReader(inFile)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF || err != nil {
			break
		}
		out <- ip.IP{Kind: ip.Normal, Data: line}
	}
}

func sendEof(name string, out chan ip.IP) {
	out <- ip.IP{Kind: ip.EOF}
}
