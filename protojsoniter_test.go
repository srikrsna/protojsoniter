package protojsoniter_test

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
	gofuzz "github.com/google/gofuzz"
	jsoniter "github.com/json-iterator/go"
	"github.com/srikrsna/goprotofuzz"
	testv1 "github.com/srikrsna/protojsoniter/internal/gen/test/v1"
	"github.com/srikrsna/protojsoniter/internal/gen/test/v1/testv1fuzz"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/testing/protocmp"
)

var (
	protoJsonMarshalOpt = &protojson.MarshalOptions{}
)

func FuzzReadWrite(f *testing.F) {
	f.Add([]byte("0"))
	f.Add([]byte("7"))
	f.Add([]byte(""))
	f.Add([]byte("\u007f\x922"))
	f.Fuzz(func(t *testing.T, data []byte) {
		f := gofuzz.NewFromGoFuzz(data).Funcs(testv1fuzz.FuzzFuncs()...).Funcs(goprotofuzz.FuzzWKT[:]...)
		var before testv1.All
		f.Fuzz(&before)
		w := jsoniter.ConfigDefault.BorrowStream(nil)
		before.WriteJSON(w)
		if w.Error != nil {
			t.Fatal(w.Error)
		}
		if !json.Valid(w.Buffer()) {
			t.Fatal("invalid json: ", string(w.Buffer()))
		}
		r := jsoniter.ConfigDefault.BorrowIterator(w.Buffer())
		var after testv1.All
		after.ReadJSON(r)
		if r.Error != nil {
			t.Error(r.Error)
		}
		if diff := cmp.Diff(&before, &after, protocmp.Transform()); diff != "" {
			t.Error("before and after did not match", diff)
			t.Error(string(w.Buffer()))
		}
	})
}

func FuzzRead(f *testing.F) {
	f.Add([]byte("0"))
	f.Fuzz(func(t *testing.T, data []byte) {
		f := gofuzz.NewFromGoFuzz(data).Funcs(testv1fuzz.FuzzFuncs()...).Funcs(goprotofuzz.FuzzWKT[:]...)
		var before testv1.All
		f.Fuzz(&before)
		jsonData, err := protoJsonMarshalOpt.Marshal(&before)
		if err != nil {
			t.Fatal("marshal error: ", err)
		}
		if !json.Valid(jsonData) {
			t.Fatal("invalid json: ", string(jsonData))
		}
		r := jsoniter.ConfigDefault.BorrowIterator(jsonData)
		var after testv1.All
		after.ReadJSON(r)
		if r.Error != nil {
			t.Error(r.Error)
		}
		if diff := cmp.Diff(&before, &after, protocmp.Transform()); diff != "" {
			t.Error("before and after did not match", diff)
			t.Error(string(jsonData))
		}
	})
}

func FuzzWrite(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		f := gofuzz.NewFromGoFuzz(data).Funcs(testv1fuzz.FuzzFuncs()...).Funcs(goprotofuzz.FuzzWKT[:]...)
		var before testv1.All
		f.Fuzz(&before)
		w := jsoniter.ConfigDefault.BorrowStream(nil)
		before.WriteJSON(w)
		if w.Error != nil {
			t.Fatal(w.Error)
		}
		if !json.Valid(w.Buffer()) {
			t.Fatal("invalid json: ", string(w.Buffer()))
		}
		var after testv1.All
		if err := protojson.Unmarshal(w.Buffer(), &after); err != nil {
			t.Fatal(err, string(w.Buffer()))
		}
		if diff := cmp.Diff(&before, &after, protocmp.Transform()); diff != "" {
			t.Error("before and after did not match", diff)
			t.Error(string(w.Buffer()))
		}
	})
}
