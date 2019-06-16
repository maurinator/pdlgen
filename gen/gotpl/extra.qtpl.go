// Code generated by qtc from "extra.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line gen/gotpl/extra.qtpl:1
package gotpl

//line gen/gotpl/extra.qtpl:1
import (
	"github.com/chromedp/cdproto-gen/gen/genutil"
	"github.com/chromedp/cdproto-gen/pdl"
)

// ExtraTimestampTemplate is a special template for the Timestamp type that
// defines its JSON unmarshaling.

//line gen/gotpl/extra.qtpl:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line gen/gotpl/extra.qtpl:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line gen/gotpl/extra.qtpl:8
func StreamExtraTimestampTemplate(qw422016 *qt422016.Writer, t *pdl.Type, d *pdl.Domain) {
//line gen/gotpl/extra.qtpl:9
	typ := t.Name
	monotonic := t.TimestampType == pdl.TimestampTypeMonotonic
	timeRes := "time.Millisecond"
	if t.TimestampType != pdl.TimestampTypeMillisecond {
		timeRes = "time.Second"
	}

//line gen/gotpl/extra.qtpl:15
	qw422016.N().S(`
`)
//line gen/gotpl/extra.qtpl:16
	if monotonic {
//line gen/gotpl/extra.qtpl:16
		qw422016.N().S(`
// `)
//line gen/gotpl/extra.qtpl:17
		qw422016.N().S(typ)
//line gen/gotpl/extra.qtpl:17
		qw422016.N().S(`Epoch is the `)
//line gen/gotpl/extra.qtpl:17
		qw422016.N().S(typ)
//line gen/gotpl/extra.qtpl:17
		qw422016.N().S(` time epoch.
var `)
//line gen/gotpl/extra.qtpl:18
		qw422016.N().S(typ)
//line gen/gotpl/extra.qtpl:18
		qw422016.N().S(`Epoch *time.Time

func init() {
	// initialize epoch
	bt := sysutil.BootTime()
	`)
//line gen/gotpl/extra.qtpl:23
		qw422016.N().S(typ)
//line gen/gotpl/extra.qtpl:23
		qw422016.N().S(`Epoch = &bt
}
`)
//line gen/gotpl/extra.qtpl:25
	}
//line gen/gotpl/extra.qtpl:25
	qw422016.N().S(`

// MarshalEasyJSON satisfies easyjson.Marshaler.
func (t `)
//line gen/gotpl/extra.qtpl:28
	qw422016.N().S(typ)
//line gen/gotpl/extra.qtpl:28
	qw422016.N().S(`) MarshalEasyJSON(out *jwriter.Writer) {
	v := `)
//line gen/gotpl/extra.qtpl:29
	if monotonic {
//line gen/gotpl/extra.qtpl:29
		qw422016.N().S(`float64(time.Time(t).Sub(*`)
//line gen/gotpl/extra.qtpl:29
		qw422016.N().S(typ)
//line gen/gotpl/extra.qtpl:29
		qw422016.N().S(`Epoch))/float64(time.Second)`)
//line gen/gotpl/extra.qtpl:29
	} else {
//line gen/gotpl/extra.qtpl:29
		qw422016.N().S(`float64(time.Time(t).UnixNano()/int64(`)
//line gen/gotpl/extra.qtpl:29
		qw422016.N().S(timeRes)
//line gen/gotpl/extra.qtpl:29
		qw422016.N().S(`))`)
//line gen/gotpl/extra.qtpl:29
	}
//line gen/gotpl/extra.qtpl:29
	qw422016.N().S(`

	out.Buffer.EnsureSpace(20)
	out.Buffer.Buf = strconv.AppendFloat(out.Buffer.Buf, v, 'f', -1, 64)
}

// MarshalJSON satisfies json.Marshaler.
func (t `)
//line gen/gotpl/extra.qtpl:36
	qw422016.N().S(typ)
//line gen/gotpl/extra.qtpl:36
	qw422016.N().S(`) MarshalJSON() ([]byte, error) {
	return easyjson.Marshal(t)
}

// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *`)
//line gen/gotpl/extra.qtpl:41
	qw422016.N().S(typ)
//line gen/gotpl/extra.qtpl:41
	qw422016.N().S(`) UnmarshalEasyJSON(in *jlexer.Lexer) {`)
//line gen/gotpl/extra.qtpl:41
	if monotonic {
//line gen/gotpl/extra.qtpl:41
		qw422016.N().S(`
	*t = `)
//line gen/gotpl/extra.qtpl:42
		qw422016.N().S(typ)
//line gen/gotpl/extra.qtpl:42
		qw422016.N().S(`(`)
//line gen/gotpl/extra.qtpl:42
		qw422016.N().S(typ)
//line gen/gotpl/extra.qtpl:42
		qw422016.N().S(`Epoch.Add(time.Duration(in.Float64()*float64(time.Second))))`)
//line gen/gotpl/extra.qtpl:42
	} else {
//line gen/gotpl/extra.qtpl:42
		qw422016.N().S(`
	*t = `)
//line gen/gotpl/extra.qtpl:43
		qw422016.N().S(typ)
//line gen/gotpl/extra.qtpl:43
		qw422016.N().S(`(time.Unix(0, int64(in.Float64()*float64(`)
//line gen/gotpl/extra.qtpl:43
		qw422016.N().S(timeRes)
//line gen/gotpl/extra.qtpl:43
		qw422016.N().S(`))))`)
//line gen/gotpl/extra.qtpl:43
	}
//line gen/gotpl/extra.qtpl:43
	qw422016.N().S(`
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *`)
//line gen/gotpl/extra.qtpl:47
	qw422016.N().S(typ)
//line gen/gotpl/extra.qtpl:47
	qw422016.N().S(`) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}
`)
//line gen/gotpl/extra.qtpl:50
}

