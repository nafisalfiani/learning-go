package main

import (
	"fmt"
	"sort"
	"strings"
)

type Klasemen struct {
	klub []string
	poin map[string]int
}

func NewKlasemen(klub []string) *Klasemen {
	poin := make(map[string]int)
	for _, k := range klub {
		poin[k] = 0
	}
	return &Klasemen{klub, poin}
}

func (k *Klasemen) catatPermainan(klubKandang, klubTandang, skor string) {
	gol := strings.Split(skor, ":")
	golKandang, golTandang := parseGol(gol)

	if golKandang > golTandang {
		k.poin[klubKandang] += 3
	} else if golKandang == golTandang {
		k.poin[klubKandang]++
		k.poin[klubTandang]++
	} else {
		k.poin[klubTandang] += 3
	}
}

func (k *Klasemen) cetakKlasemen() map[string]int {
	return k.poin
}

func (k *Klasemen) ambilPeringkat(nomorPeringkat int) string {
	klubSorted := make([]string, len(k.klub))
	copy(klubSorted, k.klub)
	sort.Slice(klubSorted, func(i, j int) bool {
		return k.poin[klubSorted[i]] > k.poin[klubSorted[j]]
	})
	return klubSorted[nomorPeringkat-1]
}

func parseGol(gol []string) (int, int) {
	golKandang := parseInt(gol[0])
	golTandang := parseInt(gol[1])
	return golKandang, golTandang
}

func parseInt(s string) int {
	var result int
	fmt.Sscanf(s, "%d", &result)
	return result
}

func main() {
	klub := []string{"Liverpool", "Chelsea", "Arsenal"}
	klasemen := NewKlasemen(klub)

	klasemen.catatPermainan("Arsenal", "Liverpool", "2:1")
	klasemen.catatPermainan("Arsenal", "Chelsea", "1:1")
	klasemen.catatPermainan("Chelsea", "Arsenal", "0:3")
	klasemen.catatPermainan("Chelsea", "Liverpool", "3:2")
	klasemen.catatPermainan("Liverpool", "Arsenal", "2:2")
	klasemen.catatPermainan("Liverpool", "Chelsea", "0:0")

	fmt.Println(klasemen.cetakKlasemen())

	fmt.Println(klasemen.ambilPeringkat(2))
}
