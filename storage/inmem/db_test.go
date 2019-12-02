package inmem

import (
	"context"
	"reflect"
	"strings"

	"testing"

	protocol "github.com/libp2p/go-libp2p-core/protocol"
	database "github.com/vgonkivs/ddd/db"

	net "github.com/libp2p/go-libp2p-core/network"

	bhost "github.com/libp2p/go-libp2p-blankhost"
	"github.com/libp2p/go-libp2p-core/peer"
	swarmt "github.com/libp2p/go-libp2p-swarm/testing"
)

var testDB = NewInmemDB(peer.ID("selfInfo"))
var testPID, testStream = genStream()
var emptyStream net.Stream

func TestStore(t *testing.T) {

	cases := []struct {
		key     peer.ID
		u       *User
		isError bool
	}{
		{peer.ID("selfInfo"), &User{stream: emptyStream}, true},
		{testPID, &User{stream: emptyStream}, true},
		{testPID, &User{stream: testStream}, false},
	}
	for caseNum, item := range cases {
		err := testDB.Store(item.key, item.u.stream)
		if item.isError && err == nil {
			t.Errorf("[%d] expected error, got nil ", caseNum)
		}
		if !item.isError && err != nil {
			t.Errorf("[%d] unexpected error: %v", caseNum, err)
		}
	}
}

func TestIsUserExist(t *testing.T) {
	cases := []struct {
		key     peer.ID
		isExist bool
	}{
		{testPID, true},
		{peer.ID("selfInfo"), false},
	}
	for caseNum, testCase := range cases {
		isExist := testDB.IsUserExist(testCase.key)
		if isExist && !testCase.isExist {
			t.Errorf("[%d] unexpected result - the user is exist", caseNum)
		}
		if !isExist && testCase.isExist {
			t.Errorf("[%d]  unexpected result - the user is not exist", caseNum)
		}
	}
}

func TestGetUser(t *testing.T) {
	tDb := (testDB).(database.DB)

	cases := []struct {
		key     peer.ID
		u       *User
		isError bool
	}{
		{peer.ID("test"), &User{name: testPID.Pretty()}, true},
		{testPID, &User{name: testPID.Pretty()}, false},
	}

	for caseNum, testCase := range cases {
		str, err := tDb.GetUser(testCase.key)

		if err == nil && testCase.isError {
			t.Errorf("[%d] expected error, got nil ", caseNum)
			continue
		}

		if err != nil && !testCase.isError {
			t.Errorf("[%d] unknown error: %v ", caseNum, err)
			continue
		}

		if !strings.EqualFold(str, testCase.u.String()) && !testCase.isError {
			t.Errorf("[%d] recevied result not matches with chosen: %s - %s ", caseNum, str, testCase.u.String())
		}
	}
}
func TestGetAllUsers(t *testing.T) {
	var newUser, stream = genStream()
	testDB.Store(newUser, stream)

	tDb := (testDB).(database.DB)
	defer tDb.DeleteUser(newUser)
	cases := []struct {
		users   []string
		isError bool
	}{
		{[]string{testPID.Pretty(), "test1"}, true},
		{[]string{testPID.Pretty()}, true},
		{[]string{testPID.Pretty(), newUser.Pretty()}, false},
	}

	users := tDb.GetAllUsers()
	for caseNum, testCase := range cases {
		if len(users) != len(testCase.users) && !testCase.isError {
			t.Errorf("[%d] invalid quantity of users: %d - %d", caseNum, len(users), len(testCase.users))
			continue
		}

		if !reflect.DeepEqual(users, testCase.users) && !testCase.isError {
			t.Errorf("[%d] received user`s list doesn`t match with actual", caseNum)
		}
	}
}
func TestGetAllStreams(t *testing.T) {
	var newUser, stream = genStream()
	var _, unknownStream = genStream()
	testDB.Store(newUser, stream)

	tDb := (testDB).(database.DB)
	defer tDb.DeleteUser(newUser)

	cases := []struct {
		streams []net.Stream
		isError bool
	}{
		{[]net.Stream{testStream}, true},
		{[]net.Stream{testStream, unknownStream}, true},
		{[]net.Stream{testStream, stream}, false},
	}

	streams := tDb.GetAllStreams()
	for caseNum, testCase := range cases {
		if len(streams) != len(testCase.streams) && !testCase.isError {
			t.Errorf("[%d] invalid quantity of streams", caseNum)
			continue
		}
		if !reflect.DeepEqual(streams, testCase.streams) && !testCase.isError {
			t.Errorf("[%d] streams does not match", caseNum)
		}
	}
}

func TestDeleteUser(t *testing.T) {
	tDb := (testDB).(database.DB)
	cases := []struct {
		key     peer.ID
		isError bool
	}{
		{testPID, false},
		{testPID, true},
	}

	for caseNum, item := range cases {
		err := tDb.DeleteUser(item.key)
		if item.isError && err == nil {
			t.Errorf("[%d] expected error, got nil ", caseNum)
		}
		if !item.isError && err != nil {
			t.Errorf("[%d] unexpected error: %v", caseNum, err)
		}
	}
}

func genStream() (pid peer.ID, stream net.Stream) {
	h1 := bhost.NewBlankHost(swarmt.GenSwarm(nil, context.TODO()))
	h2 := bhost.NewBlankHost(swarmt.GenSwarm(nil, context.TODO()))
	h1.SetStreamHandler(protocol.ID("db_test/"), func(net.Stream) {})
	h2.SetStreamHandler(protocol.ID("db_test/"), func(net.Stream) {})

	if h1.ID() != h2.ID() {
		err := h1.Connect(context.TODO(), h2.Peerstore().PeerInfo(h2.ID()))
		if err != nil {
			return "", nil
		}
		stream, err := h1.NewStream(context.TODO(), h2.ID(), protocol.ID("db_test/"))
		if err != nil {
			return "", nil
		}
		return h2.ID(), stream
	}
	return "", nil
}
