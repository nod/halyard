package halyard


type ConfigState struct {
  DbUri string
  HttpListenUri string
}


var cfg *ConfigState


func GetConfig() (*ConfigState) {
  if cfg == nil {
    cfg = &ConfigState{}
  }
  return cfg
}


