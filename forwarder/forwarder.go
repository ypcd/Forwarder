/*
*
*Open source agreement:
*   The project is based on the GPLv3 protocol.
*
*/
package main

import (
    "fmt"
    "net"
    "os"
    "time"
    "timerm"
)

func main() {
    
    service := ":8001"
    
    tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
    checkError(err)
    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)
    
    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        
        service := "127.0.0.1:8000"
        
        tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
        fmt.Println(tcpAddr)
        checkError(err)
        dst, err := net.Dial("tcp", service)
        checkError(err)
        fmt.Println("conn.")
        
        
        src := conn
        go srcTOdst( src, dst)
        go srcTOdst( dst, src)
        fmt.Println("go.")
    }
}

func srcTOdst(src net.Conn, dst net.Conn){
    
    tmr := timerm.CreateTimer(time.Second*60)
    //debug
    //outf, err := os.Create("recv.data")
    //checkError(err)
    defer src.Close()
    defer dst.Close()
    //debug
    //defer outf.Close()
    buf := make([]byte, 1024*64)
    var wbuf []byte
    
    for {
        
        //fmt.Println("read.")
        rlen, err := src.Read(buf)
        fmt.Fprintf(os.Stderr, "%d read end...", rlen)
        //checkError(err)
        if tmr.Run() {
            return
        }
        if rlen == 0 {
            return
        }        
        if err != nil{
            continue
        }
        //debug
        //outf.Write( buf[:rlen])
        tmr.Boot()
        wbuf = buf
        buf = buf[:rlen]
        for{
            if len(buf)>0 {                
                wlen, err := dst.Write(buf)
                if wlen == 0 {
                    return
                }
                if err != nil && wlen <= 0{
                    continue
                }
                if len(buf)==wlen {
                    break
                }
                buf = buf[wlen:]
            } else {
                break
            }
        }
        buf = wbuf
        
    }
}


func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}