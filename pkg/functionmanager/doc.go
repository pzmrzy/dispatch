///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////

// Package gen contains generated code
// this code needs to be considered by go, so it can't be in a file that starts with _ or .
package functionmanager

// NO TESTS

//go:generate mkdir -p gen
//go:generate go-bindata -o ./gen/bindata.go -pkg gen -prefix '../../swagger' ../../swagger
//go:generate swagger generate server -A FunctionManager -t ./gen -f ../../swagger/function-manager.yaml --exclude-main
