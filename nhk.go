package nhk

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

var ServiceName = map[string]string{
	"g1":       "ＮＨＫ総合１",
	"g2":       "ＮＨＫ総合２",
	"e1":       "ＮＨＫＥテレ１",
	"e2":       "ＮＨＫＥテレ２",
	"e3":       "ＮＨＫＥテレ３",
	"e4":       "ＮＨＫワンセグ２",
	"s1":       "ＮＨＫＢＳ１",
	"s2":       "ＮＨＫＢＳ１(１０２ｃｈ)",
	"s3":       "ＮＨＫＢＳプレミアム",
	"s4":       "ＮＨＫＢＳプレミアム(１０４ｃｈ)",
	"r1":       "ＮＨＫラジオ第1",
	"r2":       "ＮＨＫラジオ第2",
	"r3":       "ＮＨＫＦＭ",
	"n1":       "ＮＨＫネットラジオ第1",
	"n2":       "ＮＨＫネットラジオ第2",
	"n3":       "ＮＨＫネットラジオＦＭ",
	"tv":       "テレビ全て",
	"radio":    "ラジオ全て",
	"netradio": "ネットラジオ全て",
}

var AreaName = map[string]string{
	"010": "札幌",
	"011": "函館",
	"012": "旭川",
	"013": "帯広",
	"014": "釧路",
	"015": "北見",
	"016": "室蘭",
	"020": "青森",
	"030": "盛岡",
	"040": "仙台",
	"050": "秋田",
	"060": "山形",
	"070": "福島",
	"080": "水戸",
	"090": "宇都宮",
	"100": "前橋",
	"110": "さいたま",
	"120": "千葉",
	"130": "東京",
	"140": "横浜",
	"150": "新潟",
	"160": "富山",
	"170": "金沢",
	"180": "福井",
	"190": "甲府",
	"200": "長野",
	"210": "岐阜",
	"220": "静岡",
	"230": "名古屋",
	"240": "津",
	"250": "大津",
	"260": "京都",
	"270": "大阪",
	"280": "神戸",
	"290": "奈良",
	"300": "和歌山",
	"310": "鳥取",
	"320": "松江",
	"330": "岡山",
	"340": "広島",
	"350": "山口",
	"360": "徳島",
	"370": "高松",
	"380": "松山",
	"390": "高知",
	"400": "福岡",
	"401": "北九州",
	"410": "佐賀",
	"420": "長崎",
	"430": "熊本",
	"440": "大分",
	"450": "宮崎",
	"460": "鹿児島",
	"470": "沖縄",
}

type Client struct {
	apikey string
}

type NowOnAirInfo struct {
	Following Program `json:"following"`
	Present   Program `json:"present"`
	Previous  Program `json:"previous"`
}

type Program struct {
	Id        string `json:"id"`
	EventId   string `json:"event_id"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Area      struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"area"`
	Service struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		LogoS struct {
			Url    string `json:"url"`
			Width  string `json:"width"`
			Height string `json:"height"`
		} `json:"logo_s"`
		LogoM struct {
			Url    string `json:"url"`
			Width  string `json:"width"`
			Height string `json:"height"`
		} `json:"logo_m"`
		LogoL struct {
			Url    string `json:"url"`
			Width  string `json:"width"`
			Height string `json:"height"`
		} `json:"logo_l"`
	} `json:"service"`
	Title    string   `json:"title"`
	Subtitle string   `json:"subtitle"`
	Genres   []string `json:"genres"`
}

func NewClient(apikey string) *Client {
	return &Client{apikey}
}

func (c *Client) ProgramList(area, service string, date *time.Time) ([]Program, error) {
	var d string
	if date == nil {
		d = time.Now().Format("2006-01-02")
	} else {
		d = date.Format("2006-01-02")
	}
	u := fmt.Sprintf("http://api.nhk.or.jp/v1/pg/list/%s/%s/%s.json?key=%s", area, service, d, url.QueryEscape(c.apikey))
	res, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, errors.New(res.Status)
	}
	pl := make(map[string]map[string][]Program)
	err = json.NewDecoder(res.Body).Decode(&pl)
	if err != nil {
		return nil, err
	}
	return pl["list"][service], nil
}

func (c *Client) ProgramGenre(area, service string, genreLevel1, genreLevel2 int, date *time.Time) ([]Program, error) {
	var d string
	if date == nil {
		d = time.Now().Format("2006-01-02")
	} else {
		d = date.Format("2006-01-02")
	}
	u := fmt.Sprintf("http://api.nhk.or.jp/v1/pg/genre/%s/%s/%02d%02d/%s.json?key=%s", area, service, genreLevel1, genreLevel2, d, url.QueryEscape(c.apikey))
	res, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, errors.New(res.Status)
	}
	pl := make(map[string]map[string][]Program)
	err = json.NewDecoder(res.Body).Decode(&pl)
	if err != nil {
		return nil, err
	}
	return pl["list"][service], nil
}

func (c *Client) ProgramInfo(area, service, id string) (*Program, error) {
	u := fmt.Sprintf("http://api.nhk.or.jp/v1/pg/info/%s/%s/%s.json?key=%s", area, service, id, url.QueryEscape(c.apikey))
	res, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, errors.New(res.Status)
	}
	pl := make(map[string]map[string][]*Program)
	err = json.NewDecoder(res.Body).Decode(&pl)
	if err != nil {
		return nil, err
	}
	return pl["list"][service][0], nil
}

func (c *Client) NowOnAir(area, service string) (*NowOnAirInfo, error) {
	u := fmt.Sprintf("http://api.nhk.or.jp/v1/pg/now/%s/%s.json?key=%s", area, service, url.QueryEscape(c.apikey))
	res, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, errors.New(res.Status)
	}
	pl := make(map[string]map[string]*NowOnAirInfo)
	err = json.NewDecoder(res.Body).Decode(&pl)
	if err != nil {
		return nil, err
	}
	return pl["nowonair_list"][service], nil
}
