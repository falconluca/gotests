package main_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
	"time"
)

type Player struct {
	Name string
}

func TestPtr(t *testing.T) {
	//
	// Struct
	//player := Player{Name: "Allen"}
	//fmt.Printf("outside func: %p\n", &player)
	//changePlayerName := func(player Player) {
	//	fmt.Printf("inside func: %p\n", &player)
	//	player.Name = "Luca"
	//}
	//changePlayerName(player)
	//fmt.Println("Player: ", player)

	//
	//
	//player := &Player{Name: "Allen"}
	//println(player)
	//fmt.Printf("outside func: %p\n", player)
	//changePlayerName := func(player *Player) {
	//	fmt.Printf("inside func: %p\n", player)
	//	player.Name = "Luca"
	//}
	//changePlayerName(player)
	//fmt.Println("Player: ", player)

	//
	// Slice FIXME
	//names := []string{"Luca", "Allen", "Curry"}
	////names := make([]string, 0, 4)
	//fmt.Printf("outside func: %p\n", names)
	//appendName := func(ns []string) {
	//	ns = append(ns, "newName")
	//	fmt.Printf("inside func: %p\n", ns)
	//}
	//appendName(names)
	//fmt.Println("names:", names)

	//
	// Slice FIXME
	//names := make([]string, 0, 4)
	//fmt.Printf("outside func: %p\n", names)
	//appendName := func(ns *[]string) {
	//	*ns = append(*ns, "newName")
	//	fmt.Printf("inside func: %p\n", ns)
	//}
	//appendName(&names)
	//fmt.Println("names:", names)

	//
	// Map
	//trends := map[string]int{"github": 10, "weibo": 23}
	//fmt.Printf("outside func: %p\n", trends)
	//addKeyValue := func(m map[string]int) {
	//	m["baidu"] = 20
	//	fmt.Printf("inside func: %p\n", m)
	//}
	//addKeyValue(trends)
	//fmt.Println("trends:", trends)

	//
	// int
	//age := 22
	//fmt.Printf("outside func: %p\n", &age)
	//changeAge := func(a int) {
	//	a = 24
	//	fmt.Printf("inside func: %p\n", &a)
	//}
	//changeAge(age)
	//fmt.Println("age:", age)

	//
	// int(ptr)
	//age := 22
	//fmt.Printf("outside func: %p\n", &age)
	//changeAge := func(a *int) {
	//	*a = 24
	//	fmt.Printf("inside func: %p\n", a)
	//}
	//changeAge(&age)
	//fmt.Println("age:", age)

	nums := []int{1, 2, 21, 23, 4, 5}
	nums2 := nums[2:4]
	fmt.Println(len(nums2))
}

func TestTimeNow(t *testing.T) {
	Time := func() int64 {
		return time.Now().Unix()
	}
	fmt.Println(Time())
}

func TestLocalTime(t *testing.T) {
	startDay := time.Date(2021, 11, 9, 20, 32, 0, 0, time.Local)
	fmt.Println(startDay)

	fmt.Println(convToDayBegin(startDay))
	fmt.Println(convToDayEnd(startDay))
}

func convToDayBegin(startTime time.Time) time.Time {
	return time.Date(startTime.Year(), startTime.Month(), startTime.Day(),
		0, 0, 0, 0, time.Local)
}

func convToDayEnd(indexDay time.Time) time.Time {
	return time.Date(indexDay.Year(), indexDay.Month(),
		indexDay.Day(), 23, 59, 59, 0, time.Local)
}

type Booking struct {
	Username  string
	CreatedAt time.Time
	Status    int
}

func TestSortSlice(t *testing.T) {
	bookings := []*Booking{
		{
			Username:  "Luca1",
			CreatedAt: time.Date(2022, 4, 13, 0, 0, 0, 0, time.Local),
			Status:    1,
		},
		{
			Username:  "Luca2",
			CreatedAt: time.Date(2022, 4, 4, 0, 0, 0, 0, time.Local),
			Status:    1,
		},
		{
			Username:  "Luca42",
			CreatedAt: time.Date(2022, 4, 22, 0, 0, 0, 0, time.Local),
			Status:    2,
		},
		{
			Username:  "Luca8",
			CreatedAt: time.Date(2022, 4, 30, 0, 0, 0, 0, time.Local),
			Status:    2,
		},
		{
			Username:  "Luca3",
			CreatedAt: time.Date(2022, 4, 23, 0, 0, 0, 0, time.Local),
			Status:    1,
		},
		{
			Username:  "Luca",
			CreatedAt: time.Now(),
			Status:    1,
		},
	}

	// 按 type 升序，createdAt 降序
	sort.Slice(bookings, func(i, j int) bool {
		if bookings[i].Status != bookings[j].Status {
			if bookings[i].Status == 1 {
				return true
			} else if bookings[j].Status == 1 {
				return false
			} else {
				return bookings[i].Status == 1
			}
		}

		return bookings[i].CreatedAt.Unix() > bookings[j].CreatedAt.Unix()
	})
	fmt.Println(bookings)
}

