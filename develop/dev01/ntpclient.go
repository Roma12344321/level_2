package main

import (
	"github.com/beevik/ntp"
	"time"
)

type NTPClient interface {
	GetTime() (time.Time, error)
}

type RealNTPClient struct{}

func (c *RealNTPClient) GetTime() (time.Time, error) {
	return ntp.Time("0.beevik-ntp.pool.ntp.org")
}
