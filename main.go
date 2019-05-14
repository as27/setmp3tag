package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	log.Println("Starting renaming mp3 tags")
	log.Println("Expecting folders like:")
	log.Println(".../genres/album/title.mp3")
	startPath := "./."
	if len(os.Args) == 2 {
		startPath = os.Args[1]
	}
	scanDir("", "", startPath)
	/*
		filepath.Walk(startPath, func(path string, info os.FileInfo, err error) error {
			log.Println(path, info.Name())
			return nil
		})
		/*
			tag, err := id3v2.Open("test/File1.mp3", id3v2.Options{Parse: true})
			if err != nil {
				log.Fatal("Error while opening mp3 file: ", err)

			}
			defer tag.Close()
			fmt.Println(tag.Artist())
			fmt.Println(tag.Title())
			fmt.Println(tag.Album())
			tag.SetAlbum("Neues Album")
			fmt.Println(tag.Album())
			tag.Save()
	*/
	// /name/genres/album/title
}

func scanDir(genres, album, path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, info := range files {
		if info.IsDir() {
			return scanDir(album, info.Name(), filepath.Join(path, info.Name()))
		}
		log.Println(path, genres, album, info.Name())
	}

	return nil
}
