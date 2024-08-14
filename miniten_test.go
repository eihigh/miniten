package miniten

import "testing"

func TestHitTestRects(t *testing.T) {
	tests := []struct {
		ax, ay, aw, ah int
		bx, by, bw, bh int
		want           bool
	}{
		{0, 0, 10, 10, 0, 0, 10, 10, true},
		{0, 0, 10, 10, 9, 9, 10, 10, true},
		{0, 0, 10, 10, 10, 10, 10, 10, false},
		{0, 0, 10, 10, 11, 11, 10, 10, false},
		{10, 10, 10, 10, 0, 0, 10, 10, false},
		{10, 10, 10, 10, 5, 5, 10, 10, true},
		{0, 10, 10, 10, 10, 10, 10, 10, false},
		{10, 0, 10, 10, 10, 10, 10, 10, false},
		{10, 0, 10, 10, 20, 0, 10, 10, false},
	}

	for _, tt := range tests {
		got := HitTestRects(tt.ax, tt.ay, tt.aw, tt.ah, tt.bx, tt.by, tt.bw, tt.bh)
		if tt.want != got {
			t.Errorf("HitTestRects(%d, %d, %d, %d, %d, %d, %d, %d) = %v; want %v", tt.ax, tt.ay, tt.aw, tt.ah, tt.bx, tt.by, tt.bw, tt.bh, got, tt.want)
		}
	}
}
