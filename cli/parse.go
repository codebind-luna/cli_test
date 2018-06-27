package cli


import(
    "path/filepath"
    "github.com/lunaimaginea/cli_test/cli/FileParsers"
)


func GetParser(path string) (FileParsers.FileParser) {
    var filename = path
    var extension = filepath.Ext(filename)

    switch extension{
  	    case ".csv":
  		      return &FileParsers.CsvParser{FileName: filename}
        case ".xlsx":
  		      return &FileParsers.XlsxParser{FileName: filename}
  	    default:
  		  //if type is invalid, return an error
  		      return nil
	  }
}
