// package wl acts as a client for the wlr_foreign_toplevel_management_unstable_v1 wayland protocol.

// generated by wl-scanner
// https://github.com/dkolbly/wl-scanner
// from: https://raw.githubusercontent.com/swaywm/wlr-protocols/master/unstable/wlr-foreign-toplevel-management-unstable-v1.xml
// on 2021-03-11 15:02:55 +0700
package main

import (
	"sync"

	wl "github.com/dkolbly/wl"
)

type ZwlrForeignToplevelManagerV1ToplevelEvent struct {
	Toplevel *ZwlrForeignToplevelHandleV1
}

type ZwlrForeignToplevelManagerV1ToplevelHandler interface {
	HandleZwlrForeignToplevelManagerV1Toplevel(ZwlrForeignToplevelManagerV1ToplevelEvent)
}

func (p *ZwlrForeignToplevelManagerV1) AddToplevelHandler(h ZwlrForeignToplevelManagerV1ToplevelHandler) {
	if h != nil {
		p.mu.Lock()
		p.toplevelHandlers = append(p.toplevelHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZwlrForeignToplevelManagerV1) RemoveToplevelHandler(h ZwlrForeignToplevelManagerV1ToplevelHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.toplevelHandlers {
		if e == h {
			p.toplevelHandlers = append(p.toplevelHandlers[:i], p.toplevelHandlers[i+1:]...)
			break
		}
	}
}

type ZwlrForeignToplevelManagerV1FinishedEvent struct {
}

type ZwlrForeignToplevelManagerV1FinishedHandler interface {
	HandleZwlrForeignToplevelManagerV1Finished(ZwlrForeignToplevelManagerV1FinishedEvent)
}

func (p *ZwlrForeignToplevelManagerV1) AddFinishedHandler(h ZwlrForeignToplevelManagerV1FinishedHandler) {
	if h != nil {
		p.mu.Lock()
		p.finishedHandlers = append(p.finishedHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZwlrForeignToplevelManagerV1) RemoveFinishedHandler(h ZwlrForeignToplevelManagerV1FinishedHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.finishedHandlers {
		if e == h {
			p.finishedHandlers = append(p.finishedHandlers[:i], p.finishedHandlers[i+1:]...)
			break
		}
	}
}

func (p *ZwlrForeignToplevelManagerV1) Dispatch(event *wl.Event) {
	switch event.Opcode {
	case 0:
		if len(p.toplevelHandlers) > 0 {
			ev := ZwlrForeignToplevelManagerV1ToplevelEvent{}
			ev.Toplevel = event.Proxy(p.Context()).(*ZwlrForeignToplevelHandleV1)
			p.mu.RLock()
			for _, h := range p.toplevelHandlers {
				h.HandleZwlrForeignToplevelManagerV1Toplevel(ev)
			}
			p.mu.RUnlock()
		}
	case 1:
		if len(p.finishedHandlers) > 0 {
			ev := ZwlrForeignToplevelManagerV1FinishedEvent{}
			p.mu.RLock()
			for _, h := range p.finishedHandlers {
				h.HandleZwlrForeignToplevelManagerV1Finished(ev)
			}
			p.mu.RUnlock()
		}
	}
}

type ZwlrForeignToplevelManagerV1 struct {
	wl.BaseProxy
	mu               sync.RWMutex
	toplevelHandlers []ZwlrForeignToplevelManagerV1ToplevelHandler
	finishedHandlers []ZwlrForeignToplevelManagerV1FinishedHandler
}

func NewZwlrForeignToplevelManagerV1(ctx *wl.Context) *ZwlrForeignToplevelManagerV1 {
	ret := new(ZwlrForeignToplevelManagerV1)
	ctx.Register(ret)
	return ret
}

// Stop will stop sending events.
//
//
// Indicates the client no longer wishes to receive events for new toplevels.
// However the compositor may emit further toplevel_created events, until
// the finished event is emitted.
//
// The client must not send any more requests after this one.
//
func (p *ZwlrForeignToplevelManagerV1) Stop() error {
	return p.Context().SendRequest(p, 0)
}

type ZwlrForeignToplevelHandleV1TitleEvent struct {
	Title string
}

type ZwlrForeignToplevelHandleV1TitleHandler interface {
	HandleZwlrForeignToplevelHandleV1Title(ZwlrForeignToplevelHandleV1TitleEvent)
}

func (p *ZwlrForeignToplevelHandleV1) AddTitleHandler(h ZwlrForeignToplevelHandleV1TitleHandler) {
	if h != nil {
		p.mu.Lock()
		p.titleHandlers = append(p.titleHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZwlrForeignToplevelHandleV1) RemoveTitleHandler(h ZwlrForeignToplevelHandleV1TitleHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.titleHandlers {
		if e == h {
			p.titleHandlers = append(p.titleHandlers[:i], p.titleHandlers[i+1:]...)
			break
		}
	}
}

type ZwlrForeignToplevelHandleV1AppIdEvent struct {
	AppId string
}

type ZwlrForeignToplevelHandleV1AppIdHandler interface {
	HandleZwlrForeignToplevelHandleV1AppId(ZwlrForeignToplevelHandleV1AppIdEvent)
}

func (p *ZwlrForeignToplevelHandleV1) AddAppIdHandler(h ZwlrForeignToplevelHandleV1AppIdHandler) {
	if h != nil {
		p.mu.Lock()
		p.appIdHandlers = append(p.appIdHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZwlrForeignToplevelHandleV1) RemoveAppIdHandler(h ZwlrForeignToplevelHandleV1AppIdHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.appIdHandlers {
		if e == h {
			p.appIdHandlers = append(p.appIdHandlers[:i], p.appIdHandlers[i+1:]...)
			break
		}
	}
}

type ZwlrForeignToplevelHandleV1OutputEnterEvent struct {
	Output *wl.Output
}

type ZwlrForeignToplevelHandleV1OutputEnterHandler interface {
	HandleZwlrForeignToplevelHandleV1OutputEnter(ZwlrForeignToplevelHandleV1OutputEnterEvent)
}

func (p *ZwlrForeignToplevelHandleV1) AddOutputEnterHandler(h ZwlrForeignToplevelHandleV1OutputEnterHandler) {
	if h != nil {
		p.mu.Lock()
		p.outputEnterHandlers = append(p.outputEnterHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZwlrForeignToplevelHandleV1) RemoveOutputEnterHandler(h ZwlrForeignToplevelHandleV1OutputEnterHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.outputEnterHandlers {
		if e == h {
			p.outputEnterHandlers = append(p.outputEnterHandlers[:i], p.outputEnterHandlers[i+1:]...)
			break
		}
	}
}

type ZwlrForeignToplevelHandleV1OutputLeaveEvent struct {
	Output *wl.Output
}

type ZwlrForeignToplevelHandleV1OutputLeaveHandler interface {
	HandleZwlrForeignToplevelHandleV1OutputLeave(ZwlrForeignToplevelHandleV1OutputLeaveEvent)
}

func (p *ZwlrForeignToplevelHandleV1) AddOutputLeaveHandler(h ZwlrForeignToplevelHandleV1OutputLeaveHandler) {
	if h != nil {
		p.mu.Lock()
		p.outputLeaveHandlers = append(p.outputLeaveHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZwlrForeignToplevelHandleV1) RemoveOutputLeaveHandler(h ZwlrForeignToplevelHandleV1OutputLeaveHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.outputLeaveHandlers {
		if e == h {
			p.outputLeaveHandlers = append(p.outputLeaveHandlers[:i], p.outputLeaveHandlers[i+1:]...)
			break
		}
	}
}

type ZwlrForeignToplevelHandleV1StateEvent struct {
	State []int32
}

type ZwlrForeignToplevelHandleV1StateHandler interface {
	HandleZwlrForeignToplevelHandleV1State(ZwlrForeignToplevelHandleV1StateEvent)
}

func (p *ZwlrForeignToplevelHandleV1) AddStateHandler(h ZwlrForeignToplevelHandleV1StateHandler) {
	if h != nil {
		p.mu.Lock()
		p.stateHandlers = append(p.stateHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZwlrForeignToplevelHandleV1) RemoveStateHandler(h ZwlrForeignToplevelHandleV1StateHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.stateHandlers {
		if e == h {
			p.stateHandlers = append(p.stateHandlers[:i], p.stateHandlers[i+1:]...)
			break
		}
	}
}

type ZwlrForeignToplevelHandleV1DoneEvent struct {
}

type ZwlrForeignToplevelHandleV1DoneHandler interface {
	HandleZwlrForeignToplevelHandleV1Done(ZwlrForeignToplevelHandleV1DoneEvent)
}

func (p *ZwlrForeignToplevelHandleV1) AddDoneHandler(h ZwlrForeignToplevelHandleV1DoneHandler) {
	if h != nil {
		p.mu.Lock()
		p.doneHandlers = append(p.doneHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZwlrForeignToplevelHandleV1) RemoveDoneHandler(h ZwlrForeignToplevelHandleV1DoneHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.doneHandlers {
		if e == h {
			p.doneHandlers = append(p.doneHandlers[:i], p.doneHandlers[i+1:]...)
			break
		}
	}
}

type ZwlrForeignToplevelHandleV1ClosedEvent struct {
}

type ZwlrForeignToplevelHandleV1ClosedHandler interface {
	HandleZwlrForeignToplevelHandleV1Closed(ZwlrForeignToplevelHandleV1ClosedEvent)
}

func (p *ZwlrForeignToplevelHandleV1) AddClosedHandler(h ZwlrForeignToplevelHandleV1ClosedHandler) {
	if h != nil {
		p.mu.Lock()
		p.closedHandlers = append(p.closedHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZwlrForeignToplevelHandleV1) RemoveClosedHandler(h ZwlrForeignToplevelHandleV1ClosedHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.closedHandlers {
		if e == h {
			p.closedHandlers = append(p.closedHandlers[:i], p.closedHandlers[i+1:]...)
			break
		}
	}
}

type ZwlrForeignToplevelHandleV1ParentEvent struct {
	Parent *ZwlrForeignToplevelHandleV1
}

type ZwlrForeignToplevelHandleV1ParentHandler interface {
	HandleZwlrForeignToplevelHandleV1Parent(ZwlrForeignToplevelHandleV1ParentEvent)
}

func (p *ZwlrForeignToplevelHandleV1) AddParentHandler(h ZwlrForeignToplevelHandleV1ParentHandler) {
	if h != nil {
		p.mu.Lock()
		p.parentHandlers = append(p.parentHandlers, h)
		p.mu.Unlock()
	}
}

func (p *ZwlrForeignToplevelHandleV1) RemoveParentHandler(h ZwlrForeignToplevelHandleV1ParentHandler) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i, e := range p.parentHandlers {
		if e == h {
			p.parentHandlers = append(p.parentHandlers[:i], p.parentHandlers[i+1:]...)
			break
		}
	}
}

func (p *ZwlrForeignToplevelHandleV1) Dispatch(event *wl.Event) {
	switch event.Opcode {
	case 0:
		if len(p.titleHandlers) > 0 {
			ev := ZwlrForeignToplevelHandleV1TitleEvent{}
			ev.Title = event.String()
			p.mu.RLock()
			for _, h := range p.titleHandlers {
				h.HandleZwlrForeignToplevelHandleV1Title(ev)
			}
			p.mu.RUnlock()
		}
	case 1:
		if len(p.appIdHandlers) > 0 {
			ev := ZwlrForeignToplevelHandleV1AppIdEvent{}
			ev.AppId = event.String()
			p.mu.RLock()
			for _, h := range p.appIdHandlers {
				h.HandleZwlrForeignToplevelHandleV1AppId(ev)
			}
			p.mu.RUnlock()
		}
	case 2:
		if len(p.outputEnterHandlers) > 0 {
			ev := ZwlrForeignToplevelHandleV1OutputEnterEvent{}
			ev.Output = event.Proxy(p.Context()).(*wl.Output)
			p.mu.RLock()
			for _, h := range p.outputEnterHandlers {
				h.HandleZwlrForeignToplevelHandleV1OutputEnter(ev)
			}
			p.mu.RUnlock()
		}
	case 3:
		if len(p.outputLeaveHandlers) > 0 {
			ev := ZwlrForeignToplevelHandleV1OutputLeaveEvent{}
			ev.Output = event.Proxy(p.Context()).(*wl.Output)
			p.mu.RLock()
			for _, h := range p.outputLeaveHandlers {
				h.HandleZwlrForeignToplevelHandleV1OutputLeave(ev)
			}
			p.mu.RUnlock()
		}
	case 4:
		if len(p.stateHandlers) > 0 {
			ev := ZwlrForeignToplevelHandleV1StateEvent{}
			ev.State = event.Array()
			p.mu.RLock()
			for _, h := range p.stateHandlers {
				h.HandleZwlrForeignToplevelHandleV1State(ev)
			}
			p.mu.RUnlock()
		}
	case 5:
		if len(p.doneHandlers) > 0 {
			ev := ZwlrForeignToplevelHandleV1DoneEvent{}
			p.mu.RLock()
			for _, h := range p.doneHandlers {
				h.HandleZwlrForeignToplevelHandleV1Done(ev)
			}
			p.mu.RUnlock()
		}
	case 6:
		if len(p.closedHandlers) > 0 {
			ev := ZwlrForeignToplevelHandleV1ClosedEvent{}
			p.mu.RLock()
			for _, h := range p.closedHandlers {
				h.HandleZwlrForeignToplevelHandleV1Closed(ev)
			}
			p.mu.RUnlock()
		}
	case 7:
		if len(p.parentHandlers) > 0 {
			ev := ZwlrForeignToplevelHandleV1ParentEvent{}
			ev.Parent = event.Proxy(p.Context()).(*ZwlrForeignToplevelHandleV1)
			p.mu.RLock()
			for _, h := range p.parentHandlers {
				h.HandleZwlrForeignToplevelHandleV1Parent(ev)
			}
			p.mu.RUnlock()
		}
	}
}

type ZwlrForeignToplevelHandleV1 struct {
	wl.BaseProxy
	mu                  sync.RWMutex
	titleHandlers       []ZwlrForeignToplevelHandleV1TitleHandler
	appIdHandlers       []ZwlrForeignToplevelHandleV1AppIdHandler
	outputEnterHandlers []ZwlrForeignToplevelHandleV1OutputEnterHandler
	outputLeaveHandlers []ZwlrForeignToplevelHandleV1OutputLeaveHandler
	stateHandlers       []ZwlrForeignToplevelHandleV1StateHandler
	doneHandlers        []ZwlrForeignToplevelHandleV1DoneHandler
	closedHandlers      []ZwlrForeignToplevelHandleV1ClosedHandler
	parentHandlers      []ZwlrForeignToplevelHandleV1ParentHandler
}

func NewZwlrForeignToplevelHandleV1(ctx *wl.Context) *ZwlrForeignToplevelHandleV1 {
	ret := new(ZwlrForeignToplevelHandleV1)
	ctx.Register(ret)
	return ret
}

// SetMaximized will requests that the toplevel be maximized.
//
//
// Requests that the toplevel be maximized. If the maximized state actually
// changes, this will be indicated by the state event.
//
func (p *ZwlrForeignToplevelHandleV1) SetMaximized() error {
	return p.Context().SendRequest(p, 0)
}

// UnsetMaximized will requests that the toplevel be unmaximized.
//
//
// Requests that the toplevel be unmaximized. If the maximized state actually
// changes, this will be indicated by the state event.
//
func (p *ZwlrForeignToplevelHandleV1) UnsetMaximized() error {
	return p.Context().SendRequest(p, 1)
}

// SetMinimized will requests that the toplevel be minimized.
//
//
// Requests that the toplevel be minimized. If the minimized state actually
// changes, this will be indicated by the state event.
//
func (p *ZwlrForeignToplevelHandleV1) SetMinimized() error {
	return p.Context().SendRequest(p, 2)
}

// UnsetMinimized will requests that the toplevel be unminimized.
//
//
// Requests that the toplevel be unminimized. If the minimized state actually
// changes, this will be indicated by the state event.
//
func (p *ZwlrForeignToplevelHandleV1) UnsetMinimized() error {
	return p.Context().SendRequest(p, 3)
}

// Activate will activate the toplevel.
//
//
// Request that this toplevel be activated on the given seat.
// There is no guarantee the toplevel will be actually activated.
//
func (p *ZwlrForeignToplevelHandleV1) Activate(seat *wl.Seat) error {
	return p.Context().SendRequest(p, 4, seat)
}

// Close will request that the toplevel be closed.
//
//
// Send a request to the toplevel to close itself. The compositor would
// typically use a shell-specific method to carry out this request, for
// example by sending the xdg_toplevel.close event. However, this gives
// no guarantees the toplevel will actually be destroyed. If and when
// this happens, the zwlr_foreign_toplevel_handle_v1.closed event will
// be emitted.
//
func (p *ZwlrForeignToplevelHandleV1) Close() error {
	return p.Context().SendRequest(p, 5)
}

// SetRectangle will the rectangle which represents the toplevel.
//
//
// The rectangle of the surface specified in this request corresponds to
// the place where the app using this protocol represents the given toplevel.
// It can be used by the compositor as a hint for some operations, e.g
// minimizing. The client is however not required to set this, in which
// case the compositor is free to decide some default value.
//
// If the client specifies more than one rectangle, only the last one is
// considered.
//
// The dimensions are given in surface-local coordinates.
// Setting width=height=0 removes the already-set rectangle.
//
func (p *ZwlrForeignToplevelHandleV1) SetRectangle(surface *wl.Surface, x int32, y int32, width int32, height int32) error {
	return p.Context().SendRequest(p, 6, surface, x, y, width, height)
}

// Destroy will destroy the zwlr_foreign_toplevel_handle_v1 object.
//
//
// Destroys the zwlr_foreign_toplevel_handle_v1 object.
//
// This request should be called either when the client does not want to
// use the toplevel anymore or after the closed event to finalize the
// destruction of the object.
//
func (p *ZwlrForeignToplevelHandleV1) Destroy() error {
	return p.Context().SendRequest(p, 7)
}

// SetFullscreen will request that the toplevel be fullscreened.
//
//
// Requests that the toplevel be fullscreened on the given output. If the
// fullscreen state and/or the outputs the toplevel is visible on actually
// change, this will be indicated by the state and output_enter/leave
// events.
//
// The output parameter is only a hint to the compositor. Also, if output
// is NULL, the compositor should decide which output the toplevel will be
// fullscreened on, if at all.
//
func (p *ZwlrForeignToplevelHandleV1) SetFullscreen(output *wl.Output) error {
	return p.Context().SendRequest(p, 8, output)
}

// UnsetFullscreen will request that the toplevel be unfullscreened.
//
//
// Requests that the toplevel be unfullscreened. If the fullscreen state
// actually changes, this will be indicated by the state event.
//
func (p *ZwlrForeignToplevelHandleV1) UnsetFullscreen() error {
	return p.Context().SendRequest(p, 9)
}

const (
	ZwlrForeignToplevelHandleV1StateMaximized  = 0
	ZwlrForeignToplevelHandleV1StateMinimized  = 1
	ZwlrForeignToplevelHandleV1StateActivated  = 2
	ZwlrForeignToplevelHandleV1StateFullscreen = 3
)

const (
	ZwlrForeignToplevelHandleV1ErrorInvalidRectangle = 0
)
