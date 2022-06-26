package main

import (
    "encoding/csv"
    "fmt"
    "strconv"
    "sort"
    "os")

func main(){
    if len(os.Args) < 3 {
        fmt.Println("Pas de chemin passer ou de nombre d'iteration")
        return
    }

    data := read_csv_return_tuple_array(os.Args[1])
    
    set := create_set_from_link(data)

    i, err := strconv.Atoi(os.Args[2])
    if err != nil{
        fmt.Println("Con't read the number of iteration")   
        return
    }
    iterate_rank(i, set, data)

    f, err := os.Create("result.csv")
    defer f.Close()

    var csvWriter = csv.NewWriter(f)
    csvWriter.Comma = ';'
    csvWriter.Write([]string{"url", "note"})

    for url, val := range set{
        csvWriter.Write([]string{url, fmt.Sprintf("%f", val)})
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
    rest := float32(0)
    prevUrl := ""
    if len(url_link) != 0{
        prevUrl = url_link[0][0]
    }
    for _, url := range url_link{
        if prevUrl != url[0]{
            rest += url_set[prevUrl]*0.15
            url_set[prevUrl]*=0.15
        }
        url_set[url[1]] += (url_set[url[0]]*0.85)/float32(countLinkFromMe(url_link, url[0]))
        prevUrl = url[0]
    }
    if len(url_link) != 0{
        rest += url_set[prevUrl]*0.15
        url_set[prevUrl]*=0.15
    }

    return rest
}

type byFrom [][]string

func (a byFrom) Len() int           { return len(a) }
func (a byFrom) Swap(i, j int)      { a[i][0], a[j][0] = a[j][0], a[i][0] }
func (a byFrom) Less(i, j int) bool { return a[i][0] < a[j][0] }

func iterate_rank(n int, set map[string]float32, link[][]string){
    sort.Sort(byFrom(link))
    for i := 0; i < n; i++ {
        rest := rank(set, link)
        rest += reduceUnrefPage(set, link)
        distribut_rank(set, rest)
    }
}

func reduceUnrefPage(set map[string]float32, links [][]string) float32{
    rest := float32(0)
    for url, _ := range set{
        if !findElementInDoubleArrayAtOne(links, url) && !findElementInDoubleArray(links, url){
            rest += set[url]
            set[url] = 0
        }

        if findElementInDoubleArrayAtOne(links, url) && !findElementInDoubleArray(links, url){
            set[url] = 0
        }
    }
    return rest
}

func distribut_rank(set map[string]float32, rest float32){
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

func findElementInDoubleArrayAtOne(arr [][]string, element string) bool{
    for _, val := range arr{
        if val[0] == element{
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