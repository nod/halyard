#!/usr/bin/env python3

import sys
import requests

addr = sys.args[1]

r = requests.post(addr, json={slug="feh", label="cool thing"})

