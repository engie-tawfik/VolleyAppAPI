package constants

const (
	AttackPoint     string = "ap"
	AttackNeutral   string = "an"
	AttackError     string = "ae"
	BlockPoint      string = "bp"
	BlockNeutral    string = "bn"
	BlockError      string = "be"
	ServePoint      string = "sp"
	ServeNeutral    string = "sn"
	ServeError      string = "se"
	OpponentError   string = "oe"
	OpponentAttack  string = "oa"
	OpponentBlock   string = "ob"
	OpponentService string = "os"
	Error           string = "e"
	RollBack        string = "rb"
)

var SetActions = []string{
	AttackPoint,
	AttackNeutral,
	AttackError,
	BlockPoint,
	BlockNeutral,
	BlockError,
	ServePoint,
	ServeNeutral,
	ServeError,
	OpponentError,
	OpponentAttack,
	OpponentBlock,
	OpponentService,
	Error,
	RollBack,
}
