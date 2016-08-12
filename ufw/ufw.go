package ufw

import (
	"github.com/prismatik/jabba"
	"os"
	"strconv"
)

func Go() {
	dockerised, _ := strconv.ParseBool(os.Getenv("DOCKER"))
	// UFW/iptables fails inside Docker
	if !dockerised {
		jabba.RunOrDie("apt-get", "install", "-y", "ufw")
		jabba.RunOrDie("ufw", "allow", "22")
		jabba.RunOrDie("ufw", "--force", "enable")
	}
}
