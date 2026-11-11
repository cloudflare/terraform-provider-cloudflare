package observatory_scheduled_test

import (
	"testing"
)

func TestParseScheduleImportID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name       string
		id         string
		wantZone   string
		wantURL    string
		wantRegion string
		wantErr    bool
	}{
		{
			name:     "two-segment: zone/url",
			id:       "abc123/example.com",
			wantZone: "abc123",
			wantURL:  "example.com",
		},
		{
			name:       "three-segment: zone/url/region",
			id:         "abc123/example.com/europe-west4",
			wantZone:   "abc123",
			wantURL:    "example.com",
			wantRegion: "europe-west4",
		},
		{
			name:     "url with encoded slash: zone/example.com%2Fpath",
			id:       "abc123/example.com%2Fpath",
			wantZone: "abc123",
			wantURL:  "example.com/path",
		},
		{
			// When the URL contains an encoded slash (%2F) and a region is also
			// present, the last literal slash separates the region. The URL
			// segment (example.com%2Fpath) is then decoded to example.com/path.
			name:       "url with encoded slash and region: zone/example.com%2Fpath/us-east1",
			id:         "abc123/example.com%2Fpath/us-east1",
			wantZone:   "abc123",
			wantURL:    "example.com/path",
			wantRegion: "us-east1",
		},
		{
			// Region segment itself is percent-encoded (hyphen as %2D).
			// PathUnescape must decode it before passing to the API.
			name:       "percent-encoded region segment: us%2Deast1 -> us-east1",
			id:         "abc123/example.com/us%2Deast1",
			wantZone:   "abc123",
			wantURL:    "example.com",
			wantRegion: "us-east1",
		},
		{
			name:    "missing slash: no zone separator",
			id:      "abc123",
			wantErr: true,
		},
		{
			name:    "empty zone_id",
			id:      "/example.com",
			wantErr: true,
		},
		{
			name:    "empty string",
			id:      "",
			wantErr: true,
		},
		{
			// Region is passed through as-is to the API; no local validation.
			name:       "three-segment with us-central1 region",
			id:         "abc123/example.com/us-central1",
			wantZone:   "abc123",
			wantURL:    "example.com",
			wantRegion: "us-central1",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			zone, rawURL, region, err := parseScheduleImportID(tc.id)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("parseScheduleImportID(%q) = (%q, %q, %q, nil); want error", tc.id, zone, rawURL, region)
				}
				return
			}
			if err != nil {
				t.Fatalf("parseScheduleImportID(%q) unexpected error: %v", tc.id, err)
			}
			if zone != tc.wantZone {
				t.Errorf("zone: got %q, want %q", zone, tc.wantZone)
			}
			if rawURL != tc.wantURL {
				t.Errorf("url: got %q, want %q", rawURL, tc.wantURL)
			}
			if region != tc.wantRegion {
				t.Errorf("region: got %q, want %q", region, tc.wantRegion)
			}
		})
	}
}
