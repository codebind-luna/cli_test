package FileParsers


import (
    "path/filepath"
    "bufio"
    "encoding/csv"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "os"
)


type CsvParser struct{
    FileName      string
}

//The CsvParser struct implements the Parse() function
func (csv_parser *CsvParser)Parse(){
    absPath, _ := filepath.Abs(csv_parser.FileName)
    csvFile, _ := os.Open(absPath)
    reader     := csv.NewReader(bufio.NewReader(csvFile))

    var people []Person

    for {
        line, error := reader.Read()
        if error == io.EOF {
            break
        } else if error != nil {
            log.Fatal(error)
        }

        people = append(people, Person{
            Name: line[0],
            Age:  line[1],
            Address: line[2],
        })
    }
    peopleJson, _ := json.Marshal(people)
    fmt.Println(string(peopleJson))
}
