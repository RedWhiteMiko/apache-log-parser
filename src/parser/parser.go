package parser

import (
	"def/log"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Parser struct {
	Logs []*log.ApacheLog
}

func Init() *Parser {
	parser := new(Parser)
	return parser
}

// Map - Map log to ApacheLog.
func (p *Parser) Map(input string) (*log.ApacheLog, error) {
	input = strings.Trim(input, " \n\t")
	re := regexp.MustCompile(`^(\S+)\s` + // RemoteHost
		`(\S+)\s` + // RFC 1413 UserIdentifier
		`(\S+)\s` + // UserId
		`\[(\d{2}/\w{3}/\d{2}(?:\d{2}:){3}\d{2} [-+]\d{4})\]\s` + // Date [10/Oct/2000:13:55:36 -0700]
		`"(.+?)(?:\s)(.+)\s(.*)"\s` + // Request
		`(\d+)\s` + // StatusCode
		`(\d+|-)`) // RequestSize Byte or "-" (Depend on %b or %B format)
	matches := re.FindStringSubmatch(input)
	if len(matches) < 1 {
		return nil, fmt.Errorf("failed to parse log, Ignoring: %s", input)
	}

	// Convert SizeByte
	size, err := strconv.ParseUint(matches[9], 10, 64)
	if err != nil {
		size = 0
	}

	log := &log.ApacheLog{
		RemoteHost:  matches[1],
		UserIdentd:  matches[2],
		UserID:      matches[3],
		ReqTime:     matches[4],
		ReqMethod:   matches[5],
		ReqResource: matches[6],
		ReqProtocol: matches[7],
		StatusCode:  matches[8],
		SizeByte:    size,
	}
	p.Logs = append(p.Logs, log)
	return log, nil
}
