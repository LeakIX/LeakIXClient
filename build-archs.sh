GOOS=linux GOARCH=386 go build -o leakix-linux-32 ./cmd/leakix
GOOS=linux GOARCH=amd64 go build -o leakix-linux-64 ./cmd/leakix
GOOS=freebsd GOARCH=amd64 go build -o leakix-freebsd-64 ./cmd/leakix
GOOS=freebsd GOARCH=386 go build -o leakix-freebsd-32 ./cmd/leakix

GOOS=linux GOARCH=arm GOARM=7 go build -o leakix-linux-arm7 ./cmd/leakix
GOOS=linux GOARCH=arm GOARM=6 go build -o leakix-linux-arm6 ./cmd/leakix

GOOS=darwin GOARCH=amd64 go build -o leakix-osx-64 ./cmd/leakix
GOOS=darwin GOARCH=386 go build -o leakix-osx-32 ./cmd/leakix

GOOS=windows GOARCH=amd64 go build -o leakix-win64.exe ./cmd/leakix
GOOS=windows GOARCH=386 go build -o leakix-win32.exe ./cmd/leakix

GOOS=netbsd GOARCH=amd64 go build -o leakix-netbsd-64 ./cmd/leakix
GOOS=netbsd GOARCH=386 go build -o leakix-netbsd-32 ./cmd/leakix

GOOS=openbsd GOARCH=amd64 go build -o leakix-openbsd-64 ./cmd/leakix
GOOS=openbsd GOARCH=386 go build -o leakix-openbsd-32 ./cmd/leakix
