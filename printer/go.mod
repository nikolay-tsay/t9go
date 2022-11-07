module github.com/nikolay-tsay/t9printer

go 1.19

replace github.com/nikolay-tsay/t9reader => ../reader

require (
	github.com/nikolay-tsay/t9filter v0.0.0-00010101000000-000000000000
	github.com/nikolay-tsay/t9reader v0.0.0-00010101000000-000000000000
)

replace github.com/nikolay-tsay/t9filter => ../filter