type (
	Info struct {
		Name string

		Type    string
		TypePtr *string

		Gender      *Gender
		GenderValue Gender

		Tags    []string
		TagsPtr *[]string

		Meta    map[string]interface{}
		MetaPtr *map[string]interface{}
	}

	Gender struct {
		Name  string
		Value int
	}
)

func TestRefactor(t *testing.T) {
	tt := assert.New(t)

	info := &Info{
		Name: "House Rent Message",
		Type: "house_rent",
	}
	tt.Nil(info.Gender, "info gender is not nil")
	info.Gender = new(Gender)
	tt.NotNil(info.Gender, "info gender is nil")

	tt.NotNil(info.GenderValue, "info genderValue is not nil")
	tt.Equal(info.GenderValue.Value, 0, "info genderValue value is not nil")

	info2 := &Info{
		Name: "Car Transaction",
		Type: "tx",
		Gender: &Gender{
			Name:  "男",
			Value: 1,
		},
	}

	tt.Equal("tx", info2.Type, "info2 type is not equal")
	updateInfoType(info2)
	tt.Equal("house_rent", info2.Type, "info2 type is not equal")

	tt.Equal(1, info2.Gender.Value, "info2 gender value is not equal")
	updateInfoGender(info2)
	tt.Equal(10, info2.Gender.Value, "info2 gender value is not equal")

	tt.Equal("男x", info2.Gender.Name, "info2 gender name is not equal")
	updateInfoGenderValue(info2)
	tt.Equal("女", info2.Gender.Name, "info2 gender name is not equal")

	//
	// Slice
	var tags []string
	tt.Equal(tags, info2.Tags, "info2 tags is not equal")

	tt.NotEqual([]string{}, info2.Tags, "info2 tags is not equal")

	tt.Equal(0, len(info2.Tags), "info2 tags len is not 0")
	updateInfoTags(info2)
	tt.Equal(2, len(info2.Tags), "info2 tags len is not 2")
	tt.Contains(info2.Tags, "greeting", "info2 tags is not contain greeting")

	tt.Nil(info2.TagsPtr, "info tags ptr is not nil")
	//tt.Nil(len(*info2.TagsPtr), "info tags ptr is not nil") // cause NPE

	//
	// Map
	info3 := &Info{
		Name: "Stuff Transaction",
		Type: "tx",
	}
	tt.Nil(info3.Meta, "info3 meta is not nil")
	tt.Equal(len(info3.Meta), 0, "info3 is not empty")
	tt.Nil(info3.Tags, "info3 tags is not nil")
	tt.Equal(len(info3.Tags), 0, "info3 is not empty")
	tt.Nil(info3.MetaPtr, "info3 meta ptr is not nil")
	//tt.Equal(len(*info3.MetaPtr), 0, "info3 is not empty")
	tt.Nil(info3.TagsPtr, "info3 tags ptr is not nil")
	//tt.Equal(len(*info3.TagsPtr), 0, "info3 is not empty")
}

func updateInfoTags(info *Info) {
	info.Tags = []string{"great", "greeting"}
}

func updateInfoGenderValue(info *Info) {
	info.Gender.Name = "女"
}

func updateInfoGender(info *Info) {
	info.Gender = &Gender{
		Name:  "男x",
		Value: 10,
	}
}

func updateInfoType(info *Info) {
	info.Type = "house_rent"
}

