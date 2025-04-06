// MP4 Parser
package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	std_log "log"
	"os"
	"path/filepath"
	"strings"
	"time"

	mp4 "github.com/DouglasHensley/mp4streamparse"

	"golang.org/x/sync/errgroup"
)

var appName string
var fileName string
var logger *std_log.Logger
var logFile *os.File

const (
	defaultFileName = ""
	usageFileName   = "Input file path"
)

func init() {
	flag.StringVar(&fileName, "file", defaultFileName, usageFileName)
	flag.StringVar(&fileName, "f", defaultFileName, usageFileName+" (shorthand)")

	flag.Parse()

	// Initialize Logger
	appName = filepath.Base(os.Args[0])
	appName = strings.TrimSuffix(appName, filepath.Ext(appName))
	logName := fmt.Sprintf("%s.log", appName)
	// Open the log file. Create it if it doesn't exist. Append if it does.
	logFile, _ = os.OpenFile(logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	logger = std_log.New(logFile, fmt.Sprintf("%s ", appName), std_log.Ldate|std_log.Ltime|std_log.Lmicroseconds|std_log.Lshortfile)
}

func main() {
	fn := "main"

	defer logFile.Close()

	var fp *os.File = nil
	if fileName != defaultFileName {
		fp, _ = os.Open(fileName)
	}
	defer func(fp *os.File) {
		if fp != nil {
			fp.Close()
		}
	}(fp)

	fileInfo, _ := os.Stat(fileName)
	logger.Printf("BEGIN: %s(%s, %v)", appName, fileName, fileInfo.Size())
	g, appCtx := errgroup.WithContext(context.Background())

	var inputMp4Channel = make(chan []byte)
	fnMp4File := mp4.ParseFile(appCtx, inputMp4Channel, logger)

	fnFileRead := func() (rcErr error) {
		fn := "FileRead"
		logger.Printf("%s: Begin", fn)
		defer logger.Printf("%s: End", fn)

		nReads := 0
		accum := 0
		fp.Seek(0, 0)
	FileReadLoop:
		for {
			select {
			case <-appCtx.Done():
				logger.Printf("%s: App Shutdown(%v)", fn, appCtx.Err())
				rcErr = appCtx.Err()
				break FileReadLoop
			default:
			}
			readBuff := make([]byte, 1316)
			n, err := fp.Read(readBuff[0:])
			accum += n
			// Send data on inputMp4Channel
		SendLoop0:
			for {
				select {
				case <-appCtx.Done():
					logger.Printf("%s: App Shutdown(%v)", fn, appCtx.Err())
					rcErr = appCtx.Err()
					break FileReadLoop
				case inputMp4Channel <- readBuff[0:n]:
					break SendLoop0
				case <-time.After(10 * time.Millisecond):
				}
			}
			nReads++
			if err != nil {
				if err == io.EOF {
					logger.Printf("%s: EOF(%s) Read(%d)", fn, fileName, accum)
					rcErr = nil
					break FileReadLoop
				} else {
					strErr := fmt.Sprintf("%s: Read Error(%s)", fn, fileName)
					logger.Printf("%s", strErr)
					rcErr = errors.New(strErr)
					break FileReadLoop
				}
			}
		} // END: FileReadLoop
		logger.Printf("%s: Exiting File Read, Num Reads(%d) BytesRead(%d)", fn, nReads, accum)
		rcErr = nil
		logger.Printf("%s: Close(inputMp4Channel)", fn)
		close(inputMp4Channel)
		return
	}

	// Begin Timing Loop
	start := time.Now()
	// Launch Processing Chain
	g.Go(fnMp4File)
	g.Go(fnFileRead)

	logger.Printf("%s: Begin Error Group Monitoring", fn)
	if err := g.Wait(); err != nil {
		logger.Printf("%s: Processing Error(%s)", fn, err.Error())
	}
	logger.Printf("%s: End Error Group Monitoring", fn)

	end := time.Now()
	logger.Printf("END: %s(%s) Elapsed Time(%v)", appName, fileName, end.Sub(start))
}
