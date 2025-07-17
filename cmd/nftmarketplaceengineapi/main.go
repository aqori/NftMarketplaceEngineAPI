// cmd/nftmarketplaceengineapi/main.go
package main

import (
"flag"
"log"
"os"

"nftmarketplaceengineapi/internal/nftmarketplaceengineapi"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := nftmarketplaceengineapi.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
