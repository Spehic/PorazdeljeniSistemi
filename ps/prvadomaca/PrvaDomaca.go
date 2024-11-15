package main

import (
	"flag"
	"fmt"
	"sync"
	"strings"
	"unicode"
	"github.com/laspp/PS-2024/vaje/naloga-1/koda/xkcd"
)

var wg 	sync.WaitGroup
var mut sync.Mutex
var res map[string]int = make(map[string]int)

//pusti skozi le stevke in crke
func validCharacter(r rune) rune {
	if unicode.IsLetter(r) || unicode.IsNumber(r) {
		return unicode.ToLower(r)
	}

	return ' '
}

//pocisti niz v skladu z mojimi zeljami v funkciji validCharacter
func cleanString(s string) string {
	return strings.Map( validCharacter, s );	
}

//prestej pogostos besed
func count(offset int, length int) {
	defer wg.Done()	
	
	pogostost := make(map[string]int)

	for i := offset; i < offset + length; i++ {
		strip, error := xkcd.FetchComic(i)
		if error != nil {
			fmt.Println( "Napaka pri fetchanju stripa!")
		}
		
		
		// prestej besede v title, transcript ali tooltip
		totalString := cleanString( strip.Title ) + " "
		if strip.Transcript == "" {
			totalString += cleanString( strip.Tooltip )
		} else {
			totalString += cleanString( strip.Transcript )
		}

		//vse besede
		besede := strings.Split(totalString, " ")
			
		for _, beseda := range besede {
			if len(beseda) < 4 {
				continue;
			}
			pogostost[beseda] += 1	
		}

	}

	// Zdaj prenesemo naše vrednosti v globalen slovar
	mut.Lock()
	for k, v := range pogostost {
		res[k] += v
	}
	mut.Unlock()

}

type KeyValue struct{
	k string
	v int
}

func main(){
	//flags za stevilo goroutin
	goNums := flag.Int("g", 8, "num of goroutines")
	flag.Parse()

	//inicializiraj map
	
	//fetchaj strip na 0 da pridobim stevilo vseh stripov
	zadnjiStrip, error := xkcd.FetchComic(0)
	if error != nil {
		fmt.Println("Napaka pri fetchanju prvega stripa")
	}

	steviloVsehStripov := zadnjiStrip.Id
	
	//enakomerno razdeli naloge med vse gorutine
	curOffset := 1
	for i := 0; i < *goNums; i = i+1 {
		if i < steviloVsehStripov % *goNums {
			wg.Add(1)
			go count(curOffset,steviloVsehStripov / *goNums + 1)
			curOffset += steviloVsehStripov / *goNums + 1
		} else {
			wg.Add(1)
			go count(curOffset,steviloVsehStripov / *goNums)
			curOffset += steviloVsehStripov / *goNums
		}
	}

	wg.Wait()
	
	var all []KeyValue
	for k, v := range res {
		append(all, KeyValue{k,v})
	}

	fmt.Println( res )
}
