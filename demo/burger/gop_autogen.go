// Code generated by gop (Go+); DO NOT EDIT.

package main

import (
	"fmt"
	"github.com/goplus/huh"
	"github.com/qiniu/x/errors"
)

const _ = true

type result struct {
	Burger     string
	Toppings   []string
	SauceLevel int
}
//line demo/burger/burger.gop:9
func main() {
//line demo/burger/burger.gop:9:1
	ret := new(result)
//line demo/burger/burger.gop:11:1
	form := func() (_gop_ret huh.Form) {
//line demo/burger/burger.gop:11:1
		var _gop_err error
//line demo/burger/burger.gop:11:1
		_gop_ret, _gop_err = huh.New("<form>\n\t<group>\n\t\t<note\n\t\t\ttitle=\"Charmburger\"\n\t\t\tdescription=\"Welcome to _Charmburger™_.\n\nHow may we take your order?\"\n\t\t\tnext/>\n\t</group>\n\n\t<group>\n\t\t<select id=\"Burger\" title=\"Choose your burger\">\n\t\t\t<option value=\"classic\" title=\"Charmburger Classic\"/>\n\t\t\t<option value=\"chickwich\" title=\"Chickwich\"/>\n\t\t\t<option value=\"fishburger\" title=\"Fishburger\"/>\n\t\t\t<option value=\"charmpossible\" title=\"Charmpossible™ Burger\"/>\n\t\t</select>\n\n\t\t<multiselect id=\"Toppings\" title=\"Toppings\" limit=4>\n\t\t\t<option value=\"Lettuce\" selected/>\n\t\t\t<option value=\"Tomatoes\" selected/>\n\t\t\t<option value=\"Jalapeños\"/>\n\t\t\t<option value=\"Cheese\"/>\n\t\t\t<option value=\"Vegan Cheese\"/>\n\t\t\t<option value=\"Nutella\"/>\n\t\t</multiselect>\n\n\t\t<select id=\"SauceLevel\" title=\"How much Charm Sauce do you want?\">\n\t\t\t<option value=0 title=\"None\"/>\n\t\t\t<option value=1 title=\"A little\"/>\n\t\t\t<option value=2 title=\"A lot\"/>\n\t\t</select>\n\t</group>\n</form>\n", ret)
//line demo/burger/burger.gop:11:1
		if _gop_err != nil {
//line demo/burger/burger.gop:11:1
			_gop_err = errors.NewFrame(_gop_err, "huh`> ret\n<form>\n\t<group>\n\t\t<note\n\t\t\ttitle=\"Charmburger\"\n\t\t\tdescription=\"Welcome to _Charmburger™_.\n\nHow may we take your order?\"\n\t\t\tnext/>\n\t</group>\n\n\t<group>\n\t\t<select id=\"Burger\" title=\"Choose your burger\">\n\t\t\t<option value=\"classic\" title=\"Charmburger Classic\"/>\n\t\t\t<option value=\"chickwich\" title=\"Chickwich\"/>\n\t\t\t<option value=\"fishburger\" title=\"Fishburger\"/>\n\t\t\t<option value=\"charmpossible\" title=\"Charmpossible™ Burger\"/>\n\t\t</select>\n\n\t\t<multiselect id=\"Toppings\" title=\"Toppings\" limit=4>\n\t\t\t<option value=\"Lettuce\" selected/>\n\t\t\t<option value=\"Tomatoes\" selected/>\n\t\t\t<option value=\"Jalapeños\"/>\n\t\t\t<option value=\"Cheese\"/>\n\t\t\t<option value=\"Vegan Cheese\"/>\n\t\t\t<option value=\"Nutella\"/>\n\t\t</multiselect>\n\n\t\t<select id=\"SauceLevel\" title=\"How much Charm Sauce do you want?\">\n\t\t\t<option value=0 title=\"None\"/>\n\t\t\t<option value=1 title=\"A little\"/>\n\t\t\t<option value=2 title=\"A lot\"/>\n\t\t</select>\n\t</group>\n</form>\n`", "demo/burger/burger.gop", 11, "main.main")
//line demo/burger/burger.gop:11:1
			panic(_gop_err)
		}
//line demo/burger/burger.gop:11:1
		return
	}()
//line demo/burger/burger.gop:48:1
	err := form.Run()
//line demo/burger/burger.gop:49:1
	if err != nil {
//line demo/burger/burger.gop:50:1
		fmt.Println(err)
	} else {
//line demo/burger/burger.gop:52:1
		fmt.Println("Your order:")
//line demo/burger/burger.gop:53:1
		fmt.Println("Burger:", ret.Burger)
//line demo/burger/burger.gop:54:1
		fmt.Println("Toppings:", ret.Toppings)
//line demo/burger/burger.gop:55:1
		fmt.Println("Charm Sauce:", ret.SauceLevel)
	}
}
