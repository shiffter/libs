package libs

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

func ReadData(path string, p []byte) []byte {
	file, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		os.Exit(1)
	}
	for {
		n, err := file.Read(p)
		if err == io.EOF {
			break
		}
		p = p[:n]
	}
	return p
}

type DBReader interface {
	Read(p []byte, ex string) DBReader
}

type StolenDB JData
type OriginalDB JData

type JsonIngredients struct {
	Name  string `json:"ingredient_name" xml:"itemname"`
	Count string `json:"ingredient_count" xml:"itemcount"`
	Unit  string `json:"ingredient_unit" xml:"itemunit"`
}

type Cake struct {
	Name        s[tring            `json:"name" xml:"name"`
	Time        string            `json:"time" xml:"stovetime"`
	Ingredients []JsonIngredients `json:"ingredients"`
	Ing         []XmlIngredients  `xml:"ingredients"`
}

type JData struct {
	XMLName  xml.Name `xml:"recipes"`
	Database []Cake   `json:"cake"`
}

type XmlIngredients struct {
	Items []XmlIngredients `xml:"item"`
}

type Data struct {
	Database []Cake `xml:"cake"`
}

func ConvertOut(path string) DBReader {
	data := make([]byte, 2049)
	data = ReadData(path, data)
	var i DBReader
	str := (path)[len(path)-4 : len(path)]
	if str == ".xml" {
		i = OriginalDB{}
	} else if str == "json" {
		i = StolenDB{}
	}
	i = i.Read(data, str)
	return i
}

func CheckName(name string, extens []string) bool {
	if len(name) > len(extens) {
		for i, v := range extens {
			fileExt := name[len(name)-len(extens[i]) : len(name)]
			if v == fileExt {
				return true
			}
		}
	}
	return false
}

func (obj StolenDB) Read(p []byte, extens string) DBReader {
	if extens == "json" {
		fmt.Println("json")
		err := json.Unmarshal(p, &obj)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else if extens == ".xml" {
		fmt.Println("json->xml")
		err := xml.Unmarshal(p, &obj)
		data := make([]byte, 2048)
		data, err = xml.Marshal(obj)
		err = xml.Unmarshal(data, &obj)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	return obj
}

func (obj OriginalDB) Read(p []byte, extens string) DBReader {
	if extens == ".xml" {
		fmt.Println("xml")
		err := xml.Unmarshal(p, &obj)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else if extens == "json" {
		fmt.Println("xml->json")
		err := json.Unmarshal(p, &obj)
		data := make([]byte, 2048)
		data, err = xml.Marshal(obj)
		err = xml.Unmarshal(data, &obj)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	return obj
}
