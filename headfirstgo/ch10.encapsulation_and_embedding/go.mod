module github.com/headfirstgo/ch.10

go 1.17

replace geo => ./geo

require (
	calendar v0.0.0-00010101000000-000000000000
	geo v0.0.0-00010101000000-000000000000
)

replace calendar => ./calendar
