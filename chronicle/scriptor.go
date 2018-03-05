package chronicle

import (
	"github.com/rossus/tower_of_babel/common/types"
	"fmt"
	"github.com/kortschak/ct"
)

func writeOneYearLocal(chronica []types.CultureYearLocalChronicle, year int) {
	fmt.Println("--------------")
	fmt.Println("Year: ", chronica[year].Year)
	fmt.Println("Name: ", chronica[year].Name)
	fmt.Println("Stage: ", chronica[year].Stage)
	fmt.Println("Substage: ", chronica[year].SubStage)
	fmt.Println(chronica[year].Code)
	fmt.Println("--------------")
}

func eventScriptor(event int) {
	switch event {
	case 0:
		break
	case 1:
		fmt.Println("Subculture changed!")
		break
	case 2:
		fmt.Println("Local culture changed!")
		break
	case 3:
		fmt.Println("Culture changed!")
		break
	case 4:
		fmt.Println("Base culture changed!")
		break
	default:
		fmt.Println("Event error!")
	}
}

//TODO: Refactor culture tree appending
func makeCultureTree(chronicle []types.CultureYearLocalChronicle) types.CultureLocalTreeHistory {
	tree := types.CultureLocalTreeHistory{}
	originLocal := types.LocalCultureTreeHistory{chronicle[0].LocalCulture, [2]int{1, 0}}
	originCultura := types.CultureTreeHistory{chronicle[0].Culture, [2]int{1, 0}, []types.LocalCultureTreeHistory{originLocal}}
	originBase := types.BaseCultureTreeHistory{chronicle[0].BaseCulture, [2]int{1, 0}, []types.CultureTreeHistory{originCultura}}
	tree = append(tree, originBase)
	base := 0
	cul := 0
	loc := 0
	for i := 1; i <= len(chronicle)-1; i++ {
		switch chronicle[i].Event {
		case 2:
			tree[base].Cultures[cul].LocalCultures[loc].AlphaOmega[1] = i + 1
			tree[base].Cultures[cul].LocalCultures = append(tree[base].Cultures[cul].LocalCultures, types.LocalCultureTreeHistory{chronicle[i].LocalCulture, [2]int{i + 1, 0}})
			loc++
			break
		case 3:
			tree[base].Cultures[cul].LocalCultures[loc].AlphaOmega[1] = i + 1
			tree[base].Cultures[cul].AlphaOmega[1] = i + 1
			tree[base].Cultures = append(tree[base].Cultures, types.CultureTreeHistory{chronicle[i].Culture, [2]int{i + 1, 0}, []types.LocalCultureTreeHistory{{chronicle[0].LocalCulture, [2]int{i + 1, 0}}}})
			loc = 0
			cul++
			break
		case 4:
			tree[base].Cultures[cul].LocalCultures[loc].AlphaOmega[1] = i + 1
			tree[base].Cultures[cul].AlphaOmega[1] = i + 1
			tree[base].AlphaOmega[1] = i + 1
			tree = append(tree, types.BaseCultureTreeHistory{chronicle[i].BaseCulture, [2]int{i + 1, 0}, []types.CultureTreeHistory{types.CultureTreeHistory{chronicle[i].Culture, [2]int{i + 1, 0}, []types.LocalCultureTreeHistory{{chronicle[0].LocalCulture, [2]int{i + 1, 0}}}}}})
			loc = 0
			cul = 0
			base++
			break
		default:
		}
	}
	return tree
}

func WriteLocalChronicle(chronicle []types.CultureYearLocalChronicle) {
	fmt.Println("Thus begins a history")
	for i := 0; i <= len(chronicle)-1; i++ {
		eventScriptor(chronicle[i].Event)
		writeOneYearLocal(chronicle, i)
	}
	fmt.Println("Here our story ends")
}

