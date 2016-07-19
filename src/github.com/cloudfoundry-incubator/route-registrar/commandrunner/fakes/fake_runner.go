// This file was generated by counterfeiter
package fakes

import (
	"bytes"
	"sync"

	"github.com/cloudfoundry-incubator/route-registrar/commandrunner"
)

type FakeRunner struct {
	RunStub        func(outbuf, errbuff *bytes.Buffer) error
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		outbuf  *bytes.Buffer
		errbuff *bytes.Buffer
	}
	runReturns struct {
		result1 error
	}
	WaitStub        func() error
	waitMutex       sync.RWMutex
	waitArgsForCall []struct{}
	waitReturns struct {
		result1 error
	}
	KillStub        func() error
	killMutex       sync.RWMutex
	killArgsForCall []struct{}
	killReturns struct {
		result1 error
	}
}

func (fake *FakeRunner) Run(outbuf *bytes.Buffer, errbuff *bytes.Buffer) error {
	fake.runMutex.Lock()
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		outbuf  *bytes.Buffer
		errbuff *bytes.Buffer
	}{outbuf, errbuff})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub(outbuf, errbuff)
	} else {
		return fake.runReturns.result1
	}
}

func (fake *FakeRunner) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeRunner) RunArgsForCall(i int) (*bytes.Buffer, *bytes.Buffer) {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return fake.runArgsForCall[i].outbuf, fake.runArgsForCall[i].errbuff
}

func (fake *FakeRunner) RunReturns(result1 error) {
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRunner) Wait() error {
	fake.waitMutex.Lock()
	fake.waitArgsForCall = append(fake.waitArgsForCall, struct{}{})
	fake.waitMutex.Unlock()
	if fake.WaitStub != nil {
		return fake.WaitStub()
	} else {
		return fake.waitReturns.result1
	}
}

func (fake *FakeRunner) WaitCallCount() int {
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	return len(fake.waitArgsForCall)
}

func (fake *FakeRunner) WaitReturns(result1 error) {
	fake.WaitStub = nil
	fake.waitReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRunner) Kill() error {
	fake.killMutex.Lock()
	fake.killArgsForCall = append(fake.killArgsForCall, struct{}{})
	fake.killMutex.Unlock()
	if fake.KillStub != nil {
		return fake.KillStub()
	} else {
		return fake.killReturns.result1
	}
}

func (fake *FakeRunner) KillCallCount() int {
	fake.killMutex.RLock()
	defer fake.killMutex.RUnlock()
	return len(fake.killArgsForCall)
}

func (fake *FakeRunner) KillReturns(result1 error) {
	fake.KillStub = nil
	fake.killReturns = struct {
		result1 error
	}{result1}
}

var _ commandrunner.Runner = new(FakeRunner)
