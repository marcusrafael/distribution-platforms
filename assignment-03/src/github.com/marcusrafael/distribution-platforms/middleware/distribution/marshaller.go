package distribution

import "encoding/json"


type Marshaller struct {}


func (* Marshaller) Marshal(message Message) ([]byte, error){

  return json.Marshal(message)

}

func (* Marshaller) Unmarshal(data []byte, message *Message) error{

  err := json.Unmarshal(data, message)
  if(err != nil) {
    return err
  }
  return nil
}
