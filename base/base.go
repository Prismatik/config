package base

import (
	"github.com/prismatik/config/base/users"
	"github.com/prismatik/jabba"
)

func Go() {
	jabba.RunOrDie("apt-get", "install", "-y", "tmux")
	jabba.RunOrDie("apt-get", "install", "-y", "build-essential")
	jabba.RunOrDie("apt-get", "install", "-y", "collectd")

	users.Go()

	files := []jabba.File{
		logglyFile,
		sudoers,
		collectd,
	}
	for _, file := range files {
		jabba.WriteFile(file)
	}
}
