module go-prototype

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/jinzhu/gorm v1.9.16
	github.com/mattn/go-sqlite3 v1.14.4 // indirect
	github.com/smofe/go-prototype/controllers v0.0.0-20201027135620-1ca9037c2ca4
	github.com/smofe/go-prototype/models v0.0.0-20201027133741-2edf2d8755f6
	github.com/stretchr/testify v1.4.0
)

replace github.com/smofe/go-prototype/models => ./models

replace github.com/smofe/go-prototype/controllers => ./controllers