func WriteLocalCultureStages(chronicle []types.CultureYearLocalChronicle) {
	tree := makeCultureTree(chronicle)
	for i := 0; i <= len(tree)-1; i++ {
		fmt.Println(tree[i].Name, " ", tree[i].AlphaOmega[0], "-", tree[i].AlphaOmega[1])
		for j := 0; j <= len(tree[i].Cultures)-1; j++ {
			fmt.Println("|__", tree[i].Name, tree[i].Cultures[j].Stage, " ", tree[i].Cultures[j].AlphaOmega[0], "-", tree[i].Cultures[j].AlphaOmega[1])
			for k := 0; k <= len(tree[i].Cultures[j].LocalCultures)-1; k++ {
				if (i == len(tree)-1) && (j == len(tree[i].Cultures)-1) {
					fmt.Print(" ")
				} else {
					fmt.Print("|")
				}
				fmt.Println("   |__", tree[i].Name, tree[i].Cultures[j].Stage, "(", tree[i].Cultures[j].LocalCultures[k].SubStage, ")  ", tree[i].Cultures[j].LocalCultures[k].AlphaOmega[0], "-", tree[i].Cultures[j].LocalCultures[k].AlphaOmega[1])
			}
		}
	}
}

func WriteScriptorHeader() {
	var babil = (ct.Bg(ct.Blue) | ct.Fg(ct.BoldYellow)).Paint

	fmt.Println(babil("+-----------------------------------+"))
	fmt.Println(babil("|~  ~  ~  ~  ~  ~  ~  ~  ~  ~  ~  ~ |"))
	fmt.Println(babil("| TOWER OF BABEL v0.03 (05.03.2018)~|"))
	fmt.Println(babil("|~  ~  ~  ~  ~  ~  ~  ~  ~  ~  ~  ~ |"))
	//fmt.Println(babil("| ~  ~  ~  ~  ~  ~  ~  ~  ~  ~  ~  ~|"))
	fmt.Println(babil("+-----------------------------------+"))
	fmt.Println()
}

func DrawYearCultureMap(culture types.Cultures, year int, chronica types.GlobalChronicle) {
	defer fmt.Println()
	//var red = (ct.Bg(ct.Red) | ct.Bold).Paint
	//var blue = (ct.Bg(ct.Blue) | ct.Bold).Paint
	var green = (ct.Bg(ct.Green) | ct.Bold).Paint
	//var chosen = (ct.Bg(ct.Yellow) | ct.Fg(ct.Yellow)).Paint

	for i := 0; i < len(chronica.WorldMap.Tiles); i++ {
		fmt.Println()
		for j := 0; j < len(chronica.WorldMap.Tiles[i]); j++ {
			var tileCulture types.Cultures
			if culture.Type() == "BaseCulture" {
				tileCulture = chronica.WorldMap.Tiles[i][j].Chronica[year-1].BaseCulture
			} else
			if culture.Type() == "Culture" {
				tileCulture = chronica.WorldMap.Tiles[i][j].Chronica[year-1].Culture
			} else
			if culture.Type() == "LocalCulture" {
				tileCulture = chronica.WorldMap.Tiles[i][j].Chronica[year-1].LocalCulture
			} else
			if culture.Type() == "SubCulture" {
				tileCulture = chronica.WorldMap.Tiles[i][j].Chronica[year-1].SubCulture
			}

			//if (i==0)&&(j==5) {
			//	fmt.Print(chosen(worldMap.Tiles[i][j].Geography))
			//} else
			if tileCulture == culture {
				fmt.Print(green(chronica.WorldMap.Tiles[i][j].Geography))
			} else {
				fmt.Print(chronica.WorldMap.Tiles[i][j].Geography)
			}
		}
	}
}

