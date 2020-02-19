package bigip

import (
	"fmt"
	"io/ioutil"

	"github.com/lornasong/go-bigip"
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
		fmt.Println("Could not post AS3:", err)
		return
	}
	fmt.Printf(" ---- Response: ----\n%s\n", resp)

	data, err := f5.GetAs3()
	if err != nil {
		fmt.Println("Could not retrieve AS3 info:", err)
		return
	}
	fmt.Printf(" ---- AS3 information: ----\n%s\n", data)

	data, err = f5.DeleteAs3()
	if err != nil {
		fmt.Println("Could not delete AS3 resources:", err)
		return
	}
	fmt.Printf(" ---- Delete Response: ----\n%s\n", data)
}
