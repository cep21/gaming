/**
 * Date: 4/11/14
 * Time: 4:56 PM
 * @author jack 
 */
package main

import (
	"gaming"
	"blackjack"
	"flag"
	"os"
	"math/rand"
	"time"
)

func main() {
	oddsA := flag.Float64("odda", .5, "Odds of option A")
	oddsB := flag.Float64("oddb", .5, "Odds of option B")
	seed := flag.Int("seed", 0, "Random seed for simulation")
	flag.Parse()
	if *seed == 0 {
		*seed = int(time.Now().Unix())
	}
	r := rand.New(rand.NewSource(int64(*seed)))
	for i:=0;i<10000;i++ {
		v1 := r.Float64() * oddsA
		v2 := r.Float64() * oddsA
	}
}

