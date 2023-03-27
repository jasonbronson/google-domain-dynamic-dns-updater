This uses a script to allow you to update your Google domain using dynamic DNS. It uses Golang http call to get your address from Amazon's website then updates it through google.

The script runs to be lightweight and runs every hour on cron. 

Set the following variables in your environment

$DOMAIN
$USERNAME
$PASSWORD

Domain is the full dynamic dns name. You get user/pass from google's domain ui. 

Container image can be found at https://hub.docker.com/r/jbronson29/google-dynamic-dns-updater
