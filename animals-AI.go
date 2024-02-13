package animals

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-redis/redis"
)

type Animal struct {
	Name     string
	Progress int
	Color    string
}

var (
	adjectives    = []string{"Loving", "Timid", "Furious", "Shiny", "Mechanical", "Pissed", "Cuddly"}
	animalNames   = []string{"Treefloof", "Murder Mittens", "Patience Monkey", "Forest Gorgi", "Wizard Cow", "Formal Chikcen"}
	colors        = []string{"\033[31m", "\033[32m", "\033[33m", "\033[34m", "\033[35m", "\033[36m"}
	maxProgress   = 3
	redisDuration = 600 * time.Second
)

func (a *Animal) Race(rdb *redis.Client) {
	key := a.Name
	status := rdb.Set(key, a.Progress, redisDuration)
	if status.Err() != nil {
		fmt.Println("Redis error:", status.Err())
	}
	val, err := rdb.Get(key).Result()
	if err != nil {
		panic(err)
	}
	advance := rand.Intn(maxProgress)
	if a.Progress > 1 {
		advance--
	}
	a.Progress += advance
}

func Create() Animal {
	adjective := adjectives[rand.Intn(len(adjectives))]
	name := animalNames[rand.Intn(len(animalNames))]
	color := colors[rand.Intn(len(colors))]
	return Animal{
		Name:  fmt.Sprintf("%s %s", adjective, name),
		Color: color,
	}
}

func IsDuplicate(animal Animal, animalList []Animal) bool {
	for _, a := range animalList {
		if animal.Name == a.Name {
			return true
		}
	}
	return false
}
