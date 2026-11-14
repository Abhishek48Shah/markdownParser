package main

import (
	"bytes"
	"log"

	"github.com/Abhishek48Shah/markdownParser/parser"
)

func main() {
	blog := struct {
		Title   string
		Content string
	}{
		Title: `# The man who sold the world Kurt Kobin`,
		Content: `
		## Story

			The Man Who Sold the World" is a classic song written and originally performed by David Bowie, released as the title track of his third studio album in November 1970
			
		**David Bowie Original (1970)**
		- Release Date: November 4, 1970
		- Album: *The Man Who Sold The World*
		- Producer: Tony Viscounti
		- Style: Psychedelic and hard rock, featuring eerie, shifting guitar lines and mysterious, alienated lyrics. Bowie later shared that writing the song was inspired by a youthful search for identity and an interest in Buddhism.

		**Nirvana Conver (1993/1994)
		1. Recording Date: November 18, 1993 (at Sony Music Studios in New York City for *MTV Unplugged*)
		2. Album: *MTV Unplugged in New York* (released 1994; promotional single in 1995)
		3. Style:  Stripped-down acoustic rock. Kurt Cobain’s haunting rendition introduced the track to a massive new generation of fans, to the point where many younger listeners mistakenly believed it was an original Nirvana song. Bowie expressed deep admiration for Cobain’s version, describing it as having a poignant, mystical quality.
		`,
	}
	var buff bytes.Buffer
	err := parser.Convert([]byte(blog.Title), &buff)
	if err != nil {
		log.Fatalln(err)
	}
}
