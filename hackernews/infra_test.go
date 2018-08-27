package hackernews

import (
	"reflect"
	"testing"
)

func Test_hnAPI_GetFeed(t *testing.T) {
	type fields struct {
		b string
	}
	type args struct {
		n string
		p int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// Add test cases.
		{"nil", fields{}, args{}, nil, true},
		//{"ok", fields{defaultBaseUrl}, args{"news", 1}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &hnAPI{
				b: tt.fields.b,
			}
			got, err := h.GetFeed(tt.args.n, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("hnAPI.GetFeed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hnAPI.GetFeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hnAPI_GetFeed_One(t *testing.T) {
	h := NewHackerNewsAPI("")
	got, err := h.GetFeed("news", 1)
	if err != nil {
		t.Errorf("hnAPI.GetFeed() error = %v", err)
		return
	}
	var vt string
	for i, v := range got.([]feed) {
		if &v == nil {
			t.Errorf("hnAPI.GetFeed() returns contains N/A item, index %d", i)
			return
		}
		if vt == v.Title {
			t.Errorf("hnAPI.GetFeed() returns contains same items, index: %d, title: %q", i, vt)
			return
		}
		vt = v.Title
	}
}

func Test_hnAPI_GetItem(t *testing.T) {
	type fields struct {
		b string
	}
	type args struct {
		i string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// Add test cases.
		{"nil", fields{}, args{}, nil, true},
		//{"ok", fields{defaultBaseUrl}, args{"id"}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &hnAPI{
				b: tt.fields.b,
			}
			got, err := h.GetItem(tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("hnAPI.GetItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hnAPI.GetItem() = %v, want %v", got, tt.want)
			}
		})
	}
}
