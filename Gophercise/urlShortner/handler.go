package urlshort

import (
	"fmt"
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

type urlStruct struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if redirectPath, ok := pathsToUrls[r.URL.Path]; ok {
			// http.ResponseWriter(w, redirectPath)
			http.Redirect(w, r, redirectPath, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}

	//	TODO: Implement this...
	// http.Redirect(w, r, pathsToUrls[])
	// http.HandlerFunc("/urlshort-godoc", urlHandlerFunc)
	// http.HandlerFunc("/yaml-godoc", urlHandlerFunc)

	// return pathsToUrls[fallback.ServeHTTP()]
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...

	var yamlArr []urlStruct
	err := yaml.Unmarshal(yml, &yamlArr)
	if err != nil {
		fmt.Println("ERR in yaml Unmarshall")
		exit 1
	}
	fmt.Println("yamlArr:", yamlArr)

	fmt.Printf("inside yaml handler")
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("r.URL.PATH", r.URL.Path)
		fmt.Println("yml:", yml)

		for k, _ := range yamlArr {
			fmt.Println("k:", yamlArr[k].Path)
			fmt.Println("r.URL.Path:", r.URL.Path)
			if r.URL.Path == yamlArr[k].Path {
				fmt.Println("path found")
				http.Redirect(w, r, yamlArr[k].Url, http.StatusFound)
			}
		}
		fallback.ServeHTTP(w, r)
	}, nil
	// return nil, nil
	// return

}
