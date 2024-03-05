APP?=crf
RELEASE?=$(shell ./AutoVersionIncrement get)
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date +%FT%T%z)
PROJECT?=github.com/Jarover/crf

clean:
	rm -f ${APP}
	rm -f ${APP}.exe


buildwin: clean
	$(shell ./AutoVersionIncrement inc-patch)
	GOOS=windows go build \
				-o ${APP}.exe \
                -ldflags "-s -w -X ${PROJECT}/internal/config.Release=${RELEASE} \
                -X ${PROJECT}/internal/config/version.Commit=${COMMIT} -X ${PROJECT}/internal/config/config.BuildTime=${BUILD_TIME}" \
                cmd/${APP}/main.go


buildlinux:	clean
	$(shell ./AutoVersionIncrement inc-patch)
	GOOS=linux go build \
				-o ${APP} \
                -ldflags "-s -w -X ${PROJECT}/internal/config.Release=${RELEASE} \
                -X ${PROJECT}/version.Commit=${COMMIT} -X ${PROJECT}/version.BuildTime=${BUILD_TIME}" \
				cmd/${APP}/main.go

container: buildlinux
    docker build -f Dockerfile -t $(APP):$(RELEASE) .				