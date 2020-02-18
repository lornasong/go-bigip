package bigip

import (
	"fmt"
	"io/ioutil"

	"github.com/scottdware/go-bigip"
)

// Note: Had to call main something different as the package complained main was in LTM's example
func mainAs3() {
	// F5's quick start example:
	// https://clouddocs.f5.com/products/extensions/f5-appsvcs-extension/latest/userguide/quick-start.html
	as3File, err := ioutil.ReadFile("as3_example.json")
	if err != nil {
		fmt.Println("Could not read declaration file:", err)
		return
	}

	// Connect to the BIG-IP system.
	f5 := bigip.NewSession("as3.company.com", "admin", "secret", nil)
	resp, err := f5.PostAs3(as3File)
	if err != nil {
		fmt.Println("Error posting AS3:", err)
		return
	}

	fmt.Printf("Response: %s", resp)
}
