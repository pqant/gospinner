package gospinner

import "time"

// Symbols for the finishing actions
const (
	successSymbol         = "✔"
	failureSymbol         = "✖"
	warningSymbol         = "⚠"
	questionMarkRedSymbol = "❓"
	exclamationMarkSymbol = "❗"
	smileySymbol          = "㋡"
	starMarkSymbol        = "★"
	questionMarkSymbol    = "?"
	aliveSymbol           = "෴"
)

// AnimationKind represents the kind of the animation
type AnimationKind int

const (
	Ball                AnimationKind = iota
	Column
	Slash
	Square
	Triangle
	Dots
	Dots2
	Pipe
	SimpleDots
	SimpleDotsScrolling
	GrowVertical
	GrowHorizontal
	Arrow
	BouncingBar
	BouncingBall
	Pong
	ProgressBar
	World
)

// Animation represents an animation with frames and speed (recommended)
type Animation struct {
	interval time.Duration
	frames   []string
}

var animations = map[AnimationKind]Animation{
	Ball:                Animation{interval: 80 * time.Millisecond, frames: []string{"◐", "◓", "◑", "◒"}},
	Column:              Animation{interval: 80 * time.Millisecond, frames: []string{"☰", "☱", "☳", "☷", "☶", "☴"}},
	Slash:               Animation{interval: 130 * time.Millisecond, frames: []string{"-", "\\", "|", "/"}},
	Square:              Animation{interval: 110 * time.Millisecond, frames: []string{"▖", "▘", "▝", "▗"}},
	Triangle:            Animation{interval: 80 * time.Millisecond, frames: []string{"◢", "◣", "◤", "◥"}},
	Dots:                Animation{interval: 80 * time.Millisecond, frames: []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}},
	Dots2:               Animation{interval: 80 * time.Millisecond, frames: []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}},
	Pipe:                Animation{interval: 100 * time.Millisecond, frames: []string{"┤", "┘", "┴", "└", "├", "┌", "┬", "┐"}},
	SimpleDots:          Animation{interval: 400 * time.Millisecond, frames: []string{".  ", ".. ", "...", "   "}},
	SimpleDotsScrolling: Animation{interval: 200 * time.Millisecond, frames: []string{".  ", ".. ", "...", " ..", "  .", "   "}},
	GrowVertical:        Animation{interval: 120 * time.Millisecond, frames: []string{"▁", "▃", "▄", "▅", "▆", "▇", "▆", "▅", "▄", "▃"}},
	GrowHorizontal:      Animation{interval: 120 * time.Millisecond, frames: []string{"▏", "▎", "▍", "▌", "▋", "▊", "▉", "▊", "▋", "▌", "▍", "▎"}},
	Arrow:               Animation{interval: 120 * time.Millisecond, frames: []string{"▹▹▹▹▹", "▸▹▹▹▹", "▹▸▹▹▹", "▹▹▸▹▹", "▹▹▹▸▹", "▹▹▹▹▸"}},
	BouncingBar:         Animation{interval: 80 * time.Millisecond, frames: []string{"[    ]", "[   =]", "[  ==]", "[ ===]", "[====]", "[=== ]", "[==  ]", "[=   ]"}},
	BouncingBall:        Animation{interval: 80 * time.Millisecond, frames: []string{"( ●    )", "(  ●   )", "(   ●  )", "(    ● )", "(     ●)", "(    ● )", "(   ●  )", "(  ●   )", "( ●    )", "(●     )"}},
	Pong:                Animation{interval: 80 * time.Millisecond, frames: []string{"▐⠂       ▌", "▐⠈       ▌", "▐ ⠂      ▌", "▐ ⠠      ▌", "▐  ⡀     ▌", "▐  ⠠     ▌", "▐   ⠂    ▌", "▐   ⠈    ▌", "▐    ⠂   ▌", "▐    ⠠   ▌", "▐     ⡀  ▌", "▐     ⠠  ▌", "▐      ⠂ ▌", "▐      ⠈ ▌", "▐       ⠂▌", "▐       ⠠▌", "▐       ⡀▌", "▐      ⠠ ▌", "▐      ⠂ ▌", "▐     ⠈  ▌", "▐     ⠂  ▌", "▐    ⠠   ▌", "▐    ⡀   ▌", "▐   ⠠    ▌", "▐   ⠂    ▌", "▐  ⠈     ▌", "▐  ⠂     ▌", "▐ ⠠      ▌", "▐ ⡀      ▌", "▐⠠       ▌"}},
	ProgressBar:         Animation{interval: 120 * time.Millisecond, frames: []string{"▒▒▒▒▒▒▒▒▒▒", "█▒▒▒▒▒▒▒▒▒", "███▒▒▒▒▒▒▒", "█████▒▒▒▒▒", "███████▒▒▒", "██████████"}},
	World:               Animation{interval: 80 * time.Millisecond, frames: []string{"🌍", "🌎", "🌏"}},
}

