package distribution


type Invocation struct {
    Host string
    Port string
    Operation string
    Parameters map[string]string
}
