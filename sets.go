// Using go to take a look at maths for algo and system analysis 

package main 

import "fmt"

type shop_list map[string]struct{}

func (s shop_list) Permute() []map[string]interface{} {

	var ret []map[string]interface{} 
	for key,_ := range s {
		lev := make(map[string]interface{})
		lev[key] = struct{}{}
		s.permute_bfs(lev, &ret)
	}
	return ret
}

func CopyMap(m map[string]interface{}) map[string]interface{} {
    cp := make(map[string]interface{})
    for k, v := range m {
        vm, ok := v.(map[string]interface{})
        if ok {
            cp[k] = CopyMap(vm)
        } else {
            cp[k] = v
        }
    }

    return cp
}

func (s shop_list) permute_bfs(lev map[string]interface{}, ret *[]map[string]interface{}){

	if len(lev) == len(s) {
		l := CopyMap(lev)
		//Why need to copy Map : -- 
		//https://golang.org/doc/effective_go.html#maps
		/*
			Like slices, maps hold references to an underlying 
			data structure. If you pass a map to a function 
			that changes the contents of the map, the changes 
			will be visible in the caller.
		*/
		*ret = append(*ret, l)
		return
	}	
	for key,_ := range s {
		if _,ok := (lev)[key]; ok {
			continue
		} else {
			(lev)[key] = struct{}{}
			s.permute_bfs(lev, ret)
			//remove value from map
			delete(lev, key)
		}

	}
	return
}

func main(){
	//Set : Unordered collection of items
	// - a golang map is a useful DS for abstraction of a set
	// to create null values we use an empty struct -- better performance than nil 

	set := make(shop_list)
	fmt.Printf("s:%T\n", set)

	set["apples"] = struct{}{}
	set["oranges"] = struct{}{}
	set["peaches"] = struct{}{}
	set["eggs"] = struct{}{}
	set["bread"] = struct{}{}
	fmt.Printf("s: %v\n", set)

	//Permutations of set of size n --- n! (factorial(n))
	for i, val := range set.Permute() {
		fmt.Printf("comb %d :  %v\n", i+1, val)
	
	}

	//Permutations of set of size k (with repititions allowed) ---  ((n)^k)
	for i, val := range set.Permute() {
		fmt.Printf("comb %d :  %v\n", i+1, val)
	
	}

	//TO DO:  Permutations of set of size k (no repititions) ---  ()
	/*for i, val := range set.SubPermute() {
		fmt.Printf("comb %d :  %v\n", i+1, val)
	
	}*/

}