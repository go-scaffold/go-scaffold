module github.com/pasdam/go-scaffold

replace github.com/pasdam/go-scaffold/pkg => ./pkg

require (
	github.com/kr/text v0.2.0 // indirect
	github.com/pasdam/go-files-test v0.0.0-20210125082053-e1c8203c662c
	github.com/pasdam/go-io-utilx v0.0.0-20210125081100-08a73663a45e
	github.com/pasdam/go-template-map-loader v0.0.0-20230710141516-e9f048463b7e
	github.com/pasdam/mockit v0.0.0-20210125081107-43da07441f3d
	github.com/spf13/cobra v1.7.0
	github.com/stretchr/testify v1.8.4
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)

go 1.16
