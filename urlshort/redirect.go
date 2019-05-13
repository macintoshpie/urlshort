package urlshort

type Redirect struct {
	Src  string `json:"src"`
	Dest string `json:"dest"`
	Ttl  int    `json:"ttl"`
}
