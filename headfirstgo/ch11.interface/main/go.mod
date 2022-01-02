module main

go 1.17

replace github.com/headfirstgo/gadget => ../gadget

require (
	github.com/headfirstgo/gadget v0.0.0-00010101000000-000000000000
	mypkg v0.0.0-00010101000000-000000000000
)

replace mypkg => ../mypkg
