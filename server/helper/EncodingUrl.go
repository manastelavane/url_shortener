package helper

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func ConvertToAlphabetic(num int64) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	var shuffledMappingString string = os.Getenv("SHUFFLED_MAPPING_STRING")
	if num < 0 {
		return ""
	}
	result := ""
	if num == 0 {
		return "a"
	}
	for num > 0 {
		remainder := num % 62
		result = string(shuffledMappingString[remainder]) + result
		num /= 62
	}
	return result
}

// First 3 digit number at 3844
// First 9 digit number at 218340105584895
// Total numbers in this range: 218340110000000=>218,340,110,000,000 => 218 trillion

// We need nearly 100M links per month.
// So, we will need 100M*12=1.2e9
// We need that each range should exhaust per year.
// So, 218340110000000/1.2e9 = 181950.092 ranges => 181,950 ranges

// Each range size => 1.2e9