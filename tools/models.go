package main

type Quizz struct {
	V                    int             `json:"__v"`
	ID                   string          `json:"_id"`
	CurrentQuestionIndex int             `json:"currentQuestionIndex"`
	GameStarted          bool            `json:"gameStarted"`
	Pin                  int             `json:"pin"`
	Questions            []QuizzQuestion `json:"questions"`
}

type QuizzQuestion struct {
	Choices      []string `json:"choices"`
	CorrectIndex int      `json:"correctIndex"`
	Question     string   `json:"question"`
}

type Data struct {
	Card   DataCard   `json:"card" yaml:"card"`
	Kahoot DataKahoot `json:"kahoot" yaml:"kahoot"`
}

type DataCard struct {
	Access              DataCardAccess           `json:"access" yaml:"access"`
	Combined            bool                     `json:"combined" yaml:"combined"`
	CompatibilityLevel  int                      `json:"compatibility_level" yaml:"compatibility_level"`
	Cover               string                   `json:"cover" yaml:"cover"`
	CoverMetadata       DataCardCoverMetadata    `json:"coverMetadata" yaml:"coverMetadata"`
	Created             int                      `json:"created" yaml:"created"`
	Creator             string                   `json:"creator" yaml:"creator"`
	CreatorUsername     string                   `json:"creator_username" yaml:"creator_username"`
	Description         string                   `json:"description" yaml:"description"`
	Draft               bool                     `json:"draft" yaml:"draft"`
	DraftExists         bool                     `json:"draftExists" yaml:"draftExists"`
	DuplicationDisabled bool                     `json:"duplication_disabled" yaml:"duplication_disabled"`
	Featured            bool                     `json:"featured" yaml:"featured"`
	LastEdit            DataCardLastEdit         `json:"last_edit" yaml:"last_edit"`
	Locked              bool                     `json:"locked" yaml:"locked"`
	Modified            int                      `json:"modified" yaml:"modified"`
	NumberOfPlayers     int                      `json:"number_of_players" yaml:"number_of_players"`
	NumberOfPlays       int                      `json:"number_of_plays" yaml:"number_of_plays"`
	NumberOfQuestions   int                      `json:"number_of_questions" yaml:"number_of_questions"`
	QuestionTypes       []string                 `json:"question_types" yaml:"question_types"`
	SampleQuestions     []DataCardSampleQuestion `json:"sample_questions" yaml:"sample_questions"`
	Slug                string                   `json:"slug" yaml:"slug"`
	Sponsored           bool                     `json:"sponsored" yaml:"sponsored"`
	Title               string                   `json:"title" yaml:"title"`
	TotalFavourites     int                      `json:"total_favourites" yaml:"total_favourites"`
	Type                string                   `json:"type" yaml:"type"`
	UUID                string                   `json:"uuid" yaml:"uuid"`
	Visibility          int                      `json:"visibility" yaml:"visibility"`
	WriteProtection     bool                     `json:"writeProtection" yaml:"writeProtection"`
	YoungFeatured       bool                     `json:"young_featured" yaml:"young_featured"`
}

type DataCardAccess struct {
	Features []interface{} `json:"features" yaml:"features"`
}

type DataCardCoverMetadata struct {
	ID string `json:"id" yaml:"id"`
}

type DataCardLastEdit struct {
	EditTimestamp  int    `json:"editTimestamp" yaml:"editTimestamp"`
	EditorUserID   string `json:"editorUserId" yaml:"editorUserId"`
	EditorUsername string `json:"editorUsername" yaml:"editorUsername"`
}

type DataCardSampleQuestion struct {
	Image         string                              `json:"image" yaml:"image"`
	ImageMetadata DataCardSampleQuestionImageMetadata `json:"imageMetadata" yaml:"imageMetadata"`
	Title         string                              `json:"title" yaml:"title"`
	Type          string                              `json:"type" yaml:"type"`
}

type DataCardSampleQuestionImageMetadata struct {
	ContentType string                                  `json:"contentType" yaml:"contentType"`
	Crop        DataCardSampleQuestionImageMetadataCrop `json:"crop" yaml:"crop"`
	Effects     []interface{}                           `json:"effects" yaml:"effects"`
	Height      int                                     `json:"height" yaml:"height"`
	ID          string                                  `json:"id" yaml:"id"`
	Width       int                                     `json:"width" yaml:"width"`
}

type DataCardSampleQuestionImageMetadataCrop struct {
	Origin DataCardSampleQuestionImageMetadataCropOrigin `json:"origin" yaml:"origin"`
	Target DataCardSampleQuestionImageMetadataCropTarget `json:"target" yaml:"target"`
}

type DataCardSampleQuestionImageMetadataCropOrigin struct {
	X int `json:"x" yaml:"x"`
	Y int `json:"y" yaml:"y"`
}

type DataCardSampleQuestionImageMetadataCropTarget struct {
	X int `json:"x" yaml:"x"`
	Y int `json:"y" yaml:"y"`
}

