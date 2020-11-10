package gopydns

import (
        "os/exec"
        "bytes"
        "strings"
        "errors"
)

func runCommand(args ...string) (string, error) {
        cmd := exec.Command(args[0], args[1:]...)

        var out bytes.Buffer
        var err bytes.Buffer

        cmd.Stdout = &out 
        cmd.Stderr = &err
        cmd.Run()

        // convert err to an error type if there is an error returned
        var e error
        if err.String() != "" {
                e = errors.New(err.String())
        }

        return strings.TrimRight(out.String(), "\r\n"), e
}

func RunPyDnsCommandCreate(name string, username string, password string, server string, zone string, ip string, dnspy string ) (string, error) {
        var pycmd string
        pycmd = "/usr/bin/python3 " + dnspy + " --name " + name + "--zone " + zone + " --ip " + ip + " --computer " + server + " --user " + username + "--password " + password
        out, err := runCommand(pycmd) 

        return out, err
}
