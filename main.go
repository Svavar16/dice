package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var (
	size   int
	amount int
)

func main() {
	flag.Parse()

	if amount > 1 {
		fmt.Printf("you will roll: %d times\n", amount)
	}
	for i := 1; i <= amount; i++ {
		fmt.Printf("%s::you rolled: %d\n", time.Now().Format(time.Stamp), dice(size))
	}
}

func init() {
	fUsage(flag.CommandLine)
	flag.IntVar(&size, "size", 6, "size of the dice")
	flag.IntVar(&amount, "amount", 1, "how often you want to roll the dice")
	flag.IntVar(&size, "s", 6, "size of the dice")
	flag.IntVar(&amount, "a", 1, "how often you want to roll the dice/how many dices you want to roll")
}

func dice(size int) int {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(size)
	if value > size {
		return size
	}
	return value + 1
}

func fUsage(f *flag.FlagSet) {
	f.Usage = func() {
		fmt.Println("usage: dice [options] [value] ...")
		fmt.Println("-a, -amount, --amount int")
		fmt.Println("    how often you want to roll the dice/how many dices you want to roll (default 1)")
		fmt.Println("-s, -size, --size int")
		fmt.Println("    size of the dice (default 6)")
		os.Exit(0)
	}
}