//line gen/gotpl/extra.qtpl:50
func WriteExtraTimestampTemplate(qq422016 qtio422016.Writer, t *pdl.Type, d *pdl.Domain) {
//line gen/gotpl/extra.qtpl:50
	qw422016 := qt422016.AcquireWriter(qq422016)
//line gen/gotpl/extra.qtpl:50
	StreamExtraTimestampTemplate(qw422016, t, d)
//line gen/gotpl/extra.qtpl:50
	qt422016.ReleaseWriter(qw422016)
//line gen/gotpl/extra.qtpl:50
}

//line gen/gotpl/extra.qtpl:50
func ExtraTimestampTemplate(t *pdl.Type, d *pdl.Domain) string {
//line gen/gotpl/extra.qtpl:50
	qb422016 := qt422016.AcquireByteBuffer()
//line gen/gotpl/extra.qtpl:50
	WriteExtraTimestampTemplate(qb422016, t, d)
//line gen/gotpl/extra.qtpl:50
	qs422016 := string(qb422016.B)
//line gen/gotpl/extra.qtpl:50
	qt422016.ReleaseByteBuffer(qb422016)
//line gen/gotpl/extra.qtpl:50
	return qs422016
//line gen/gotpl/extra.qtpl:50
}

// ExtraFrameTemplate is a special template for the Page.Frame type, adding FrameState.

//line gen/gotpl/extra.qtpl:53
func StreamExtraFrameTemplate(qw422016 *qt422016.Writer) {
//line gen/gotpl/extra.qtpl:53
	qw422016.N().S(`
// FrameState is the state of a Frame.
type FrameState uint16

// FrameState enum values.
const (
    FrameDOMContentEventFired FrameState = 1 << (15 - iota)
    FrameLoadEventFired
    FrameAttached
    FrameNavigated
    FrameLoading
    FrameScheduledNavigation
)

// frameStateNames are the names of the frame states.
var frameStateNames = map[FrameState]string{
    FrameDOMContentEventFired: "DOMContentEventFired",
    FrameLoadEventFired:       "LoadEventFired",
    FrameAttached:             "Attached",
    FrameNavigated:            "Navigated",
	FrameLoading:			   "Loading",
    FrameScheduledNavigation:  "ScheduledNavigation",
}

// String satisfies stringer interface.
func (fs FrameState) String() string {
    var s []string
    for k, v := range frameStateNames {
        if fs&k != 0 {
            s = append(s, v)
        }
    }
    return "[" + strings.Join(s, " ") + "]"
}

// EmptyFrameID is the "non-existent" frame id.
const EmptyFrameID = FrameID("")
`)
//line gen/gotpl/extra.qtpl:90
}

