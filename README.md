# goclone

A simple tool I made to making cloning repos to the appropriate place in the $GOPATH in the event it is not something you can or want to use "go get".

My primary reason for this was I was downloading a lot of go-flutter examples and each time I would have to go into $GOPATH/src/github.com/ and then create a folder for the creators user name, then clone to it manually each time. This tool parses the url for the username, creates the directory, then clones to it.
If you build it, you can just rename it goclone or whatever you would like then place it in /usr/local/bin/

The -f flag signifies which folder you want as the root. 

```-f go``` will make the root $GOPATH  
```-f git``` will use a new ENV variable I created $GITPATH which I pointed to a folder I use for non Go related repos, but I don't have that one parse the username, it just auto clones to the $GITPATH folder 

Example usage: ```$ goclone -f go -u https://github.com/instance-id/goclone.git``` is the equivalent of ```mkdir $GOPATH/src/github.com/username && cd $GOPATH/src/github.com/username && git clone https://github.com/username/repo.git```