package tgbotapi

import (
	"fmt"
	"sort"
	"strings"
)

// Language is a canonical Prism/libprisma language tag (lowercase), e.g. "go", "python", "markup".
// Use these constants when rendering <pre><code class="language-...">...</code></pre> (HTML)
// and fenced code blocks with language tags in MarkdownV2 (```go).
type Language string

// Below are canonical Language constants derived from libprisma's "Supported languages" table.
// Each constant's value equals the canonical language key used by libprisma/Prism.
// For aliases (e.g., "html", "xml" → "markup"), use ResolveLanguage.
const (
	// Markup: html, xml, svg, mathml, ssml, atom, rss
	LangMarkup Language = "markup"
	// CSS
	LangCSS Language = "css"
	// C-like
	LangCLike Language = "clike"
	// Regex
	LangRegex Language = "regex"
	// JavaScript (alias: js)
	LangJavaScript Language = "javascript"
	// ABAP
	LangABAP Language = "abap"
	// ABNF
	LangABNF Language = "abnf"
	// ActionScript
	LangActionScript Language = "actionscript"
	// Ada
	LangAda Language = "ada"
	// Agda
	LangAgda Language = "agda"
	// AL
	LangAL Language = "al"
	// ANTLR4 (alias: g4)
	LangANTLR4 Language = "antlr4"
	// Apache Configuration
	LangApacheConf Language = "apacheconf"
	// SQL
	LangSQL Language = "sql"
	// Apex
	LangApex Language = "apex"
	// APL
	LangAPL Language = "apl"
	// AppleScript
	LangAppleScript Language = "applescript"
	// AQL
	LangAQL Language = "aql"
	// C
	LangC Language = "c"
	// C++
	LangCPP Language = "cpp"
	// Arduino (alias: ino)
	LangArduino Language = "arduino"
	// ARFF
	LangARFF Language = "arff"
	// ARM Assembly (alias: arm-asm)
	LangARMAsm Language = "armasm"
	// Bash (aliases: sh, shell)
	LangBash Language = "bash"
	// YAML (alias: yml)
	LangYAML Language = "yaml"
	// Markdown (alias: md)
	LangMarkdown Language = "markdown"
	// Arturo (alias: art)
	LangArturo Language = "arturo"
	// AsciiDoc (alias: adoc)
	LangAsciiDoc Language = "asciidoc"
	// C# (aliases: cs, dotnet)
	LangCSharp Language = "csharp"
	// ASP.NET (C#)
	LangASPNet Language = "aspnet"
	// 6502 Assembly
	LangAsm6502 Language = "asm6502"
	// Atmel AVR Assembly
	LangAsmAtmel Language = "asmatmel"
	// AutoHotkey
	LangAutoHotkey Language = "autohotkey"
	// AutoIt
	LangAutoIt Language = "autoit"
	// AviSynth (alias: avs)
	LangAviSynth Language = "avisynth"
	// Avro IDL (alias: avdl)
	LangAvroIDL Language = "avro-idl"
	// AWK (alias: gawk)
	LangAWK Language = "awk"
	// BASIC
	LangBASIC Language = "basic"
	// Batch
	LangBatch Language = "batch"
	// BBcode (alias: shortcode)
	LangBBCode Language = "bbcode"
	// BBj
	LangBBj Language = "bbj"
	// Bicep
	LangBicep Language = "bicep"
	// Birb
	LangBirb Language = "birb"
	// Bison
	LangBison Language = "bison"
	// BNF (alias: rbnf)
	LangBNF Language = "bnf"
	// BQN
	LangBQN Language = "bqn"
	// Brainfuck
	LangBrainfuck Language = "brainfuck"
	// BrightScript
	LangBrightScript Language = "brightscript"
	// Bro
	LangBro Language = "bro"
	// CFScript (alias: cfc)
	LangCFScript Language = "cfscript"
	// ChaiScript
	LangChaiScript Language = "chaiscript"
	// CIL
	LangCIL Language = "cil"
	// Cilk/C
	LangCilkC Language = "cilkc"
	// Cilk/C++ (aliases: cilkcpp, cilk-cpp, cilk)
	LangCilkCPP Language = "cilkcpp"
	// Clojure
	LangClojure Language = "clojure"
	// CMake
	LangCMake Language = "cmake"
	// COBOL
	LangCOBOL Language = "cobol"
	// CoffeeScript (alias: coffee)
	LangCoffeeScript Language = "coffeescript"
	// Concurnas (alias: conc)
	LangConcurnas Language = "concurnas"
	// Content-Security-Policy
	LangCSP Language = "csp"
	// Cooklang
	LangCooklang Language = "cooklang"
	// Ruby (alias: rb)
	LangRuby Language = "ruby"
	// Crystal
	LangCrystal Language = "crystal"
	// CSV
	LangCSV Language = "csv"
	// CUE
	LangCUE Language = "cue"
	// Cypher
	LangCypher Language = "cypher"
	// D
	LangD Language = "d"
	// Dart
	LangDart Language = "dart"
	// DataWeave
	LangDataWeave Language = "dataweave"
	// DAX
	LangDAX Language = "dax"
	// Dhall
	LangDhall Language = "dhall"
	// Diff
	LangDiff Language = "diff"
	// Markup templating
	LangMarkupTmpl Language = "markup-templating"
	// Django/Jinja2 (alias: jinja2)
	LangDjango Language = "django"
	// DNS zone file (alias: dns-zone)
	LangDNSZone Language = "dns-zone-file"
	// Docker (alias: dockerfile)
	LangDocker Language = "docker"
	// DOT (Graphviz) (alias: gv)
	LangDOT Language = "dot"
	// EBNF
	LangEBNF Language = "ebnf"
	// EditorConfig
	LangEditorConfig Language = "editorconfig"
	// Eiffel
	LangEiffel Language = "eiffel"
	// EJS (alias: eta)
	LangEJS Language = "ejs"
	// Elixir
	LangElixir Language = "elixir"
	// Elm
	LangElm Language = "elm"
	// Lua
	LangLua Language = "lua"
	// Embedded Lua templating
	LangETLua Language = "etlua"
	// ERB
	LangERB Language = "erb"
	// Erlang
	LangErlang Language = "erlang"
	// Excel Formula (aliases: xlsx, xls)
	LangExcelFormula Language = "excel-formula"
	// F#
	LangFSharp Language = "fsharp"
	// Factor
	LangFactor Language = "factor"
	// False
	LangFalse Language = "false"
	// Fift
	LangFift Language = "fift"
	// Firestore security rules
	LangFirestoreRules Language = "firestore-security-rules"
	// Flow
	LangFlow Language = "flow"
	// Fortran
	LangFortran Language = "fortran"
	// FreeMarker Template Language
	LangFTL Language = "ftl"
	// FunC
	LangFunC Language = "func"
	// GameMaker Language (alias: gamemakerlanguage)
	LangGML Language = "gml"
	// GAP (CAS)
	LangGAP Language = "gap"
	// G-code
	LangGCode Language = "gcode"
	// GDScript
	LangGDScript Language = "gdscript"
	// GEDCOM
	LangGEDCOM Language = "gedcom"
	// gettext (alias: po)
	LangGettext Language = "gettext"
	// Git
	LangGit Language = "git"
	// GLSL
	LangGLSL Language = "glsl"
	// GN (alias: gni)
	LangGN Language = "gn"
	// GNU Linker Script (alias: ld)
	LangLinkerScript Language = "linker-script"
	// Go
	LangGo Language = "go"
	// Go module (alias: go-mod)
	LangGoModule Language = "go-module"
	// Gradle
	LangGradle Language = "gradle"
	// GraphQL
	LangGraphQL Language = "graphql"
	// Groovy
	LangGroovy Language = "groovy"
	// Less
	LangLess Language = "less"
	// Sass (SCSS)
	LangSCSS Language = "scss"
	// Textile
	LangTextile Language = "textile"
	// Haml
	LangHaml Language = "haml"
	// Handlebars (aliases: hbs, mustache)
	LangHandlebars Language = "handlebars"
	// Haskell (alias: hs)
	LangHaskell Language = "haskell"
	// Haxe
	LangHaxe Language = "haxe"
	// HCL
	LangHCL Language = "hcl"
	// HLSL
	LangHLSL Language = "hlsl"
	// Hoon
	LangHoon Language = "hoon"
	// HTTP Public-Key-Pins
	LangHPKP Language = "hpkp"
	// HTTP Strict-Transport-Security
	LangHSTS Language = "hsts"
	// JSON (alias: webmanifest)
	LangJSON Language = "json"
	// URI (alias: url)
	LangURI Language = "uri"
	// HTTP
	LangHTTP Language = "http"
	// IchigoJam
	LangIchigoJam Language = "ichigojam"
	// Icon
	LangIcon Language = "icon"
	// ICU Message Format
	LangICUMessage Language = "icu-message-format"
	// Idris (alias: idr)
	LangIdris Language = "idris"
	// .ignore (aliases: gitignore, hgignore, npmignore)
	LangIgnore Language = "ignore"
	// Inform 7
	LangInform7 Language = "inform7"
	// Ini
	LangINI Language = "ini"
	// Io
	LangIo Language = "io"
	// J
	LangJ Language = "j"
	// Java
	LangJava Language = "java"
	// Scala
	LangScala Language = "scala"
	// PHP
	LangPHP Language = "php"
	// JavaDoc-like
	LangJavadocLike Language = "javadoclike" // JavaDoc
	LangJavadoc     Language = "javadoc"
	// Java stack trace
	LangJavaStack Language = "javastacktrace"
	// Jolie
	LangJolie Language = "jolie"
	// JQ
	LangJQ Language = "jq"
	// TypeScript (alias: ts)
	LangTypeScript Language = "typescript"
	// JSDoc
	LangJSDoc Language = "jsdoc"
	// N4JS (alias: n4jsd)
	LangN4JS Language = "n4js"
	// JSON5
	LangJSON5 Language = "json5"
	// JSONP
	LangJSONP Language = "jsonp"
	// JS stack trace
	LangJSStack Language = "jsstacktrace"
	// Julia
	LangJulia Language = "julia"
	// Keepalived Configure
	LangKeepalived Language = "keepalived"
	// Keyman
	LangKeyman Language = "keyman"
	// Kotlin (aliases: kt, kts)
	LangKotlin Language = "kotlin"
	// Kusto
	LangKusto Language = "kusto"
	// LaTeX (aliases: tex, context)
	LangLaTeX Language = "latex"
	// Latte
	LangLatte Language = "latte"
	// Scheme
	LangScheme Language = "scheme"
	// LilyPond (alias: ly)
	LangLilyPond Language = "lilypond"
	// Liquid
	LangLiquid Language = "liquid"
	// Lisp (aliases: emacs, elisp, emacs-lisp)
	LangLisp Language = "lisp"
	// LiveScript
	LangLiveScript Language = "livescript"
	// LLVM IR
	LangLLVM Language = "llvm"
	// Log file
	LangLog Language = "log"
	// LOLCODE
	LangLOLCODE Language = "lolcode"
	// Magma (CAS)
	LangMagma Language = "magma"
	// Makefile
	LangMakefile Language = "makefile"
	// Mata
	LangMata Language = "mata"
	// MATLAB
	LangMATLAB Language = "matlab"
	// MAXScript
	LangMAXScript Language = "maxscript"
	// MEL
	LangMEL Language = "mel"
	// Mermaid
	LangMermaid Language = "mermaid"
	// METAFONT
	LangMETAFONT Language = "metafont"
	// Mizar
	LangMizar Language = "mizar"
	// MongoDB
	LangMongoDB Language = "mongodb"
	// Monkey
	LangMonkey Language = "monkey"
	// MoonScript (alias: moon)
	LangMoonScript Language = "moonscript"
	// N1QL
	LangN1QL Language = "n1ql"
	// Nand To Tetris HDL
	LangNand2Tetris Language = "nand2tetris-hdl"
	// Naninovel Script (alias: nani)
	LangNaninovel Language = "naniscript"
	// NASM
	LangNASM Language = "nasm"
	// NEON
	LangNEON Language = "neon"
	// Nevod
	LangNevod Language = "nevod"
	// nginx
	LangNginx Language = "nginx"
	// Nim
	LangNim Language = "nim"
	// Nix
	LangNix Language = "nix"
	// NSIS
	LangNSIS Language = "nsis"
	// Objective-C (alias: objc)
	LangObjectiveC Language = "objectivec"
	// OCaml
	LangOCaml Language = "ocaml"
	// Odin
	LangOdin Language = "odin"
	// OpenCL
	LangOpenCL Language = "opencl"
	// OpenQasm (alias: qasm)
	LangOpenQasm Language = "openqasm"
	// Oz
	LangOz Language = "oz"
	// PARI/GP
	LangPARIGP Language = "parigp"
	// Parser
	LangParser Language = "parser"
	// Pascal (alias: objectpascal)
	LangPascal Language = "pascal"
	// Pascaligo
	LangPascaligo Language = "pascaligo"
	// PATROL Scripting Language
	LangPSL Language = "psl"
	// PC-Axis (alias: px)
	LangPCAxis Language = "pcaxis"
	// PeopleCode (alias: pcode)
	LangPeopleCode Language = "peoplecode"
	// Perl
	LangPerl Language = "perl"
	// PHPDoc
	LangPHPDoc Language = "phpdoc"
	// PlantUML (alias: plantuml)
	LangPlantUML Language = "plant-uml"
	// PL/SQL
	LangPLSQL Language = "plsql"
	// PowerQuery (aliases: pq, mscript)
	LangPowerQuery Language = "powerquery"
	// PowerShell
	LangPowerShell Language = "powershell"
	// Processing
	LangProcessing Language = "processing"
	// Prolog
	LangProlog Language = "prolog"
	// PromQL
	LangPromQL Language = "promql"
	// .properties
	LangProperties Language = "properties"
	// Protocol Buffers
	LangProtoBuf Language = "protobuf"
	// Stylus
	LangStylus Language = "stylus"
	// Twig
	LangTwig Language = "twig"
	// Pug
	LangPug Language = "pug"
	// Puppet
	LangPuppet Language = "puppet"
	// PureBasic (alias: pbfasm)
	LangPureBasic Language = "purebasic"
	// Python (alias: py)
	LangPython Language = "python"
	// Q# (alias: qs)
	LangQSharp Language = "qsharp"
	// Q (kdb+)
	LangQ Language = "q"
	// QML
	LangQML Language = "qml"
	// Qore
	LangQore Language = "qore"
	// R
	LangR Language = "r"
	// Racket (alias: rkt)
	LangRacket Language = "racket"
	// Razor C# (alias: razor)
	LangRazorCS Language = "cshtml"
	// React JSX
	LangJSX Language = "jsx"
	// React TSX
	LangTSX Language = "tsx"
	// Reason
	LangReason Language = "reason"
	// Rego
	LangRego Language = "rego"
	// Ren'py (alias: rpy)
	LangRenPy Language = "renpy"
	// ReScript (alias: res)
	LangReScript Language = "rescript"
	// reST (reStructuredText)
	LangReST Language = "rest"
	// Rip
	LangRip Language = "rip"
	// Roboconf
	LangRoboconf Language = "roboconf"
	// Robot Framework (alias: robot)
	LangRobot Language = "robotframework"
	// Rust
	LangRust Language = "rust"
	// SAS
	LangSAS Language = "sas"
	// Sass (Sass)
	LangSass Language = "sass"
	// Shell session (aliases: sh-session, shellsession)
	LangShellSession Language = "shell-session"
	// Smali
	LangSmali Language = "smali"
	// Smalltalk
	LangSmalltalk Language = "smalltalk"
	// Smarty
	LangSmarty Language = "smarty"
	// SML (alias: smlnj)
	LangSML Language = "sml"
	// Solidity (Ethereum) (alias: sol)
	LangSolidity Language = "solidity"
	// Solution file (alias: sln)
	LangSolution Language = "solution-file"
	// Soy (Closure Template)
	LangSoy Language = "soy"
	// Splunk SPL
	LangSplunkSPL Language = "splunk-spl"
	// SQF (Arma 3)
	LangSQF Language = "sqf"
	// Squirrel
	LangSquirrel Language = "squirrel"
	// Stan
	LangStan Language = "stan"
	// Stata Ado
	LangStata Language = "stata"
	// Structured Text (IEC 61131-3)
	LangIECST Language = "iecst"
	// SuperCollider (alias: sclang)
	LangSuperCollider Language = "supercollider"
	// Swift
	LangSwift Language = "swift"
	// Systemd configuration file
	LangSystemd Language = "systemd"
	// Tact
	LangTact Language = "tact"
	// T4 templating
	LangT4Tmpl Language = "t4-templating"
	// T4 Text Templates (C#) (alias: t4)
	LangT4CS Language = "t4-cs"
	// VB.Net
	LangVBNet Language = "vbnet"
	// T4 Text Templates (VB)
	LangT4VB Language = "t4-vb"
	// TAP
	LangTAP Language = "tap"
	// Tcl
	LangTcl Language = "tcl"
	// Template Toolkit 2
	LangTT2 Language = "tt2"
	// TOML
	LangTOML Language = "toml"
	// Tremor (aliases: trickle, troy)
	LangTremor Language = "tremor"
	// Type Language
	LangTL Language = "tl"
	// Type Language - Binary
	LangTLB Language = "tlb"
	// TypoScript (alias: tsconfig)
	LangTypoScript Language = "typoscript"
	// UnrealScript (aliases: uscript, uc)
	LangUnreal Language = "unrealscript"
	// UO Razor Script
	LangUORazor Language = "uorazor"
	// V
	LangV Language = "v"
	// Vala
	LangVala Language = "vala"
	// Velocity
	LangVelocity Language = "velocity"
	// Verilog
	LangVerilog Language = "verilog"
	// VHDL
	LangVHDL Language = "vhdl"
	// vim
	LangVim Language = "vim"
	// Visual Basic (aliases: vb, vba)
	LangVisualBasic Language = "visual-basic"
	// WarpScript
	LangWarpScript Language = "warpscript"
	// WebAssembly
	LangWASM Language = "wasm"
	// Web IDL (alias: webidl)
	LangWebIDL Language = "web-idl"
	// WGSL
	LangWGSL Language = "wgsl"
	// Wiki markup
	LangWiki Language = "wiki"
	// Wolfram language (aliases: mathematica, nb, wl)
	LangWolfram Language = "wolfram"
	// Wren
	LangWren Language = "wren"
	// Xeora (alias: xeoracube)
	LangXeora Language = "xeora"
	// Xojo (REALbasic)
	LangXojo Language = "xojo"
	// XQuery
	LangXQuery Language = "xquery"
	// YANG
	LangYANG Language = "yang"
	// Zig
	LangZig Language = "zig"
)

