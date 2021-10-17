#!/bin/bash

exec /usr/bin/supervisord -c /etc/supervisord.conf -n

chmod 777 /var/run/supervisor/supervisor.sock

