module github.com/abhaytiket/gomonorepo/services/app1

go 1.20

replace github.com/abhaytiket/gomonorepo/lib/msg => ../../lib/msg

require (
	github.com/abhaytiket/gomonorepo/lib/msg v0.0.0-00010101000000-000000000000
	github.com/go-chi/chi/v5 v5.0.10
)