// LanguageAliases maps any alias (lowercase) to its canonical Language.
// It is populated from libprisma’s Supported languages table at init().
var LanguageAliases map[string]Language

// SupportedLanguages returns a sorted list of all canonical languages known to this build.
func SupportedLanguages() []Language {
	out := make([]Language, 0, len(allCanonLangs))
	out = append(out, allCanonLangs...)
	sort.Slice(out, func(i, j int) bool { return out[i] < out[j] })
	return out
}

// ResolveLanguage normalizes an input (canonical or alias) to a canonical Language.
// Returns "" if the language is not supported.
func ResolveLanguage(input string) Language {
	key := strings.ToLower(strings.TrimSpace(input))
	if key == "" {
		return ""
	}
	if can, ok := LanguageAliases[key]; ok {
		return can
	}
	return ""
}

// --- internal: build aliases from the embedded table ---

// allCanonLangs is used by SupportedLanguages().
var allCanonLangs = []Language{
	// keep in sync with init() parse below
	LangMarkup, LangCSS, LangCLike, LangRegex, LangJavaScript, LangABAP, LangABNF, LangActionScript,
	LangAda, LangAgda, LangAL, LangANTLR4, LangApacheConf, LangSQL, LangApex, LangAPL, LangAppleScript,
	LangAQL, LangC, LangCPP, LangArduino, LangARFF, LangARMAsm, LangBash, LangYAML, LangMarkdown,
	LangArturo, LangAsciiDoc, LangCSharp, LangASPNet, LangAsm6502, LangAsmAtmel, LangAutoHotkey, LangAutoIt,
	LangAviSynth, LangAvroIDL, LangAWK, LangBASIC, LangBatch, LangBBCode, LangBBj, LangBicep, LangBirb,
	LangBison, LangBNF, LangBQN, LangBrainfuck, LangBrightScript, LangBro, LangCFScript, LangChaiScript,
	LangCIL, LangCilkC, LangCilkCPP, LangClojure, LangCMake, LangCOBOL, LangCoffeeScript, LangConcurnas,
	LangCSP, LangCooklang, LangRuby, LangCrystal, LangCSV, LangCUE, LangCypher, LangD, LangDart,
	LangDataWeave, LangDAX, LangDhall, LangDiff, LangMarkupTmpl, LangDjango, LangDNSZone, LangDocker,
	LangDOT, LangEBNF, LangEditorConfig, LangEiffel, LangEJS, LangElixir, LangElm, LangLua, LangETLua,
	LangERB, LangErlang, LangExcelFormula, LangFSharp, LangFactor, LangFalse, LangFift, LangFirestoreRules,
	LangFlow, LangFortran, LangFTL, LangFunC, LangGML, LangGAP, LangGCode, LangGDScript, LangGEDCOM,
	LangGettext, LangGit, LangGLSL, LangGN, LangLinkerScript, LangGo, LangGoModule, LangGradle, LangGraphQL,
	LangGroovy, LangLess, LangSCSS, LangTextile, LangHaml, LangHandlebars, LangHaskell, LangHaxe, LangHCL,
	LangHLSL, LangHoon, LangHPKP, LangHSTS, LangJSON, LangURI, LangHTTP, LangIchigoJam, LangIcon,
	LangICUMessage, LangIdris, LangIgnore, LangInform7, LangINI, LangIo, LangJ, LangJava, LangScala,
	LangPHP, LangJavadocLike, LangJavadoc, LangJavaStack, LangJolie, LangJQ, LangTypeScript, LangJSDoc,
	LangN4JS, LangJSON5, LangJSONP, LangJSStack, LangJulia, LangKeepalived, LangKeyman, LangKotlin,
	LangKusto, LangLaTeX, LangLatte, LangScheme, LangLilyPond, LangLiquid, LangLisp, LangLiveScript,
	LangLLVM, LangLog, LangLOLCODE, LangMagma, LangMakefile, LangMata, LangMATLAB, LangMAXScript, LangMEL,
	LangMermaid, LangMETAFONT, LangMizar, LangMongoDB, LangMonkey, LangMoonScript, LangN1QL, LangNand2Tetris,
	LangNaninovel, LangNASM, LangNEON, LangNevod, LangNginx, LangNim, LangNix, LangNSIS, LangObjectiveC,
	LangOCaml, LangOdin, LangOpenCL, LangOpenQasm, LangOz, LangPARIGP, LangParser, LangPascal, LangPascaligo,
	LangPSL, LangPCAxis, LangPeopleCode, LangPerl, LangPHPDoc, LangPlantUML, LangPLSQL, LangPowerQuery,
	LangPowerShell, LangProcessing, LangProlog, LangPromQL, LangProperties, LangProtoBuf, LangStylus,
	LangTwig, LangPug, LangPuppet, LangPureBasic, LangPython, LangQSharp, LangQ, LangQML, LangQore, LangR,
	LangRacket, LangRazorCS, LangJSX, LangTSX, LangReason, LangRego, LangRenPy, LangReScript, LangReST,
	LangRip, LangRoboconf, LangRobot, LangRust, LangSAS, LangSass, LangShellSession, LangSmali, LangSmalltalk,
	LangSmarty, LangSML, LangSolidity, LangSolution, LangSoy, LangSplunkSPL, LangSQF, LangSquirrel, LangStan,
	LangStata, LangIECST, LangSuperCollider, LangSwift, LangSystemd, LangTact, LangT4Tmpl, LangT4CS, LangVBNet,
	LangT4VB, LangTAP, LangTcl, LangTT2, LangTOML, LangTremor, LangTL, LangTLB, LangTypoScript, LangUnreal,
	LangUORazor, LangV, LangVala, LangVelocity, LangVerilog, LangVHDL, LangVim, LangVisualBasic, LangWarpScript,
	LangWASM, LangWebIDL, LangWGSL, LangWiki, LangWolfram, LangWren, LangXeora, LangXojo, LangXQuery, LangYANG, LangZig,
}

