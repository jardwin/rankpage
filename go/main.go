package main

import (
    "encoding/csv"
    "fmt"
    //"log"
    "sort"
    "os")

func main(){
    if len(os.Args) < 2 {
        fmt.Println("Pas de chemin passer")
        return
    }

    data := read_csv_return_tuple_array(os.Args[1])
    
    set := create_set_from_link(data)

    //rank(set,data)

    for k, v := range set{
        fmt.Println(k," => ",v)
    }
}


func read_csv_return_tuple_array(path string)[][]string{
    f, err := os.Open(path)
    if err != nil{
        fmt.Println("Chemin eroner")
        return [][]string{}
    }
    defer f.Close()

    var csvReader = csv.NewReader(f)
    csvReader.Comma = ';'
    csvReader.Read()
    data, err := csvReader.ReadAll()

    if err != nil {
        fmt.Println("Fichier n'est pas un csv")
    }
    return data
}


func create_set_from_link(links [][]string) map[string]float32{
    var filtered = make(map[string]float32)
    for _, link := range links{
        if _, ok := filtered[link[0]]; !ok{
                filtered[link[0]] = 0.00
        }
        if _, ok := filtered[link[1]]; !ok{
            filtered[link[1]] = 0.00
        }
    }

    for k, _ := range filtered{
        filtered[k]=1.0/float32(len(filtered))
    }
    return filtered
}

func rank(url_set map[string]float32, url_link [][]string) float32{
    for url, _ := range url_set{
        for _,link := range findLinkToMe(url_link, url){
            url_set[url] += (url_set[link[0]]*0.85)/float32(countLinkFromMe(url_link, link[0]))
        }
    }
}

type byFrom [][]string

func (a byFrom) Len() int           { return len(a) }
func (a byFrom) Swap(i, j int)      { a[i][0], a[j][0] = a[j][0], a[i][0] }
func (a byFrom) Less(i, j int) bool { return a[i][0] < a[j][0] }

func iterate_rank(n int, set map[string]float32, link[][]string){
    sort.Sort(byFrom(link))
    for i := 0; i < n; i++ {
        rest = rank(set, link)
        if len(link)!=0{
            distribut_rank(set, link)
        }
    }
}

func distribut_rank(set map[string]float32, links [][]string){
    rest := float32(0)
    for _, link := range links{
        rest+=float32(set[link[0]])
        set[link[0]]=0
    }
    for page, _ := range set{
        set[page]+=rest/float32(len(set))
    }
}

func findLinkToMe(urlLinks [][]string, me string)[][]string{
    var linked = [][]string{}
    for _,url := range urlLinks{
        if url[1] == me && findElementInDoubleArray(linked,me){
            linked = append(linked, url)
        }
    }
    return linked
}

func countLinkFromMe(urlLinks [][]string, me string) int{
    var cmp = 0
    for _,url := range urlLinks{
        if url[0] == me{
            cmp+=1
        }
    }
    return cmp
}

func findElementInDoubleArray(arr [][]string, element string) bool{
    for _, val := range arr{
        if val[1] == element{
            return true
        }
    }
    return false
}

func findElementIn(arr []string, element string) bool{
    for _, val := range arr{
        if val == element{
            return true
        }
    }
    return false
}