# rpi-pinger

Simple utility designed run on a rasperry pi from startup and once every hour ping a remote server with the current value of the Pi:s IP address. 
This allows a user to connect to the pi via ssh without first hooking up the pi to a screen.

Requers the follwing to run:

On the server
Go 1.8
Docker 1.13 + 

On the pi
python 2.7 
pip
