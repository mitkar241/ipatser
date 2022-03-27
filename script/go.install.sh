#!/bin/env bash

<<DESC
@ FileName: go.install.sh
@ Description: Installation of latest GoLang
@ Usage:
  - chmod +x go.install.sh
  - ./go.install.sh
@ Documentation: https://go.dev/doc/install
@ Installation:
  - sudo snap install go
  - sudo apt install golang-go
  - sudo apt install gccgo-go
@ Additional Details:
  - sudo apt search golang-go
  - snap info go
DESC

##########
# DOWNLOAD GO TAR
##########
GOVER=$(curl -s https://go.dev/VERSION?m=text)
#PLATFORM=$(echo $OSTYPE)
UNAME=$(uname)
PLATFORM=$(echo "${UNAME,,}")
DPKGARCH=$(dpkg-architecture -q DEB_BUILD_ARCH)

GOTAR="${GOVER}.${PLATFORM}-${DPKGARCH}.tar.gz"
echo "Downloading ${GOTAR}"
curl -s -O -L "https://go.dev/dl/${GOTAR}"

##########
# CHECK IF GO TAR IS DOWNLOADED
##########
if [[ -f "$GOTAR" ]]; then
  echo "$GOTAR exists."
else 
  echo "$GOTAR does not exist."
fi

##########
# CLEANUP EXISTING GO BIN
##########
GOLOC="/usr/local/go"
if [[ -d "$GOLOC" ]]; then
  echo "$GOLOC exists."
  sudo rm -rf /usr/local/go
  echo "Removed $GOLOC."
else 
  echo "$GOLOC does not exist."
fi

##########
# UNTAR GO TAR
##########
sudo tar -C /usr/local -xzf ${GOTAR}

##########
# ADD GO BIN TO PATH
##########
# temporarily add go bin to PATH
#export PATH=$PATH:/usr/local/go/bin

profile_count_go=$(cat .profile | grep 'PATH' | grep -c '/usr/local/go/bin')

if [ $profile_count_go -eq 0 ]; then
  echo "/usr/local/go/bin does not exist in PATH."
  cat << EOF >>.profile

# set PATH so it includes golang bin if it exists
if [ -d "/usr/local/go/bin" ] ; then
    PATH="/usr/local/go/bin:\$PATH"
fi
EOF
  echo "Added /usr/local/go/bin in PATH."
else
  echo "/usr/local/go/bin exists in PATH."
fi

source .profile

##########
# CHECK GO VERSION
##########
# log in again to not use `source`.
#su -l $USER
# else run `source .profile` manually

go version
