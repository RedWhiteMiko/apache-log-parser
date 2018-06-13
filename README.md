# Apache Log Parser - UBS

Apache Log Parser

## Compile Instruction
- Requirements: `golang 1.8`
- Execute Test: `make test`
- Execute Test & Build: `make build`
- Execute Test, Build, & Run: `make run`

`$GOPATH` is set in the `Makefile`, 
Tested environment: Linux Fedora 25

## Structure
- `src/def`: contain classes definition for Logs
- `src/scanner`: Scan log string to parser
    - `src/scanner/file`: File scanner
- `src/parser`: Parser for string to ApacheLog struct
- `bin/`: Output of binary file, `bin/apache_test` also copied to root directory

## Binary execution
- `./apache_test [file_name] [is_exit_of_fail]`
    - `file_name`: Log file to be processed, default: `example.txt`
    - `is_exit_of_fail`: Should test exit if one test Fail, default: `true`, accept: `true, false, 0, 1, T, F`
    - Example: `./apache_test example.txt false`

## Assumption
- UserID (`%u`) and UserIdent (`%l`) doesn't contain spaces
- Logs is already sorted by date ascending
- In case logs unable to be parsed, it will be ignored
- In case `POST` doesn't find `PUT` before the logs, test `3. PUT before POST` will `Fail`
- For test `4. Suspicious activity`, `The log must have less than five 401 responses from any host` assumed as `Log must have less than five 401 responses from each Remote Hosts`

## Reference Links
- http://httpd.apache.org/docs/current/logs.html#common
  - To check apache log format
- https://stackoverflow.com/questions/30305542/using-positive-lookahead-regex-with-re2
  - Fix for case URL contain non urlencoded space
- https://golang.org/pkg/
  - API references