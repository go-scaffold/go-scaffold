module github.com/pasdam/go-scaffold

replace github.com/pasdam/go-scaffold/pkg => ./pkg

require (
	bou.ke/monkey v1.0.2
	github.com/jessevdk/go-flags v1.5.0
	github.com/kr/text v0.2.0 // indirect
	github.com/manifoldco/promptui v0.9.0
	github.com/otiai10/copy v1.11.0
	github.com/pasdam/go-files-test v0.0.0-20210125082053-e1c8203c662c
	github.com/pasdam/go-io-utilx v0.0.0-20210125081100-08a73663a45e
	github.com/pasdam/mockit v0.0.0-20210125081107-43da07441f3d
	github.com/stretchr/testify v1.8.2
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v2 v2.4.0
)

go 1.16
