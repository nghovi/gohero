module vietanh.com/gohero

go 1.17

replace vietanh.com/gohero/mappings => ./mappings

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/go-gorp/gorp v2.2.0+incompatible
	github.com/go-sql-driver/mysql v1.6.0
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
)

require (
	github.com/flosch/pongo2/v5 v5.0.0
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-contrib/static v0.0.1
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/golang/protobuf v1.3.3 // indirect
	github.com/howeyc/fsnotify v0.9.0 // indirect
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/lib/pq v1.10.4 // indirect
	github.com/mattn/go-colorable v0.1.4 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/mattn/go-sqlite3 v1.14.10 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v0.0.0-20180701023420-4b7aa43c6742 // indirect
	github.com/pilu/config v0.0.0-20131214182432-3eb99e6c0b9a // indirect
	github.com/pilu/fresh v0.0.0-20190826141211-0fa698148017 // indirect
	github.com/poy/onpar v1.1.2 // indirect
	github.com/ugorji/go/codec v1.1.7 // indirect
	github.com/ziutek/mymysql v1.5.4 // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 // indirect
	golang.org/x/sys v0.0.0-20200116001909-b77594299b42 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
	gorm.io/driver/mysql v1.2.3
	gorm.io/gorm v1.22.5
)

replace vietanh.com/gohero/controllers => ./controllers

replace vietanh.com/gohero/models => ./models
