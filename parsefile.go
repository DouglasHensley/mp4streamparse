package mp4streamparse

import (
	"context"
	std_log "log"
)

// ParseFile toplevel mp4 atom parse
func ParseFile(ctx context.Context, chInbytes chan []byte, logger *std_log.Logger) func() error {
	fn := "ParseFile"

	outFn := func() (rcErr error) {
		logger.Printf("%s: Begin", fn)
		defer logger.Printf("%s: End", fn)

		BoxReport := make([]string, 0)
		var workBuff []byte
		var workBuffCnt uint64 = 0
		var n uint32
	TopLoop:
		for inBytes := range chInbytes {
			select {
			case <-ctx.Done():
				logger.Printf("%s: App Shutdown(%v)", fn, ctx.Err())
				rcErr = nil
				break TopLoop
			default:
			}
			workBuffCnt += uint64(len(inBytes))
			workBuff = append(workBuff, inBytes...)

		BoxLoop:
			for {
				select {
				case <-ctx.Done():
					logger.Printf("%s: App Shutdown(%v)", fn, ctx.Err())
					rcErr = nil
					break TopLoop
				default:
				}
				if int64(len(workBuff)) < BoxHeaderSize {
					break BoxLoop
				}
				offset, _, success := FindNextBox(workBuff)
				if !success {
					break BoxLoop // get more data
				}
				workBuff = workBuff[offset:]
				boxSize, boxType := ParseHeader(workBuff)
				if boxSize == 0 {
					logger.Printf("%s: Final atom", fn)
					break TopLoop
				}
				logger.Printf("%s: boxType(%s) boxSize(%d) workBuff(%d)", fn, boxType, boxSize, len(workBuff))

				n, BoxReport = ReadBoxes(workBuff[:boxSize], BoxReport)
				if n != boxSize {
					logger.Printf("%s: BuffSize(%d) != ReadBoxes(%d)", fn, n, boxSize)
				}
				workBuff = workBuff[boxSize:] // Shift boxSize bytes out of working buffer
			} // END: BoxLoop
		} // END: TopLoop
		logger.Printf("%s: Total Bytes Processed(%d)", fn, workBuffCnt)
		logger.Printf("%s: MP4 BoxReport(%d):\n%v", fn, len(BoxReport), BoxReport)
		return
	}
	return outFn
}

/*
// ParseFile toplevel mp4 atom parse
func ParseFile(ctx context.Context, chInbytes chan []byte, logger *std_log.Logger) (func() error, chan TopBox) {
	fn := "ParseFile"

	chBoxTree := make(chan TopBox, 1)

	outFn := func() (rcErr error) {
		logger.Printf("%s: Begin", fn)
		defer logger.Printf("%s: End", fn)

		defer func() {
			close(chBoxTree)
		}()

		BoxTree := TopBox{}
		var workBuff []byte
		var workBuffCnt uint64 = 0
	TopLoop:
		for {
			select {
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
						rcErr = nil
						break TopLoop
					default:
					}
					if int64(len(workBuff)) < BoxHeaderSize {
						if !chInOpen {
							logger.Printf("%s: chInbytes Closed", fn)
							rcErr = nil
							break TopLoop
						}
						break BoxLoop
					}
					offset, _, success := FindNextBox(workBuff)
					if !success {
						break BoxLoop // get more data
					}
					workBuff = workBuff[offset:]
					boxSize, boxType := ParseHeader(workBuff)
					if boxSize == 0 {
						logger.Printf("%s: Final atom", fn)
						break TopLoop
					}
					logger.Printf("%s: boxType(%s) boxSize(%d) workBuff(%d)", fn, boxType, boxSize, len(workBuff))

					n := ReadBoxes(workBuff[:boxSize], &BoxTree)
					if n != boxSize {
						logger.Printf("%s: BuffSize(%d) != ReadBoxes(%d)", fn, n, boxSize)
					}
					workBuff = workBuff[boxSize:] // Shift boxSize bytes out of working buffer
				} // END: BoxLoop
			} // END: Select
		} // END: TopLoop
		logger.Printf("%s: Send BoxTree on buffered channel", fn)
		// chBoxTree <- BoxTree
		logger.Printf("%s: Total Bytes Processed(%d)", fn, workBuffCnt)
		return
	}
	return outFn, chBoxTree
}
*/
