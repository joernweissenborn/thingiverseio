
/*
 * generated by event_generator
 *
 * DO NOT EDIT
 */

package thingiverseio

import "github.com/joernweissenborn/eventual2go"



type StreamEventCompleter struct {
	*eventual2go.Completer
}

func NewStreamEventCompleter() *StreamEventCompleter {
	return &StreamEventCompleter{eventual2go.NewCompleter()}
}

func (c *StreamEventCompleter) Complete(d StreamEvent) {
	c.Completer.Complete(d)
}

func (c *StreamEventCompleter) Future() *StreamEventFuture {
	return &StreamEventFuture{c.Completer.Future()}
}

type StreamEventFuture struct {
	*eventual2go.Future
}

func (f *StreamEventFuture) Result() StreamEvent {
	return f.Future.Result().(StreamEvent)
}

type StreamEventCompletionHandler func(StreamEvent) StreamEvent

func (ch StreamEventCompletionHandler) toCompletionHandler() eventual2go.CompletionHandler {
	return func(d eventual2go.Data) eventual2go.Data {
		return ch(d.(StreamEvent))
	}
}

func (f *StreamEventFuture) Then(ch StreamEventCompletionHandler) *StreamEventFuture {
	return &StreamEventFuture{f.Future.Then(ch.toCompletionHandler())}
}

func (f *StreamEventFuture) AsChan() chan StreamEvent {
	c := make(chan StreamEvent, 1)
	cmpl := func(d chan StreamEvent) StreamEventCompletionHandler {
		return func(e StreamEvent) StreamEvent {
			d <- e
			close(d)
			return e
		}
	}
	ecmpl := func(d chan StreamEvent) eventual2go.ErrorHandler {
		return func(error) (eventual2go.Data, error) {
			close(d)
			return nil, nil
		}
	}
	f.Then(cmpl(c))
	f.Err(ecmpl(c))
	return c
}

type StreamEventStreamController struct {
	*eventual2go.StreamController
}

func NewStreamEventStreamController() *StreamEventStreamController {
	return &StreamEventStreamController{eventual2go.NewStreamController()}
}

func (sc *StreamEventStreamController) Add(d StreamEvent) {
	sc.StreamController.Add(d)
}

func (sc *StreamEventStreamController) Join(s *StreamEventStream) {
	sc.StreamController.Join(s.Stream)
}

func (sc *StreamEventStreamController) JoinFuture(f *StreamEventFuture) {
	sc.StreamController.JoinFuture(f.Future)
}

func (sc *StreamEventStreamController) Stream() *StreamEventStream {
	return &StreamEventStream{sc.StreamController.Stream()}
}

type StreamEventStream struct {
	*eventual2go.Stream
}

type StreamEventSubscriber func(StreamEvent)

func (l StreamEventSubscriber) toSubscriber() eventual2go.Subscriber {
	return func(d eventual2go.Data) { l(d.(StreamEvent)) }
}

func (s *StreamEventStream) Listen(ss StreamEventSubscriber) *eventual2go.Completer {
	return s.Stream.Listen(ss.toSubscriber())
}

func (s *StreamEventStream) ListenNonBlocking(ss StreamEventSubscriber) *eventual2go.Completer {
	return s.Stream.ListenNonBlocking(ss.toSubscriber())
}

type StreamEventFilter func(StreamEvent) bool

func (f StreamEventFilter) toFilter() eventual2go.Filter {
	return func(d eventual2go.Data) bool { return f(d.(StreamEvent)) }
}

func toStreamEventFilterArray(f ...StreamEventFilter) (filter []eventual2go.Filter){

	filter = make([]eventual2go.Filter, len(f))
	for i, el := range f {
		filter[i] = el.toFilter()
	}
	return
}

func (s *StreamEventStream) Where(f ...StreamEventFilter) *StreamEventStream {
	return &StreamEventStream{s.Stream.Where(toStreamEventFilterArray(f...)...)}
}

func (s *StreamEventStream) WhereNot(f ...StreamEventFilter) *StreamEventStream {
	return &StreamEventStream{s.Stream.WhereNot(toStreamEventFilterArray(f...)...)}
}

func (s *StreamEventStream) TransformWhere(t eventual2go.Transformer, f ...StreamEventFilter) *eventual2go.Stream {
	return s.Stream.TransformWhere(t, toStreamEventFilterArray(f...)...)
}

func (s *StreamEventStream) Split(f StreamEventFilter) (*StreamEventStream, *StreamEventStream)  {
	return s.Where(f), s.WhereNot(f)
}

func (s *StreamEventStream) First() *StreamEventFuture {
	return &StreamEventFuture{s.Stream.First()}
}

func (s *StreamEventStream) FirstWhere(f... StreamEventFilter) *StreamEventFuture {
	return &StreamEventFuture{s.Stream.FirstWhere(toStreamEventFilterArray(f...)...)}
}

func (s *StreamEventStream) FirstWhereNot(f ...StreamEventFilter) *StreamEventFuture {
	return &StreamEventFuture{s.Stream.FirstWhereNot(toStreamEventFilterArray(f...)...)}
}

func (s *StreamEventStream) AsChan() (c chan StreamEvent, stop *eventual2go.Completer) {
	c = make(chan StreamEvent)
	stop = s.Listen(pipeToStreamEventChan(c))
	stop.Future().Then(closeStreamEventChan(c))
	return
}

func pipeToStreamEventChan(c chan StreamEvent) StreamEventSubscriber {
	return func(d StreamEvent) {
		c <- d
	}
}

func closeStreamEventChan(c chan StreamEvent) eventual2go.CompletionHandler {
	return func(d eventual2go.Data) eventual2go.Data {
		close(c)
		return nil
	}
}

type StreamEventCollector struct {
	*eventual2go.Collector
}

func NewStreamEventCollector() *StreamEventCollector {
	return &StreamEventCollector{eventual2go.NewCollector()}
}

func (c *StreamEventCollector) Add(d StreamEvent) {
	c.Collector.Add(d)
}

func (c *StreamEventCollector) AddFuture(f *StreamEventFuture) {
	c.Collector.Add(f.Future)
}

func (c *StreamEventCollector) AddStream(s *StreamEventStream) {
	c.Collector.AddStream(s.Stream)
}

func (c *StreamEventCollector) Get() StreamEvent {
	return c.Collector.Get().(StreamEvent)
}

func (c *StreamEventCollector) Preview() StreamEvent {
	return c.Collector.Preview().(StreamEvent)
}

type StreamEventObservable struct {
	*eventual2go.Observable
}

func NewStreamEventObservable (value StreamEvent) (o *StreamEventObservable) {
	return &StreamEventObservable{eventual2go.NewObservable(value)}
}

func (o *StreamEventObservable) Value() StreamEvent {
	return o.Observable.Value().(StreamEvent)
}

func (o *StreamEventObservable) Change(value StreamEvent) {
	o.Observable.Change(value)
}

func (o *StreamEventObservable) OnChange(s StreamEventSubscriber) (cancel *eventual2go.Completer) {
	return o.Observable.OnChange(s.toSubscriber())
}

func (o *StreamEventObservable) Stream() (*StreamEventStream) {
	return &StreamEventStream{o.Observable.Stream()}
}


func (o *StreamEventObservable) AsChan() (c chan StreamEvent, cancel *eventual2go.Completer) {
	return o.Stream().AsChan()
}

func (o *StreamEventObservable) NextChange() (f *StreamEventFuture) {
	return o.Stream().First()
}
