package handlers

import "testing"

func TestAddNum(T *testing.T) {
	if AddNum(2) != 4 {
		T.Error("Expected 2 + 2 equals four")
	}
}
