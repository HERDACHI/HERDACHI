#!/bin/bash
source .env.sh
source venv/bin/activate
python mysql2postgres.py
