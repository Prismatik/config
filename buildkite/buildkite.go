package buildkite

import (
	"github.com/prismatik/jabba"
	"github.com/prismatik/secrets"
	"strconv"
)

func Go() {
	jabba.AddUser(buildkiteUser)
	jabba.RunOrDie("sudo", "apt-get", "install", "mcrypt", "-y")
	jabba.RunOrDie("sudo", "usermod", "--append", "--groups", "sudo", "buildkite-agent")
	jabba.RunOrDie("sudo", "usermod", "--append", "--groups", "docker", "buildkite-agent")
	jabba.RunOrDie("sudo", "sh", "-c", "'echo deb https://apt.buildkite.com/buildkite-agent stable main > /etc/apt/sources.list.d/buildkite-agent.list'")
	jabba.RunOrDie("sudo", "apt-key", "adv", "--keyserver hkp://keyserver.ubuntu.com:80", "--recv-keys 32A37959C2FA5C3C99EFBC32A79206696452D198")
	jabba.RunOrDie("sudo", "apt-get", "update", "&&", "sudo", "apt-get", "install", "-y", "buildkite-agent")
	jabba.WriteFile(idRsaFile)
	jabba.WriteFile(secretFile)
	jabba.WriteFile(agentFile)
	jabba.RunOrDie("sudo", "systemctl", "enable", "buildkite-agent", "&&", "sudo", "systemctl", "start", "buildkite-agent")
	for i := 1; i < 3; i++ {
		s := strconv.Itoa(i)
		jabba.RunOrDie("sudo", "cp", "/lib/systemd/system/buildkite-agent", "/lib/systemd/system/buildkite-agent-"+s)
		jabba.RunOrDie("sudo", "systemctl", "enable", "buildkite-agent-"+s, "&&", "sudo", "systemctl", "start", "buildkite-agent-"+s)
	}
	jabba.RunOrDie("ssh", "-T", "git@github.com")
}

var buildkiteSecrets = map[string]string{
	"idRsaKey":       secrets.Get("buildkite", "idRsaKey"),
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
	Template: `{{.idRsaKey}}`,
}

var secretFile = jabba.File{
	Path:     "/var/lib/buildkite-agent/super_secret_key",
	Perm:     0644,
	Vars:     buildkiteSecrets,
	Template: `{{.superSecretKey}}`,
}

var agentFile = jabba.File{
	Path: "/etc/buildkite-agent/buildkite-agent.cfg",
	Perm: 0600,
	Vars: buildkiteSecrets,
	Template: `# The token from your Buildkite "Agents" page
token="{{.agentToken}}"

# The name of the agent
name="BuildABot"

# The priority of the agent (higher priorities are assigned work first)
# priority=1

# Meta-data for the agent (default is "queue=default")
# meta-data="key1=val2,key2=val2"

# Include the host's EC2 meta-data (instance-id, instance-type, and ami-id) as meta-data
# meta-data-ec2=true

# Include the host's EC2 tags as meta-data
# meta-data-ec2-tags=true

# Path to the bootstrap script. You should avoid changing this file as it will
# be overridden when you update your agent. If you need to make changes to this
# file: use the hooks provided, or copy the file and reference it here.
bootstrap-script="/usr/share/buildkite-agent/bootstrap.sh"

# Path to where the builds will run from
build-path="/var/lib/buildkite-agent/builds"

# Directory where the hook scripts are found
hooks-path="/etc/buildkite-agent/hooks"

# Do not run jobs within a pseudo terminal
# no-pty=true

# Don't automatically verify SSH fingerprints
# no-automatic-ssh-fingerprint-verification=true

# Don't allow this agent to run arbitrary console commands
# no-command-eval=true

# Enable debug mode
# debug=true

# Don't show colors in logging
# no-color=true`,
}