//line gen/gotpl/extra.qtpl:90
func WriteExtraFrameTemplate(qq422016 qtio422016.Writer) {
//line gen/gotpl/extra.qtpl:90
	qw422016 := qt422016.AcquireWriter(qq422016)
//line gen/gotpl/extra.qtpl:90
	StreamExtraFrameTemplate(qw422016)
//line gen/gotpl/extra.qtpl:90
	qt422016.ReleaseWriter(qw422016)
//line gen/gotpl/extra.qtpl:90
}

//line gen/gotpl/extra.qtpl:90
func ExtraFrameTemplate() string {
//line gen/gotpl/extra.qtpl:90
	qb422016 := qt422016.AcquireByteBuffer()
//line gen/gotpl/extra.qtpl:90
	WriteExtraFrameTemplate(qb422016)
//line gen/gotpl/extra.qtpl:90
	qs422016 := string(qb422016.B)
//line gen/gotpl/extra.qtpl:90
	qt422016.ReleaseByteBuffer(qb422016)
//line gen/gotpl/extra.qtpl:90
	return qs422016
//line gen/gotpl/extra.qtpl:90
}

// ExtraNodeTemplate is a special template for the DOM.Node type, adding NodeState.

//line gen/gotpl/extra.qtpl:93
func StreamExtraNodeTemplate(qw422016 *qt422016.Writer) {
//line gen/gotpl/extra.qtpl:93
	qw422016.N().S(`
// AttributeValue returns the named attribute for the node.
func (n *Node) AttributeValue(name string) string {
	n.RLock()
	defer n.RUnlock()

	for i := 0; i < len(n.Attributes); i+=2 {
		if n.Attributes[i] == name  {
			return n.Attributes[i+1]
		}
	}

	return ""
}

// xpath builds the xpath string.
func (n *Node) xpath(stopAtDocument, stopAtID bool) string {
	n.RLock()
	defer n.RUnlock()

	p, pos, id := "", "", n.AttributeValue("id")
	switch {
	case n.Parent == nil:
		return n.LocalName

	case stopAtDocument && n.NodeType == NodeTypeDocument:
		return ""

	case stopAtID && id != "":
		p = "/"
		pos = `)
//line gen/gotpl/extra.qtpl:93
	qw422016.N().S("`")
//line gen/gotpl/extra.qtpl:93
	qw422016.N().S(`[@id='`)
//line gen/gotpl/extra.qtpl:93
	qw422016.N().S("`")
//line gen/gotpl/extra.qtpl:93
	qw422016.N().S(`+id+`)
//line gen/gotpl/extra.qtpl:93
	qw422016.N().S("`")
//line gen/gotpl/extra.qtpl:93
	qw422016.N().S(`']`)
//line gen/gotpl/extra.qtpl:93
	qw422016.N().S("`")
//line gen/gotpl/extra.qtpl:93
	qw422016.N().S(`

	case n.Parent != nil:
		var i int
		var found bool

		n.Parent.RLock()
		for j := 0; j < len(n.Parent.Children); j++ {
			if n.Parent.Children[j].LocalName == n.LocalName {
				i++
			}
			if n.Parent.Children[j].NodeID == n.NodeID {
				found = true
				break
			}
		}
		n.Parent.RUnlock()

		if found {
			pos = "["+strconv.Itoa(i)+"]"
		}

		p = n.Parent.xpath(stopAtDocument, stopAtID)
	}

	localName := n.LocalName
	if n.IsSVG {
		localName = `)
//line gen/gotpl/extra.qtpl:93
	qw422016.N().S("`")
//line gen/gotpl/extra.qtpl:93
	qw422016.N().S(`*[local-name()='`)
//line gen/gotpl/extra.qtpl:93
	qw422016.N().S("`")
//line gen/gotpl/extra.qtpl:93
	qw422016.N().S(` + localName + `)
//line gen/gotpl/extra.qtpl:93
	qw422016.N().S("`")
//line gen/gotpl/extra.qtpl:93
	qw422016.N().S(`']`)
//line gen/gotpl/extra.qtpl:93
	qw422016.N().S("`")
//line gen/gotpl/extra.qtpl:93
	qw422016.N().S(`
	}
	return  p + "/" + localName + pos
}

// PartialXPathByID returns the partial XPath for the node, stopping at the
// first parent with an id attribute or at nearest parent document node.
func (n *Node) PartialXPathByID() string {
	return n.xpath(true, true)
}

// PartialXPath returns the partial XPath for the node, stopping at the nearest
// parent document node.
func (n *Node) PartialXPath() string {
	return n.xpath(true, false)
}

// FullXPathByID returns the full XPath for the node, stopping at the top most
// document root or at the closest parent node with an id attribute.
func (n *Node) FullXPathByID() string {
	return n.xpath(false, true)
}

// FullXPath returns the full XPath for the node, stopping only at the top most
// document root.
func (n *Node) FullXPath() string {
	return n.xpath(false, false)
}

// Dump builds a printable string representation of the node and its children.
func (n *Node) Dump(prefix, indent string, nodeIDs bool) string {
	if n == nil {
		return prefix + "<nil>"
	}

	n.RLock()
	defer n.RUnlock()

	s := n.LocalName
	if s == "" {
		s = n.NodeName
	}

	for i := 0; i < len(n.Attributes); i += 2 {
		if strings.ToLower(n.Attributes[i]) == "id" {
			s += "#" + n.Attributes[i+1]
			break
		}
	}

	if n.NodeType != NodeTypeElement && n.NodeType != NodeTypeText {
		s += fmt.Sprintf(" <%s>", n.NodeType)
	}

	if n.NodeType == NodeTypeText {
		v := n.NodeValue
		if len(v) > 15 {
			v = v[:15] + "..."
		}
		s += fmt.Sprintf(" %q", v)
	}

	if n.NodeType == NodeTypeElement && len(n.Attributes) > 0 {
		attrs := ""
		for i := 0; i < len(n.Attributes); i += 2 {
			if strings.ToLower(n.Attributes[i]) == "id" {
				continue
			}
			if attrs != "" {
				attrs += " "
			}
			attrs += fmt.Sprintf("%s=%q", n.Attributes[i], n.Attributes[i+1])
		}
		if attrs != "" {
			s += " [" + attrs + "]"
		}
	}

	if nodeIDs {
		s += fmt.Sprintf(" (%d)", n.NodeID)
	}

	for i := 0; i < len(n.Children); i++ {
		s += "\n" + n.Children[i].Dump(prefix+indent, indent, nodeIDs)
	}

	return prefix + s
}

// NodeState is the state of a DOM node.
type NodeState uint8

// NodeState enum values.
const (
    NodeReady NodeState = 1 << (7 - iota)
	NodeVisible
	NodeHighlighted
)

// nodeStateNames are the names of the node states.
var nodeStateNames = map[NodeState]string{
    NodeReady:		 "Ready",
    NodeVisible:     "Visible",
    NodeHighlighted: "Highlighted",
}

// String satisfies stringer interface.
func (ns NodeState) String() string {
    var s []string
    for k, v := range nodeStateNames {
        if ns&k != 0 {
            s = append(s, v)
        }
    }
    return "[" + strings.Join(s, " ") + "]"
}

// EmptyNodeID is the "non-existent" node id.
const EmptyNodeID = NodeID(0)
`)
//line gen/gotpl/extra.qtpl:269
}

