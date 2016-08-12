package base

import (
  "github.com/prismatik/jabba"
)

var sudoers = jabba.File{
  Path: "/etc/sudoers.d/ubuntu",
  Perm: 0400,
  Template: `# User rules for group ubuntu
%ubuntu ALL=(ALL) NOPASSWD:ALL`,
}