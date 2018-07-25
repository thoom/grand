env GOOS=linux GOARCH=386 go build -ldflags "-X main.buildVersion=$TRAVIS_BRANCH-linux386" -o goron-linux-386
env GOOS=linux GOARCH=amd64 go build -ldflags "-X main.buildVersion=$TRAVIS_BRANCH-linux64" -o goron-linux-amd64
env GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.buildVersion=$TRAVIS_BRANCH-darwin" -o goron-darwin
env GOOS=windows GOARCH=amd64 go build -ldflags "-X main.buildVersion=$TRAVIS_BRANCH-windows" -o goron-windows