package session

import "testing"

var (
	user1 = &User{"Tom", 18}
	user2 = &User{"Jack", 20}
	User3 = &User{"Lily", 25}
)

func testRecordInit(t *testing.T) *Session {
	t.Helper()
	s := NewSession().Model(&User{})
	err1 := s.DropTable()
	err2 := s.CreateTable()
	_, err3 := s.Insert(user1, user2)
	if err1 != nil || err2 != nil || err3 != nil {
		t.Fatal("Failed to init test records")
	}
	return s
}

func TestSession_Find(t *testing.T) {
	s := testRecordInit(t)
	var users []User
	err := s.Limit(1).Find(&User{})
	if err != nil || len(users) != 1 {
		t.Fatal("Failed to query with limit record")
	}
}
