# install cobra
go get -u github.com/spf13/cobra/cobra

# https://www.youtube.com/watch?v=-tO7zSv80UY
mkdir dadjoke
cd dadjoke
go mod init github.com/example/dadjoke
cobra init
go mod tidy
go run main.go
# change short and long description in cmd/root.go
cobra add random
