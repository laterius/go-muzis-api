package stringutil

import "fmt"

func Leaning(){
	//silence
	/*silence1 := []int {1,2,3}
	silence2 := make ([]int, 2)
	copy(silence2, silence1)
	fmt.Println(silence2)*/

	//map
	/*testmap :=make(map[string] int)
	testmap["key"] = 10
	fmt.Println(testmap)*/

	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	    elements["He"] = "Helium"
	    elements["Li"] = "Lithium"
	    elements["Be"] = "Beryllium"
	    elements["B"] = "Boron"
	    elements["C"] = "Carbon"
	    elements["N"] = "Nitrogen"
	    elements["O"] = "Oxygen"
	    elements["F"] = "Fluorine"
	    elements["Ne"] = "Neon"

	if name,ok := elements["ok"]; ok{
		fmt.Println(name, ok)
	}

	if name,ok := elements["Be"]; ok{
		fmt.Println(name, ok)
	}

	elements2 := map[string]map[string]string{
            "H": map[string]string{
            "name":"Hydrogen",
            "state":"gas",
        },
        "He": map[string]string{
            "name":"Helium",
            "state":"gas",
        },
        "Li": map[string]string{
            "name":"Lithium",
            "state":"solid",
        },
        "Be": map[string]string{
            "name":"Beryllium",
            "state":"solid",
        },
        "B":  map[string]string{
            "name":"Boron",
            "state":"solid",
        },
        "C":  map[string]string{
            "name":"Carbon",
            "state":"solid",
        },
        "N":  map[string]string{
            "name":"Nitrogen",
            "state":"gas",
        },
        "O":  map[string]string{
            "name":"Oxygen",
            "state":"gas",
        },
        "F":  map[string]string{
            "name":"Fluorine",
            "state":"gas",
        },
        "Ne":  map[string]string{
            "name":"Neon",
            "state":"gas",
        },
}
	if el, ok := elements2["Ne"]; ok{
		fmt.Println(el["name"], el["state"])
	}

}




