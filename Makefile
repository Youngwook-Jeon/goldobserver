BINARY_NAME=GoldObserver.exe
APP_NAME=GoldObserver
VERSION=1.0.1
BUILD_NO=2
APP_ID=com.project.goldobserver
ICON=icon.png

## build: build binary and package app
build:
	del ${BINARY_NAME}
	fyne package -os windows -icon ${ICON} -appID ${APP_ID} -appVersion ${VERSION} -appBuild ${BUILD_NO} -name ${APP_NAME} -release

## run: builds and runs the application
run:
	set DB_PATH=%cd%\sql.db & go run .

## clean: runs go clean and deletes binaries
clean:
	@echo "Cleaning..."
	@go clean
	@del ${BINARY_NAME}
	@echo "Cleaned!"

## test: runs all tests
test:
	go test -v ./...