package base

import (
  "github.com/prismatik/jabba"
)

jabba.RunOrDie("echo", "\"#include /etc/sudoers.local\"", ">>", "/etc/sudoers")

var sudoers = jabba.File{
  Path: "/etc/sudoers.d/ubuntu",
  Perm: 0400,
  Template: `# User rules for group ubuntu
%sudo ALL=(ALL) NOPASSWD:ALL`,
}