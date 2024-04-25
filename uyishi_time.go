// package main

// import(
// 	"fmt"
// 	"time"
// )
// func main(){
// 	n := time.Now()
// 	f := time.Date(2024,04,16,11,47,45,100,time.Local)
// 	fmt.Println(n)
// 	fmt.Println(f)
// }



package main

import (
	"fmt"
	"time"
)

type Event struct {
	Title     string
	StartTime time.Time
	EndTime   time.Time
}

type User struct {
	Name   string
	Events []Event
}

func (u *User) AddEvent(event Event) {
	u.Events = append(u.Events, event)
}

func (u *User) RemoveEvent(index int) {
	if index >= 0 && index < len(u.Events) {
		u.Events = append(u.Events[:index], u.Events[index+1:]...)
	}
}

func (u *User) UpdateEvent(index int, event Event) {
	if index >= 0 && index < len(u.Events) {
		u.Events[index] = event
	}
}

func main() {
	var name string
	fmt.Print("Foydalanuvchi ismini kiriting: ")
	fmt.Scan(&name)

	People := User{
		Name: name,
	}

	var count int
	fmt.Print("Qancha event qoshasiz? ")
	fmt.Scan(&count)

	for i := 0; i < count; i++ {
		var title string
		var year, month, day, hour, min int

		fmt.Printf("Event #%d nomini kiriting: ", i+1)
		fmt.Scan(&title)

		fmt.Printf("Event #%d boshlangich vaqtni kiriting (yil oy kun soat daqiqa): ", i+1)
		fmt.Scan(&year, &month, &day, &hour, &min)

		startTime := time.Date(year, time.Month(month), day, hour, min, 0, 0, time.Local)

		fmt.Printf("Event #%d tugash vaqtni kiriting (yil oy kun soat daqiqa): ", i+1)
		fmt.Scan(&year, &month, &day, &hour, &min)

		endTime := time.Date(year, time.Month(month), day, hour, min, 0, 0, time.Local)

		event := Event{
			Title:     title,
			StartTime: startTime,
			EndTime:   endTime,
		}

		People.AddEvent(event)
	}

	fmt.Printf("%s ning eventlari:\n", People.Name)
	for i, event := range People.Events {
		fmt.Printf("%d. %s: %s - %s\n", i+1, event.Title, event.StartTime, event.EndTime)
	}

	var removeIndex int
	fmt.Print("Ochirmoqchi bolgan event raqamini kiriting: ")
	fmt.Scan(&removeIndex)

	People.RemoveEvent(removeIndex - 1)

	var updateIndex int
	fmt.Print("Ozgartirmoqchi bolgan event raqamini kiriting: ")
	fmt.Scan(&updateIndex)

	var updatedTitle string
	var updatedStartTime, updatedEndTime time.Time

	fmt.Print("Yangi nom: ")
	fmt.Scan(&updatedTitle)

	year := 0
	month := 0
	day := 0
	hour := 0
	min := 0

	fmt.Print("Yangi boshlangich vaqtni kiriting (yil oy kun soat daqiqa): ")
	fmt.Scan(&year, &month, &day, &hour, &min)
	updatedStartTime = time.Date(year, time.Month(month), day, hour, min, 0, 0, time.Local)

	fmt.Print("Yangi tugash vaqtni kiriting (yil oy kun soat daqiqa): ")
	fmt.Scan(&year, &month, &day, &hour, &min)
	updatedEndTime = time.Date(year, time.Month(month), day, hour, min, 0, 0, time.Local)

	updatedEvent := Event{
		Title:     updatedTitle,
		StartTime: updatedStartTime,
		EndTime:   updatedEndTime,
	}

	People.UpdateEvent(updateIndex-1, updatedEvent)

	
	fmt.Println(updatedStartTime)
	fmt.Println(updatedEndTime)
}