package main

import (
	"log"
	"os"

	flacvorbis "github.com/go-flac/flacvorbis"
	"github.com/go-flac/go-flac"
)

func main() {
	f, err := flac.ParseFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	var cmt *flacvorbis.MetaDataBlockVorbisComment
	for _, meta := range f.Meta {
		if meta.Type == flac.VorbisComment {
			cmt, err = flacvorbis.ParseFromMetaDataBlock(*meta)
			if err != nil {
				panic(err)
			}
			log.Printf("found comments %v.\n", cmt.Comments)
			cmt.Comments = nil
			meta.Data = cmt.Marshal().Data
			log.Printf("removed comments.")

		}
	}
	f.Save(os.Args[1])
}
