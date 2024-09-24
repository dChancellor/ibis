package main

import (
	"context"
    "errors"
    "fmt"
)


type App struct {
	ctx context.Context
    messages MessagesStruct
}

type ReturnMessage string

type MessagesStruct struct {
    SkillExistsMessage ReturnMessage
    SkillAddedMessage  ReturnMessage
}

type Skill struct {
    id int
    Name string
    SVG string
}

var Messages = MessagesStruct{
    SkillExistsMessage: "Skill already exists",
    SkillAddedMessage: "Skill added successfully",
}


// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) AddSkill(name string, svg string) error {
    var ErrSkillExists = errors.New("This skill already exists")
    var exists bool
    err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM skills WHERE name = ?)", name).Scan(&exists)
    if err != nil {
        return err
    }

    // If the skill already exists, return without inserting
    if exists {
        return ErrSkillExists
    }

    fmt.Println("SVG", svg)
    // If the skill doesn't exist, proceed to insert
    statement, err := db.Prepare("INSERT INTO skills (name, svg) VALUES (?, ?)")
    if err != nil {
        return err
    }
    _, err = statement.Exec(name, svg)
    return err
}

func (a *App) GetSkills() ([]Skill, error) {
    rows, err := db.Query("SELECT * FROM skills")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var skills = []Skill{}
    for rows.Next() {
        var skill Skill
        // rows.Scan(&skill.Name)
        // rows.Scan(&skill.SVG)
        err = rows.Scan(&skill.id, &skill.Name, &skill.SVG)
        fmt.Println("ERROR", err)
        skills = append(skills, skill)
    }
    fmt.Println("SKILLS", skills)
    
    return skills, nil
}

func (a *App) GetConstants() (MessagesStruct) {
    return Messages
}

func (a *App) DeleteSkill(name string) ([]Skill, error) {
    var exists bool

    // Check if the skill exists
    err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM skills WHERE name = ?)", name).Scan(&exists)
    if err != nil {
        return nil, err
    }

    // Delete the skill from the database
    _, err = db.Exec("DELETE FROM skills WHERE name = ?", name)
    if err != nil {
        return nil, err
    }

    // Retrieve and return the remaining skills
    return a.GetSkills()
}

func (a *App) DeleteAllSkills() error {
	// Prepare the DELETE statement
	statement, err := db.Prepare("DELETE FROM skills")
	if err != nil {
		return err
	}

	// Execute the DELETE statement to remove all rows
	_, err = statement.Exec()
	if err != nil {
		return err
	}

	return nil
}