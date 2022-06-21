#!/bin/bash

go test -v application/go/application_test.go -bench=. -run=none -benchmem .
