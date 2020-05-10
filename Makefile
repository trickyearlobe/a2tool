# Builds the project

# Make a build number from Habitat pkg_prefix or use the current date
ifdef pkg_release
  LDFLAGS:="-X main.Version=$(pkg_release)"
else
	LDFLAGS:="-X main.Version=`date +"%Y%m%d%H%M%S"`"
endif

build:
	go get
	go build -ldflags $(LDFLAGS)

# Installs our project: copies binaries
install:
	echo LDFLAGS=
	go install -ldflags $(LDFLAGS)

# Cleans our project: deletes binaries
clean:
	if [ -f a2tool ] ; then rm a2tool ; fi

.PHONY: clean install