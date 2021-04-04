module github.com/mysa-fika/go/authentication-service

go 1.16

replace (
	github.com/mysa-fika/go/domain => ../domain
	github.com/mysa-fika/go/infrastructure => ../infrastructure
)

require (
	github.com/go-ozzo/ozzo-dbx v1.5.0
	github.com/go-ozzo/ozzo-routing v2.1.4+incompatible
	github.com/go-ozzo/ozzo-validation/v4 v4.3.0
	github.com/golang/gddo v0.0.0-20190904175337-72a348e765d2 // indirect
	github.com/google/go-cmp v0.3.1 // indirect
	github.com/google/uuid v1.2.0
	github.com/mysa-fika/go/domain v0.0.0-00010101000000-000000000000
	github.com/mysa-fika/go/infrastructure v0.0.0-00010101000000-000000000000
	github.com/rs/zerolog v1.21.0
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0 // indirect
)