//line gen/gotpl/extra.qtpl:269
func WriteExtraNodeTemplate(qq422016 qtio422016.Writer) {
//line gen/gotpl/extra.qtpl:269
	qw422016 := qt422016.AcquireWriter(qq422016)
//line gen/gotpl/extra.qtpl:269
	StreamExtraNodeTemplate(qw422016)
//line gen/gotpl/extra.qtpl:269
	qt422016.ReleaseWriter(qw422016)
//line gen/gotpl/extra.qtpl:269
}

//line gen/gotpl/extra.qtpl:269
func ExtraNodeTemplate() string {
//line gen/gotpl/extra.qtpl:269
	qb422016 := qt422016.AcquireByteBuffer()
//line gen/gotpl/extra.qtpl:269
	WriteExtraNodeTemplate(qb422016)
//line gen/gotpl/extra.qtpl:269
	qs422016 := string(qb422016.B)
//line gen/gotpl/extra.qtpl:269
	qt422016.ReleaseByteBuffer(qb422016)
//line gen/gotpl/extra.qtpl:269
	return qs422016
//line gen/gotpl/extra.qtpl:269
}

// ExtraFixStringUnmarshaler is a template that forces values to be parsed properly.

//line gen/gotpl/extra.qtpl:272
func StreamExtraFixStringUnmarshaler(qw422016 *qt422016.Writer, typ, parseFunc, extra string) {
//line gen/gotpl/extra.qtpl:272
	qw422016.N().S(`
// UnmarshalEasyJSON satisfies easyjson.Unmarshaler.
func (t *`)
//line gen/gotpl/extra.qtpl:274
	qw422016.N().S(typ)
//line gen/gotpl/extra.qtpl:274
	qw422016.N().S(`) UnmarshalEasyJSON(in *jlexer.Lexer) {
	buf := in.Raw()
	if l := len(buf); l > 2 && buf[0] == '"' && buf[l-1] == '"' {
		buf = buf[1:l-1]
	}
`)
//line gen/gotpl/extra.qtpl:279
	if parseFunc != "" {
//line gen/gotpl/extra.qtpl:279
		qw422016.N().S(`
	v, err := strconv.`)
//line gen/gotpl/extra.qtpl:280
		qw422016.N().S(parseFunc)
//line gen/gotpl/extra.qtpl:280
		qw422016.N().S(`(string(buf)`)
//line gen/gotpl/extra.qtpl:280
		qw422016.N().S(extra)
//line gen/gotpl/extra.qtpl:280
		qw422016.N().S(`)
	if err != nil {
		in.AddError(err)
	}
`)
//line gen/gotpl/extra.qtpl:284
	}
//line gen/gotpl/extra.qtpl:284
	qw422016.N().S(`
	*t = `)
//line gen/gotpl/extra.qtpl:285
	qw422016.N().S(typ)
//line gen/gotpl/extra.qtpl:285
	qw422016.N().S(`(`)
//line gen/gotpl/extra.qtpl:285
	if parseFunc != "" {
//line gen/gotpl/extra.qtpl:285
		qw422016.N().S(`v`)
//line gen/gotpl/extra.qtpl:285
	} else {
//line gen/gotpl/extra.qtpl:285
		qw422016.N().S(`buf`)
//line gen/gotpl/extra.qtpl:285
	}
//line gen/gotpl/extra.qtpl:285
	qw422016.N().S(`)
}

// UnmarshalJSON satisfies json.Unmarshaler.
func (t *`)
//line gen/gotpl/extra.qtpl:289
	qw422016.N().S(typ)
//line gen/gotpl/extra.qtpl:289
	qw422016.N().S(`) UnmarshalJSON(buf []byte) error {
	return easyjson.Unmarshal(buf, t)
}
`)
//line gen/gotpl/extra.qtpl:292
}

