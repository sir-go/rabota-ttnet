package store

import (
	"database/sql/driver"
	"reflect"
	"testing"
)

func TestJSON_Equals(t *testing.T) {
	type args struct {
		j1 JSON
	}
	tests := []struct {
		name string
		j    JSON
		args args
		want bool
	}{
		{"nil", nil, args{nil}, true},
		{"empty", JSON{}, args{JSON{}}, true},
		{"eq", JSON{1, 4, 5, 3, 6}, args{JSON{1, 4, 5, 3, 6}}, true},
		{"non-eq", JSON{1, 4, 3, 6}, args{JSON{1, 4, 5, 3, 6}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.j.Equals(tt.args.j1); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSON_IsNull(t *testing.T) {
	tests := []struct {
		name string
		j    JSON
		want bool
	}{
		{"nil", nil, true},
		{"non-nil", JSON{}, true},
		{"null", JSON("null"), true},
		{"non-null", JSON("non-null"), false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.j.IsNull(); got != tt.want {
				t.Errorf("IsNull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSON_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		j       JSON
		want    []byte
		wantErr bool
	}{
		{"nil", nil, nil, false},
		{"empty", JSON{}, []byte{}, false},
		{"ok", JSON{3, 5, 6, 2, 0}, []byte{3, 5, 6, 2, 0}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.j.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSON_Scan(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantObj JSON
		wantErr bool
	}{
		{"nil", args{nil}, nil, false},
		{"empty", args{[]byte{}}, JSON{}, false},
		{"ok", args{[]byte{4, 5, 2, 'f', 0, 1}}, JSON{4, 5, 2, 'f', 0, 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := JSON{}
			if err := j.Scan(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Scan() error = %v, want %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(j, tt.wantObj) {
				t.Errorf("after Scan() JSON = %v, want %v", j, tt.wantObj)
			}
		})
	}
}

func TestJSON_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		wantObj JSON
		wantErr bool
	}{
		{"nil", args{nil}, JSON{}, false},
		{"empty", args{[]byte{}}, JSON{}, false},
		{"ok", args{[]byte("ok")}, JSON{111, 107}, false},
	}
	for _, tt := range tests {
		j := JSON{}
		t.Run(tt.name, func(t *testing.T) {
			if err := j.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(j, tt.wantObj) {
				t.Errorf("after UnmarshalJSON() JSON = %v, want %v", j, tt.wantObj)
			}
		})
	}
}

func TestJSON_Value(t *testing.T) {
	tests := []struct {
		name    string
		j       JSON
		want    driver.Value
		wantErr bool
	}{
		{"nil", nil, nil, false},
		{"empty", JSON{}, nil, false},
		{"ok", JSON("records"), "records", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.j.Value()
			if (err != nil) != tt.wantErr {
				t.Errorf("Value() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Value() got = %v, want %v", got, tt.want)
			}
		})
	}
}
