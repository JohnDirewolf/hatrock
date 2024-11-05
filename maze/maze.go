package maze

//This handles the maze and compass specific functions.

import (
	"fmt"
	"html/template"
	"log"
	"strings"

	"database/sql"

	"github.com/JohnDirewolf/capstone/database"
	"github.com/JohnDirewolf/capstone/shared/types"
)

// Global variables
// var theMaze map[int]types.RoomData
var playerLocation int

func Init() {
	//rest the database to the initial state.
	//Clear the database
	database.Clear()
	//Add items.
	database.InsertItem(types.ItemData{
		Id:          1,
		Name:        "Golden Key",
		Article:     "a ",
		Description: "A delicate golden ward key with a ruby in the bow.",
		CurLocation: 8,
	})
	database.InsertItem(types.ItemData{
		Id:          2,
		Name:        "Magic Sword",
		Article:     "a ",
		Description: "A sword that glows blue. Along the blad are runes that says \"Goblin Scourge\"",
		CurLocation: 15,
	})
	//This sets up a lucky coin in the player's default inventory as item 0. Player inventory is location "-1"
	database.InsertItem(types.ItemData{
		Id:          0,
		Name:        "Lucky Coin",
		Article:     "a ",
		Description: "A ancient coin you found long ago and you feel has brought you luck.",
		CurLocation: -1,
	})
	//Set up rooms to initial state
	database.InsertRoom(types.RoomData{
		Id:          0,
		Title:       "Bark Room",
		Description: "The room is covered in various types of tree bark. The smell is musky. The only exist is the way you entered to the East.",
		Discovered:  false,
		North:       types.DoorData{Exists: false, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		South:       types.DoorData{Exists: false, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		West:        types.DoorData{Exists: false, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		East:        types.DoorData{Exists: true, Locked: false, KeyID: sql.NullInt16{Valid: false}},
	})
	database.InsertRoom(types.RoomData{
		Id:          1,
		Title:       "Leaf Room",
		Description: "You find a room overgrown with leaves. Pushing through the folliage you find exits to the North, West, and East.",
		Discovered:  false,
		North:       types.DoorData{Exists: true, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		South:       types.DoorData{Exists: false, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		West:        types.DoorData{Exists: true, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		East:        types.DoorData{Exists: true, Locked: false, KeyID: sql.NullInt16{Valid: false}},
	})
	database.InsertRoom(types.RoomData{
		Id:          2,
		Title:       "Stone Room",
		Description: "You are in a room made of cleaved stones. You see doors to the North, West, and East.",
		Discovered:  true,
		North:       types.DoorData{Exists: true, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		South:       types.DoorData{Exists: false, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		West:        types.DoorData{Exists: true, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		East:        types.DoorData{Exists: true, Locked: false, KeyID: sql.NullInt16{Valid: false}},
	})
	database.InsertRoom(types.RoomData{
		Id:          3,
		Title:       "Grass Room",
		Description: "This room is filled with grass and smells natural and clean. You can make out doors to the North and West.",
		Discovered:  false,
		North:       types.DoorData{Exists: true, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		South:       types.DoorData{Exists: false, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		West:        types.DoorData{Exists: true, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		East:        types.DoorData{Exists: false, Locked: false, KeyID: sql.NullInt16{Valid: false}},
	})
	database.InsertRoom(types.RoomData{
		Id:          4,
		Title:       "Dirt Room",
		Description: "You enter a room that is empty, the foor being ony dirt. There are exits to the North and East.",
		Discovered:  false,
		North:       types.DoorData{Exists: true, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		South:       types.DoorData{Exists: false, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		West:        types.DoorData{Exists: false, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		East:        types.DoorData{Exists: true, Locked: false, KeyID: sql.NullInt16{Valid: false}},
	})
	database.InsertRoom(types.RoomData{
		Id:          5,
		Title:       "Water Room",
		Description: "As you enter you fall into a pool of warm water. Swimming about you find exits to the North, South and West.",
		Discovered:  false,
		North:       types.DoorData{Exists: true, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		South:       types.DoorData{Exists: true, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		West:        types.DoorData{Exists: true, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		East:        types.DoorData{Exists: false, Locked: false, KeyID: sql.NullInt16{Valid: false}},
	})
	database.InsertRoom(types.RoomData{
		Id:          6,
		Title:       "Oil Room",
		Description: "The air is hard to breath here as the room is knee deep in black oil. Wading through the room you find doors to the North and South.",
		Discovered:  false,
		North:       types.DoorData{Exists: true, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		South:       types.DoorData{Exists: true, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		West:        types.DoorData{Exists: false, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		East:        types.DoorData{Exists: false, Locked: false, KeyID: sql.NullInt16{Valid: false}},
	})
	database.InsertRoom(types.RoomData{
		Id:          7,
		Title:       "Wood Room",
		Description: "All around is worked wood creating a cosy cabin feel. There is only a door to the South.",
		Discovered:  false,
		North:       types.DoorData{Exists: false, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		South:       types.DoorData{Exists: true, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		West:        types.DoorData{Exists: false, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		East:        types.DoorData{Exists: false, Locked: false, KeyID: sql.NullInt16{Valid: false}},
	})
	database.InsertRoom(types.RoomData{
		Id:          8,
		Title:       "Lava Room",
		Description: "The heat in this room is nearly unbarable as the floor is mostly lava. There seems to be no way forward, only the exit to the South.",
		Discovered:  false,
		North:       types.DoorData{Exists: false, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		South:       types.DoorData{Exists: true, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		West:        types.DoorData{Exists: false, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		East:        types.DoorData{Exists: false, Locked: false, KeyID: sql.NullInt16{Valid: false}},
	})
	database.InsertRoom(types.RoomData{
		Id:          9,
		Title:       "Boil Room",
		Description: "The air is full of steam and the sound of roiling water. A simple bridge over the boiling water allows exit to the North and South.",
		Discovered:  false,
		North:       types.DoorData{Exists: true, Locked: true, KeyID: sql.NullInt16{Int16: 1, Valid: true}},
		South:       types.DoorData{Exists: true, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		West:        types.DoorData{Exists: false, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		East:        types.DoorData{Exists: false, Locked: false, KeyID: sql.NullInt16{Valid: false}},
	})
	database.InsertRoom(types.RoomData{
		Id:          10,
		Title:       "Mud Room",
		Description: "You find a huge expanse of cracked mud, desperate for water. You can see exits South and East.",
		Discovered:  false,
		North:       types.DoorData{Exists: false, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		South:       types.DoorData{Exists: true, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		West:        types.DoorData{Exists: false, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		East:        types.DoorData{Exists: true, Locked: false, KeyID: sql.NullInt16{Valid: false}},
	})
	database.InsertRoom(types.RoomData{
		Id:          11,
		Title:       "Rust Room",
		Description: "There is the sound of clanking metals and steam flowing through old pipes in this room filed with rusted machinery. There are doors to the North and West.",
		Discovered:  false,
		North:       types.DoorData{Exists: true, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		South:       types.DoorData{Exists: false, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		West:        types.DoorData{Exists: true, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		East:        types.DoorData{Exists: false, Locked: false, KeyID: sql.NullInt16{Valid: false}},
	})
	database.InsertRoom(types.RoomData{
		Id:          12,
		Title:       "Copper Room - Goal!",
		Description: "Huzzah! Entering this room made of copper and metal you see a large portal open, and show the way out. You can go back to the maze to the East.",
		Discovered:  false,
		North:       types.DoorData{Exists: false, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		South:       types.DoorData{Exists: false, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		West:        types.DoorData{Exists: false, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		East:        types.DoorData{Exists: true, Locked: false, KeyID: sql.NullInt16{Valid: false}},
	})
	database.InsertRoom(types.RoomData{
		Id:          13,
		Title:       "Gold Room",
		Description: "The room glitters with gold in all shapes and sizes, then you realize it is just fool's gold. You see doors through the faux horde to the South, West and East.",
		Discovered:  false,
		North:       types.DoorData{Exists: false, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		South:       types.DoorData{Exists: true, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		West:        types.DoorData{Exists: true, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		East:        types.DoorData{Exists: true, Locked: false, KeyID: sql.NullInt16{Valid: false}},
	})
	database.InsertRoom(types.RoomData{
		Id:          14,
		Title:       "Magma Room",
		Description: "A powerful heat radiates from this room, magma slowly shifing. There is no other exit then the door you came in to the West.",
		Discovered:  false,
		North:       types.DoorData{Exists: false, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		South:       types.DoorData{Exists: false, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		West:        types.DoorData{Exists: true, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		East:        types.DoorData{Exists: false, Locked: false, KeyID: sql.NullInt16{Valid: false}},
	})
	database.InsertRoom(types.RoomData{
		Id:          15,
		Title:       "Granite Room",
		Description: "You enter a quarry of granite and stone. You only see a door back South.",
		Discovered:  false,
		North:       types.DoorData{Exists: false, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		South:       types.DoorData{Exists: true, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		West:        types.DoorData{Exists: false, Locked: false, KeyID: sql.NullInt16{Valid: false}},
		East:        types.DoorData{Exists: false, Locked: false, KeyID: sql.NullInt16{Valid: false}},
	})

	//Set player location to starting room, 2
	playerLocation = 2
}

func Move(direction string) {
	//Each direction changes the playerLocation by a different value. If there is no direction, then playerLocation does not change.
	//Get a RoomData from the Database
	room, err := database.GetRoom(playerLocation)
	if err != nil {
		log.Printf("Maze, Move, Error getting room: %v\n", err)
	}
	switch types.UrlAction(direction) {
	case types.North:
		if room.North.Exists {
			playerLocation += 4
		}
	case types.South:
		if room.South.Exists {
			playerLocation -= 4
		}
	case types.West:
		if room.West.Exists {
			playerLocation -= 1
		}
	case types.East:
		if room.East.Exists {
			playerLocation += 1
		}
	}

	//This sets the room the player is in to discovered, regardless if the player moved or not.
	database.DiscoverRoom(playerLocation)
}

func GetItems() {
	//move all the items from the room into the player's inventory.
	itemList, err := database.GetItemsByLocation(playerLocation)
	if err != nil {
		log.Printf("Maze, GetItems, Error getting list of items in room: %v\n", err)
	}
	for _, item := range itemList {
		database.MoveItemToLocation(item.Id, -1)
	}
}

func GenerateKnownMap() template.HTML {
	//This runs through the map and all rooms showing discovered have their container added to the string.
	var knownMap strings.Builder
	roomsDiscovered, err := database.GetDiscoveredRooms()
	if err != nil {
		log.Printf("Maze, GenerateKnownMap, Error getting list of discovered rooms: %v\n", err)
	}
	for _, roomID := range roomsDiscovered {
		fmt.Fprintf(&knownMap, "<div class='room%d'><img src='/assets/images/r%d.png' alt='Maze Room' width='200' height='200' /></div>\n", roomID, roomID)
	}
	//Add the Player
	playerLocationTop := ((15 - playerLocation) / 4) * 200
	playerLocationLeft := (playerLocation % 4) * 200
	//log.Printf("playerLocationLeft: %d", playerLocationLeft)
	fmt.Fprintf(&knownMap, "<div style='position:absolute;top:%dpx;left:%dpx'><img src='/assets/images/player.png' alt='Player!' width='200' height='200' /></div>\n", playerLocationTop, playerLocationLeft)
	return template.HTML(knownMap.String())
}

func GetPageInfo() types.PageData {
	var compass strings.Builder
	var inventory strings.Builder
	var description strings.Builder
	var action strings.Builder
	room, err := database.GetRoom(playerLocation)
	//log.Printf("Room returned to GenerateCompass is: %v", room)
	if err != nil {
		log.Printf("Maze, GenerateCompass, Error getting room: %v\n", err)
	}
	if room.North.Exists {
		fmt.Fprint(&compass, `<div class="n-arrow"><a href="/app?action=north"><img class="n-arrow" src="/assets/images/green_arrow_n.png" alt="Green Arrow North" width="100" height="100" /></a></div>`)
	} else {
		fmt.Fprint(&compass, `<div class="n-arrow"><img class="n-arrow" src="/assets/images/red_arrow_n.png" alt="Red Arrow North" width="100" height="100" /></div>`)
	}
	if room.South.Exists {
		fmt.Fprint(&compass, `<div class="s-arrow"><a href="/app?action=south"><img class="s-arrow" src="/assets/images/green_arrow_s.png" alt="Green Arrow South" width="100" height="100" /></a></div>`)
	} else {
		fmt.Fprint(&compass, `<div class="s-arrow"><img class="s-arrow" src="/assets/images/red_arrow_s.png" alt="Red Arrow South" width="100" height="100" /></div>`)
	}
	if room.West.Exists {
		fmt.Fprint(&compass, `<div class="w-arrow"><a href="/app?action=west"><img class="w-arrow" src="/assets/images/green_arrow_w.png" alt="Green Arrow West" width="100" height="100" /></a></div>`)
	} else {
		fmt.Fprint(&compass, `<div class="w-arrow"><img class="w-arrow" src="/assets/images/red_arrow_w.png" alt="Red Arrow West" width="100" height="100" /></div>`)
	}
	if room.East.Exists {
		fmt.Fprint(&compass, `<div class="e-arrow"><a href="/app?action=east"><img class="e-arrow" src="/assets/images/green_arrow_e.png" alt="Green Arrow East" width="100" height="100" /></a></div>`)
	} else {
		fmt.Fprint(&compass, `<div class="e-arrow"><img class="e-arrow" src="/assets/images/red_arrow_e.png" alt="Red Arrow East" width="100" height="100" /></div>`)
	}
	//Show character inventory (Location -1)
	itemsInInventory, err := database.GetItemsByLocation(-1)
	if err != nil {
		log.Printf("Maze, GetePageInfo, Error getting inventory items: %v\n", err)
	}
	for index, item := range itemsInInventory {
		if index >= 1 {
			fmt.Fprint(&inventory, "<br />")
		}
		fmt.Fprint(&inventory, item.Name)
	}

	//Checking for items in the room. If there are item(s) in the room we add an action, get item, and the item(s) are added to the description.
	fmt.Fprint(&description, room.Description)
	itemsInLocation, err := database.GetItemsByLocation(playerLocation)
	if err != nil {
		log.Printf("Maze, GetePageInfo, Error getting room items: %v\n", err)
	}
	if len(itemsInLocation) >= 1 {
		fmt.Fprint(&description, "<br />In the room you see: ")
		for index, item := range itemsInLocation {
			if index >= 1 {
				fmt.Fprint(&description, ", ")
			}
			fmt.Fprint(&description, item.Article)
			fmt.Fprint(&description, item.Name)
		}
		//Action to get items
		fmt.Fprint(&action, `<div class="action"><a class="action" href="/app?action=get"><span class="action">`)
		if len(itemsInLocation) == 1 {
			//Single item in room.
			fmt.Fprint(&action, `Get `)
			fmt.Fprint(&action, itemsInLocation[0].Name)
		} else {
			//More then one item in the room.
			fmt.Fprint(&action, `Get all items.`)
		}
		fmt.Fprint(&action, `</span></a></div>`)
	}

	pageData := types.PageData{
		Title:       room.Title,
		Rooms:       GenerateKnownMap(),
		Compass:     template.HTML(compass.String()),
		Inventory:   template.HTML(inventory.String()),
		Description: template.HTML(description.String()),
		Action:      template.HTML(action.String()),
	}

	return pageData
}

func GetInstructions() template.HTML {
	//This sends the Instructions HTML when we want to display them, specifically on the start of the game.
	return template.HTML(`
		<div class="instructions">
			<p class="instructions">Click on the Compass Arrow in the direction you wish to go!</p>
		</div>`)
}
