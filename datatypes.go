package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

////////////////////////////////////////////////
// From Positions
////////////////////////////////////////////////

// Receive
type ExerPosition struct {
	ImageSetID    string  `bson:"imageset"`
	Hardcoded     bool    `bson:"hardcoded"`
	HardcodedSecs float32 `bson:"hardcodedsecs"`
	MaxSecs       float32 `bson:"maxsecs"` // ?
	PercentSecs   float32 `bson:"percentsecs"`
}

type StrPosition struct {
	ImageSetID  string  `bson:"imageset"`
	PercentSecs float32 `bson:"percentsecs"`
}

type ExerciseSCR struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	BackendID      string             `bson:"backendID"`
	Name           string             `bson:"name"`
	Parent         string             `bson:"parent"`
	MaxSecs        float32            `bson:"maxsecs"`
	MinSecs        float32            `bson:"minsecs"`
	ImageSetID0    string             `bson:"imageset0"`
	PositionSlice1 []ExerPosition     `bson:"positions1"`
	PositionSlice2 []ExerPosition     `bson:"positions2"`
	SampleID       string             `bson:"sampleid"`
}

type DynamicStr struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	BackendID      string             `bson:"backendID"`
	Name           string             `bson:"name"`
	Secs           float32            `bson:"secs"`
	SeparateSets   bool               `bson:"separate"`
	PositionSlice1 []StrPosition      `bson:"positions1"`
	PositionSlice2 []StrPosition      `bson:"positions2"`
	SampleID       string             `bson:"sampleid"`
}

type StaticStr struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	BackendID   string             `bson:"backendID"`
	Name        string             `bson:"name"`
	ImageSetID1 string             `bson:"imageset1"`
	ImageSetID2 string             `bson:"imageset2"`
	SampleID    string             `bson:"sampleid"`
}

type TransitionMatrix struct {
	ID            primitive.ObjectID    `bson:"_id,omitempty"`
	FastMatrix    [11][11]TransitionRep `bson:"fastmatrix"`
	RegularMatrix [11][11]TransitionRep `bson:"regularmatrix"`
	SlowMatrix    [11][11]TransitionRep `bson:"slowmatrix"`
}

// Return
type Rep struct {
	Positions []string  `bson:"positions"`
	Times     []float32 `bson:"times"`
	FullTime  float32   `bson:"fulltime"`
}

type Set struct {
	RepSlice        []Rep   `bson:"reps"`
	RepSequence     []int   `bson:"repsequence"`
	RepCount        int     `bson:"repcount"`
	PositionInit    string  `bson:"positioninit"`
	PositionEnd     string  `bson:"positionend"`
	SeparateStretch bool    `bson:"sepstretch"`
	FullTime        float32 `bson:"fulltime"`
}

type WORound struct {
	SetSlice     []Set    `bson:"sets"`
	SetSequence  []int    `bson:"setsequence"`
	SetCount     int      `bson:"setcount"`
	FullTime     float32  `bson:"fulltime"`
	RestPerRound float32  `bson:"restround"`
	RestPerSet   float32  `bson:"restset"`
	ExerPerSet   float32  `bson:"exerset"`
	Type         string   `bson:"type"`
	Names        []string `bson:"names"`
	Reps         []int    `bson:"reps"`
	SplitPairs   [2]bool  `bson:"splitpairs"`
	SampleIDs    []string `bson:"samples"`
	RestPosition string   `bson:"restposition"`
}

type StretchWorkoutSCR struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	BackendID        string             `bson:"backendID"`
	DynamicSlice     []Set              `bson:"dynamics"`
	StaticSlice      []Set              `bson:"statics"`
	DynamicNames     []string           `bson:"dynamicnames"`
	StaticNames      []string           `bson:"staticnames"`
	DynamicSamples   []string           `bson:"dynamicsamples"`
	StaticSamples    []string           `bson:"staticsamples"`
	CongratsPosition string             `bson:"congratspos"`
	StandingPosition string             `bson:"standingpos"`
	RoundTime        float32            `bson:"roundtime"`
}

type WorkoutSCR struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	BackendID        string             `bson:"backendID"`
	DynamicSlice     []Set              `bson:"dynamics"`
	StaticSlice      []Set              `bson:"statics"`
	StaticTime       float32            `bson:"statictime"`
	DynamicTime      float32            `bson:"dynamictime"`
	DynamicRest      float32            `bson:"dynamicrest"`
	DynamicNames     []string           `bson:"dynamicnames"`
	StaticNames      []string           `bson:"staticnames"`
	DynamicSamples   []string           `bson:"dynamicsamples"`
	StaticSamples    []string           `bson:"staticsamples"`
	CongratsPosition string             `bson:"congratspos"`
	StandingPosition string             `bson:"standingpos"`
	Exercises        [9]WORound         `bson:"exercises"`
}

type TransitionRep struct {
	ImageSetIDs []string  `bson:"imagesetids"`
	Times       []float32 `bson:"times"`
	FullTime    float32   `bson:"fulltime"`
}

////////////////////////////////////////////////
// From Backend
////////////////////////////////////////////////

