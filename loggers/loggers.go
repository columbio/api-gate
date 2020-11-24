package loggers

import (
	"io"
	"log"
)

var (
	// Trace - logger handler for trace messages
	Trace *log.Logger
	// Info - logger handler for info messages
	Info *log.Logger
	// Warning - logger handler for Warning messages
	Warning *log.Logger
	// Error - logger handler for Error messages
	Error *log.Logger
)

// Init - init logging operations
func Init(traceHandle io.Writer, infoHandle io.Writer, warningHandle io.Writer, errorHandle io.Writer) {
	Trace = log.New(traceHandle, "API-GATE TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(infoHandle, "API-GATE INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(warningHandle, "API-GATE WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(errorHandle, "API-GATE ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
