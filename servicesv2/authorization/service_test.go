package authorization_test

import (
	"context"
	"errors"
	"io/ioutil"
	"os"
	"testing"

	influxdb "github.com/influxdata/influxdb/servicesv2"
	"github.com/influxdata/influxdb/servicesv2/authorization"
	"github.com/influxdata/influxdb/servicesv2/bolt"
	"github.com/influxdata/influxdb/servicesv2/kv"
	"github.com/influxdata/influxdb/servicesv2/tenant"
	influxdbtesting "github.com/influxdata/influxdb/servicesv2/testing"
	"go.uber.org/zap/zaptest"
)

func initBoltAuthService(f influxdbtesting.AuthorizationFields, t *testing.T) (influxdb.AuthorizationService, string, func()) {
	s, closeBolt, err := NewTestBoltStore(t)
	if err != nil {
		t.Fatalf("failed to create new kv store: %v", err)
	}

	svc, closeSvc := initAuthService(s, f, t)
	return svc, "service_auth", func() {
		closeSvc()
		closeBolt()
	}
}

func initAuthService(s kv.Store, f influxdbtesting.AuthorizationFields, t *testing.T) (influxdb.AuthorizationService, func()) {
	st := tenant.NewStore(s)
	ts := tenant.NewService(st)
	storage, err := authorization.NewStore(s)
	if err != nil {
		t.Fatal(err)
	}
	svc := authorization.NewService(storage, ts)

	for _, u := range f.Users {
		if err := ts.CreateUser(context.Background(), u); err != nil {
			t.Fatalf("error populating users: %v", err)
		}
	}

	for _, o := range f.Orgs {
		if err := ts.CreateOrganization(context.Background(), o); err != nil {
			t.Fatalf("failed to populate organizations: %s", err)
		}
	}

	for _, m := range f.Authorizations {
		if err := svc.CreateAuthorization(context.Background(), m); err != nil {
			t.Fatalf("failed to populate authorizations: %v", err)
		}
	}

	return svc, func() {
		for _, m := range f.Authorizations {
			if err := svc.DeleteAuthorization(context.Background(), m.ID); err != nil {
				t.Logf("failed to remove user resource mapping: %v", err)
			}
		}
	}
}

func NewTestBoltStore(t *testing.T) (kv.Store, func(), error) {
	f, err := ioutil.TempFile("", "influxdata-bolt-")
	if err != nil {
		return nil, nil, errors.New("unable to open temporary boltdb file")
	}
	f.Close()

	path := f.Name()
	ctx := context.Background()
	logger := zaptest.NewLogger(t)

	s := bolt.NewKVStore(logger, path)

	if err := s.Open(ctx); err != nil {
		return nil, nil, err
	}

	// if err := all.Up(ctx, logger, s); err != nil {
	// 	return nil, nil, err
	// }

	close := func() {
		s.Close()
		os.Remove(path)
	}

	return s, close, nil
}

func TestBoltAuthService(t *testing.T) {
	t.Parallel()
	influxdbtesting.AuthorizationService(initBoltAuthService, t)
}