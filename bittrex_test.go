package bittrex

import (
"testing"
"net"
"net/http"
"net/url"
"time"
)



func TestBittrex_GetWithdrawStatus(t *testing.T) {
    var (
        apiKey    = ""
        secretKey = ""

        proxyString = "https://127.0.0.1:1087"
        proxyUrl, _ = url.Parse(proxyString)
        trans           = &http.Transport{
            Dial: (&net.Dialer{
                Timeout:   60 * time.Second,
                KeepAlive: 30 * time.Second,
            }).Dial,
            // We use ABSURDLY large keys, and should probably not.
            TLSHandshakeTimeout: 60 * time.Second,
            Proxy:               http.ProxyURL(proxyUrl),
        }
        c0 = &http.Client{
            Transport: trans,
            Timeout:   15 * time.Second,
        }

        c = NewWithCustomHttpClient(apiKey,secretKey,c0)
    )
    a,e := c.GetPendingWithdrawals("BTC")
    t.Log(a,e)
}
