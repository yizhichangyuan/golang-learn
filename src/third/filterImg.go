package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"learn/src/third/xckd"
	"log"
	"os"
)

var (
	f = flag.Bool("f", false, "")
	n = flag.Int("n", 100, "")
)

func main() {
	flag.Parse()
	if *f {
		if *n > xckd.MaxNum {
			log.Fatalf("%d can't bigger than %d", xckd.MaxNum)
		}
		fetch(*n)
	} else {
		search(flag.Args())
	}
}

func fetch(n int) {
	var cartons []*xckd.Carton
	for num := xckd.MinNum; num < n; num++ {
		c, err := xckd.Get(num)
		if err != nil {
			log.Fatal(err)
		}
		cartons = append(cartons, c)
	}
	index := xckd.New()
	index.Comics = cartons
	out, err := json.MarshalIndent(index, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", out)
}

func search(keywords []string) {
	in, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	index := xckd.New()
	if err := json.Unmarshal(in, &index); err != nil {
		log.Fatal(err)
	}
	result := xckd.Filter(keywords, index.Comics)
	for _, c := range result {
		fmt.Println(*c)
	}
}
