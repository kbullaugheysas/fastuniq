package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

type Args struct {
	Count    bool
	MinCount int
}

var args = Args{}

func init() {
	log.SetFlags(0)
	flag.BoolVar(&args.Count, "count", false, "count how many times each string occurs")
	flag.IntVar(&args.MinCount, "min-count", 1, "when counting, only output entries with this min count")
}

func main() {
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	seen := make(map[string]int)
	for scanner.Scan() {
		line := scanner.Text()
		_, ok := seen[line]
		if !ok {
			seen[line] = 0
			if !args.Count {
				fmt.Println(line)
			}
		}
		seen[line] += 1
	}
	if args.Count {
		for k, v := range seen {
			if v >= args.MinCount {
				fmt.Printf("%s\t%d\n", k, v)
			}
		}
	}
}
