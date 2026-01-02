package accounts

import (
	"testing"

	"github.com/anuragcarret/djang-drf-go/drf/serializers"
)

func TestAccountSerializer_Validation(t *testing.T) {
	t.Run("validates account location length", func(t *testing.T) {
		account := &Account{}
		serializer := serializers.NewSerializer(account)

		data := map[string]interface{}{
			"user_id":  uint64(1),
			"location": "this_location_is_definitely_way_too_long_for_the_one_hundred_character_limit_defined_in_the_model_tags_so_it_should_fail",
		}

		if serializer.IsValid(data) {
			t.Error("Expected IsValid to be false for long location")
		}

		if _, ok := serializer.Errors()["location"]; !ok {
			t.Error("Expected error for location field")
		}
	})
}
