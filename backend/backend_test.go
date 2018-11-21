package backend_test

import (
	"time"

	"github.com/beaukode/gohound/backend"
	. "gopkg.in/check.v1"
)

type interfaceTests struct {
	backend backend.Interface
}

func (b *interfaceTests) TestGetNextTodo(c *C) {
	todo, err := b.backend.GetNextTodo(10)
	c.Assert(err, IsNil)
	c.Assert(todo, HasLen, 4)
	c.Assert(todo[0].Probetype, Equals, "tcp-connect")
	c.Assert(todo[1].Probetype, Equals, "http-response")
	c.Assert(todo[2].Probetype, Equals, "tcp-connect")
	c.Assert(todo[3].Probetype, Equals, "tcp-connect")
}

func (b *interfaceTests) TestGetNextTodoUseLimit(c *C) {
	todo, err := b.backend.GetNextTodo(1)
	c.Assert(err, IsNil)
	c.Assert(todo, HasLen, 1)
}

func (b *interfaceTests) TestGetNextPreventConcurrentAccess(c *C) {
	fbe := b.backend
	// Consume 1st probe
	todo, err := fbe.GetNextTodo(1)
	c.Assert(err, IsNil)
	c.Assert(todo, HasLen, 1)
	c.Assert(todo[0].Probetype, Equals, "tcp-connect")

	// Consume next 3 probes
	todo, err = fbe.GetNextTodo(3)
	c.Assert(err, IsNil)
	c.Assert(todo, HasLen, 3)
	c.Assert(todo[0].Probetype, Equals, "http-response")

	// No more probe
	todo, err = fbe.GetNextTodo(10)
	c.Assert(err, IsNil)
	c.Assert(todo, HasLen, 0)
}

func (b *interfaceTests) TestUpdate(c *C) {
	fbe := b.backend

	todo, err := fbe.GetNextTodo(10)
	c.Assert(err, IsNil)
	c.Assert(todo, HasLen, 4)

	probe := todo[0]
	firstLockuid := probe.Lockuid
	firstLocktime := probe.Locktime
	firstProbetype := probe.Probetype

	probe.Lockuid = ""
	probe.Locktime = time.Time{}
	fbe.Update(probe)

	// Interval not elapsed
	todo, err = fbe.GetNextTodo(10)
	c.Assert(err, IsNil)
	c.Assert(todo, HasLen, 0)

	time.Sleep(time.Second)

	// Interval elapsed probe must came back with new lock uid & time
	todo, err = fbe.GetNextTodo(10)
	c.Assert(err, IsNil)
	c.Assert(todo, HasLen, 1)
	c.Assert(todo[0].Lockuid, Not(Equals), firstLockuid)
	c.Assert(todo[0].Locktime, Not(Equals), firstLocktime)
	c.Assert(todo[0].Probetype, Equals, firstProbetype)
}
