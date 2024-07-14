#! /bin/bash

cd .venv/lib/python3.12/site-packages
zip -r ../../../../lambda.zip .

cd ../../../../lambda/
zip ../lambda.zip main.py

cd ../