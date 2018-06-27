package FileParsers


type Person struct {
    Name     string       `json:"name"`
    Age      string       `json:"age"`
    Address  string       `json:"address"`
}



type FileParser interface{
    Parse()
}
