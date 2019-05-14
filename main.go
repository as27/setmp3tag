package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/bogem/id3v2"
)

func main() {
	log.Println("Starting renaming mp3 tags")
	log.Println("Expecting folders like:")
	log.Println(".../genres/artist/album/title.mp3")
	startPath := ""
	if len(os.Args) == 2 {
		startPath = os.Args[1]
	}
	scanDir("", "", "", startPath)
}

func scanDir(genres, artist, album, path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, info := range files {
		if info.IsDir() {
			err := scanDir(artist, album, info.Name(), filepath.Join(path, info.Name()))
			if err != nil {
				log.Println("Error calling scanDir: ", err)
			}
			continue
		}
		if useFile(info.Name()) {
			log.Println(path, genres, album, info.Name())
			tag, err := id3v2.Open(
				filepath.Join(path, info.Name()),
				id3v2.Options{Parse: true})
			if err != nil {
				log.Println("Error while opening mp3 file: ", err)
			}
			tag.SetGenre(escapeName(genres))
			tag.SetArtist(escapeName(artist))
			tag.SetAlbum(escapeName(album))
			tag.SetTitle(escapeName(info.Name()))
			tag.Save()
			tag.Close()
		}
	}
	return nil
}

func useFile(fileName string) bool {
	exts := []string{".mp3"}
	for _, ext := range exts {
		if strings.EqualFold(filepath.Ext(fileName), ext) {
			return true
		}
	}
	return false
}

func escapeName(name string) string {
	escChars := []struct {
		old string
		new string
	}{
		{"Ä", "Ae"},
		{"Ö", "Oe"},
		{"Ü", "Ue"},
		{"ä", "ae"},
		{"ö", "oe"},
		{"ü", "ue"},
		{" ", "_"},
		{"ß", "ss"},
		{"Á", "A"},
		{"á", "a"},
		{"É", "E"},
		{"é", "e"},
		{"Í", "I"},
		{"í", "i"},
		{"Ó", "O"},
		{"ó", "o"},
		{"Ú", "U"},
		{"ú", "u"},
		{"", ""},
	}
	for _, esc := range escChars {
		name = strings.ReplaceAll(name, esc.old, esc.new)
	}
	return name
}
