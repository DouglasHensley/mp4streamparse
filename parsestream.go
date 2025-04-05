package mp4streamparse

import (
	"context"
	std_log "log"
	"time"
)

// ParseStream top level mp4 stream parse
func ParseStream(ctx context.Context, chInbytes chan []byte, logger *std_log.Logger) (func() error, chan *TimestampBox) {
	fn := "ParseStream"

	chOutTSBox := make(chan *TimestampBox)

	outFn := func() (rcErr error) {
		logger.Printf("%s: Begin", fn)
		defer logger.Printf("%s: End", fn)

		defer func() {
			close(chOutTSBox)
		}()

		var workBuff []byte
		var workBuffCnt uint64 = 0
	TopLoop:
		for {
			select {
			case <-ctx.Done():
				logger.Printf("%s: App Shutdown(%v)", fn, ctx.Err())
				break TopLoop
			case inBytes, chInOpen := <-chInbytes:
				if !chInOpen && len(inBytes) == 0 {
					logger.Printf("%s: chInbytes Closed", fn)
					rcErr = nil
					break TopLoop
				}
				workBuffCnt += uint64(len(inBytes))
				workBuff = append(workBuff, inBytes...)

			BoxLoop:
				for {
					select {
					case <-ctx.Done():
						logger.Printf("%s: App Shutdown(%v)", fn, ctx.Err())
						rcErr = ctx.Err()
						break TopLoop
					default:
					}
					if int64(len(workBuff)) < BoxHeaderSize {
						if !chInOpen {
							rcErr = nil
							break TopLoop
						}
						break BoxLoop
					}
					offset, boxsize, success := FindNextBox(workBuff)
					if !success {
						break BoxLoop // get more data
					}
					if boxsize == 0 {
						if chInOpen {
							break BoxLoop
						}
						logger.Printf("%s: Final Atom", fn)
					}
					workBuff = workBuff[offset:]

					n, tsBox := ReadBox(workBuff[:])
					if tsBox.Mp4Box != nil {
						warnCnt := 0
					SendLoop0:
						for {
							select {
							case <-ctx.Done():
								logger.Printf("%s: App Shutdown(%v)", fn, ctx.Err())
								rcErr = ctx.Err()
								break TopLoop
							case chOutTSBox <- tsBox:
								dTime := time.Now().Sub(time.Unix(0, tsBox.TimestampUnixNano)).Milliseconds()
								if dTime > 50 {
									logger.Printf("%s: TimestampBox Timestamp(%s) - dMSec(%d) - FAILED SEND", fn,
										time.Unix(0, tsBox.TimestampUnixNano).UTC().Format("20060102T15:04:05.000Z07:00"), dTime)
								}
								break SendLoop0
							case <-time.After(time.Millisecond):
								warnCnt++
								if warnCnt > 25 {
									logger.Printf("%s: TimestampBox Timestamp(%s) - ATTEMPTING SEND", fn,
										time.Unix(0, tsBox.TimestampUnixNano).UTC().Format("20060102T15:04:05.000Z07:00"))
									warnCnt = 0
								}
							}
						}
					}
					workBuff = workBuff[n:] // Shift processed bytes out of working buffer
				} // END: BoxLoop
			} // END: Select
		} // END: TopLoop
		logger.Printf("%s: Total Bytes Processed(%d)", fn, workBuffCnt)
		return
	}
	return outFn, chOutTSBox
}
