package application

import (
	"context"
    "errors"
    "ibis/backend/db"
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
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) AddSkill(name string, svg string) error {
    var ErrSkillExists = errors.New("This skill already exists")
    var exists bool
    err := db.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM skills WHERE name = ?)", name).Scan(&exists)
    if err != nil {
        return err
    }

    // If the skill already exists, return without inserting
    if exists {
        return ErrSkillExists
    }

    // If the skill doesn't exist, proceed to insert
    statement, err := db.DB.Prepare("INSERT INTO skills (name, svg) VALUES (?, ?)")
    if err != nil {
        return err
    }
    _, err = statement.Exec(name, svg)
    return err
}

func (a *App) GetSkills() ([]Skill, error) {
    rows, err := db.DB.Query("SELECT * FROM skills")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var skills = []Skill{}
    for rows.Next() {
        var skill Skill
        err = rows.Scan(&skill.id, &skill.Name, &skill.SVG)
        skills = append(skills, skill)
    }
    return skills, nil
}

func (a *App) GetConstants() (MessagesStruct) {
    return Messages
}

func (a *App) DeleteSkill(name string) ([]Skill, error) {
    var exists bool

    // Check if the skill exists
    err := db.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM skills WHERE name = ?)", name).Scan(&exists)
    if err != nil {
        return nil, err
    }

    // Delete the skill from the database
    _, err = db.DB.Exec("DELETE FROM skills WHERE name = ?", name)
    if err != nil {
        return nil, err
    }

    // Retrieve and return the remaining skills
    return a.GetSkills()
}

func (a *App) DeleteAllSkills() error {
	// Prepare the DELETE statement
	statement, err := db.DB.Prepare("DELETE FROM skills")
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