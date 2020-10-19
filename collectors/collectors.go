/*
Package collectors defines collectors for app
Copyright © 2020 Vishnu Rajendran vishnraj@umich.edu

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package collectors

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

var shoeSources = map[string]func(string) []ShoeData{
	"nike": nikeCollect,
}

// ShoeData contains fields from sites that we want to collect
type ShoeData struct {
	Name  string
	Price float64
}

func getSupportedSources() []string {
	supported := make([]string, 0)
	for k := range shoeSources {
		supported = append(supported, k)
	}

	return supported
}

func nikeCollect(shoeType string) []ShoeData {
	data := make([]ShoeData, 0)
	c := colly.NewCollector()

	c.OnHTML("div.product-card__body", func(e *colly.HTMLElement) {
		var name string
		var price float64

		e.ForEach("div.product-card__title", func(i int, h *colly.HTMLElement) {
			name = h.Text
		})
		e.ForEach("div.product-price.is--current-price", func(i int, h *colly.HTMLElement) {
			var err error
			tmp := strings.Replace(h.Text, "$", "", -1)
			price, err = strconv.ParseFloat(tmp, 64)
			if err != nil {
				fmt.Printf("For price: %s encountered error: %s", h.Text, err.Error())
			}
		})

		data = append(data, ShoeData{name, price})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	fmt.Printf("Requesting data for shoe type: %s\n", shoeType)
	c.Visit("https://www.nike.com/w?q=" + shoeType)

	return data
}

// GenerateShoeData generates shoe data from shoe source
func GenerateShoeData(shoeSource string, shoeType string, outfile string) error {
	f, ok := shoeSources[shoeSource]
	if !ok {
		return fmt.Errorf("Error: source %s not supported, only sources:\n%v\nare supported", shoeSource, getSupportedSources())
	}

	data := f(shoeType)
	dump, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	fmt.Printf("Writing output to %s", outfile)
	ioutil.WriteFile(outfile, dump, 0644)

	return nil
}
