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
	switch event{
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
func makeCultureTree(chronicle []types.CultureYearLocalChronicle) types.CultureLocalTreeHistory{
	tree:=types.CultureLocalTreeHistory{}
	originLocal:=types.LocalCultureTreeHistory{chronicle[0].LocalCulture, [2]int{1, 0}}
	originCultura:=types.CultureTreeHistory{chronicle[0].Culture, [2]int{1, 0}, []types.LocalCultureTreeHistory{originLocal}}
	originBase:=types.BaseCultureTreeHistory{chronicle[0].BaseCulture, [2]int{1, 0}, []types.CultureTreeHistory{originCultura}}
	tree=append(tree, originBase)
	base := 0
	cul := 0
	loc:=0
	for i:=1; i<=len(chronicle)-1; i++ {
		switch chronicle[i].Event {
		case 2:
			tree[base].Cultures[cul].LocalCultures[loc].AlphaOmega[1]=i+1
			tree[base].Cultures[cul].LocalCultures=append(tree[base].Cultures[cul].LocalCultures, types.LocalCultureTreeHistory{chronicle[i].LocalCulture, [2]int{i+1, 0}})
			loc++
			break
		case 3:
			tree[base].Cultures[cul].LocalCultures[loc].AlphaOmega[1]=i+1
			tree[base].Cultures[cul].AlphaOmega[1]=i+1
			tree[base].Cultures=append(tree[base].Cultures, types.CultureTreeHistory{chronicle[i].Culture, [2]int{i+1, 0}, []types.LocalCultureTreeHistory{{chronicle[0].LocalCulture, [2]int{i+1, 0}}}})
			loc=0
			cul++
			break
		case 4:
			tree[base].Cultures[cul].LocalCultures[loc].AlphaOmega[1]=i+1
			tree[base].Cultures[cul].AlphaOmega[1]=i+1
			tree[base].AlphaOmega[1]=i+1
			tree=append(tree, types.BaseCultureTreeHistory{chronicle[i].BaseCulture, [2]int{i+1, 0}, []types.CultureTreeHistory{types.CultureTreeHistory{chronicle[i].Culture, [2]int{i+1, 0}, []types.LocalCultureTreeHistory{{chronicle[0].LocalCulture, [2]int{i+1, 0}}}}}})
			loc=0
			cul=0
			base++
			break
		default:
		}
	}
	return tree
}

func WriteLocalChronicle(chronicle []types.CultureYearLocalChronicle) {
	fmt.Println("Thus begins a history")
	for i:=0; i<=len(chronicle)-1; i++ {
		eventScriptor(chronicle[i].Event)
		writeOneYearLocal(chronicle, i)
	}
	fmt.Println("Here our story ends")
}

func WriteLocalCultureStages(chronicle []types.CultureYearLocalChronicle) {
	tree:=makeCultureTree(chronicle)
	for i:=0; i<=len(tree)-1; i++ {
		fmt.Println(tree[i].Name, " ", tree[i].AlphaOmega[0], "-", tree[i].AlphaOmega[1])
		for j:=0; j<=len(tree[i].Cultures)-1; j++ {
			fmt.Println("|__", tree[i].Name, tree[i].Cultures[j].Stage, " ", tree[i].Cultures[j].AlphaOmega[0], "-", tree[i].Cultures[j].AlphaOmega[1])
			for k:=0; k<=len(tree[i].Cultures[j].LocalCultures)-1; k++ {
				if (i==len(tree)-1)&&(j==len(tree[i].Cultures)-1) {
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
	fmt.Println(babil("| TOWER OF BABEL v0.02 (15.01.2018)~|"))
	fmt.Println(babil("|~  ~  ~pre-0.03 (27.02.2018)  ~  ~ |"))
	fmt.Println(babil("| ~  ~  ~  ~  ~  ~  ~  ~  ~  ~  ~  ~|"))
	fmt.Println(babil("+-----------------------------------+"))
	fmt.Println()
}

func DrawYearCultureMap (culture types.Cultures, year int, chronica types.GlobalChronicle) {
	//var red = (ct.Bg(ct.Red) | ct.Bold).Paint
	//var blue = (ct.Bg(ct.Blue) | ct.Bold).Paint
	var green = (ct.Bg(ct.Green) | ct.Bold).Paint
	//var chosen = (ct.Bg(ct.Yellow) | ct.Fg(ct.Yellow)).Paint

	for i:=0; i<len(chronica.WorldMap.Tiles); i++ {
		fmt.Println()
		for j:=0; j<len(chronica.WorldMap.Tiles[i]); j++ {
			var tileCulture types.Cultures
			if culture.Type()=="BaseCulture" {tileCulture=chronica.WorldMap.Tiles[i][j].Chronica[year-1].BaseCulture} else
			if culture.Type()=="Culture" {tileCulture=chronica.WorldMap.Tiles[i][j].Chronica[year-1].Culture} else
			if culture.Type()=="LocalCulture" {tileCulture=chronica.WorldMap.Tiles[i][j].Chronica[year-1].LocalCulture} else
			if culture.Type()=="SubCulture" {tileCulture=chronica.WorldMap.Tiles[i][j].Chronica[year-1]}

			//if (i==0)&&(j==5) {
			//	fmt.Print(chosen(worldMap.Tiles[i][j].Geography))
			//} else
			if tileCulture==culture {
				fmt.Print(green(chronica.WorldMap.Tiles[i][j].Geography))
			} else {
				fmt.Print(chronica.WorldMap.Tiles[i][j].Geography)
			}
		}
	}
}