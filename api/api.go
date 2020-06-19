package api

import "fmt"

const (
	DOMAIN    = ""
	VERSION   = "v1"
	PORT      = ":7070"
	LOCAL     = "http://127.0.0.1" + PORT
	APIConfig = "WELLS_FAR_GO_CONFIG"
	BANNER    = `
dP   dP   dP          dP dP                    88888888b                             .88888.          
88   88   88          88 88                    88                                   d8'    88
88  .8P  .8P .d8888b. 88 88 .d8888b.          a88aaaa    .d8888b. 88d888b.          88        .d8888b.
88  d8'  d8' 88ooood8 88 88 Y8ooooo. 88888888  88        88'   88 88'   88 88888888 88   YP88 88'   88
88.d8P8.d8P  88.  ... 88 88       88           88        88.  .88 88                Y8.   .88 88.  .88
8888' Y88'    88888P' dP dP  88888P'           dP         88888P8 dP                  88888'   88888P'
`
)

type API struct{}

// PrintBanner prints 'Wells-Far-Go' Banner on console.
func PrintBanner() {
	fmt.Println(BANNER)
}