// libprismaTable is the “Supported languages” table copied from
// https://github.com/TelegramMessenger/libprisma#supported-languages
// Format: One row per line: <DisplayName>\t<aliases csv>  (aliases list includes the canonical key).
const libprismaTable = `
Markup	markup,markup,html,xml,svg,mathml,ssml,atom,rss
CSS	css
C-like	clike
Regex	regex
JavaScript	javascript,js
ABAP	abap
ABNF	abnf
ActionScript	actionscript
Ada	ada
Agda	agda
AL	al
ANTLR4	antlr4,g4
Apache Configuration	apacheconf
SQL	sql
Apex	apex
APL	apl
AppleScript	applescript
AQL	aql
C	c
C++	cpp
Arduino	arduino,ino
ARFF	arff
ARM Assembly	armasm,arm-asm
Bash	bash,bash,sh,shell
YAML	yaml,yml
Markdown	markdown,md
Arturo	arturo,art
AsciiDoc	asciidoc,adoc
C#	csharp,csharp,cs,dotnet
ASP.NET (C#)	aspnet
6502 Assembly	asm6502
Atmel AVR Assembly	asmatmel
AutoHotkey	autohotkey
AutoIt	autoit
AviSynth	avisynth,avs
Avro IDL	avro-idl,avdl
AWK	awk,gawk
BASIC	basic
Batch	batch
BBcode	bbcode,shortcode
BBj	bbj
Bicep	bicep
Birb	birb
Bison	bison
BNF	bnf,rbnf
BQN	bqn
Brainfuck	brainfuck
BrightScript	brightscript
Bro	bro
CFScript	cfscript,cfc
ChaiScript	chaiscript
CIL	cil
Cilk/C	cilkc,cilk-c
Cilk/C++	cilkcpp,cilkcpp,cilk-cpp,cilk
Clojure	clojure
CMake	cmake
COBOL	cobol
CoffeeScript	coffeescript,coffee
Concurnas	concurnas,conc
Content-Security-Policy	csp
Cooklang	cooklang
Ruby	ruby,rb
Crystal	crystal
CSV	csv
CUE	cue
Cypher	cypher
D	d
Dart	dart
DataWeave	dataweave
DAX	dax
Dhall	dhall
Diff	diff
Markup templating	markup-templating
Django/Jinja2	django,jinja2
DNS zone file	dns-zone-file,dns-zone
Docker	docker,dockerfile
DOT (Graphviz)	dot,gv
EBNF	ebnf
EditorConfig	editorconfig
Eiffel	eiffel
EJS	ejs,eta
Elixir	elixir
Elm	elm
Lua	lua
Embedded Lua templating	etlua
ERB	erb
Erlang	erlang
Excel Formula	excel-formula,excel-formula,xlsx,xls
F#	fsharp
Factor	factor
False	false
Fift	fift
Firestore security rules	firestore-security-rules
Flow	flow
Fortran	fortran
FreeMarker Template Language	ftl
FunC	func
GameMaker Language	gml,gamemakerlanguage
GAP (CAS)	gap
G-code	gcode
GDScript	gdscript
GEDCOM	gedcom
gettext	gettext,po
Git	git
GLSL	glsl
GN	gn,gni
GNU Linker Script	linker-script,ld
Go	go
Go module	go-module,go-mod
Gradle	gradle
GraphQL	graphql
Groovy	groovy
Less	less
Sass (SCSS)	scss
Textile	textile
Haml	haml
Handlebars	handlebars,handlebars,hbs,mustache
Haskell	haskell,hs
Haxe	haxe
HCL	hcl
HLSL	hlsl
Hoon	hoon
HTTP Public-Key-Pins	hpkp
HTTP Strict-Transport-Security	hsts
JSON	json,webmanifest
URI	uri,url
HTTP	http
IchigoJam	ichigojam
Icon	icon
ICU Message Format	icu-message-format
Idris	idris,idr
.ignore	ignore,ignore,gitignore,hgignore,npmignore
Inform 7	inform7
Ini	ini
Io	io
J	j
Java	java
Scala	scala
PHP	php
JavaDoc-like	javadoclike
JavaDoc	javadoc
Java stack trace	javastacktrace
Jolie	jolie
JQ	jq
TypeScript	typescript,ts
JSDoc	jsdoc
N4JS	n4js,n4jsd
JSON5	json5
JSONP	jsonp
JS stack trace	jsstacktrace
Julia	julia
Keepalived Configure	keepalived
Keyman	keyman
Kotlin	kotlin,kotlin,kt,kts
Kusto	kusto
LaTeX	latex,latex,tex,context
Latte	latte
Scheme	scheme
LilyPond	lilypond,ly
Liquid	liquid
Lisp	lisp,lisp,emacs,elisp,emacs-lisp
LiveScript	livescript
LLVM IR	llvm
Log file	log
LOLCODE	lolcode
Magma (CAS)	magma
Makefile	makefile
Mata	mata
MATLAB	matlab
MAXScript	maxscript
MEL	mel
Mermaid	mermaid
METAFONT	metafont
Mizar	mizar
MongoDB	mongodb
Monkey	monkey
MoonScript	moonscript,moon
N1QL	n1ql
Nand To Tetris HDL	nand2tetris-hdl
Naninovel Script	naniscript,nani
NASM	nasm
NEON	neon
Nevod	nevod
nginx	nginx
Nim	nim
Nix	nix
NSIS	nsis
Objective-C	objectivec,objc
OCaml	ocaml
Odin	odin
OpenCL	opencl
OpenQasm	openqasm,qasm
Oz	oz
PARI/GP	parigp
Parser	parser
Pascal	pascal,objectpascal
Pascaligo	pascaligo
PATROL Scripting Language	psl
PC-Axis	pcaxis,px
PeopleCode	peoplecode,pcode
Perl	perl
PHPDoc	phpdoc
PlantUML	plant-uml,plantuml
PL/SQL	plsql
PowerQuery	powerquery,powerquery,pq,mscript
PowerShell	powershell
Processing	processing
Prolog	prolog
PromQL	promql
.properties	properties
Protocol Buffers	protobuf
Stylus	stylus
Twig	twig
Pug	pug
Puppet	puppet
PureBasic	purebasic,pbfasm
Python	python,py
Q#	qsharp,qs
Q (kdb+ database)	q
QML	qml
Qore	qore
R	r
Racket	racket,rkt
Razor C#	cshtml,razor
React JSX	jsx
React TSX	tsx
Reason	reason
Rego	rego
Ren'py	renpy,rpy
ReScript	rescript,res
reST (reStructuredText)	rest
Rip	rip
Roboconf	roboconf
Robot Framework	robotframework,robot
Rust	rust
SAS	sas
Sass (Sass)	sass
Shell session	shell-session,shell-session,sh-session,shellsession
Smali	smali
Smalltalk	smalltalk
Smarty	smarty
SML	sml,smlnj
Solidity (Ethereum)	solidity,sol
Solution file	solution-file,sln
Soy (Closure Template)	soy
Splunk SPL	splunk-spl
SQF: Status Quo Function (Arma 3)	sqf
Squirrel	squirrel
Stan	stan
Stata Ado	stata
Structured Text (IEC 61131-3)	iecst
SuperCollider	supercollider,sclang
Swift	swift
Systemd configuration file	systemd
Tact	tact
T4 templating	t4-templating
T4 Text Templates (C#)	t4-cs,t4
VB.Net	vbnet
T4 Text Templates (VB)	t4-vb
TAP	tap
Tcl	tcl
Template Toolkit 2	tt2
TOML	toml
Tremor	tremor,tremor,trickle,troy
Type Language	tl
Type Language - Binary	tlb
TypoScript	typoscript,tsconfig
UnrealScript	unrealscript,unrealscript,uscript,uc
UO Razor Script	uorazor
V	v
Vala	vala
Velocity	velocity
Verilog	verilog
VHDL	vhdl
vim	vim
Visual Basic	visual-basic,visual-basic,vb,vba
WarpScript	warpscript
WebAssembly	wasm
Web IDL	web-idl,webidl
WGSL	wgsl
Wiki markup	wiki
Wolfram language	wolfram,wolfram,mathematica,nb,wl
Wren	wren
Xeora	xeora,xeoracube
Xojo (REALbasic)	xojo
XQuery	xquery
YANG	yang
Zig	zig
`