//line gen/gotpl/extra.qtpl:292
func WriteExtraFixStringUnmarshaler(qq422016 qtio422016.Writer, typ, parseFunc, extra string) {
//line gen/gotpl/extra.qtpl:292
	qw422016 := qt422016.AcquireWriter(qq422016)
//line gen/gotpl/extra.qtpl:292
	StreamExtraFixStringUnmarshaler(qw422016, typ, parseFunc, extra)
//line gen/gotpl/extra.qtpl:292
	qt422016.ReleaseWriter(qw422016)
//line gen/gotpl/extra.qtpl:292
}

//line gen/gotpl/extra.qtpl:292
func ExtraFixStringUnmarshaler(typ, parseFunc, extra string) string {
//line gen/gotpl/extra.qtpl:292
	qb422016 := qt422016.AcquireByteBuffer()
//line gen/gotpl/extra.qtpl:292
	WriteExtraFixStringUnmarshaler(qb422016, typ, parseFunc, extra)
//line gen/gotpl/extra.qtpl:292
	qs422016 := string(qb422016.B)
//line gen/gotpl/extra.qtpl:292
	qt422016.ReleaseByteBuffer(qb422016)
//line gen/gotpl/extra.qtpl:292
	return qs422016
//line gen/gotpl/extra.qtpl:292
}

