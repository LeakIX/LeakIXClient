GOOS=linux GOARCH=386 go build -ldflags="-s -w" -o bin/leakix-ns-linux-32 ./cmd/leakix-ns &
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/leakix-ns-linux-64 ./cmd/leakix-ns &
GOOS=freebsd GOARCH=amd64 go build -ldflags="-s -w" -o bin/leakix-ns-freebsd-64 ./cmd/leakix-ns &
GOOS=freebsd GOARCH=386 go build -ldflags="-s -w" -o bin/leakix-ns-freebsd-32 ./cmd/leakix-ns &

GOOS=linux GOARCH=arm GOARM=7 go build -ldflags="-s -w" -o bin/leakix-ns-linux-arm7 ./cmd/leakix-ns &
GOOS=linux GOARCH=arm GOARM=6 go build -ldflags="-s -w" -o bin/leakix-ns-linux-arm6 ./cmd/leakix-ns &

GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o bin/leakix-ns-osx ./cmd/leakix-ns &

GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o bin/leakix-ns-win64.exe ./cmd/leakix-ns &
GOOS=windows GOARCH=386 go build -ldflags="-s -w" -o bin/leakix-ns-win32.exe ./cmd/leakix-ns &

GOOS=netbsd GOARCH=amd64 go build -ldflags="-s -w" -o bin/leakix-ns-netbsd-64 ./cmd/leakix-ns &
GOOS=netbsd GOARCH=386 go build -ldflags="-s -w" -o bin/leakix-ns-netbsd-32 ./cmd/leakix-ns &

GOOS=openbsd GOARCH=amd64 go build -ldflags="-s -w" -o bin/leakix-ns-openbsd-64 ./cmd/leakix-ns &
GOOS=openbsd GOARCH=386 go build -ldflags="-s -w" -o bin/leakix-ns-openbsd-32 ./cmd/leakix-ns &

GOOS=linux GOARCH=386 go build -ldflags="-s -w" -o bin/leakix-linux-32 ./cmd/leakix &
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/leakix-linux-64 ./cmd/leakix &
GOOS=freebsd GOARCH=amd64 go build -ldflags="-s -w" -o bin/leakix-freebsd-64 ./cmd/leakix &
GOOS=freebsd GOARCH=386 go build -ldflags="-s -w" -o bin/leakix-freebsd-32 ./cmd/leakix &

GOOS=linux GOARCH=arm GOARM=7 go build -ldflags="-s -w" -o bin/leakix-linux-arm7 ./cmd/leakix &
GOOS=linux GOARCH=arm GOARM=6 go build -ldflags="-s -w" -o bin/leakix-linux-arm6 ./cmd/leakix &

GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o bin/leakix-osx ./cmd/leakix &

GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o bin/leakix-win64.exe ./cmd/leakix &
GOOS=windows GOARCH=386 go build -ldflags="-s -w" -o bin/leakix-win32.exe ./cmd/leakix &

GOOS=netbsd GOARCH=amd64 go build -ldflags="-s -w" -o bin/leakix-netbsd-64 ./cmd/leakix &
GOOS=netbsd GOARCH=386 go build -ldflags="-s -w" -o bin/leakix-netbsd-32 ./cmd/leakix &

GOOS=openbsd GOARCH=amd64 go build -ldflags="-s -w" -o bin/leakix-openbsd-64 ./cmd/leakix &
GOOS=openbsd GOARCH=386 go build -ldflags="-s -w" -o bin/leakix-openbsd-32 ./cmd/leakix &
