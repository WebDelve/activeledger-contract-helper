/*
* MIT License (MIT)
* Copyright (c) 2023 WebDelve Ltd
*
* Permission is hereby granted, free of charge, to any person obtaining a copy
* of this software and associated documentation files (the "Software"), to deal
* in the Software without restriction, including without limitation the rights
* to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
* copies of the Software, and to permit persons to whom the Software is
* furnished to do so, subject to the following conditions:
*
* The above copyright notice and this permission notice shall be included in all
* copies or substantial portions of the Software.
*
* THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
* IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
* FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
* AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
* LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
* OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
* SOFTWARE.
 */

package main

import (
	"flag"
	"log"
	"os"
)

type SetArgs struct {
	build     bool
	run       bool
	inputFile string
}

func main() {
	log.Println("Starting")
	_ = loadConfig()

	_ = handleArgs()
}

func handleArgs() SetArgs {

	buildPtr := flag.Bool("b", false, "Build and sign the transaction")
	runPtr := flag.Bool("r", false, "Build, sign, and run the transaction")
	inputFilePtr := flag.String("i", "tx-data/tx.json", "A file to pull transaction data from")

	flag.Parse()

	args := SetArgs{
		build:     *buildPtr,
		run:       *runPtr,
		inputFile: *inputFilePtr,
	}

	// Show arg options and exit if none provided
	if !args.build && !args.run {
		log.Println(flag.ErrHelp)
		os.Exit(0)
	}

	return args
}
