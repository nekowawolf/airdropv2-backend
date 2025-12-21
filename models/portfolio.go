package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Portfolio struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`

	Hero         HeroProfile   `bson:"hero" json:"hero"`
	Certificates []Certificate `bson:"certificates" json:"certificates"`
	Designs      []Design      `bson:"designs" json:"designs"`
	Projects     []Project     `bson:"projects" json:"projects"`
	Experience   []Experience  `bson:"experience" json:"experience"`
	Education    []Education   `bson:"education" json:"education"`
	Skills       Skills        `bson:"skills" json:"skills"`
}

type HeroProfile struct {
	Name      string `bson:"name" json:"name"`
	Summary   string `bson:"summary" json:"summary"`
	AvatarURL string `bson:"avatar_url" json:"avatar_url"`
	CVURL     string `bson:"cv_url" json:"cv_url"`
	Github    string `bson:"github" json:"github"`
	Twitter   string `bson:"twitter" json:"twitter"`
	LinkedIn  string `bson:"linkedin" json:"linkedin"`
	Email     string `bson:"email" json:"email"`
}

type BaseItem struct {
	ID string `bson:"id" json:"id"`
}

type Certificate struct {
	BaseItem `bson:",inline"`
	Title    string `bson:"title" json:"title"`
	ImageURL string `bson:"image_url" json:"image_url"`
}

type Design struct {
	BaseItem    `bson:",inline"`
	Title       string           `bson:"title" json:"title"`
	Description string           `bson:"description" json:"description"`
	ImageURL    string           `bson:"image_url" json:"image_url"`
	Link        string           `bson:"link" json:"link"`

	Category    string           `bson:"category" json:"category"`
	Tools       []string         `bson:"tools,omitempty" json:"tools,omitempty"`
	Tags        []string         `bson:"tags,omitempty" json:"tags,omitempty"`
	Screenshots []ScreenshotItem `bson:"screenshots,omitempty" json:"screenshots,omitempty"`
}

type ScreenshotItem struct {
	ImageURL    string `bson:"image_url" json:"image_url"`
	Description string `bson:"description,omitempty" json:"description,omitempty"`
}

type Project struct {
	BaseItem    `bson:",inline"`
	Title        string       `bson:"title" json:"title"`
	Description  string       `bson:"description" json:"description"`
	ImageURL     string       `bson:"image_url" json:"image_url"` 
	Link         string       `bson:"link" json:"link"`           
	GitHubURL    string       `bson:"github_url,omitempty" json:"github_url,omitempty"` 
	Screenshots  []string     `bson:"screenshots,omitempty" json:"screenshots,omitempty"`
	VideoURL     string       `bson:"video_url,omitempty" json:"video_url,omitempty"`
	UseCase      DiagramItem  `bson:"use_case,omitempty" json:"use_case,omitempty"`
	Activity     DiagramItem  `bson:"activity,omitempty" json:"activity,omitempty"`
	ERD          DiagramItem  `bson:"erd,omitempty" json:"erd,omitempty"`
	Flowchart    DiagramItem  `bson:"flowchart,omitempty" json:"flowchart,omitempty"`
	Stack        []string     `bson:"stack,omitempty" json:"stack,omitempty"`
}

type DiagramItem struct {
	ImageURL string `bson:"image_url" json:"image_url"`
	Description  string       `bson:"description" json:"description"`
}

type Experience struct {
	BaseItem   `bson:",inline"`
	Title       string `bson:"title" json:"title"`
	Description string `bson:"description" json:"description"`
	Subjects    []string `bson:"subjects" json:"subjects"`
	StartYear   int    `bson:"start_year" json:"start_year"`
	EndYear     int    `bson:"end_year" json:"end_year"`
}

type Education struct {
	BaseItem   `bson:",inline"`
	Title       string `bson:"title" json:"title"` 
	Description string   `bson:"description" json:"description"`
	Subjects    []string `bson:"subjects" json:"subjects"`
	StartYear   int    `bson:"start_year" json:"start_year"`
	EndYear     int    `bson:"end_year" json:"end_year"`
}

type Skills struct {
	Tech   []SkillItem `bson:"tech" json:"tech"`
	Design []SkillItem `bson:"design" json:"design"`
}

type SkillItem struct {
	BaseItem `bson:",inline"`
	Name     string `bson:"name" json:"name"`
	IconURL string `bson:"icon_url" json:"icon_url"`
}