// ExtraExecutorTemplate is the additional shared executor interface for all
// the domains.

//line gen/gotpl/extra.qtpl:296
func StreamExtraExecutorTemplate(qw422016 *qt422016.Writer) {
//line gen/gotpl/extra.qtpl:296
	qw422016.N().S(`
// Executor is the common interface for executing a command.
type Executor interface {
	// Execute executes the command.
	Execute(string, easyjson.Marshaler, easyjson.Unmarshaler) error
}

// contextKeyType is the context key type.
type contextKeyType int

// context keys.
const (
	ContextKey contextKeyType = iota
)

// FromContext returns the message executor for the context.
func FromContext(ctx context.Context) Executor {
	return ctx.Value(ContextKey).(Executor)
}

// Execute uses the context's message executor to send a command or event
// method marshaling the provided parameters, and unmarshaling to res.
func Execute(ctx context.Context, method string, params easyjson.Marshaler, res easyjson.Unmarshaler) error {
	if executor := ctx.Value(ContextKey); executor != nil {
		return executor.(Executor).Execute(method, params, res)
	}
	return ErrInvalidContext
}

// Error is a error.
type Error string

// Error values.
const (
	// ErrInvalidContext is the invalid context error.
	ErrInvalidContext Error = "invalid context"

	// ErrMsgMissingParamsOrResult is the msg missing params or result error.
	ErrMsgMissingParamsOrResult Error = "msg missing params or result"
)

// Error satisfies the error interface.
func (err Error) Error() string {
	return string(err)
}

// ErrUnknownCommandOrEvent is an unknown command or event error.
type ErrUnknownCommandOrEvent string

// Error satisfies the error interface.
func (err ErrUnknownCommandOrEvent) Error() string {
	return fmt.Sprintf("unknown command or event %q", string(err))
}

`)
//line gen/gotpl/extra.qtpl:350
}

//line gen/gotpl/extra.qtpl:350
func WriteExtraExecutorTemplate(qq422016 qtio422016.Writer) {
//line gen/gotpl/extra.qtpl:350
	qw422016 := qt422016.AcquireWriter(qq422016)
//line gen/gotpl/extra.qtpl:350
	StreamExtraExecutorTemplate(qw422016)
//line gen/gotpl/extra.qtpl:350
	qt422016.ReleaseWriter(qw422016)
//line gen/gotpl/extra.qtpl:350
}

//line gen/gotpl/extra.qtpl:350
func ExtraExecutorTemplate() string {
//line gen/gotpl/extra.qtpl:350
	qb422016 := qt422016.AcquireByteBuffer()
//line gen/gotpl/extra.qtpl:350
	WriteExtraExecutorTemplate(qb422016)
//line gen/gotpl/extra.qtpl:350
	qs422016 := string(qb422016.B)
//line gen/gotpl/extra.qtpl:350
	qt422016.ReleaseByteBuffer(qb422016)
//line gen/gotpl/extra.qtpl:350
	return qs422016
//line gen/gotpl/extra.qtpl:350
}