func TestTimeBefore(t *testing.T) {
	type Order struct {
		CreatedAt time.Time
	}

	orders := []*Order{
		{time.Date(2022, 10, 3, 0, 0, 0, 0, time.Local)},
		{time.Date(2022, 10, 2, 0, 0, 0, 0, time.Local)},
		{time.Date(2022, 10, 5, 0, 0, 0, 0, time.Local)},
		{time.Date(2022, 10, 8, 0, 0, 0, 0, time.Local)},
	}
	sort.Slice(orders, func(i, j int) bool {
		return orders[i].CreatedAt.Before(orders[j].CreatedAt)
	})
	fmt.Println(orders)
}

func TestSortIntSlice(t *testing.T) {
	type Order struct {
		GMV  int
		Type int
	}
	//nums := []Order{10, 2, 23, 13, 33, 3, 5}

	orders := []Order{
		{GMV: 10, Type: 1},
		{GMV: 2, Type: 1},
		{GMV: 23, Type: 2},
		{GMV: 13, Type: 1},
		{GMV: 33, Type: 1},
		{GMV: 3, Type: 1},
		{GMV: 5, Type: 2},
		{GMV: 10, Type: 2},
	}

	less := func(i, j int) bool { // i less than j
		return orders[i].GMV < orders[j].GMV
	}
	sort.Slice(orders, less)
	fmt.Printf("%+v\n", orders)
}

func TestTimeUtils(t *testing.T) {
	// create time.Time
	t1 := time.Date(2022, 4, 6, 0, 0, 0, 0, time.Local)

	// parse time.Time to string
	_ = t1.Format("2006-01-02 15:04:05")

	// parse string back to time.Time
	s1 := "2022-04-20 23:59"
	tt, err := time.ParseInLocation("2006-01-02 15:04", s1, time.Local)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(tt)

	// parse time.Time to int64
	unix := t1.Unix()

	// parse int64 back to time.Time
	_ = time.Unix(unix, 0)
	// Output:
	// 2022-04-01 00:00:00 +0800 CST

	// time.Before
	now := time.Now()
	tBeforeNow := time.Date(2022, 4, 15, 0, 0, 0, 0, time.Local)
	if ok := tBeforeNow.Before(now); ok {
		fmt.Println("tBeforeNow is before now")
	}

	// sort by time
	orders := []struct {
		CreatedAt time.Time
	}{
		{CreatedAt: time.Now().AddDate(0, 0, 7)},
		{CreatedAt: time.Now()},
		{CreatedAt: time.Now().AddDate(0, 0, 5)},
	}
	sort.Slice(orders, func(i, j int) bool {
		return orders[i].CreatedAt.After(orders[j].CreatedAt) // i after j mean i is less then j
	})
	fmt.Printf("CreatedAt After: %v\n", orders)
	// CreatedAt After: [{2022-04-22 15:30:26.012152 +0800 CST} {2022-04-20 15:30:26.012153 +0800 CST} {2022-04-15 15:30:26.012153 +0800 CST m=+0.001989710}]

	sort.Slice(orders, func(i, j int) bool {
		return orders[i].CreatedAt.Before(orders[j].CreatedAt) // i before j mean i is less than j
	})
	fmt.Printf("CreatedAt Before: %v\n", orders)
	// CreatedAt Before: [{2022-04-15 15:30:26.012153 +0800 CST m=+0.001989710} {2022-04-20 15:30:26.012153 +0800 CST} {2022-04-22 15:30:26.012152 +0800 CST}]
}

type (
	Snapshot struct {
		CreateAt   time.Time
		RoomDetail *RoomDetail
	}

	RoomDetail struct {
		RoomId      string
		CreatedAt   *time.Time
		Description string
	}
)

// go test -v --run=TestUpdatePtrInForRange ./a/
func TestUpdatePtrInForRange(t *testing.T) {
	now := time.Now()
	snapshots := []*Snapshot{
		{
			CreateAt: time.Now(),
			RoomDetail: &RoomDetail{
				RoomId:      "1000",
				CreatedAt:   &now,
				Description: "Crazy Cat",
			},
		},
		{
			CreateAt: time.Now(),
			RoomDetail: &RoomDetail{
				RoomId:      "1002",
				CreatedAt:   &now,
				Description: "Crazy Dog",
			},
		},
	}

	for _, snapshot := range snapshots {
		func(s *Snapshot) {
			s.CreateAt = s.CreateAt.AddDate(1, 0, 0)
			//s.RoomDetail = nil
			s.RoomDetail.CreatedAt.AddDate(1, 0, 0)
		}(snapshot)
	}
	fmt.Println(snapshots)
}
