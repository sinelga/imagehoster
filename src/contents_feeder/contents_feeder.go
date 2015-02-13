package main

import (
	"contents_feeder/find_adv_phone"
	"contents_feeder/find_all_characters"
	"contents_feeder/find_names"
	"contents_feeder/find_regions"
	"contents_feeder/update_characters"
	"findfreeparagraph"
	"flag"
	"fmt"
	"math/rand"
	"startones"
	"strings"
	"time"
)

const APP_VERSION = "0.1"

// The flag package provides a default help printer via -h switch
var versionFlag *bool = flag.Bool("v", false, "Print the version number.")

func main() {
	flag.Parse() // Scan the arguments list

	if *versionFlag {
		fmt.Println("Version:", APP_VERSION)
	}

	locale := "fi_FI"
	themes := "porno"

	golog, config := startones.Start()

	adv_phone_id := find_adv_phone.FindAll(golog, config)
	regions_id := find_regions.FindAll(golog, config)

	names := find_names.FindAll(golog, config)

	characters := find_all_characters.FindAll(golog, config)
	
	

	for n, character := range characters {

		if character.Moto == "" || character.Description == "" {

			paragraph := findfreeparagraph.FindFromQ(golog, locale, themes, config)

			if character.Moto == "" {

				characters[n].Moto = paragraph.Ptitle

			}
			if character.Description == "" {
				s := []string{paragraph.Pphrase, paragraph.Sentences[0], paragraph.Sentences[1], paragraph.Sentences[2]}
				description := strings.Join(s, " ")
				characters[n].Description = description

			}

		}

		rand.Seed(time.Now().UTC().UnixNano())

		if character.Adv_phone_id == 0 {

			i := rand.Intn(len(adv_phone_id))
			characters[n].Adv_phone_id = adv_phone_id[i]

		}

		if character.Region_id == 0 {

			i := rand.Intn(len(regions_id))
			characters[n].Region_id = regions_id[i]

		}

		if character.Name == "" {

			i := rand.Intn(len(names))
			characters[n].Name = names[i]

		}

	}

	update_characters.Update(golog, config, characters)

}
