#!/bin/sh
#
# Starts and stops eponine. Toss this in your /etc/init.d.

# The path to the eponine directory
path=/home/pi/apps/eponine

cd $path

case "$1" in
  start)
    ./eponine &
    ;;
  stop)
    sudo killall eponine
    ;;
  *)
    echo "Usage:  {start|stop}"
    exit 1
    ;;
esac
exit $?
