#!/bin/bash

pkill chess
nohup ./chess -conf config/test.yml &