// Receive
type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	Name              string             `bson:"name"`
	Username          string             `bson:"username"`
	Paying            bool               `bson:"paying"`
	Provider          string             `bson:"provider"`
	Level             float32            `bson:"level"`
	BannedExercises   []string           `bson:"bannedExer"`
	BannedStretches   []string           `bson:"bannedStr"`
	BannedParts       []int              `bson:"bannedParts"`
	PlyoTolerance     int                `bson:"plyoToler"`
	ExerFavoriteRates map[string]float32 `bson:"exerfavs"`
	ExerModifications map[string]float32 `bson:"exermods"`
	TypeModifications map[string]float32 `bson:"typemods"`
	RoundEndurance    map[int]float32    `bson:"roundendur"`
	TimeEndurance     map[int]float32    `bson:"timeendur"`
	PushupSetting     string             `bson:"pushsetting"`
	LastMinutes       float32            `bson:"lastmins"`
	LastDifficulty    int                `bson:"lastdiff"`
	Assessed          bool               `bson:"assessed"`
}

type Exercise struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	Parent       string             `bson:"parent"`
	MinLevel     float32            `bson:"minlevel"`
	MaxLevel     float32            `bson:"maxlevel"`
	MinReps      int                `bson:"minreps"`
	PlyoRating   int                `bson:"plyorating"`
	StartQuality float32            `bson:"startquality"`
	BodyParts    []int              `bson:"bodyparts"`
	RepVars      [3]float32         `bson:"repvars"`
	InSplits     bool               `bson:"insplits"`
	InPairs      bool               `bson:"inpairs"`
	UnderCombos  bool               `bson:"undercombos"`
	CardioRating float32            `bson:"cardiorating"`
	PushupType   string             `bson:"pushuptype"`
	GeneralType  []string           `bson:"generaltype"`
}

type Stretch struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	MinLevel  float32            `bson:"minlevel"`
	Status    string             `bson:"status"`
	BodyParts []int              `bson:"bodyparts"`
	InPairs   bool               `bson:"inpairs"`
}

type TypeMatrix struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Matrix [11][11]float32    `bson:"matrix"`
}

// Return
type Workout struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Name            string             `bson:"name"`
	UserID          string             `bson:"userid"`
	Username        string             `bson:"username"`
	Date            primitive.DateTime `bson:"date"`
	Status          string             `bson:"status"`
	Minutes         float32            `bson:"minutes"`
	StretchTimes    StretchTimes       `bson:"stretchtimes"`
	PausedTime      float32            `bson:"paused"`
	LevelAtStart    float32            `bson:"level"`
	Difficulty      int                `bson:"difficulty"`
	Dynamics        []string           `bson:"dynamics"`
	Statics         []string           `bson:"statics"`
	Exercises       [9]WorkoutRound    `bson:"exercises"`
	CardioRatings   [9]float32         `bson:"cardioratings"`
	CardioRating    float32            `bson:"cardiorating"`
	GeneralTypeVals [3]float32         `bson:"gentypevals"`
	IsIntro         bool               `bson:"intro"`
}

type StretchTimes struct {
	DynamicPerSet []float32 `bson:"dynamicperset"`
	StaticPerSet  []float32 `bson:"staticperset"`
	DynamicSets   int       `bson:"dynamicsets"`
	StaticSets    int       `bson:"staticsets"`
	DynamicRest   float32   `bson:"dynamicrest"`
	FullRound     float32   `bson:"fullround"`
}

type ExerciseTimes struct {
	ExercisePerSet float32 `bson:"exerciseperset"`
	RestPerSet     float32 `bson:"restperset"`
	Sets           int     `bson:"sets"`
	RestPerRound   float32 `bson:"restperround"`
	FullRound      float32 `bson:"fullround"`
	ComboExers     int     `bson:"comboexers"`
}

type WorkoutRound struct {
	ExerciseIDs []string      `bson:"exerids"`
	Reps        []float32     `bson:"reps"`
	Pairs       []bool        `bson:"pairs"`
	Status      string        `bson:"status"`
	Times       ExerciseTimes `bson:"times"`
	Rating      float32       `bson:"rating"`
}

type StretchWorkout struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name"`
	UserID       string             `bson:"userid"`
	Date         primitive.DateTime `bson:"date"`
	Status       string             `bson:"status"`
	StretchTimes StretchTimes       `bson:"stretchtimes"`
	LevelAtStart float32            `bson:"level"`
	PausedTime   float32            `bson:"paused"`
	Minutes      float32            `bson:"minutes"`
	Dynamics     []string           `bson:"dynamics"`
	Statics      []string           `bson:"statics"`
}

type AnyWorkout interface {
	Display()
}

func (t Workout) Display() {
	fmt.Println("Workout: ")
	fmt.Println(t)
}

func (t StretchWorkout) Display() {
	fmt.Println("Stretch Workout: ")
	fmt.Println(t)
}

type ReceiveWorkoutJSON struct {
	Exercises        []Exercise
	Stretches        []Stretch
	TypeMatrix       TypeMatrix
	User             User
	PastWOs          []Workout
	ExerciseSCRs     []ExerciseSCR
	DynamicSCRs      []DynamicStr
	StaticSCRs       []StaticStr
	TransitionMatrix TransitionMatrix
}

type ReceiveStrWorkoutJSON struct {
	Stretches   []Stretch
	User        User
	DynamicSCRs []DynamicStr
	StaticSCRs  []StaticStr
}
