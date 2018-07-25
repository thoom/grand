docker build -t goron -f Dockerfile --no-cache --build-arg BUILD_VERSION=$TRAVIS_BRANCH .

docker login -u $DOCKER_USER -p $DOCKER_PASS
docker tag goron thoom/goron:latest
docker tag goron thoom/goron:$TRAVIS_BRANCH

docker push thoom/goron