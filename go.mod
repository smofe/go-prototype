module go-prototype

go 1.15

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/jinzhu/gorm v1.9.16
	github.com/mattn/go-sqlite3 v1.14.4 // indirect
	github.com/smofe/go-prototype/models v0.0.0-00010101000000-000000000000
)

replace github.com/smofe/go-prototype/models => ./models
