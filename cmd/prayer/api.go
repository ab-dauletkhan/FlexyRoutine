package prayer

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// getPrayerTimes fetches prayer times for a given city and year
func GetPrayerTimes(year int) {
	url := fmt.Sprintf(
		"https://api.muftyat.kz/prayer-times/%d/%s/%s",
		year,
		strconv.FormatFloat(Location.Lat, 'f', 6, 64),
		strconv.FormatFloat(Location.Lng, 'f', 6, 64),
	)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making HTTP request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Non-OK HTTP status: %s\n", resp.Status)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	fmt.Println(string(body))
}
