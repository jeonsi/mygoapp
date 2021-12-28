module testapp

go 1.17

replace greeting => ../greeting

require (
	greeting v0.0.0-00010101000000-000000000000
	keyboard v0.0.0-00010101000000-000000000000
)

replace keyboard => ../keyboard
