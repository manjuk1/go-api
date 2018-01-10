

This project was setup on Ubuntu

Download the archive from official go download page

sudo tar -C /usr/local -xzf ~/Downloads/go1.9.2.linux-amd64.tar.gz

Add /usr/local/go/bin to the PATH environment variable. You can do this by adding this line to your /etc/profile (for a system-wide installation) or $HOME/.profile:

export PATH=$PATH:/usr/local/go/bin

Set the GOPATH to your project workspace root - I have set it in ~/.profile as shown below

export GOPATH="${HOME}/projects/go-projects/go-api"