// new comments
// has been added!
// v.2
var CharSets = map[int][]string{
	0:  {"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"},
	1:  {"▁", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃", "▁"},
	2:  {"▖", "▘", "▝", "▗"},
	3:  {"┤", "┘", "┴", "└", "├", "┌", "┬", "┐"},
	4:  {"◢", "◣", "◤", "◥"},
	5:  {"◰", "◳", "◲", "◱"},
	6:  {"◴", "◷", "◶", "◵"},
	7:  {"◐", "◓", "◑", "◒"},
	8:  {".", "o", "O", "@", "*"},
	9:  {"|", "/", "-", "\\"},
	10: {"◡◡", "⊙⊙", "◠◠"},
	11: {"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"},
	12: {">))'>", " >))'>", "  >))'>", "   >))'>", "    >))'>", "   <'((<", "  <'((<", " <'((<"},
	13: {"⠁", "⠂", "⠄", "⡀", "⢀", "⠠", "⠐", "⠈"},
	14: {"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
	15: {"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"},
	16: {"▉", "▊", "▋", "▌", "▍", "▎", "▏", "▎", "▍", "▌", "▋", "▊", "▉"},
	17: {"■", "□", "▪", "▫"},
	18: {"←", "↑", "→", "↓"},
	19: {"╫", "╪"},
	20: {"⇐", "⇖", "⇑", "⇗", "⇒", "⇘", "⇓", "⇙"},
	21: {"⠁", "⠁", "⠉", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠤", "⠄", "⠄", "⠤", "⠠", "⠠", "⠤", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋", "⠉", "⠈", "⠈"},
	22: {"⠈", "⠉", "⠋", "⠓", "⠒", "⠐", "⠐", "⠒", "⠖", "⠦", "⠤", "⠠", "⠠", "⠤", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋", "⠉", "⠈"},
	23: {"⠁", "⠉", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠤", "⠄", "⠄", "⠤", "⠴", "⠲", "⠒", "⠂", "⠂", "⠒", "⠚", "⠙", "⠉", "⠁"},
	24: {"⠋", "⠙", "⠚", "⠒", "⠂", "⠂", "⠒", "⠲", "⠴", "⠦", "⠖", "⠒", "⠐", "⠐", "⠒", "⠓", "⠋"},
	25: {"ｦ", "ｧ", "ｨ", "ｩ", "ｪ", "ｫ", "ｬ", "ｭ", "ｮ", "ｯ", "ｱ", "ｲ", "ｳ", "ｴ", "ｵ", "ｶ", "ｷ", "ｸ", "ｹ", "ｺ", "ｻ", "ｼ", "ｽ", "ｾ", "ｿ", "ﾀ", "ﾁ", "ﾂ", "ﾃ", "ﾄ", "ﾅ", "ﾆ", "ﾇ", "ﾈ", "ﾉ", "ﾊ", "ﾋ", "ﾌ", "ﾍ", "ﾎ", "ﾏ", "ﾐ", "ﾑ", "ﾒ", "ﾓ", "ﾔ", "ﾕ", "ﾖ", "ﾗ", "ﾘ", "ﾙ", "ﾚ", "ﾛ", "ﾜ", "ﾝ"},
	26: {".", "..", "..."},
	27: {"▁", "▂", "▃", "▄", "▅", "▆", "▇", "█", "▉", "▊", "▋", "▌", "▍", "▎", "▏", "▏", "▎", "▍", "▌", "▋", "▊", "▉", "█", "▇", "▆", "▅", "▄", "▃", "▂", "▁"},
	28: {".", "o", "O", "°", "O", "o", "."},
	29: {"+", "x"},
	30: {"v", "<", "^", ">"},
	31: {">>--->", " >>--->", "  >>--->", "   >>--->", "    >>--->", "    <---<<", "   <---<<", "  <---<<", " <---<<", "<---<<"},
	32: {"|", "||", "|||", "||||", "|||||", "|||||||", "||||||||", "|||||||", "||||||", "|||||", "||||", "|||", "||", "|"},
	33: {"[          ]", "[=         ]", "[==        ]", "[===       ]", "[====      ]", "[=====     ]", "[======    ]", "[=======   ]", "[========  ]", "[========= ]", "[==========]"},
	34: {"(*---------)", "(-*--------)", "(--*-------)", "(---*------)", "(----*-----)", "(-----*----)", "(------*---)", "(-------*--)", "(--------*-)", "(---------*)"},
	35: {"█▒▒▒▒▒▒▒▒▒", "███▒▒▒▒▒▒▒", "█████▒▒▒▒▒", "███████▒▒▒", "██████████"},
	36: {"[                    ]", "[=>                  ]", "[===>                ]", "[=====>              ]", "[======>             ]", "[========>           ]", "[==========>         ]", "[============>       ]", "[==============>     ]", "[================>   ]", "[==================> ]", "[===================>]"},
	39: {"🌍", "🌎", "🌏"},
	40: {"◜", "◝", "◞", "◟"},
	41: {"⬒", "⬔", "⬓", "⬕"},
	42: {"⬖", "⬘", "⬗", "⬙"},
	43: {"[>>>          >]", "[]>>>>        []", "[]  >>>>      []", "[]    >>>>    []", "[]      >>>>  []", "[]        >>>>[]", "[>>          >>]"},
}
