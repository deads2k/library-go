package manifestclient

import (
	"reflect"
	"testing"
)

func TestDifferenceOfSerializedRequests(t *testing.T) {
	type args struct {
		lhs []FileOriginatedSerializedRequest
		rhs []TrackedSerializedRequest
	}
	tests := []struct {
		name string
		args args
		want []FileOriginatedSerializedRequest
	}{
		{
			name: "different types, no diff",
			args: args{
				lhs: []FileOriginatedSerializedRequest{
					{
						BodyFilename:    "foo.yaml",
						OptionsFilename: "foo-options.yaml",
						SerializedRequest: SerializedRequest{
							Namespace: "foo-ns",
							Name:      "bar",
							Options:   nil,
							Body:      []byte("content"),
						},
					},
				},
				rhs: []TrackedSerializedRequest{
					{
						RequestNumber: 6,
						SerializedRequest: SerializedRequest{
							Namespace: "foo-ns",
							Name:      "bar",
							Options:   nil,
							Body:      []byte("content"),
						},
					},
				},
			},
			want: []FileOriginatedSerializedRequest{},
		},
		{
			name: "different types with diff",
			args: args{
				lhs: []FileOriginatedSerializedRequest{
					{
						BodyFilename:    "foo.yaml",
						OptionsFilename: "foo-options.yaml",
						SerializedRequest: SerializedRequest{
							Namespace: "foo-ns",
							Name:      "bar",
							Options:   []byte("options!"),
							Body:      []byte("content"),
						},
					},
				},
				rhs: []TrackedSerializedRequest{
					{
						RequestNumber: 6,
						SerializedRequest: SerializedRequest{
							Namespace: "foo-ns",
							Name:      "bar",
							Options:   nil,
							Body:      []byte("content"),
						},
					},
				},
			},
			want: []FileOriginatedSerializedRequest{
				{
					BodyFilename:    "foo.yaml",
					OptionsFilename: "foo-options.yaml",
					SerializedRequest: SerializedRequest{
						Namespace: "foo-ns",
						Name:      "bar",
						Options:   []byte("options!"),
						Body:      []byte("content"),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DifferenceOfSerializedRequests(tt.args.lhs, tt.args.rhs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DifferenceOfSerializedRequests() = %v, want %v", got, tt.want)
			}
		})
	}
}
