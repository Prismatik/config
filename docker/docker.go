package docker

import (
	"github.com/prismatik/jabba"
)

var distro = "trusty"

func Go() {
	distro = jabba.DistroString()
	jabba.RunOrDie("sudo", "apt-get", "install", "apt-transport-https", "ca-certificates")
	jabba.RunOrDie("sudo", "apt-key", "adv", "--keyserver", "hkp://p80.pool.sks-keyservers.net:80", "--recv-keys", "58118E89F3A912897C070ADBF76221572C52609D")
	jabba.WriteFile(dockerList)
	jabba.RunOrDie("apt-cache", "policy", "docker-engine")
	jabba.RunOrDie("sudo", "apt-get", "update")
	jabba.RunOrDie("sudo", "apt-get", "install", "linux-image-extra-$(uname", "-r)", "linux-image-extra-virtual")
	jabba.RunOrDie("sudo", "apt-get", "install", "docker-engine")
	jabba.RunOrDie("sudo", "service", "docker", "start")
	jabba.RunOrDie("sudo", "curl", "-L", "https://github.com/docker/compose/releases/download/1.8.0/docker-compose-`uname", "-s`-`uname", "-m`", ">", "/usr/local/bin/docker-compose")
}

var dockerList = jabba.File{
	Path: "/etc/apt/sources.list.d/docker.list",
	Perm: 0644,
	Vars: map[string]string{
		"distro": distro,
	},
	Template: "deb https://apt.dockerproject.org/repo ubuntu-{{.distro}} main",
}