// ExtraMethodTypeTemplate generates the additional MethodType funcs and consts.

//line gen/gotpl/extra.qtpl:353
func StreamExtraMethodTypeTemplate(qw422016 *qt422016.Writer, domains []*pdl.Domain) {
//line gen/gotpl/extra.qtpl:353
	qw422016.N().S(`
// Domain returns the Chrome DevTools Protocol domain of the event or command.
func (t MethodType) Domain() string {
	return string(t[:strings.IndexByte(string(t), '.')])
}

// MethodType values.
const (`)
//line gen/gotpl/extra.qtpl:360
	for _, d := range domains {
//line gen/gotpl/extra.qtpl:360
		for _, c := range d.Commands {
//line gen/gotpl/extra.qtpl:360
			qw422016.N().S(`
	`)
//line gen/gotpl/extra.qtpl:361
			qw422016.N().S(CommandMethodType(c, d))
//line gen/gotpl/extra.qtpl:361
			qw422016.N().S(` = `)
//line gen/gotpl/extra.qtpl:361
			qw422016.N().S(genutil.PackageName(d))
//line gen/gotpl/extra.qtpl:361
			qw422016.N().S(`.`)
//line gen/gotpl/extra.qtpl:361
			qw422016.N().S(CommandMethodType(c, nil))
//line gen/gotpl/extra.qtpl:361
		}
//line gen/gotpl/extra.qtpl:361
		for _, e := range d.Events {
//line gen/gotpl/extra.qtpl:361
			qw422016.N().S(`
	`)
//line gen/gotpl/extra.qtpl:362
			qw422016.N().S(EventMethodType(e, d))
//line gen/gotpl/extra.qtpl:362
			qw422016.N().S(` = `)
//line gen/gotpl/extra.qtpl:362
			qw422016.N().Q(ProtoName(e, d))
//line gen/gotpl/extra.qtpl:362
		}
//line gen/gotpl/extra.qtpl:362
	}
//line gen/gotpl/extra.qtpl:362
	qw422016.N().S(`)
`)
//line gen/gotpl/extra.qtpl:363
}

//line gen/gotpl/extra.qtpl:363
func WriteExtraMethodTypeTemplate(qq422016 qtio422016.Writer, domains []*pdl.Domain) {
//line gen/gotpl/extra.qtpl:363
	qw422016 := qt422016.AcquireWriter(qq422016)
//line gen/gotpl/extra.qtpl:363
	StreamExtraMethodTypeTemplate(qw422016, domains)
//line gen/gotpl/extra.qtpl:363
	qt422016.ReleaseWriter(qw422016)
//line gen/gotpl/extra.qtpl:363
}

//line gen/gotpl/extra.qtpl:363
func ExtraMethodTypeTemplate(domains []*pdl.Domain) string {
//line gen/gotpl/extra.qtpl:363
	qb422016 := qt422016.AcquireByteBuffer()
//line gen/gotpl/extra.qtpl:363
	WriteExtraMethodTypeTemplate(qb422016, domains)
//line gen/gotpl/extra.qtpl:363
	qs422016 := string(qb422016.B)
//line gen/gotpl/extra.qtpl:363
	qt422016.ReleaseByteBuffer(qb422016)
//line gen/gotpl/extra.qtpl:363
	return qs422016
//line gen/gotpl/extra.qtpl:363
}

// ExtraMessageTemplate generates the additional Message funcs.

