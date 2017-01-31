
/*
 * generated by event_generator
 *
 * DO NOT EDIT
 */

package message

import "github.com/joernweissenborn/eventual2go"



type ResultCompleter struct {
	*eventual2go.Completer
}

func NewResultCompleter() *ResultCompleter {
	return &ResultCompleter{eventual2go.NewCompleter()}
}

func (c *ResultCompleter) Complete(d *Result) {
	c.Completer.Complete(d)
}

func (c *ResultCompleter) Future() *ResultFuture {
	return &ResultFuture{c.Completer.Future()}
}

type ResultFuture struct {
	*eventual2go.Future
}

func (f *ResultFuture) Result() *Result {
	return f.Future.Result().(*Result)
}

type ResultCompletionHandler func(*Result) *Result

func (ch ResultCompletionHandler) toCompletionHandler() eventual2go.CompletionHandler {
	return func(d eventual2go.Data) eventual2go.Data {
		return ch(d.(*Result))
	}
}

func (f *ResultFuture) Then(ch ResultCompletionHandler) *ResultFuture {
	return &ResultFuture{f.Future.Then(ch.toCompletionHandler())}
}

func (f *ResultFuture) AsChan() chan *Result {
	c := make(chan *Result, 1)
	cmpl := func(d chan *Result) ResultCompletionHandler {
		return func(e *Result) *Result {
			d <- e
			close(d)
			return e
		}
	}
	ecmpl := func(d chan *Result) eventual2go.ErrorHandler {
		return func(error) (eventual2go.Data, error) {
			close(d)
			return nil, nil
		}
	}
	f.Then(cmpl(c))
	f.Err(ecmpl(c))
	return c
}

type ResultStreamController struct {
	*eventual2go.StreamController
}

func NewResultStreamController() *ResultStreamController {
	return &ResultStreamController{eventual2go.NewStreamController()}
}

func (sc *ResultStreamController) Add(d *Result) {
	sc.StreamController.Add(d)
}

func (sc *ResultStreamController) Join(s *ResultStream) {
	sc.StreamController.Join(s.Stream)
}

func (sc *ResultStreamController) JoinFuture(f *ResultFuture) {
	sc.StreamController.JoinFuture(f.Future)
}

func (sc *ResultStreamController) Stream() *ResultStream {
	return &ResultStream{sc.StreamController.Stream()}
}

type ResultStream struct {
	*eventual2go.Stream
}

type ResultSubscriber func(*Result)

func (l ResultSubscriber) toSubscriber() eventual2go.Subscriber {
	return func(d eventual2go.Data) { l(d.(*Result)) }
}

func (s *ResultStream) Listen(ss ResultSubscriber) *eventual2go.Completer {
	return s.Stream.Listen(ss.toSubscriber())
}

type ResultFilter func(*Result) bool

func (f ResultFilter) toFilter() eventual2go.Filter {
	return func(d eventual2go.Data) bool { return f(d.(*Result)) }
}

func toResultFilterArray(f ...ResultFilter) (filter []eventual2go.Filter){

	filter = make([]eventual2go.Filter, len(f))
	for i, el := range f {
		filter[i] = el.toFilter()
	}
	return
}

func (s *ResultStream) Where(f ...ResultFilter) *ResultStream {
	return &ResultStream{s.Stream.Where(toResultFilterArray(f...)...)}
}

func (s *ResultStream) WhereNot(f ...ResultFilter) *ResultStream {
	return &ResultStream{s.Stream.WhereNot(toResultFilterArray(f...)...)}
}

func (s *ResultStream) Split(f ResultFilter) (*ResultStream, *ResultStream)  {
	return s.Where(f), s.WhereNot(f)
}

func (s *ResultStream) First() *ResultFuture {
	return &ResultFuture{s.Stream.First()}
}

func (s *ResultStream) FirstWhere(f... ResultFilter) *ResultFuture {
	return &ResultFuture{s.Stream.FirstWhere(toResultFilterArray(f...)...)}
}

func (s *ResultStream) FirstWhereNot(f ...ResultFilter) *ResultFuture {
	return &ResultFuture{s.Stream.FirstWhereNot(toResultFilterArray(f...)...)}
}

func (s *ResultStream) AsChan() (c chan *Result, stop *eventual2go.Completer) {
	c = make(chan *Result)
	stop = s.Listen(pipeToResultChan(c))
	stop.Future().Then(closeResultChan(c))
	return
}

func pipeToResultChan(c chan *Result) ResultSubscriber {
	return func(d *Result) {
		c <- d
	}
}

func closeResultChan(c chan *Result) eventual2go.CompletionHandler {
	return func(d eventual2go.Data) eventual2go.Data {
		close(c)
		return nil
	}
}

type ResultCollector struct {
	*eventual2go.Collector
}

func NewResultCollector() *ResultCollector {
	return &ResultCollector{eventual2go.NewCollector()}
}

func (c *ResultCollector) Add(d *Result) {
	c.Collector.Add(d)
}

func (c *ResultCollector) AddFuture(f *ResultFuture) {
	c.Collector.Add(f.Future)
}

func (c *ResultCollector) AddStream(s *ResultStream) {
	c.Collector.AddStream(s.Stream)
}

func (c *ResultCollector) Get() *Result {
	return c.Collector.Get().(*Result)
}

func (c *ResultCollector) Preview() *Result {
	return c.Collector.Preview().(*Result)
}