func TellMeAStory(chronica types.GlobalChronicle) {
	final:=len(chronica.Chronica)
	fmt.Println()
	origin := chronica.WorldMap.Tiles[0][0].Chronica[0].SubCulture
	fmt.Printf("This story began at the year 1, when entire world was of %s culture...", origin.Name)
	DrawYearCultureMap(origin.LocalCulture, 1, chronica)
	fmt.Println()
	year, _ := origin.Area(chronica)
	if year[len(year)-1]==final {
		fmt.Printf("And this subculture still at it's highest point at the year %v", final)
	} else {
		fmt.Printf("Original subculture began to decay at the year %v...", year[len(year)-1]+1)
		DrawYearCultureMap(origin, year[len(year)-1]+1, chronica)
		fmt.Println()
		yearEnd := origin.Ended(chronica)
		if yearEnd==0 {
			fmt.Printf("This subculture still exists at the year %v", final)
			DrawYearCultureMap(origin, final, chronica)
		} else {
			fmt.Printf("The year %v was the last year for that subculture...", yearEnd)
			DrawYearCultureMap(origin, yearEnd, chronica)
		}
	}
	fmt.Println()
	year, _ = origin.LocalCulture.Area(chronica)
	if year[len(year)-1]==final {
		fmt.Printf("This local culture still at it's highest point at the year %v", final)
	} else {
		fmt.Printf("Original local culture began to decay at the year %v...", year[len(year)-1]+1)
		DrawYearCultureMap(origin.LocalCulture, year[len(year)-1]+1, chronica)
		fmt.Println()
		yearEnd := origin.LocalCulture.Ended(chronica)
		if yearEnd==0 {
			fmt.Printf("This local culture still exists at the year %v", final)
			DrawYearCultureMap(origin.LocalCulture, final, chronica)
		} else {
			fmt.Printf("The year %v was the last year for that local culture...", yearEnd)
			DrawYearCultureMap(origin.LocalCulture, yearEnd, chronica)
		}
	}
	fmt.Println()
	year, _ = origin.Culture.Area(chronica)
	if year[len(year)-1]==final {
		fmt.Printf("This culture still at it's highest point at the year %v", final)
	} else {
		fmt.Printf("Original culture began to decay at the year %v...", year[len(year)-1]+1)
		DrawYearCultureMap(origin.Culture, year[len(year)-1]+1, chronica)
		fmt.Println()
		yearEnd := origin.Culture.Ended(chronica)
		if yearEnd==0 {
			fmt.Printf("This culture still exists at the year %v", final)
			DrawYearCultureMap(origin.Culture, final, chronica)
		} else {
			fmt.Printf("The year %v was the last year for that culture...", yearEnd)
			DrawYearCultureMap(origin.Culture, yearEnd, chronica)
		}
	}
	fmt.Println()
	year, _ = origin.BaseCulture.Area(chronica)
	if year[len(year)-1]==final {
		fmt.Printf("This base culture still at it's highest point at the year %v", final)
	} else {
		fmt.Printf("Original base culture began to decay at the year %v...", year[len(year)-1]+1)
		DrawYearCultureMap(origin.BaseCulture, year[len(year)-1]+1, chronica)
		fmt.Println()
		yearEnd := origin.BaseCulture.Ended(chronica)
		if yearEnd==0 {
			fmt.Printf("This base culture still exists at the year %v", final)
			DrawYearCultureMap(origin.BaseCulture, final, chronica)
		} else {
			fmt.Printf("The year %v was the last year for that base culture...", yearEnd)
			DrawYearCultureMap(origin.BaseCulture, yearEnd, chronica)
		}
	}
	fmt.Println()
}


func Atlas(chronica types.GlobalChronicle) {
	fmt.Println()
	fmt.Print("Original culture at the year 500:")
	DrawYearCultureMap(chronica.WorldMap.Tiles[0][0].Chronica[0].BaseCulture, 500, chronica)
	fmt.Println()
	fmt.Print("Original culture at the year 1000:")
	DrawYearCultureMap(chronica.WorldMap.Tiles[0][0].Chronica[0].BaseCulture, 1000, chronica)
	fmt.Println()
	fmt.Print("Original culture at the year 2000:")
	DrawYearCultureMap(chronica.WorldMap.Tiles[0][0].Chronica[0].BaseCulture, 2000, chronica)
	fmt.Println()
	fmt.Print("Original culture at the year 3000:")
	DrawYearCultureMap(chronica.WorldMap.Tiles[0][0].Chronica[0].BaseCulture, 3000, chronica)
	fmt.Println()
	fmt.Print("Original culture at the year 4000:")
	DrawYearCultureMap(chronica.WorldMap.Tiles[0][0].Chronica[0].BaseCulture, 4000, chronica)
	fmt.Println()
	fmt.Print("Original culture at the year 5000:")
	DrawYearCultureMap(chronica.WorldMap.Tiles[0][0].Chronica[0].BaseCulture, 5000, chronica)

	fmt.Println()
}