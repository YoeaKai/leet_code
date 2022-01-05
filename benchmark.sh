#!/bin/bash

go test -v application/application_test.go -bench=. -run=none -benchmem .
