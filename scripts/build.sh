env GOOS=linux GOARCH=386 go build -ldflags "-X main.buildVersion=$RELEASE_VERSION" -o grand && tar cfz grand.linux-386.tar.gz grand
env GOOS=linux GOARCH=amd64 go build -ldflags "-X main.buildVersion=$RELEASE_VERSION" -o grand && tar cfz grand.linux-amd64.tar.gz grand
env GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.buildVersion=$RELEASE_VERSION" -o grand && tar cfz grand.darwin-amd64.tar.gz grand
env GOOS=darwin GOARCH=arm64 go build -ldflags "-X main.buildVersion=$RELEASE_VERSION" -o grand && tar cfz grand.darwin-arm64.tar.gz grand
env GOOS=windows GOARCH=amd64 go build -ldflags "-X main.buildVersion=$RELEASE_VERSION" -o grand && zip grand.windows.zip grand