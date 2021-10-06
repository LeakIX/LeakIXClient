
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags="-s -w  -extldflags '-static'" -o bin/leakix-linux-32 ./cmd/leakix &
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -extldflags '-static'" -o bin/leakix-linux-64 ./cmd/leakix &
CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build -ldflags="-s -w -extldflags '-static'" -o bin/leakix-freebsd-64 ./cmd/leakix &
CGO_ENABLED=0 GOOS=freebsd GOARCH=386 go build -ldflags="-s -w -extldflags '-static'" -o bin/leakix-freebsd-32 ./cmd/leakix &

CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=7 go build -ldflags="-s -w -extldflags '-static'" -o bin/leakix-linux-arm7 ./cmd/leakix &
CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=6 go build -ldflags="-s -w -extldflags '-static'" -o bin/leakix-linux-arm6 ./cmd/leakix &

CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w -extldflags '-static'" -o bin/leakix-osx ./cmd/leakix &

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -extldflags '-static'" -o bin/leakix-win64.exe ./cmd/leakix &
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags="-s -w -extldflags '-static'" -o bin/leakix-win32.exe ./cmd/leakix &

CGO_ENABLED=0 GOOS=netbsd GOARCH=amd64 go build -ldflags="-s -w -extldflags '-static'" -o bin/leakix-netbsd-64 ./cmd/leakix &
CGO_ENABLED=0 GOOS=netbsd GOARCH=386 go build -ldflags="-s -w -extldflags '-static'" -o bin/leakix-netbsd-32 ./cmd/leakix &

CGO_ENABLED=0 GOOS=openbsd GOARCH=amd64 go build -ldflags="-s -w -extldflags '-static'" -o bin/leakix-openbsd-64 ./cmd/leakix &
CGO_ENABLED=0 GOOS=openbsd GOARCH=386 go build -ldflags="-s -w -extldflags '-static'" -o bin/leakix-openbsd-32 ./cmd/leakix &