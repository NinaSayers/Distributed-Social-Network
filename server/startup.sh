#!/bin/sh
ip route del default 
ip route add default via 10.0.11.254
