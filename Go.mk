# file: Go.mk
# Description: Golang makefile baseline to build and package files


### gobuild function
##  $1 = name of build
## $2 path to source files,
## $3 = CommitId,
## $4 = Branch
gobuild = go build \
	-trimpath \
	-race \
	-ldflags '-X main.CommitID=$(3) -X main.Branch=$(4)' \
	-o $1 \
	$2


### Strip function
## $1 binary name
gostrip = strip -s $1


### UPX function
## $1 = binary name
goupx = upx $1

