## Config

This is Prismatik's machine configuration. It uses Jabba

https://github.com/prismatik/jabba

It also uses Glide to manage dependencies. You should download Glide and put the binary somewhere in your $PATH

https://github.com/Masterminds/glide/releases

You can then run `glide install` and you're good to go.

### Adding new configs

Write some code. Organise it in subfolders and with go packages.

To check you haven't done anything really silly, `go build main.go` and see if you get any compiler errors

To see whether the code you wrote actually does what you wanted it to:

```
docker build -t config .
docker run -i config
```

You are now at the bash prompt of an Ubuntu machine that has had the config binary run on it. Poke around and see if things are like you expect them to be. Are the users in place? Does the config file look right? Did the output from the docker build look sane?

Protip: You might want to comment out the lines about installing new packages, particularly build-essential. The build process will download all the packages fresh each time and 40MB of downloads each test run gets tiresome.

The Dockerfile is only intended for testing purposes. I don't expect this binary to ever be useful in a dockerised form.

### Roles

You can specify roles as a command line flag, like so:

```
config --roles=base,ufw
```

Additional roles are added as part of the switch statement
