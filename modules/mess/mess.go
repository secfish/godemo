package mess

import (
	"fmt"
)

//  return message
func GetMessage() string {

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("GetMessage Welcome!")
	return message
}
