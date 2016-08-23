package buildkite

import (
	"github.com/prismatik/jabba"
	"github.com/prismatik/secrets"
)

func Go() {
	jabba.AddUser(buildkiteUser)
	jabba.RunOrDie("sudo", "usermod", "--append", "--groups", "sudo", "buildkite-agent")
	jabba.RunOrDie("sudo usermod", "--append", "--groups", "docker", "buildkite-agent")
	jabba.RunOrDie("sudo", "sh", "-c", "'echo deb https://apt.buildkite.com/buildkite-agent stable main > /etc/apt/sources.list.d/buildkite-agent.list'")
	jabba.RunOrDie("sudo", "apt-key", "adv", "--keyserver hkp://keyserver.ubuntu.com:80", "--recv-keys 32A37959C2FA5C3C99EFBC32A79206696452D198")
	jabba.RunOrDie("sudo", "apt-get", "update", "&&", "sudo", "apt-get", "install", "-y", "buildkite-agent")
	jabba.WriteFile(idRsaFile)
	jabba.WriteFile(secretFile)
	jabba.RunOrDie("sudo", "sed", "-i", "\"s/xxx/"+secrets.Get("buildkite", "agentToken")+"/g\"", "/etc/buildkite-agent/buildkite-agent.cfg")
	jabba.RunOrDie("sudo", "systemctl", "enable", "buildkite-agent", "&&", "sudo", "systemctl", "start", "buildkite-agent")
}

var buildkiteSecrets = map[string]string{
	"idRsaFile":      secrets.Get("buildkite", "idRsaFile"),
	"agentToken":     secrets.Get("buildkite", "agentToken"),
	"superSecretKey": secrets.Get("buildkite", "superSecretKey"),
}

var buildkiteUser = jabba.User{
	Username: "buildkite-agent",
	Sudo:     true,
	Key:      "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDJ3JOjlmXmTB6v51QUQe5px4VIOYk0DwEUH9/59wfjz9GNG6sLW2kqXfmHdLb0RGAlt+xIvPHT1PeY3QI9rdxnfAzH0UjBJpBdn7lXT2afPB4NffOH/iAOyP3wPoZ48T20UK3zZn29lqiy+03j0tRmeoCHFXgUiOnYJss/W9cLR48LuoCAUgiaeFMfR+EFdRK7zn7rBw1cvahu3zILybBT2CLH4kakfLwjr6MDJMX3VOdMHCu8v4NU/pHhX4/EVF6lnDXR6mXQzX5i4werS4B5kshqz5T7PHytKgw7fDaL+rGIFSxVW27zjhTRTEl8/XMszIWEWpkJSe4JEMFnZXFd",
}

var idRsaFile = jabba.File{
	Path:     "/home/buildkite-agent/.ssh/id_rsa",
	Perm:     0600,
	Vars:     buildkiteSecrets,
	Template: `{{.idRsaFile}}`,
}

var secretFile = jabba.File{
	Path:     "/var/lib/buildkite-agent/super_secret_key",
	Perm:     0644,
	Vars:     buildkiteSecrets,
	Template: `{{.superSecretKey}}`,
}
