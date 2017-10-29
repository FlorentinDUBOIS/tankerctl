package core

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// XMLContainer deserializer
type XMLContainer struct {
	XMLName  xml.Name     `xml:"pdv_liste"`
	Stations []XMLStation `xml:"pdv"`
}

// ToSensision transform stations into sensision format
func (s *XMLContainer) ToSensision() ([]string, error) {
	sensisions := make([]string, 0)
	for _, station := range s.Stations {
		metrics, err := station.ToSensision()
		if err != nil {
			return nil, err
		}

		sensisions = append(sensisions, metrics...)
	}

	return sensisions, nil
}

// XMLStation deserializer
type XMLStation struct {
	XMLName    xml.Name   `xml:"pdv" json:"-"`
	ID         int64      `xml:"id,attr" json:"id"`
	Address    string     `xml:"adresse" json:"address"`
	City       string     `xml:"ville" json:"city"`
	Latitude   float64    `xml:"latitude,attr" json:"-"`
	Longitude  float64    `xml:"longitude,attr" json:"-"`
	PostalCode int32      `xml:"cp,attr" json:"postal_code"`
	Services   []string   `xml:"services>service" json:"services"`
	Prices     []XMLPrice `xml:"prix" json:"-"`
}

func (s *XMLStation) toSensision() (string, error) {
	serialized, err := json.Marshal(s)
	if err != nil {
		return "", err
	}

	station := string(serialized)
	station = strings.Replace(station, "\n", "", -1)

	t := time.Now()
	ts := strconv.FormatInt(t.UnixNano()/1000, 10)

	return fmt.Sprintf("%s/%f:%f/ od.station{id=%d} \"%s\"", ts, s.Latitude, s.Longitude, s.ID, strings.Replace(station, "\"", "\\\"", -1)), nil
}

// ToSensision transform the station and prices into sensision format
func (s *XMLStation) ToSensision() ([]string, error) {

	sensisions := make([]string, 0)
	station, err := s.toSensision()
	if err != nil {
		return nil, err
	}

	sensisions = append(sensisions, station)
	for _, price := range s.Prices {
		sensision, err := price.ToSensision(s.Latitude, s.Longitude)
		if err != nil {
			return nil, err
		}

		sensisions = append(sensisions, sensision)
	}

	return sensisions, nil
}

// XMLPrice deserializer
type XMLPrice struct {
	XMLName xml.Name `xml:"prix"`
	ID      int64    `xml:"id,attr"`
	Name    string   `xml:"nom,attr"`
	Update  string   `xml:"maj,attr"`
	Value   float64  `xml:"valeur,attr"`
}

// ToSensision transform the price into sensision format
func (s *XMLPrice) ToSensision(lat, long float64) (string, error) {
	t, err := s.ParseTime()
	if err != nil {
		return "", err
	}

	ts := strconv.FormatInt(t.UnixNano()/1000, 10)

	return fmt.Sprintf("%s/%f:%f/ od.station.price{id=%d} %f", ts, lat, long, s.ID, s.Value), nil
}

// ParseTime parse the update field of price
func (s *XMLPrice) ParseTime() (time.Time, error) {

	location, _ := time.LoadLocation("Local")
	return time.ParseInLocation("2006-01-02 15:04:05", s.Update, location)
}
