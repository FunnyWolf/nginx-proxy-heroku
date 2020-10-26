#!/bin/sh

# Download and install V2Ray
mkdir /tmp/portfwd
curl -L -H "Cache-Control: no-cache" -o /tmp/portfwd/portfwd.zip https://github.com/FunnyWolf/portfwd-heroku/releases/download/v1.0/releases.zip
unzip /tmp/portfwd/portfwd.zip -d /tmp/portfwd
install -m 755 /tmp/portfwd/portfwd /usr/local/bin/portfwd

# Remove temporary directory
rm -rf /tmp/portfwd

# Run portfwd
/usr/local/bin/portfwd -target "$UUID"
