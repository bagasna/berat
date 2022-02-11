package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Data []map[string]interface{}
	// tanggal   string
	// max       int
	// min       int
	// perbedaan int
}

var result Response

func showData() Response {
	return Response{
		Data: []map[string]interface{}{},
	}
}

func (res *Response) ShowResultAll() {
	for _, res := range res.Data {
		hasil, _ := json.MarshalIndent(res, "", " ")
		fmt.Println(string(hasil))
	}
}

func (res *Response) Create(tanggal string, max int, min int) {

	perbedaan := max - min

	hasil := map[string]interface{}{
		"tanggal":   tanggal,
		"max":       max,
		"min":       min,
		"perbedaan": perbedaan,
	}

	res.Data = append(res.Data, hasil)
}

func (res *Response) ShowResultSingle(tanggal string) {
	for x, index := range res.Data {
		if index["tanggal"] == tanggal {
			hasil, _ := json.MarshalIndent(res.Data[x], "", " ")
			fmt.Println(string(hasil))
		}
	}
}

func (res *Response) Edit(tanggal string) {
	for i, index := range res.Data {
		if index["tanggal"] == tanggal {
			delete(res.Data[i], tanggal)
		}
	}
}

func (res *Response) Delete(tanggal string) {
	for i, index := range res.Data {
		if index["tanggal"] == tanggal {
			// fmt.Println(res.Data[i])
			delete(res.Data[i], tanggal)
		}
	}
}

func main() {
	res := showData()
	res.Create("21-02-2022", 55, 50)
	res.Create("22-02-2022", 53, 50)
	res.Create("23-02-2022", 52, 50)
	// res.Delete("21-02-2022")
	res.ShowResultSingle("22-02-2022")
	fmt.Println("=====================================")
	res.ShowResultAll()

}