func init() {
	LanguageAliases = make(map[string]Language, 256)

	lines := strings.Split(strings.TrimSpace(libprismaTable), "\n")
	for _, line := range lines {
		// each line: "<DisplayName>\t<aliases csv>"
		parts := strings.Split(line, "\t")
		if len(parts) != 2 {
			continue
		}
		aliasesCSV := parts[1]
		aliases := strings.Split(aliasesCSV, ",")
		var canonical Language
		for i, a := range aliases {
			key := strings.ToLower(strings.TrimSpace(a))
			if key == "" {
				continue
			}
			// First alias that matches any known canonical name becomes canonical key,
			// otherwise we use the first token as canonical.
			if i == 0 {
				canonical = Language(key)
			}
			// map alias -> canonical
			LanguageAliases[key] = canonical
		}
	}

	// Manually ensure we map well-known alternates (harmless if present already)
	LanguageAliases["html"] = LangMarkup
	LanguageAliases["xml"] = LangMarkup
	LanguageAliases["svg"] = LangMarkup
	LanguageAliases["js"] = LangJavaScript
	LanguageAliases["ts"] = LangTypeScript
	LanguageAliases["py"] = LangPython
	LanguageAliases["go-mod"] = LangGoModule
	LanguageAliases["go-module"] = LangGoModule
	LanguageAliases["c++"] = LangCPP
}

// String implements fmt.Stringer for logging/debug purposes.
func (l Language) String() string { return string(l) }

// MustResolveLanguage resolves an input to a canonical Language or panics with a descriptive error.
// Prefer ResolveLanguage in production code; use MustResolveLanguage in tests.
func MustResolveLanguage(input string) Language {
	if lang := ResolveLanguage(input); lang != "" {
		return lang
	}
	panic(fmt.Errorf("unsupported language %q (see libprisma Supported languages)", input))
}
