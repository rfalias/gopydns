#!/bin/env python
import argparse
from pypsrp.client import Client
import getpass
par = argparse.ArgumentParser(description='Create a DNS Record on a target domain')
par.add_argument('--name', type=str, help='Name to create')
par.add_argument('--zone', type=str, help='Zone to create record in')
par.add_argument('--ip', type=str, help='IP Address of the record')
par.add_argument('--computer',type=str, help='Remote computer to run against, usually a DC')
par.add_argument('--user', type=str, help='Username')
par.add_argument('--password', type=str, help='Password')
par.add_argument('--ask-pass', help='Ask for password', action='store_true')
args = par.parse_args()
print(args)

if args.ask_pass:
    password = getpass.getpass("Enter password: ")
else:
    if not args.password:
        print("Specify password if ask-pass not setup")
        exit(-1)
    password = args.password
ms = Client(args.computer, username=args.user,password=password, ssl=False)
ps_cmd = "Add-dnsserverresourcerecorda -name '%s' -ZoneName '%s' -IPV4Address '%s' -AllowUpdateAny" % (args.name, args.zone, args.ip)
output, streams, had_errors = ms.execute_ps(ps_cmd)
print("Status: %s, Success: %s" % (output, not had_errors))

f = open("./test", "a")
f.write("Testing!!")
f.close()
