package main

import (
	"fmt"
	"testing" 
"sort")

func TestReadCsv(t *testing.T){
	data := read_csv_return_tuple_array("../testdataset.csv")
	if data[0][0] != "http://toto.com"{
		t.Fatalf("Read csv as fail, exepted : %s , actual : %s", "http://toto.com", data[0][0])
	}
	if data[0][1] != "http://titi.com"{
		t.Fatalf("Read csv as fail, exepted : %s , actual : %s", "http://titi.com", data[0][1])
	}
}

func TestBuildSetFromLink(t *testing.T){
	data := read_csv_return_tuple_array("../testdataset.csv")
	set := create_set_from_link(data)
	if len(set) != 2{
		t.Fatalf("Set as not the good lenght, exepted : %d , actual : %d", 2, len(set))
	}

	if set["http://toto.com"] != 0.5{
		t.Fatalf("The toto.com entry as not the good score, exepted : %f , actual : %f", 0.5, set["http://toto.com"])
	}

	if set["http://titi.com"] != 0.5{
		t.Fatalf("The titi.com entry as not the good score, exepted : %f , actual : %f", 0.5, set["http://titi.com"])
	}
}

func Test_rank_one_entry(t *testing.T){
	data := map[string]float32{"toto.com":1.0}
	iterate_rank(1, data, [][]string{})
	if data["toto.com"] != 1.0{
		t.Fatalf("toto.com has not the good score, exepted : %f , actual : %f", 1.0, data["toto.com"])
	}
}

func Test_rank_two_entry_unref(t *testing.T){

	data := map[string]float32{"toto.com":0.5,"titi.com":0.5}
	iterate_rank(1,data, [][]string{})
	if data["toto.com"] != 0.5{
		t.Fatalf("toto.com has not the good score, exepted : %f , actual : %f", 0.5, data["toto.com"])
	}

	if data["titi.com"] != 0.5{
		t.Fatalf("titi.com has not the good score, exepted : %f , actual : %f", 0.5, data["titi.com"])
	}
}

func Test_rank_two_entry_with_ref(t *testing.T){
	data := map[string]float32{"toto.com":0.5,"titi.com":0.5}
	iterate_rank(1,data, [][]string{{"toto.com", "titi.com"}})
	var toto float32 = 0.0375
	var titi float32 = 0.9625

	if fmt.Sprintf("%f",data["toto.com"]) != fmt.Sprintf("%f",toto){
		t.Fatalf("toto.com has not the good score, exepted : %f , actual : %f", toto, data["toto.com"])
	}

	if  fmt.Sprintf("%f",data["titi.com"]) != fmt.Sprintf("%f",titi){
		t.Fatalf("titi.com has not the good score, exepted : %f , actual : %f", titi, data["titi.com"])
	}
}

func Test_rank_two_entry_with_ref_three_iteration(t *testing.T){
	data := map[string]float32{"toto.com":0.5,"titi.com":0.5}
	iterate_rank(3, data, [][]string{{"toto.com", "titi.com"}})
	var toto float32 = 0.0002109375
	var titi float32 = 0.9997890625

	if fmt.Sprintf("%f",data["toto.com"]) != fmt.Sprintf("%f",toto){
		t.Fatalf("toto.com has not the good score, exepted : %f , actual : %f", toto, data["toto.com"])
	}

	if fmt.Sprintf("%f",data["titi.com"]) != fmt.Sprintf("%f",titi){
		t.Fatalf("titi.com has not the good score, exepted : %f , actual : %f", titi, data["titi.com"])
	}
}

func TestSortLink(t *testing.T){
	data := [][]string{{"b","a"},{"c","a"},{ "b", "c"},{"a","b"},{"b","d"}}
	sort.Sort(byFrom(data))

	if data[0][0] != "a"{
		t.Fatal("the first element may be a B")
	}
	if data[1][0] != "b"{
		t.Fatal("the second element may be a B")
	}
	if data[2][0] != "b"{
		t.Fatal("the three element may be a B")
	}
	if data[3][0] != "b"{
		t.Fatal("the four element may be a B")
	}
	if data[4][0] != "c"{
		t.Fatal("the five element may be a C")
	}
}