//line gen/gotpl/extra.qtpl:366
func StreamExtraMessageTemplate(qw422016 *qt422016.Writer, domains []*pdl.Domain) {
//line gen/gotpl/extra.qtpl:366
	qw422016.N().S(`
type empty struct{}
var emptyVal = &empty{}

// UnmarshalMessage unmarshals the message result or params.
func UnmarshalMessage(msg *Message) (interface{}, error) {
	var v easyjson.Unmarshaler
	switch msg.Method {`)
//line gen/gotpl/extra.qtpl:373
	for _, d := range domains {
//line gen/gotpl/extra.qtpl:373
		for _, c := range d.Commands {
//line gen/gotpl/extra.qtpl:373
			qw422016.N().S(`
	case `)
//line gen/gotpl/extra.qtpl:374
			qw422016.N().S(CommandMethodType(c, d))
//line gen/gotpl/extra.qtpl:374
			qw422016.N().S(`:`)
//line gen/gotpl/extra.qtpl:374
			if len(c.Returns) == 0 {
//line gen/gotpl/extra.qtpl:374
				qw422016.N().S(`
		return emptyVal, nil`)
//line gen/gotpl/extra.qtpl:375
			} else {
//line gen/gotpl/extra.qtpl:375
				qw422016.N().S(`
		v = new(`)
//line gen/gotpl/extra.qtpl:376
				qw422016.N().S(genutil.PackageName(d))
//line gen/gotpl/extra.qtpl:376
				qw422016.N().S(`.`)
//line gen/gotpl/extra.qtpl:376
				qw422016.N().S(CommandReturnsType(c))
//line gen/gotpl/extra.qtpl:376
				qw422016.N().S(`)`)
//line gen/gotpl/extra.qtpl:376
			}
//line gen/gotpl/extra.qtpl:376
			qw422016.N().S(`
	`)
//line gen/gotpl/extra.qtpl:377
		}
//line gen/gotpl/extra.qtpl:377
		for _, e := range d.Events {
//line gen/gotpl/extra.qtpl:377
			qw422016.N().S(`
	case `)
//line gen/gotpl/extra.qtpl:378
			qw422016.N().S(EventMethodType(e, d))
//line gen/gotpl/extra.qtpl:378
			qw422016.N().S(`:
		v = new(`)
//line gen/gotpl/extra.qtpl:379
			qw422016.N().S(genutil.PackageName(d))
//line gen/gotpl/extra.qtpl:379
			qw422016.N().S(`.`)
//line gen/gotpl/extra.qtpl:379
			qw422016.N().S(EventType(e))
//line gen/gotpl/extra.qtpl:379
			qw422016.N().S(`)
	`)
//line gen/gotpl/extra.qtpl:380
		}
//line gen/gotpl/extra.qtpl:380
	}
//line gen/gotpl/extra.qtpl:380
	qw422016.N().S(`
	default:
		return nil, cdp.ErrUnknownCommandOrEvent(msg.Method)
	}

	var buf easyjson.RawMessage
	switch {
	case msg.Params != nil:
		buf = msg.Params

	case msg.Result != nil:
		buf = msg.Result

	default:
		return nil, cdp.ErrMsgMissingParamsOrResult
	}

	err := easyjson.Unmarshal(buf, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}
`)
//line gen/gotpl/extra.qtpl:404
}

//line gen/gotpl/extra.qtpl:404
func WriteExtraMessageTemplate(qq422016 qtio422016.Writer, domains []*pdl.Domain) {
//line gen/gotpl/extra.qtpl:404
	qw422016 := qt422016.AcquireWriter(qq422016)
//line gen/gotpl/extra.qtpl:404
	StreamExtraMessageTemplate(qw422016, domains)
//line gen/gotpl/extra.qtpl:404
	qt422016.ReleaseWriter(qw422016)
//line gen/gotpl/extra.qtpl:404
}

//line gen/gotpl/extra.qtpl:404
func ExtraMessageTemplate(domains []*pdl.Domain) string {
//line gen/gotpl/extra.qtpl:404
	qb422016 := qt422016.AcquireByteBuffer()
//line gen/gotpl/extra.qtpl:404
	WriteExtraMessageTemplate(qb422016, domains)
//line gen/gotpl/extra.qtpl:404
	qs422016 := string(qb422016.B)
//line gen/gotpl/extra.qtpl:404
	qt422016.ReleaseByteBuffer(qb422016)
//line gen/gotpl/extra.qtpl:404
	return qs422016
//line gen/gotpl/extra.qtpl:404
}
