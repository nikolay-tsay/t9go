module github.com/nikolay-tsay/t9search

go 1.19

replace github.com/nikolay-tsay/t9printer => /printer

replace github.com/nikolay-tsay/t9reader => /reader

require github.com/nikolay-tsay/t9printer v0.0.0-00010101000000-000000000000

require (
	github.com/nikolay-tsay/t9filter v0.0.0-00010101000000-000000000000 // indirect
	github.com/nikolay-tsay/t9reader v0.0.0-00010101000000-000000000000 // indirect
)

replace github.com/nikolay-tsay/t9filter => /filter
