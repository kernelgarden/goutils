package config

import (
	"os"
	"testing"
)

func TestRead(t *testing.T) {
	curPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	c := Config{}
	err = Read(curPath, "sample", &c)

	if err != nil {
		t.Logf("Failed to Read: %s\n", err)
		t.Fail()
	}

	if c.Debug != true {
		t.Logf("Failed c.Debug: %v\n", c.Debug)
		t.Fail()
	}

	if c.Database.Driver != "mysql" {
		t.Fail()
	}
}

type Config struct {
	Database struct {
		Driver string
		Connection string
		Logger string
	}

	Behaviorlog struct {
		Kafka string
	}

	Debug bool
	Service string
	Httpport string
}