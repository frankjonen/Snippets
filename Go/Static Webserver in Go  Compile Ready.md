```go

package main

import (
	"log"
	"fmt"
	"net/http"
	"os/user"
)

func main() {
// Get the user folder.
// We'll be serving the site from there.
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

// Website
	fs := http.FileServer(http.Dir(usr.HomeDir+"/where/your/stuff/is/at"))
	http.Handle("/", fs)

	fmt.Println("\r\n")
	log.Println("| Serving the static has begun! \n")
	http.ListenAndServe(":8080", nil)
}

```
