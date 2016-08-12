package codescreen

import (
	"github.com/prismatik/jabba"
	"github.com/prismatik/secrets"
)

func Go() {
	jabba.AddUser(codescreenUser)
	jabba.WriteFile(codescreenUpstartFile)
	jabba.RunOrDie("sudo", "apt", "install", "-y", "wget", "unzip")
	jabba.RunOrDie("mkdir", "-p", "/opt/codescreen/src")
	jabba.RunOrDie("wget", "https://github.com/Prismatik/codescreen/archive/1.1.0.zip")
	jabba.RunOrDie("unzip", "-d", "/opt/codescreen/src", "1.1.0.zip")
	jabba.RunOrDie("chown", "-R", "codescreen:codescreen", "/opt/codescreen")
	jabba.RunOrDie("service", "codescreen", "start")
}

var codescreenUser = jabba.User{
	Username: "codescreen",
	Sudo:     false,
	Key:      "",
	Shell:    "/dev/null",
}

var codescreenUpstartFile = jabba.File{
	Path: "/etc/init/codescreen.conf",
	Perm: 0644,
	Vars: map[string]string{
		"secret":    secrets.Get("codescreen", "secret"),
		"domain":    secrets.Get("codescreen", "domain"),
		"emailUser": secrets.Get("codescreen", "emailUser"),
		"emailPass": secrets.Get("codescreen", "emailPass"),
		"emailFrom": secrets.Get("codescreen", "emailFrom"),
		"emailTo":   secrets.Get("codescreen", "emailTo"),
	},
	Template: ` description "Start codescreen"
author "@davidbanham"

start on runlevel [2345]
stop on runlevel [2345]

respawn
respawn limit 5 60

script
	chdir /opt/codescreen/src/codescreen-1.1.0
	sudo -E -u codescreen npm install 2>&1 | /usr/bin/env logger -t codescreen
	export PORT=3000
	export CODESCREEN_SECRET={{.secret}}
	export DOMAIN={{.domain}}
	export EMAIL_USER={{.emailUser}}
	export EMAIL_PASS={{.emailPass}}
	export EMAIL_FROM={{.emailFrom}}
	export EMAIL_TO={{.emailTo}}
	sudo -E -u codescreen /usr/bin/env node index.js 2>&1 | /usr/bin/env logger -t codescreen
end script`,
}
