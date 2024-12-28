name="gotify-wechat"

fullName="$name"_linux_amd64
CGO_ENABLED=1 CGO_CFLAGS=-D__USE_MINGW_ANSI_STDIO=1 CGO_LDFLAGS=-lmsvcrt GOOS=linux GOARCH=amd64 go build -buildmode=plugin -x -v -ldflags="-w -s" -o ./bin/$fullName.so
echo "$fullName 编译完成..."


fullName="$name"_linux_arm
CGO_ENABLED=1 CGO_CFLAGS=-D__USE_MINGW_ANSI_STDIO=1 CGO_LDFLAGS=-lmsvcrt GOOS=linux GOARCH=arm GOARM=7 go build -buildmode=plugin -x -v -ldflags="-w -s" -o ./bin/$fullName.so
echo "$fullName 编译完成..."

sleep 8
