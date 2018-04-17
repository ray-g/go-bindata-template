# go-bindata-template

## About go-bindata-template

`go-bindata-template` extends Go's built-in [html/template](https://godoc.org/html/template)
to load and parse embeded template content instead of filesystem packaged by `go-bindata`
[jteeuwen/go-bindata](https://github.com/jteeuwen/go-bindata)
or a little newer [kevenburke/go-bindata](https://github.com/kevinburke/go-bindata).

Modified `Parse` and `ParseFiles` to parse embeded contents.
Added `ParseDir` and `ParseAll` for convenice.

## Example

(Suppose your data are under `data` folder and after running `go-bindata data/...`)

```golang
import (
    template "github.com/ray-g/go-bindata-template"
)

func someHandler(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.New("index", &template.BinData{
        Asset: Asset,
        AssetDir: AssetDir,
        AssetNames: AssetNames,
    }).Parse("data/index.tmpl")

    if err != nil {
        log.Fatalf("error parsing template: %s", err)
    }

    err = tmpl.Execute(w)

    if err != nil {
        log.Fatalf("error execute template: %s", err)
    }
}
```
