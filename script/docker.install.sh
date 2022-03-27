#!/bin/env bash

# First, updating existing list of packages
sudo apt update

# installing a few prerequisite packages
# which let `apt` use packages over HTTPS
sudo apt install -y apt-transport-https ca-certificates curl software-properties-common

#PLATFORM=$(echo $OSTYPE)
UNAME=$(uname)
PLATFORM=$(echo "${UNAME,,}")
DPKGARCH=$(dpkg-architecture -q DEB_BUILD_ARCH)
source /etc/os-release
DISTRO="$ID"
OSVERSION="$VERSION_CODENAME"

PKGTYPE="unknown"
if [ "$ID_LIKE" == "debian" ]; then
  PKGTYPE="deb"
else
  echo "Package type is ${PKGTYPE}."
  exit
fi

# Adding the GPG key for the official Docker repository to your system
curl -fsSL https://download.docker.com/${PLATFORM}/${DISTRO}/gpg | sudo apt-key add -

# Adding the Docker repository to APT sources
sudo add-apt-repository "${PKGTYPE} [arch=${DPKGARCH}] https://download.docker.com/${PLATFORM}/${DISTRO} ${OSVERSION} stable"

# updating the package database with the Docker packages from the newly added repo
sudo apt update

# ensuring to install from the Docker repo instead of the default Ubuntu repo
apt-cache policy docker-ce

# Finally, installing Docker
sudo apt install -y docker-ce

# Checking if docker is running
sudo systemctl status docker

# Adding current username to the docker group
sudo usermod -aG docker ${USER}

# Confirming current user is added to the docker group
id -nG

# To apply the new group membership, log back in
su -l ${USER}

# Checking docker version
docker version
