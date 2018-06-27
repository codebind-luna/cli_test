package FileParsers


import (
    "path/filepath"
    "encoding/json"
    "fmt"
    "log"
    "github.com/tealeg/xlsx"
)

type structTest struct {
		Name      string     `xlsx:"0"`
		Age       string     `xlsx:"1"`
		Address   string     `xlsx:"2"`
}

type XlsxParser struct{
    FileName      string
}

//The XlsxParser struct implements the Parse() function
func (xlsx_parser *XlsxParser)Parse(){
    absPath, _ := filepath.Abs(xlsx_parser.FileName)
    xlFile, err := xlsx.OpenFile(absPath)

    if err != nil {
        log.Fatal(err)
    }
    var people []Person
    for _, sheet := range xlFile.Sheets {
        readStruct := &structTest{}
            for k, row := range sheet.Rows {
                if k!= 0{
                    row.ReadStruct(readStruct)

                    people = append(people, Person{
                        Name: readStruct.Name,
                        Age:  readStruct.Age,
                        Address: readStruct.Address,
                    })
                }
            }
    }
    peopleJson, _ := json.Marshal(people)
    fmt.Println(string(peopleJson))
}
