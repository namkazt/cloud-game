package config

import (
	"flag"
	"time"
)

const defaultoverlord = "ws://localhost:9000/wso"
const DefaultSTUNTURN = `[{"urls":"stun:stun-turn.webgame2d.com:3478"},{"urls":"turn:stun-turn.webgame2d.com:3478","username":"root","credential":"root"}]`
const CODEC_VP8 = "VP8"
const CODEC_H264 = "H264"

var IsDebug = flag.Bool("debug", false, "Is game running in debug mode?")
var OverlordHost = flag.String("overlordhost", defaultoverlord, "Specify the path for overlord. If the flag is `overlord`, the server will be run as overlord")
var Port = flag.String("port", "8000", "Port of the game")
var IsMonitor = flag.Bool("monitor", false, "Turn on monitor")
var FrontendSTUNTURN = flag.String("stunturn", DefaultSTUNTURN, "Frontend STUN TURN servers")
var Mode = flag.String("mode", "dev", "Environment")
var IsRetro = flag.Bool("isretro", true, "Is retro")
var StunTurnTemplate = `[{"urls":"stun:stun.l.google.com:19302"},{"urls":"stun:%s:3478"},{"urls":"turn:%s:3478","username":"root","credential":"root"}]`

var WSWait = 20 * time.Second
var MatchWorkerRandom = false
var ProdEnv = "prod"

const NumKeys = 10

var FileTypeToEmulator = map[string]string{
	"gba": "gba",
	"cue": "pcsx",
	"zip": "mame",
	"nes": "nes",
	"smc": "snes",
	"sfc": "snes",
	"swc": "snes",
	"fig": "snes",
	"bs":  "snes",
}

// There is no good way to determine main width and height of the emulator.
// When game run, frame width and height can scale abnormally.
type EmulatorMeta struct {
	Path            string
	Width           int
	Height          int
	AudioSampleRate int
	Fps             int
}

var EmulatorConfig = map[string]EmulatorMeta{
	"gba": EmulatorMeta{
		Path:   "libretro/cores/mgba_libretro.so",
		Width:  240,
		Height: 160,
	},
	"pcsx": EmulatorMeta{
		Path:   "libretro/cores/mednafen_psx_libretro.so",
		Width:  350,
		Height: 240,
	},
	"nes": EmulatorMeta{
		Path:   "libretro/cores/nestopia_libretro.so",
		Width:  256,
		Height: 240,
	},
	"snes": EmulatorMeta{
		Path:   "libretro/cores/mednafen_snes_libretro.so",
		Width:  256,
		Height: 224,
	},
	"mame": EmulatorMeta{
		Path:   "libretro/cores/mame2016_libretro.so",
		Width:  0,
		Height: 0,
	},
}
