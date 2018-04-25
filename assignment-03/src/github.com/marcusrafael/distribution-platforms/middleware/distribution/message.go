package distribution

type RequestBody struct {
    Parameters map[string]string
}

type ReplyBody struct {
    Result map[string]string
}

type ReplyHeader struct {
    ServiceContext string
    RequestId      int
    ReplyStatus    int
}

type RequestHeader struct {
    Context          string
    RequestId        int
    ResponseExpected bool
    ObjectKey        int
    Operation        string
}

type MessageBody struct {
    RequestHeader RequestHeader
    RequestBody   RequestBody
    ReplyHeader   ReplyHeader
    ReplyBody     ReplyBody
}

type MessageHeader struct {
    Magic       string
    Version     int
    ByteOrder   bool
    MessageType int
    MessageSize int
}

type Message struct {
    Header MessageHeader
    Body   MessageBody
}
