# NeuronOS
Codes for NeuronOS interview


# Build Instruction on Mac
    1. Download files from https://github.com/gnarychkine/NeuronOS
    2. cd <path-to-source-dir>
    3. > build.sh

    Script will generate executable ./dist/GN_Int and installation package ./GN_Int.pkg


# API
    This sample supports execution for two commands(commander.go)
        func (c *Commander) Ping(host string) (PingResult, error)
        func (c *Commander) GetSystemInfo() (SystemInfo, error)

        Datatypes:
            type Commander interface {
                Ping(host string) (PingResult, error)
                GetSystemInfo() (SystemInfo, error)
            }
            type PingResult struct {
                Successful bool
                Time       time.Duration
            }
            type SystemInfo struct {
                Hostname  string
                IPAddress string
            }

        Returns results via HTTP:
            type CommandResponse struct {
                Success bool        `json:"success"`
                Data    interface{} `json:"data"`
                Error   string      `json:"error,omitempty"`
            }


# Program installation on Mac
    Launch GN_Int.pkg and it will install GN_Int executable into Application directory


# Testing (GO)
    cd <path-to-source-dir>
    > go test

    expected result like:
    PASS
    ok  	gn.neu.com	2.623s

# Testing (curl)
    Launch GN_Int executable from Application directory
    open terminal window and try
    > curl -X GET "http://localhost:8080/execute?command=ping&host=www.google.com"
    > curl -X POST "http://localhost:8080/execute?command=ping&host=www.google.com"
    > curl -X GET "http://localhost:8080/execute?command=sysinfo"
    > curl -X POST "http://localhost:8080/execute?command=sysinfo"


