package main

import (
	"log"
	"os"

	"github.com/Anirudh12345678/torrent-client/torrentfiles"
)

func main() {
	inPath := os.Args[1]
	outPath := os.Args[2]

	tf, err := torrentfiles.Open(inPath)
	if err != nil {
		log.Fatal(err)
	}

	err = tf.DownloadToFile(outPath)
	if err != nil {
		log.Fatal(err)
	}
}
