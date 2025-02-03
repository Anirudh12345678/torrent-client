package main

import (
	"log"

	"github.com/Anirudh12345678/torrent-client/torrentfiles"
)

func main() {
	tf, err := torrentfiles.Open("./puppy.torrent")
	if err != nil {
		log.Fatal(err)
	}

	err = tf.DownloadToFile("./result")
	if err != nil {
		log.Fatal(err)
	}
}
