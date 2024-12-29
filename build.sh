name="wechat"

fullName="$name"_linux_amd64
CGO_ENABLED=1 CC=gcc CXX=g++ go build -buildmode=plugin -v -ldflags="-w -s" -o ./bin/$fullName.so
echo "$fullName build done..."


fullName="$name"_linux_arm
CGO_ENABLED=1 CC=arm-linux-gnueabi-gcc CXX=arm-linux-gnueabi-g++ GOARCH=arm GOARM=7 go build -buildmode=plugin -v -ldflags="-w -s" -o ./bin/$fullName.so
echo "$fullName build done..."

fullName="$name"_linux_arm64
CGO_ENABLED=1 CC=aarch64-linux-gnu-gcc CXX=arm-linux-gnueabi-g++ GOARCH=arm64 go build -buildmode=plugin -v -ldflags="-w -s" -o ./bin/$fullName.so
echo "$fullName build done..."

sleep 8
