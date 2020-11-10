package gopydns

import (
        "os/exec"
        "bytes"
        "strings"
        "errors"
        "fmt"
)

func runCommand(args []string) (string, error) {
        cmd := exec.Command("/usr/bin/python3", args...)

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
        dnscmd := []string{"create-dns.py","--name",name,"--zone",zone,"--ip",ip,"--computer",server,"--user",username,"--password",password}
        //pycmd = "/usr/bin/python3" + dnspy + " --name " + name + "--zone " + zone + " --ip " + ip + " --computer " + server + " --user " + username + "--password " + password
        out, err := runCommand(dnscmd) 
        fmt.Println(out)
        fmt.Println(err)
        return out, err
}


func RunPyDnsCommandRemove(name string, username string, password string, server string, zone string, ip string, dnspy string ) (string, error) {
        dnscmd := []string{"create-dns.py","--name",name,"--zone",zone,"--ip",ip,"--computer",server,"--user",username,"--password",password, "--remove"}
        //pycmd = "/usr/bin/python3" + dnspy + " --name " + name + "--zone " + zone + " --ip " + ip + " --computer " + server + " --user " + username + "--password " + password
        out, err := runCommand(dnscmd)
        fmt.Println(out)
        fmt.Println(err)
        return out, err
}