type DataKahoot struct {
	Audience            string                  `json:"audience" yaml:"audience"`
	CompatibilityLevel  int                     `json:"compatibilityLevel" yaml:"compatibilityLevel"`
	Cover               string                  `json:"cover" yaml:"cover"`
	CoverMetadata       DataKahootCoverMetadata `json:"coverMetadata" yaml:"coverMetadata"`
	Created             int                     `json:"created" yaml:"created"`
	Creator             string                  `json:"creator" yaml:"creator"`
	CreatorPrimaryUsage string                  `json:"creator_primary_usage" yaml:"creator_primary_usage"`
	CreatorUsername     string                  `json:"creator_username" yaml:"creator_username"`
	Description         string                  `json:"description" yaml:"description"`
	FolderID            string                  `json:"folderId" yaml:"folderId"`
	Language            string                  `json:"language" yaml:"language"`
	LobbyVideo          DataKahootLobbyVideo    `json:"lobby_video" yaml:"lobby_video"`
	Metadata            DataKahootMetadata      `json:"metadata" yaml:"metadata"`
	Modified            int                     `json:"modified" yaml:"modified"`
	Questions           []DataKahootQuestion    `json:"questions" yaml:"questions"`
	QuizType            string                  `json:"quizType" yaml:"quizType"`
	Resources           string                  `json:"resources" yaml:"resources"`
	Slug                string                  `json:"slug" yaml:"slug"`
	Title               string                  `json:"title" yaml:"title"`
	Type                string                  `json:"type" yaml:"type"`
	UUID                string                  `json:"uuid" yaml:"uuid"`
	Visibility          int                     `json:"visibility" yaml:"visibility"`
}

type DataKahootCoverMetadata struct {
	ID string `json:"id" yaml:"id"`
}

type DataKahootLobbyVideo struct {
	Youtube DataKahootLobbyVideoYoutube `json:"youtube" yaml:"youtube"`
}

type DataKahootLobbyVideoYoutube struct {
	EndTime   int    `json:"endTime" yaml:"endTime"`
	FullURL   string `json:"fullUrl" yaml:"fullUrl"`
	ID        string `json:"id" yaml:"id"`
	Service   string `json:"service" yaml:"service"`
	StartTime int    `json:"startTime" yaml:"startTime"`
}

type DataKahootMetadata struct {
	Access                DataKahootMetadataAccess   `json:"access" yaml:"access"`
	DuplicationProtection bool                       `json:"duplicationProtection" yaml:"duplicationProtection"`
	LastEdit              DataKahootMetadataLastEdit `json:"lastEdit" yaml:"lastEdit"`
}

type DataKahootMetadataAccess struct {
	Features []interface{} `json:"features" yaml:"features"`
}

type DataKahootMetadataLastEdit struct {
	EditTimestamp  int    `json:"editTimestamp" yaml:"editTimestamp"`
	EditorUserID   string `json:"editorUserId" yaml:"editorUserId"`
	EditorUsername string `json:"editorUsername" yaml:"editorUsername"`
}

type DataKahootQuestion struct {
	Choices          []DataKahootQuestionChoice      `json:"choices" yaml:"choices"`
	Image            string                          `json:"image" yaml:"image"`
	ImageMetadata    DataKahootQuestionImageMetadata `json:"imageMetadata" yaml:"imageMetadata"`
	Layout           string                          `json:"layout" yaml:"layout"`
	NumberOfAnswers  int                             `json:"numberOfAnswers" yaml:"numberOfAnswers"`
	Points           bool                            `json:"points" yaml:"points"`
	PointsMultiplier int                             `json:"pointsMultiplier" yaml:"pointsMultiplier"`
	Question         string                          `json:"question" yaml:"question"`
	QuestionFormat   int                             `json:"questionFormat" yaml:"questionFormat"`
	Resources        string                          `json:"resources" yaml:"resources"`
	Time             int                             `json:"time" yaml:"time"`
	Type             string                          `json:"type" yaml:"type"`
	Video            DataKahootQuestionVideo         `json:"video" yaml:"video"`
}

type DataKahootQuestionChoice struct {
	Answer  string `json:"answer" yaml:"answer"`
	Correct bool   `json:"correct" yaml:"correct"`
}

type DataKahootQuestionImageMetadata struct {
	ContentType string                              `json:"contentType" yaml:"contentType"`
	Crop        DataKahootQuestionImageMetadataCrop `json:"crop" yaml:"crop"`
	Effects     []interface{}                       `json:"effects" yaml:"effects"`
	Height      int                                 `json:"height" yaml:"height"`
	ID          string                              `json:"id" yaml:"id"`
	Width       int                                 `json:"width" yaml:"width"`
}

type DataKahootQuestionImageMetadataCrop struct {
	Origin DataKahootQuestionImageMetadataCropOrigin `json:"origin" yaml:"origin"`
	Target DataKahootQuestionImageMetadataCropTarget `json:"target" yaml:"target"`
}

type DataKahootQuestionImageMetadataCropOrigin struct {
	X int `json:"x" yaml:"x"`
	Y int `json:"y" yaml:"y"`
}

type DataKahootQuestionImageMetadataCropTarget struct {
	X int `json:"x" yaml:"x"`
	Y int `json:"y" yaml:"y"`
}

type DataKahootQuestionVideo struct {
	EndTime   int    `json:"endTime" yaml:"endTime"`
	FullURL   string `json:"fullUrl" yaml:"fullUrl"`
	ID        string `json:"id" yaml:"id"`
	Service   string `json:"service" yaml:"service"`
	StartTime int    `json:"startTime" yaml:"startTime"`
}

