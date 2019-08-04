package limiter

type counter struct { }
func newCounter(rule *Rule) *counter {
	return &counter{}
}

func (counter *counter) allow() bool {
	return false
}

